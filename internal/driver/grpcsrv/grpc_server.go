package grpcsrv

import (
	"context"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc/codes"

	"github.com/golang-jwt/jwt/v5"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc/status"

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

func NewTwitter(feedManager usecase.FeedManager, userProfiler usecase.UserProfiler, authSecret string) *grpc.Server {
	twitterImpl := &twitter{fm: feedManager, up: userProfiler, secret: authSecret}
	grpcSrv := grpc.NewServer(grpc.ChainUnaryInterceptor(
		takeObservation, recoverPanic,
		validator.UnaryServerInterceptor(),
		twitterImpl.authorize,
	))

	proto.RegisterTwitterServer(grpcSrv, twitterImpl)
	reflection.Register(grpcSrv)
	return grpcSrv
}

type twitter struct {
	proto.UnimplementedTwitterServer
	fm usecase.FeedManager
	up usecase.UserProfiler

	// secret used for jwt tokens creation
	secret string
}

var requestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name: "grpc_request_duration_ms",
	Help: "Time (in milliseconds) spent serving GRPC requests",
}, []string{"method", "status_code"})

// takeObservation will create trace_id for each request, log and gather metric
func takeObservation(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
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
func recoverPanic(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
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

// authorize will extract jwt token from request, validate and prepare user_id for rpc work
func (t *twitter) authorize(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	// we have only 2 methods that doesn't use jwt tokens
	if strings.HasSuffix(info.FullMethod, "Register") || strings.HasSuffix(info.FullMethod, "Login") {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("Authorization")) == 0 {
		return nil, status.Error(codes.Unauthenticated, "have no Authorization metadata")
	}

	jwtToken := strings.TrimPrefix(md.Get("Authorization")[0], "Bearer ")
	userID, err := t.extractUserID(jwtToken)
	if err != nil {
		return nil, err
	}
	ctx = withUserID(ctx, userID)

	return handler(ctx, req)
}

type ctxAuthKey uint

const authKey = ctxAuthKey(0)

func withUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, authKey, userID)
}

func getUserID(ctx context.Context) int {
	return ctx.Value(authKey).(int)
}

func (t *twitter) extractUserID(jwtToken string) (int, error) {
	token, err := jwt.Parse(jwtToken, func(*jwt.Token) (any, error) {
		return []byte(t.secret), nil
	})
	if err != nil {
		return 0, err
	}
	if !token.Valid {
		return 0, status.Error(codes.Unauthenticated, "jwt token is not valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, status.Error(codes.Unauthenticated, "have no claims in jwt token")
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, status.Error(codes.Unauthenticated, "user_id should be added to jwt token")
	}

	return int(userID), nil
}
