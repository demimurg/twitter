package tests

import (
    "github.com/demimurg/twitter/pkg/proto"
)

func (s *endToEndTestSuite) TestFeed() {
	var elonID int64
	s.Run("register elon musk", func() {
		resp, err := s.cli.Register(ctx, &proto.UserProfile{
			FullName:    "Elon Musk",
			Email:       "elonID@tesla.us",
			DateOfBirth: date(1971, 06, 28),
		})
		s.NoError(err)
		elonID = resp.UserId
	})

	elonTweet := "hey guys, should i buy twitter?"
	s.Run("first elonID tweet", func() {
		_, err := s.cli.AddTweet(ctx, &proto.AddTweetRequest{
			Text:   elonTweet,
			UserId: elonID,
		})
		s.NoError(err)
	})

	s.Run("unregistered user can't send tweet", func() {
		_, err := s.cli.AddTweet(ctx, &proto.AddTweetRequest{
			Text:   "scam message",
			UserId: -1, // who the f*ck are you??
		})
		s.Error(err)
	})

	var amberID int64
	s.Run("elonID have new follower", func() {
		resp, err := s.cli.Register(ctx, &proto.UserProfile{
			FullName:    "Amber Heard",
			Email:       "beach@club.com",
			DateOfBirth: date(1986, 04, 22),
		})
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
			UserId: amberID, Limit:  10,
		})
		s.NoError(err)

		s.Len(resp.Tweets, 1)
		s.Equal(elonTweet, resp.Tweets[0])
	})

    s.Run("amber wants to unfollow elon", func() {
        _, err := s.cli.Unfollow(ctx, &proto.UnfollowRequest{
            UserId: elonID,
            OldFollowerId: amberID,
       })
        s.Require().NoError(err)

        resp, err := s.cli.GetNewsFeed(ctx, &proto.GetNewsFeedRequest{
            UserId: amberID, Limit:  10,
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
