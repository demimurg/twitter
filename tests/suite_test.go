package tests

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/demimurg/twitter/internal/adapter/inmem"
	"github.com/demimurg/twitter/internal/adapter/scamdetector"
	"github.com/demimurg/twitter/internal/driver/grpcsrv"
	"github.com/demimurg/twitter/internal/usecase"
	"github.com/demimurg/twitter/pkg/grace"
	"github.com/demimurg/twitter/pkg/proto"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestE2E(t *testing.T) {
	suite.Run(t, &endToEndTestSuite{})
}

type endToEndTestSuite struct {
	suite.Suite
	// twitter client connected to grpc server
	cli proto.TwitterClient
	// this is the running process of twitter grpc server
	srv grace.Process
}

func (s *endToEndTestSuite) SetupSuite() {
	if s.srv != nil {
		err := s.srv.Shutdown()
		s.Require().NoError(err, "shutdown server")
	}

	scamClient := scamdetector.NewDummyClient()
	userRepo := inmem.NewUserRepository()
	tweetRepo := inmem.NewTweetRepository()
	followerRepo := inmem.NewFollowerRepository()

	feedManager := usecase.NewFeedManager(userRepo, followerRepo, tweetRepo)
	userProfiler := usecase.NewUserProfiler(userRepo, scamClient)

	s.srv = grace.GRPC(grpcsrv.NewTwitter(feedManager, userProfiler), ":8080")
	go func() {
		err := s.srv.Run()
		s.Require().NoError(err, "shutdown server")
	}()

	<-time.After(50 * time.Millisecond)

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.Require().NoError(err, "dial server")
	s.cli = proto.NewTwitterClient(conn)
}

// EqualProto makes a correct comparisons for protobuf structs
func (s *endToEndTestSuite) EqualProto(expected, actual any, msg string) {
	// stretchr equal func additionally compare private fields
	// for protobuf it means that encoder temprorary data can break diff
	expectJSON, _ := json.Marshal(expected)
	actualJSON, _ := json.Marshal(actual)
	s.JSONEq(string(expectJSON), string(actualJSON), msg)
}

var ctx = context.Background()

// date is a helper for easy creation of the protobuf timestamps
func date(year, month, day int) *timestamppb.Timestamp {
	return timestamppb.New(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
}
