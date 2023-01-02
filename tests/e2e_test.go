package tests

import (
	"context"
	"testing"
	"time"

	"github.com/demimurg/twitter/internal/adapter/inmem"
	"github.com/demimurg/twitter/internal/adapter/scamdetector"
	"github.com/demimurg/twitter/internal/driver/grpcsrv"
	"github.com/demimurg/twitter/internal/usecase"
	"github.com/demimurg/twitter/pkg/grace"
	"github.com/demimurg/twitter/pkg/proto"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var ctx = context.Background()

func TestE2E(t *testing.T) {
	scamClient := scamdetector.NewDummyClient()
	userRepo := inmem.NewUserRepository()
	tweetRepo := inmem.NewTweetRepository()
	followerRepo := inmem.NewFollowerRepository()

	feedManager := usecase.NewFeedManager(userRepo, followerRepo, tweetRepo)
	userRegistrator := usecase.NewUserRegistrator(userRepo, scamClient)

	srv := grpcsrv.NewTwitter(feedManager, userRegistrator)
	go func() {
		err := grace.GRPC(srv, ":8080").Run()
		require.NoError(t, err)
	}()
	<-time.After(100 * time.Millisecond)

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	suite.Run(t, &endToEndTestSuite{cli: proto.NewTwitterClient(conn)})
}

type endToEndTestSuite struct {
	suite.Suite
	cli proto.TwitterClient
}

func (s *endToEndTestSuite) TestBasic() {
	var (
		elon, amber int64 = 0, 1
		elonTweet         = "hey guys, should i buy twitter?"
	)

	s.Run("register elon musk", func() {
		_, err := s.cli.Register(ctx, &proto.RegisterRequest{
			FullName:    "Elon Musk",
			Email:       "elon@tesla.us",
			DateOfBirth: "1971-06-28",
		})
		s.NoError(err)
	})

	s.Run("first elon tweet", func() {
		_, err := s.cli.AddTweet(ctx, &proto.AddTweetRequest{
			Text:   elonTweet,
			UserId: elon,
		})
		s.NoError(err)
	})

	s.Run("elon have new follower", func() {
		_, err := s.cli.Register(ctx, &proto.RegisterRequest{
			FullName:    "Amber Heard",
			Email:       "beach@club.com",
			DateOfBirth: "1986-04-22",
		})
		s.NoError(err)

		_, err = s.cli.Follow(ctx, &proto.FollowRequest{
			UserId:        elon,
			NewFollowerId: amber,
		})
		s.NoError(err)
	})

	s.Run("amber checks the news feed", func() {
		resp, err := s.cli.GetNewsFeed(ctx, &proto.GetNewsFeedRequest{
			Limit:  10,
			Offset: 0,
			UserId: amber,
		})
		s.NoError(err)

		s.Len(resp.Tweets, 1)
		s.Equal(elonTweet, resp.Tweets[0])
	})
}

func (s *endToEndTestSuite) TestNegative() {
	s.Run("unregistered user can't send tweet", func() {
		_, err := s.cli.AddTweet(ctx, &proto.AddTweetRequest{
			Text:   "scam message",
			UserId: -1, // who the f*ck are you??
		})
		s.Error(err)
	})

}
