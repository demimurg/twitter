package grpcsrv

import (
	"context"
	"errors"

	"github.com/demimurg/twitter/internal/entity"

	"github.com/demimurg/twitter/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (t *twitter) AddTweet(ctx context.Context, req *proto.AddTweetRequest) (*proto.AddTweetResponse, error) {
	tweetID, err := t.fm.AddTweet(ctx, getUserID(ctx), req.Text)
	if err != nil {
		if errors.Is(err, entity.ErrValidationFailed) {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, err
	}
	return &proto.AddTweetResponse{TweetId: int64(tweetID)}, nil
}

func (t *twitter) AddComment(ctx context.Context, req *proto.AddCommentRequest) (*proto.AddCommentResponse, error) {
	commentID, err := t.fm.AddComment(ctx, getUserID(ctx), int(req.TweetId), req.Text)
	if err != nil {
		return nil, err
	}
	return &proto.AddCommentResponse{CommentId: int64(commentID)}, nil
}
func (t *twitter) UpdateTweet(ctx context.Context, req *proto.UpdateTweetRequest) (*proto.UpdateTweetResponse, error) {
	err := t.fm.EditTweet(ctx, int(req.TweetId), req.NewText)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateTweetResponse{}, nil
}
func (t *twitter) UpdateComment(ctx context.Context, req *proto.UpdateCommentRequest) (*proto.UpdateCommentResponse, error) {
	err := t.fm.EditComment(ctx, int(req.CommentId), req.NewText)
	if err != nil {
		return nil, err
	}
	return &proto.UpdateCommentResponse{}, nil
}

func (t *twitter) GetNewsFeed(ctx context.Context, _ *proto.GetNewsFeedRequest) (*proto.GetNewsFeedResponse, error) {
	feed, err := t.fm.GetNewsFeed(ctx, getUserID(ctx))
	if err != nil {
		return nil, err
	}

	resp := &proto.GetNewsFeedResponse{Tweets: make([]string, len(feed))}
	for i, tweet := range feed {
		resp.Tweets[i] = tweet.Text
	}
	return resp, nil
}

func (t *twitter) Follow(ctx context.Context, req *proto.FollowRequest) (*proto.FollowResponse, error) {
	err := t.fm.AddFollower(ctx, getUserID(ctx), int(req.FolloweeId))
	if err != nil {
		return nil, err
	}
	return &proto.FollowResponse{}, nil
}

func (t *twitter) Unfollow(ctx context.Context, req *proto.UnfollowRequest) (*proto.UnfollowResponse, error) {
	err := t.fm.RemoveFollower(ctx, getUserID(ctx), int(req.FolloweeId))
	if err != nil {
		return nil, err
	}
	return &proto.UnfollowResponse{}, nil
}

func (t *twitter) RecommendUsers(ctx context.Context, _ *proto.RecommendUsersRequest) (*proto.RecommendUsersResponse, error) {
	users, err := t.fm.GetRecommendedUsers(ctx, getUserID(ctx))
	if err != nil {
		return nil, err
	}
	return &proto.RecommendUsersResponse{
		Users: convertToUserProfiles(users),
	}, nil
}

func (t *twitter) GetFollowing(ctx context.Context, req *proto.GetFollowingRequest) (*proto.GetFollowingResponse, error) {
	users, err := t.fm.GetFollowing(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}
	return &proto.GetFollowingResponse{
		Users: convertToUserProfiles(users),
	}, nil
}

func (t *twitter) GetFollowers(ctx context.Context, req *proto.GetFollowersRequest) (*proto.GetFollowersResponse, error) {
	users, err := t.fm.GetFollowers(ctx, int(req.UserId))
	if err != nil {
		return nil, err
	}
	return &proto.GetFollowersResponse{
		Users: convertToUserProfiles(users),
	}, nil
}

func convertToUserProfiles(users []entity.User) []*proto.UserProfile {
	protoUsers := make([]*proto.UserProfile, 0, len(users))
	for _, user := range users {
		protoUsers = append(protoUsers, convertToUserProfile(&user))
	}
	return protoUsers
}
