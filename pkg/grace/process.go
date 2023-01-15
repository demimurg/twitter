package grace

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/fullstorydev/grpcui/standalone"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

const (
	SoftLimit = 5 * time.Second  // cancel requests context
	HardLimit = 10 * time.Second // immediately stop server
)

// Process is a simple interface for long running processes
type Process interface {
	// Name will be logged on start
	Name() string
	// Run blocks until application stops
	Run() error
	// Shutdown will gracefully shutdown process
	Shutdown() error
}

// -------------------------- GRPC --------------------------

// GRPC adapts grpc.Server to the Process interface, accepts port in the ":8080" form.
// Cancel all requests (unary and stream) context after 5s soft limit, and waits rpcs
func GRPC(srv *grpc.Server, addr string) Process {
	return &grpcProc{srv, addr}
}

type grpcProc struct {
	*grpc.Server
	addr string
}

func (g *grpcProc) Name() string {
	return "GRPC " + g.addr
}

func (g *grpcProc) Run() error {
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

func (g *grpcProc) Shutdown() error {
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

// HTTP adapts http.Server to the Process interface, accepts port in the ":8080" form.
//
// All requests context will be done after 5s calling shutdown method.
// hijackedWG is an optional waitgroup to wait for a graceful shutdown
// of websockets or server-sent events. HTTP server doesn't care about them itself
func HTTP(srv *http.Server, addr string, hijackedWG ...*sync.WaitGroup) Process {
	if hijackedWG != nil {
		return &httpProc{Server: srv, addr: addr, hijackedConns: hijackedWG[0]}
	}
	return &httpProc{Server: srv, addr: addr}
}

type httpProc struct {
	*http.Server
	addr              string
	cancelRequestsCtx func()          // call when you need all request context to be done
	hijackedConns     *sync.WaitGroup // waitgroup for all hijacked connections (websockets, server sent events)
}

func (h *httpProc) Name() string {
	return "HTTP " + h.addr
}

func (h *httpProc) Run() error {
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

func (h *httpProc) Shutdown() error {
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

// Prometheus add prometheus http server, accepts port in the ":8080" form.
//
// It design to used with "github.com/prometheus/client_golang/prometheus/promauto"
// You can create promauto metrics and this handler will catch them automatically.
//
// You can curl text file with the all metrics on host "addr" and "/" path (e.g. curl localhost:9090)
func Prometheus(addr string) Process {
	return &prometheusProc{httpProc{
		Server: &http.Server{Handler: promhttp.Handler()},
		addr:   addr,
	}}
}

type prometheusProc struct {
	// embed to change process type in logging (%T used)
	httpProc
}

func (p *prometheusProc) Name() string {
	return "Prometheus " + p.httpProc.addr
}

// ------------------------ GRPC UI ------------------------

// GRPCUI will create http server that serve front-end
// for grpc requests to application (analogue swagger)
//
// Read more: https://github.com/fullstorydev/grpcui
func GRPCUI(addr, appServerAddr string) Process {
	return &grpcuiProc{
		appServerAddr: appServerAddr,
		httpProc: httpProc{
			// will init handler in Run
			Server: &http.Server{},
			addr:   addr,
		},
	}
}

type grpcuiProc struct {
	httpProc
	appServerAddr string
}

func (g *grpcuiProc) Name() string {
	return "GRPC UI " + g.httpProc.addr
}

func (g *grpcuiProc) Run() error {
	ctx := context.Background()
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.DialContext(ctx, g.appServerAddr, creds)
	if err != nil {
		return fmt.Errorf("can't dial app server: %w", err)
	}

	handler, err := standalone.HandlerViaReflection(ctx, cc, cc.Target())
	if err != nil {
		return fmt.Errorf("grpcui can't init handler")
	}
	g.Server.Handler = handler

	return g.httpProc.Run()
}
