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

func (t *twitter) AddTweet(ctx context.Context, req *proto.AddTweetRequest) (*emptypb.Empty, error) {
	err := t.fm.AddTweet(ctx, int(req.UserId), req.Text)
	if err != nil {
		if errors.Is(err, usecase.ErrValidationFailed) {
			return &emptypb.Empty{}, status.Error(codes.InvalidArgument, err.Error())
		}
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (t *twitter) GetNewsFeed(ctx context.Context, req *proto.GetNewsFeedRequest) (*proto.GetNewsFeedResponse, error) {
	feed, err := t.fm.GetNewsFeed(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}

	resp := &proto.GetNewsFeedResponse{Tweets: make([]string, len(feed))}
	for i, tweet := range feed {
		resp.Tweets[i] = tweet.Text
	}
	return resp, nil
}

func (t *twitter) Register(ctx context.Context, req *proto.UserProfile) (*proto.RegisterResponse, error) {
	user, err := t.ur.Register(ctx, req.FullName, req.Email, req.DateOfBirth.AsTime())
	if err != nil {
		if errors.Is(err, usecase.ErrValidationFailed) {
			err = status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, err
	}
	return &proto.RegisterResponse{UserId: int64(user.ID)}, nil
}

func (t *twitter) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	user, err := t.ur.Login(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{
		UserId:      int64(user.ID),
		UserProfile: convertToUserProfile(user),
	}, nil
}

func convertToUserProfile(user *entity.User) *proto.UserProfile {
	return &proto.UserProfile{
		Email:       user.Email,
		FullName:    user.FullName,
		Caption:     user.Caption,
		DateOfBirth: timestamppb.New(user.BirthDate),
	}
}

func (t *twitter) Follow(ctx context.Context, req *proto.FollowRequest) (*emptypb.Empty, error) {
	err := t.fm.AddFollower(ctx, int(req.NewFollowerId), int(req.UserId))
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (t *twitter) Unfollow(ctx context.Context, req *proto.UnfollowRequest) (*emptypb.Empty, error) {
	err := t.fm.RemoveFollower(ctx, int(req.OldFollowerId), int(req.UserId))
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (t *twitter) RecommendUsers(ctx context.Context, req *proto.RecommendUsersRequest) (*proto.RecommendUsersResponse, error) {
	users, err := t.fm.GetRecommendedUsers(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}

	protoUsers := make([]*proto.UserProfile, 0, len(users))
	for _, user := range users {
		protoUsers = append(protoUsers, convertToUserProfile(&user))
	}
	return &proto.RecommendUsersResponse{Users: protoUsers}, nil
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
