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
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestE2E(t *testing.T) {
	scamClient := scamdetector.NewDummyClient()
	userRepo := inmem.NewUserRepository()
	tweetRepo := inmem.NewTweetRepository()
	followerRepo := inmem.NewFollowerRepository()

	feedManager := usecase.NewFeedManager(userRepo, followerRepo, tweetRepo)
	userProfiler := usecase.NewUserProfiler(userRepo, scamClient)

	srv := grpcsrv.NewTwitter(feedManager, userProfiler)
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

var ctx = context.Background()

// date is a helper for easy creation of the protobuf timestamps
func date(year, month, day int) *timestamppb.Timestamp {
	return timestamppb.New(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
}
