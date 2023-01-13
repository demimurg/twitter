package main

import (
	"context"
	"database/sql"
	"github.com/caarlos0/env/v6"

	"github.com/demimurg/twitter/internal/adapter/postgres"
	"github.com/demimurg/twitter/internal/adapter/scamdetector"
	"github.com/demimurg/twitter/internal/driver/grpcsrv"
	"github.com/demimurg/twitter/internal/usecase"
	"github.com/demimurg/twitter/pkg/grace"
	"github.com/demimurg/twitter/pkg/log"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

var cfg struct {
	LogLevel      string `env:"LOG_LEVEL" envDefault:"debug"`
	PostgresqlDSN string `env:"POSTGRESQL_DSN" envDefault:"host=localhost user=postgres"`
	MigrationsDir string `env:"MIGRATIONS_DIR" envDefault:"./migrations"`
}

func main() {
	db, err := sql.Open("pgx", cfg.PostgresqlDSN)
	handle(err, "connect to postgres")
	defer db.Close()

	err = goose.Up(db, cfg.MigrationsDir, goose.WithAllowMissing())
	handle(err, "can't migrate to last schema")

	scamClient := scamdetector.NewDummyClient()
	userRepo := postgres.NewUserRepository(db)
	tweetRepo := postgres.NewTweetRepository(db)
	followerRepo := postgres.NewFollowerRepository(db)

	feedManager := usecase.NewFeedManager(userRepo, followerRepo, tweetRepo)
	userProfiler := usecase.NewUserProfiler(userRepo, scamClient)

	srv := grpcsrv.NewTwitter(feedManager, userProfiler)
	grace.Run(
		grace.GRPC(srv, ":80"),
		grace.GRPCUI(":81", "localhost:80"),
		grace.Prometheus(":82"),
	)
}

func init() {
    err := env.Parse(&cfg)
    handle(err, "parse config")
    log.SetLevel(cfg.LogLevel)
}

func handle(err error, msg string) {
	if err != nil {
		// we use panic, to be able to stop execution
		// and not break defer calls, as Fatal will did
		log.Panic(context.Background(), msg,
			"error", err,
			"config", cfg)
	}
}
