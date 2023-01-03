package main

import (
	"github.com/demimurg/twitter/internal/adapter/inmem"
	"github.com/demimurg/twitter/internal/adapter/scamdetector"
	"github.com/demimurg/twitter/internal/driver/grpcsrv"
	"github.com/demimurg/twitter/internal/usecase"
	"github.com/demimurg/twitter/pkg/grace"
)

func main() {
	scamClient := scamdetector.NewDummyClient()
	userRepo := inmem.NewUserRepository()
	tweetRepo := inmem.NewTweetRepository()
	followerRepo := inmem.NewFollowerRepository()

	feedManager := usecase.NewFeedManager(userRepo, followerRepo, tweetRepo)
	userRegistrator := usecase.NewUserRegistrator(userRepo, scamClient)

	srv := grpcsrv.NewTwitter(feedManager, userRegistrator)
	grace.Run(
		grace.GRPC(srv, ":80"),
        grace.GRPCUI(":81", "localhost:80"),
		grace.PromHTTP(":82"),
	)
}
