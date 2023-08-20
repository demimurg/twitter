//go:build e2e

package tests

import (
	"context"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/demimurg/twitter/pkg/proto"
)

// TestFeed for basic operations, subtests can't be run separate
// it should be understood like one related story
func (s *endToEndTestSuite) TestFeed() {
	var (
		elonID      int64
		elonJWT     string
		elonProfile *proto.UserProfile
		elonTweetID int64
		elonTweet   = "hey guys, should i buy twitter?"
	)

	s.Run("first elon musk tweet", func() {
		elonProfile = &proto.UserProfile{
			FullName:    "Elon Musk",
			Email:       "elonID@tesla.us",
			DateOfBirth: date(1971, 06, 28),
		}
		resp, err := s.cli.Register(ctx, &proto.RegisterRequest{User: elonProfile, Password: "AE A-XII"})
		s.NoError(err)
		elonID = resp.UserId
		elonJWT = resp.Jwt

		tResp, err := s.cli.AddTweet(withToken(ctx, elonJWT), &proto.AddTweetRequest{Text: elonTweet})
		s.NoError(err)
		elonTweetID = tResp.TweetId
	})

	s.Run("elon updated his tweet", func() {
		updatedTweet := "i think there too much fake accounts and stocks too high"
		_, err := s.cli.UpdateTweet(withToken(ctx, elonJWT), &proto.UpdateTweetRequest{
			TweetId: elonTweetID, NewText: updatedTweet,
		})
		s.NoError(err)
		elonTweet = updatedTweet
	})

	s.Run("unregistered user can't send tweet", func() {
		_, err := s.cli.AddTweet(withToken(ctx, "fake jwt"), &proto.AddTweetRequest{
			Text: "scam message",
		})
		s.Error(err) // who the f*ck are you??
	})

	var (
		amberID  int64
		amberJWT string
	)
	s.Run("elon have new follower amber", func() {
		amberProfile := &proto.UserProfile{
			FullName:    "Amber Heard",
			Email:       "makeyou@cry.com",
			DateOfBirth: date(1986, 04, 22),
		}
		resp, err := s.cli.Register(ctx, &proto.RegisterRequest{
			User: amberProfile, Password: "bye-bye-jack",
		})
		s.Require().NoError(err)
		amberID = resp.UserId
		amberJWT = resp.Jwt
		s.NotEqual(elonID, amberID, "different users should have different ids")

		amberCtx := withToken(ctx, amberJWT)
		_, err = s.cli.Follow(amberCtx, &proto.FollowRequest{
			FolloweeId: elonID,
		})
		s.Require().NoError(err, "ember send follow to elon")

		elonCtx := withToken(ctx, elonJWT)
		fresp, err := s.cli.GetFollowers(elonCtx, &proto.GetFollowersRequest{UserId: elonID})
		s.Require().NoError(err, "get elon followers")
		s.Require().Len(fresp.Users, 1, "elon followed only be amber")
		s.EqualProto(amberProfile, fresp.Users[0], "follower user equal amber profile")
	})

	s.Run("amber checks the news feed", func() {
		ctx := withToken(ctx, amberJWT)
		fresp, err := s.cli.GetFollowing(ctx, &proto.GetFollowingRequest{UserId: amberID})
		s.Require().NoError(err, "get users that amber following")
		s.Require().Len(fresp.Users, 1, "amber should folow only elon")
		s.EqualProto(elonProfile, fresp.Users[0], "followed user equal elon profile")

		resp, err := s.cli.GetNewsFeed(ctx, &proto.GetNewsFeedRequest{Limit: 10})
		s.NoError(err)

		s.Len(resp.Tweets, 1)
		s.Equal(elonTweet, resp.Tweets[0])
	})

	s.Run("amber adds comment to elon tweet and update it", func() {
		ctx := withToken(ctx, amberJWT)
		resp, err := s.cli.AddComment(ctx, &proto.AddCommentRequest{
			TweetId: elonTweetID, Text: "wanna go on a date?",
		})
		s.Require().NoError(err)

		<-time.After(time.Millisecond) // waits elon response
		_, err = s.cli.UpdateComment(ctx, &proto.UpdateCommentRequest{
			CommentId: resp.CommentId, NewText: "nevermind, idiot.",
		})
		s.NoError(err)
	})

	s.Run("amber wants to unfollow elon", func() {
		ctx := withToken(ctx, amberJWT)
		_, err := s.cli.Unfollow(ctx, &proto.UnfollowRequest{
			FolloweeId: elonID,
		})
		s.Require().NoError(err)

		resp, err := s.cli.GetNewsFeed(ctx, &proto.GetNewsFeedRequest{
			Limit: 10,
		})
		s.Require().NoError(err)
		s.Len(resp.Tweets, 0) // amber have no friends😢
	})

	s.Run("amber checks user recommendations", func() {
		ctx := withToken(ctx, amberJWT)
		resp, err := s.cli.RecommendUsers(ctx, &proto.RecommendUsersRequest{})
		s.Require().NoError(err)

		s.GreaterOrEqual(len(resp.Users), 1)
		s.T().Log(resp.Users)
	})
}

func withToken(ctx context.Context, jwtTok string) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "Authorization", "Bearer "+jwtTok)
}
