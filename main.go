package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/rwidjojo/HanchanManager/internal/api"
	"github.com/rwidjojo/HanchanManager/internal/repository"
)

func main() {
	ctx := context.Background()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://mahjong:mahjong@host.containers.internal:5432/mahjong?sslmode=disable"
	}

	pool, err := repository.Connect(ctx, dsn)
	if err != nil {
		slog.Error("failed to connect to database", "err", err)
		os.Exit(1)
	}
	defer pool.Close()
	slog.Info("connected to database")

	// ToDo: if we want to use a repository/db.go then we need to change the code like this
	// db := repository.NewDB(pool)
	// router := api.NewRouter(db)
	router := api.NewRouter(pool)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	slog.Info("server starting", "addr", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("server error", "err", err)
		os.Exit(1)
	}
}
