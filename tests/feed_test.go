package tests

import (
	"github.com/demimurg/twitter/pkg/proto"
	"time"
)

// TestFeed for basic operations, subtests can't be run separate
// it should be understand like one related story
func (s *endToEndTestSuite) TestFeed() {
	var (
		elonID      int64
		elonTweetID int64
		elonTweet   = "hey guys, should i buy twitter?"
	)
	s.Run("first elon musk tweet", func() {
		resp, err := s.cli.Register(ctx, &proto.RegisterRequest{User: &proto.UserProfile{
			FullName:    "Elon Musk",
			Email:       "elonID@tesla.us",
			DateOfBirth: date(1971, 06, 28),
		}})
		s.NoError(err)
		elonID = resp.UserId

		tResp, err := s.cli.AddTweet(ctx, &proto.AddTweetRequest{
			Text: elonTweet, UserId: elonID,
		})
		s.NoError(err)
		elonTweetID = tResp.TweetId
	})

	s.Run("elon updated his tweet", func() {
		updatedTweet := "i think there too much fake accounts and stocks too high"
		_, err := s.cli.UpdateTweet(ctx, &proto.UpdateTweetRequest{
			UserId: elonID, TweetId: elonTweetID, NewText: updatedTweet,
		})
		s.NoError(err)
		elonTweet = updatedTweet
	})

	s.Run("unregistered user can't send tweet", func() {
		_, err := s.cli.AddTweet(ctx, &proto.AddTweetRequest{
			Text: "scam message", UserId: -1, // who the f*ck are you??
		})
		s.Error(err)
	})

	var amberID int64
	s.Run("elon have new follower amber", func() {
		resp, err := s.cli.Register(ctx, &proto.RegisterRequest{User: &proto.UserProfile{
			FullName:    "Amber Heard",
			Email:       "makeyou@cry.com",
			DateOfBirth: date(1986, 04, 22),
		}})
		s.NoError(err)
		amberID = resp.UserId
		s.NotEqual(elonID, amberID, "different users should have different ids")

		_, err = s.cli.Follow(ctx, &proto.FollowRequest{
			UserId:        elonID,
			NewFollowerId: amberID,
		})
		s.NoError(err)
	})

	s.Run("amber checks the news feed", func() {
		resp, err := s.cli.GetNewsFeed(ctx, &proto.GetNewsFeedRequest{
			UserId: amberID, Limit: 10,
		})
		s.NoError(err)

		s.Len(resp.Tweets, 1)
		s.Equal(elonTweet, resp.Tweets[0])
	})

	s.Run("amber adds comment to elon tweet and update it", func() {
		resp, err := s.cli.AddComment(ctx, &proto.AddCommentRequest{
			UserId: amberID, TweetId: elonTweetID, Text: "wanna go on a date?",
		})
		s.Require().NoError(err)

		<-time.After(time.Millisecond) // waits elon response
		_, err = s.cli.UpdateComment(ctx, &proto.UpdateCommentRequest{
			UserId: amberID, CommentId: resp.CommentId, NewText: "nevermind, idiot.",
		})
		s.NoError(err)
	})

	s.Run("amber wants to unfollow elon", func() {
		_, err := s.cli.Unfollow(ctx, &proto.UnfollowRequest{
			UserId:        elonID,
			OldFollowerId: amberID,
		})
		s.Require().NoError(err)

		resp, err := s.cli.GetNewsFeed(ctx, &proto.GetNewsFeedRequest{
			UserId: amberID, Limit: 10,
		})
		s.Require().NoError(err)
		s.Len(resp.Tweets, 0) // amber have no friends😢
	})

	s.Run("amber checks user recommendations", func() {
		resp, err := s.cli.RecommendUsers(ctx, &proto.RecommendUsersRequest{UserId: amberID})
		s.Require().NoError(err)

		s.GreaterOrEqual(len(resp.Users), 1)
		s.T().Log(resp.Users)
	})
}
