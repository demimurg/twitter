package main

import (
	"context"
	"database/sql"

	"github.com/demimurg/twitter/internal/adapter/postgres"
	"github.com/demimurg/twitter/internal/adapter/scamdetector"
	"github.com/demimurg/twitter/internal/driver/grpcsrv"
	"github.com/demimurg/twitter/internal/usecase"
	"github.com/demimurg/twitter/pkg/grace"
	"github.com/demimurg/twitter/pkg/log"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	db, err := sql.Open("pgx", "host=localhost user=postgres")
	handle(err, "connect to postgres")
	defer db.Close()

	err = goose.Up(db, "/migrations", goose.WithAllowMissing())
	handle(err, "can't migrate to last schema")

	scamClient := scamdetector.NewDummyClient()
	userRepo := postgres.NewUserRepository(db)
	tweetRepo := postgres.NewTweetRepository(db)
	followerRepo := postgres.NewFollowerRepository(db)

	feedManager := usecase.NewFeedManager(userRepo, followerRepo, tweetRepo)
	userRegistrator := usecase.NewUserRegistrator(userRepo, scamClient)

	srv := grpcsrv.NewTwitter(feedManager, userRegistrator)
	grace.Run(
		grace.GRPC(srv, ":80"),
		grace.GRPCUI(":81", "localhost:80"),
		grace.PromHTTP(":82"),
	)
}

func handle(err error, msg string) {
	if err != nil {
		log.Panic(context.Background(), msg,
			"error", err)
	}
}
