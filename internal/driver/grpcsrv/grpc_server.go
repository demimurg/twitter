package grpcsrv

import (
    "context"
    "github.com/prometheus/client_golang/prometheus"
    "google.golang.org/grpc/status"
    "strings"
    "time"

    "github.com/demimurg/twitter/internal/usecase"
    "github.com/demimurg/twitter/pkg/log"
    "github.com/demimurg/twitter/pkg/proto"
    "github.com/google/uuid"
    validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
    "github.com/prometheus/client_golang/prometheus/promauto"
    "go.uber.org/zap"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

func NewTwitter(feedManager usecase.FeedManager, userProfiler usecase.UserProfiler) *grpc.Server {
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(takeObservation, recoverPanic, validator.UnaryServerInterceptor()),
	)
	proto.RegisterTwitterServer(srv, &twitter{
		fm: feedManager, up: userProfiler,
	})
	reflection.Register(srv)
	return srv
}

type twitter struct {
	proto.UnimplementedTwitterServer
	fm usecase.FeedManager
	up usecase.UserProfiler
}

var requestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
    Name: "grpc_request_duration_ms",
    Help: "Time (in milliseconds) spent serving GRPC requests",
}, []string{"method", "status_code"})



// takeObservation will create trace_id for each request, log and gather metric
func takeObservation(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
    start := time.Now()
    uuidParts := strings.Split(uuid.NewString(), "-")
	ctx = log.With(ctx, "trace_id", uuidParts[len(uuidParts)-1])

	method := info.FullMethod
	if i := strings.LastIndex(info.FullMethod, "/"); i != -1 && i != len(info.FullMethod)-1 {
		method = info.FullMethod[i+1:]
	}
    log.Info(ctx, "received request",
        "method", method)
	resp, err = handler(ctx, req)
	if err != nil {
		log.Error(ctx, "returned error in response",
			"method", method,
			"error", err)
	}

    requestDuration.
        WithLabelValues(method, status.Code(err).String()).
        Observe(float64(time.Since(start).Milliseconds()))

	return resp, err
}

// recoverPanic will save our app from crash in case of some exception will occure while serving request
func recoverPanic(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if msg := recover(); msg != nil {
			// print stack trace of panic in json
			// the first frames will be with runtime package info
			// but the next one is a place where panic appeared
			log.Error(ctx, "rpc throw panic",
				"panic", msg,
				zap.StackSkip("stacktrace", 1), // skip current func
				"todo", "copy stacktrace and print with newlines using `echo -e <stactrace>`, search problem line from up to down")
		}
	}()
	return handler(ctx, req)
}
