package grpcsrv

import (
	"context"
	"errors"
	"strings"

	"github.com/demimurg/twitter/internal/entity"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/demimurg/twitter/internal/usecase"
	"github.com/demimurg/twitter/pkg/log"
	"github.com/demimurg/twitter/pkg/proto"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewTwitter(feedManager usecase.FeedManager, userRegistrator usecase.UserRegistrator) *grpc.Server {
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(logRequest, recoverPanic),
	)
	proto.RegisterTwitterServer(srv, &twitter{
		fm: feedManager, ur: userRegistrator,
	})
	reflection.Register(srv)
	return srv
}

type twitter struct {
	proto.UnimplementedTwitterServer
	fm usecase.FeedManager
	ur usecase.UserRegistrator
}

func logRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
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
	return resp, err
}

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
