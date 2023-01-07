package grpcsrv

import (
    "context"
    "errors"
    "github.com/demimurg/twitter/internal/usecase"
    "github.com/demimurg/twitter/pkg/proto"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    "google.golang.org/protobuf/types/known/emptypb"
)

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
