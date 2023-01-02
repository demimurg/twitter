package grace

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

const (
	SoftLimit = 5 * time.Second  // cancel requests context
	HardLimit = 10 * time.Second // immediately stop server
)

// Routine is a simple interface for different server type
type Routine interface {
	// Run blocks until application stops
	Run() error
	// Shutdown will gracefully shutdown routine
	Shutdown() error
}

// -------------------------- GRPC --------------------------

// GRPC adapts grpc.Server to the Routine interface accepts port in the ":8080" form.
// Cancel all requests (unary and stream) context after 5s soft limit, and waits rpcs
func GRPC(srv *grpc.Server, addr string) Routine {
	return &grpcRoutine{srv, addr}
}

type grpcRoutine struct {
	*grpc.Server
	addr string
}

func (g *grpcRoutine) Run() error {
	lis, err := net.Listen("tcp", g.addr)
	if err != nil {
		return err
	}

	// always returns error. ErrServerStopped on graceful close
	err = g.Serve(lis)
	if err != nil && err != grpc.ErrServerStopped {
		return err
	}
	return nil
}

func (g *grpcRoutine) Shutdown() error {
	done := make(chan struct{})
	go func() {
		g.GracefulStop()
		close(done)
	}()

	select {
	case <-time.After(SoftLimit):
		g.Stop()                          // cancel all contexts and unblock graceful stop
		time.Sleep(HardLimit - SoftLimit) // give rpc time to finish
	case <-done:
	}

	return nil
}

// -------------------------- HTTP --------------------------

// HTTP adapts http.Server to the Routine interface, accepts port in the ":8080" form.
//
// All requests context will be done after 5s calling shutdown method.
// hijackedWG is an optional waitgroup to wait for a graceful shutdown
// of websockets or server-sent events. HTTP server doesn't care about them itself
func HTTP(srv *http.Server, addr string, hijackedWG ...*sync.WaitGroup) Routine {
	if hijackedWG != nil {
		return &httpRoutine{Server: srv, addr: addr, hijackedConns: hijackedWG[0]}
	}
	return &httpRoutine{Server: srv, addr: addr}
}

type httpRoutine struct {
	*http.Server
	addr              string
	cancelRequestsCtx func()          // call when you need all request context to be done
	hijackedConns     *sync.WaitGroup // waitgroup for all hijacked connections (websockets, server sent events)
}

func (h *httpRoutine) Run() error {
	lis, err := net.Listen("tcp", h.addr)
	if err != nil {
		return err
	}

	// by default http server doesn't cancel request
	// context on shutdown, so we add this capability
	ctx, cancel := context.WithCancel(context.Background())
	h.cancelRequestsCtx = cancel
	h.BaseContext = func(l net.Listener) context.Context {
		return ctx
	}

	// always returns error. ErrServerClosed on graceful close
	err = h.Serve(lis)
	if err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (h *httpRoutine) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), HardLimit)

	go func() {
		select {
		case <-time.After(SoftLimit):
			h.cancelRequestsCtx()
		case <-ctx.Done():
			return
		}
	}()

	err := h.Server.Shutdown(ctx)

	if h.hijackedConns != nil {
		go func() {
			h.hijackedConns.Wait()
			cancel()
		}()
	} else {
		cancel()
	}

	<-ctx.Done()
	return err
}

// ------------------------ Prometheus ------------------------

// PromHTTP add prometheus http server, accepts port in the ":8080" form.
//
// It design to used with "github.com/prometheus/client_golang/prometheus/promauto"
// You can create promauto metrics and this handler will catch them automatically.
//
// You can curl text file with the all metrics on host "addr" and "/" path (e.g. curl localhost:9090)
func PromHTTP(addr string) Routine {
	return &prometheusRoutine{httpRoutine{
		Server: &http.Server{Handler: promhttp.Handler()},
		addr:   addr,
	}}
}

type prometheusRoutine struct {
	// embed to change routine type in logging (%T used)
	httpRoutine
}
