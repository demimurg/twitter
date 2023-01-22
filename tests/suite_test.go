//go:build e2e

package tests

import (
	"context"
	"database/sql"
	"encoding/json"
	"testing"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/demimurg/twitter/internal/adapter/postgres"
	"github.com/pressly/goose/v3"
	"github.com/stretchr/testify/require"

	"github.com/demimurg/twitter/internal/adapter/scamdetector"
	"github.com/demimurg/twitter/internal/driver/grpcsrv"
	"github.com/demimurg/twitter/internal/usecase"
	"github.com/demimurg/twitter/pkg/grace"
	"github.com/demimurg/twitter/pkg/proto"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestE2E(t *testing.T) {
	db, err := sql.Open("pgx", "host=localhost user=postgres")
	require.NoError(t, err, "connect to postgres")
	defer func() {
		require.NoError(t, db.Close(), "close db connection")
	}()

	err = retry.Do(func() error {
		return goose.Up(db, "../migrations")
	})
	require.NoError(t, err, "migrate to last schema")

	scamClient := scamdetector.NewDummyClient()
	userRepo := postgres.NewUserRepository(db)
	tweetRepo := postgres.NewTweetRepository(db)
	followerRepo := postgres.NewFollowerRepository(db)

	feedManager := usecase.NewFeedManager(userRepo, followerRepo, tweetRepo)
	userProfiler := usecase.NewUserProfiler(userRepo, scamClient)

	srv := grace.GRPC(grpcsrv.NewTwitter(feedManager, userProfiler), ":8080")
	go func() {
		require.NoError(t, srv.Run(), "shutdown server")
	}()
	defer func() {
		require.NoError(t, srv.Shutdown(), "send shutdown signal to server")
	}()

	<-time.After(50 * time.Millisecond)

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err, "dial server")
	cli := proto.NewTwitterClient(conn)

	suite.Run(t, &endToEndTestSuite{cli: cli, db: db})
}

type endToEndTestSuite struct {
	suite.Suite
	// twitter client connected to grpc server
	cli proto.TwitterClient
	// connection to database, to clean up state after suite
	db *sql.DB
}

func (s *endToEndTestSuite) TearDownSuite() {
	_, err := s.db.Exec("TRUNCATE TABLE comment, tweet, follower, users")
	s.Require().NoError(err, "truncate db tables")

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
