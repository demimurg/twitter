package main

import (
    "context"
    "github.com/demimurg/twitter/internal/adapter/inmem"
    "github.com/demimurg/twitter/internal/driver/grpcsrv"
    "github.com/demimurg/twitter/internal/usecase"
    "github.com/demimurg/twitter/pkg/log"
    "net"
)

func main() {
    userRepo := inmem.NewUserRepository()
    tweetRepo := inmem.NewTweetRepository()
    followerRepo := inmem.NewFollowerRepository()

    feedManager := usecase.NewFeedManager(userRepo, followerRepo, tweetRepo, nil)
    userRegistrator := usecase.NewUserRegistrator(userRepo, nil)

    srv := grpcsrv.NewTwitter(feedManager, userRegistrator)
    l, err := net.Listen("tcp", ":80")
    handle(err, "can't use port :80")
    err = srv.Serve(l)
    handle(err, "can't start server")
}

func handle(err error, msg string) {
    if err != nil {
        log.Panic(context.Background(), msg,
            "error", err)
    }
}
