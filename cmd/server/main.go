package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gettmure/go-ntn-bot/internal/config"
	"github.com/gettmure/go-ntn-bot/internal/domain/repo"
	"github.com/gettmure/go-ntn-bot/internal/http-server/handlers"
	"github.com/gettmure/go-ntn-bot/internal/slogger"
	"github.com/gettmure/go-ntn-bot/internal/storage/mongodb"
	"github.com/gettmure/go-ntn-bot/internal/webhook"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/exp/slog"
)

type Flags struct {
	WebhookURL string
	BotToken   string
}

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting application...")
	log.Debug("debug messages are enabled")

	log.Info("establishing mongodb connection...")
	db, err := mongodb.NewClient(cfg.MongoDB.Path)
	if err != nil {
		log.Error("failed to create mongodb client", slogger.FormatError(err))
		os.Exit(1)
	}

	if err := db.Ping(context.Background(), nil); err != nil {
		log.Error("failed to ping mongodb", slogger.FormatError(err))
		os.Exit(1)
	}

	URL := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := setupServer(cfg, log, db)

	log.Info("creating tcp connection...")
	conn, err := net.Listen("tcp", URL)
	if err != nil {
		log.Error("failed to create tcp connection", slogger.FormatError(err))
		os.Exit(1)
	}

	log.Info("registering telegram webhook telegram...")
	body, err := webhook.RegisterTelegramWebhook(cfg.Telegram.WebhookURL, cfg.Telegram.BotToken)
	if err != nil {
		log.Error("failed to register telegram webhook",
			slogger.FormatError(err),
			slogger.FormatBody(body),
		)
		os.Exit(1)
	}

	log.Info(fmt.Sprintf("listening to incoming requests at %s", URL))
	if err := srv.Serve(conn); err != nil {
		log.Error("failed to start server", slogger.FormatError(err))
	}

	log.Error("server shutdown")
}

func setupLogger(env string) slogger.Logger {
	var logger slogger.Logger

	switch env {
	case "local":
		logger = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		logger = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	default:
		log.Fatalf("this env is not supported: %s", env)
	}

	return logger
}

func setupServer(cfg *config.Config, log slogger.Logger, db *mongodb.Client) *http.Server {
	router := chi.NewRouter()

	notionUserRepo := repo.NewNotionUserRepository(db)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/test", handlers.Test)
	router.Get("/notion/oauth/callback", handlers.OAuthRedirectURI(log, cfg.Notion, notionUserRepo))
	router.Post("/telegram/webhook/callback", handlers.TelegramWebhook(log, cfg.Telegram.BotToken))

	srv := &http.Server{
		Handler:      router,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 5,
	}

	return srv
}
