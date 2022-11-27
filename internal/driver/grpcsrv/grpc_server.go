package grpcsrv

import (
	"context"
	"errors"
	"github.com/demimurg/twitter/internal/usecase"
	"github.com/demimurg/twitter/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func NewTwitter(feedManager usecase.FeedManager, userRegistrator usecase.UserRegistrator) *grpc.Server {
	srv := grpc.NewServer()
	proto.RegisterTwitterServer(srv, &twitter{
		fm: feedManager, ur: userRegistrator,
	})
	return srv
}

type twitter struct {
	proto.UnimplementedTwitterServer
	fm usecase.FeedManager
	ur usecase.UserRegistrator
}

func (t *twitter) AddTweet(ctx context.Context, req *proto.AddTweetRequest) (*emptypb.Empty, error) {
	err := t.fm.AddNewTweet(ctx, int(req.UserId), req.Text)
	if err != nil {
		if errors.Is(err, usecase.ErrValidationFailed) {
			return &emptypb.Empty{}, status.Error(codes.InvalidArgument, err.Error())
		}
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (t *twitter) GetNewsFeed(ctx context.Context, req *proto.GetNewsFeedRequest) (*proto.GetNewsFeedResponse, error) {
	feed, err := t.fm.GiveNewsFeed(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}

	resp := &proto.GetNewsFeedResponse{Tweets: make([]string, len(feed))}
	for i, tweet := range feed {
		resp.Tweets[i] = tweet.Text
	}
	return resp, nil
}

func (t *twitter) Register(ctx context.Context, req *proto.RegisterRequest) (*emptypb.Empty, error) {
	_, err := t.ur.Register(ctx, req.FullName, req.Email, req.DateOfBirth)
	if err != nil {
		if errors.Is(err, usecase.ErrValidationFailed) {
			return &emptypb.Empty{}, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	return &emptypb.Empty{}, nil
}

func (t *twitter) Follow(ctx context.Context, req *proto.FollowRequest) (*emptypb.Empty, error) {
	err := t.fm.AddFollower(ctx, int(req.UserId), int(req.NewFollowerId))
	if err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
