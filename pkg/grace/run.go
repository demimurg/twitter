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

// Run will start Process's in parallel and take care about graceful shutdown.
// If one process stops every other will stop too
func Run(processes ...Process) {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer stop()

	ctx = log.With(ctx, "package", "grace")

	var wg sync.WaitGroup
	wg.Add(len(processes) * 2)

	for _, proc := range processes {
		proc := proc // clone for using in goroutine
        ctx := log.With(ctx, "name", proc.Name())

		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Error(ctx, "panic recovered",
						"error", fmt.Errorf("panic: %v", r),
						"stack", debug.Stack())
				}
				stop() // first stopped process will cancel context and other processes will shut down
				wg.Done()
			}()

			log.Info(ctx, "start process")
			if err := proc.Run(); err != nil {
				log.Error(ctx, "process unexpectedly stopped",
					"error", err)
			}
		}()
	}

	<-ctx.Done()
	for _, proc := range processes {
		proc := proc
		ctx := log.With(ctx, "name", proc.Name())

		go func() {
			defer wg.Done()
			err := proc.Shutdown()
			if err != nil {
				log.Error(ctx, "stop process",
					"error", err)
			}
			log.Info(ctx, "stop process")
		}()
	}

	wg.Wait()
}
