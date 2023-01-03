package tests

import (
	"context"
    "google.golang.org/protobuf/types/known/timestamppb"
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

func (s *endToEndTestSuite) TestAuth() {
    s.Run("register greta thunberg", func() {
		_, err := s.cli.Register(ctx, &proto.UserProfile{
            FullName:    "Greta Thunberg",
            Email:       "smalldickenergy@getalife.com",
            DateOfBirth: date(2003, 01, 03),
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
}

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
            Limit:  10,
            Offset: 0,
            UserId: amberID,
            })
        s.NoError(err)

        s.Len(resp.Tweets, 1)
        s.Equal(elonTweet, resp.Tweets[0])
    })
}

func date(year, month, day int) *timestamppb.Timestamp {
    return timestamppb.New(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
}
