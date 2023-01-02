package grace

import (
	"context"
	"fmt"
	"os/signal"
	"runtime/debug"
	"sync"
	"syscall"

	"github.com/demimurg/twitter/pkg/log"
)

// Run will start Routine's in parallel and take care about graceful shutdown.
// If one routine stops every other will stop too
func Run(routines ...Routine) {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer stop()

	ctx = log.With(ctx, "package", "graceful")

	var wg sync.WaitGroup
	wg.Add(len(routines) * 2)

	for _, r := range routines {
		r := r // clone for using in goroutine
		ctx := log.With(ctx, "type", fmt.Sprintf("%T", r))

		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Error(ctx, "panic recovered",
						"error", fmt.Errorf("panic: %v", r),
						"stack", debug.Stack())
				}
				stop() // first stopped routine will cancel context and other routines will shut down
				wg.Done()
			}()

			log.Info(ctx, "start routine")
			if err := r.Run(); err != nil {
				log.Error(ctx, "routine unexpectedly stopped",
					"error", err)
			}
		}()
	}

	<-ctx.Done()
	for _, r := range routines {
		r := r
		ctx := log.With(ctx, "type", fmt.Sprintf("%T", r))

		go func() {
			defer wg.Done()
			err := r.Shutdown()
			if err != nil {
				log.Error(ctx, "stop routine",
					"error", err)
			}
			log.Info(ctx, "stop routine")
		}()
	}

	wg.Wait()
}
