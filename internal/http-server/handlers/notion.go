package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gettmure/go-ntn-bot/internal/config"
	"github.com/gettmure/go-ntn-bot/internal/domain/repo"
	"github.com/gettmure/go-ntn-bot/internal/slogger"
	"github.com/gettmure/go-ntn-bot/pkg/notion"
)

func OAuthRedirectURI(
	log slogger.Logger,
	cfg config.Notion,
	repo repo.NotionUserRepository,
) func(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.notion.OAuthRedirectURI"
	log.With("op", op)

	client := notion.InitClient(cfg.OAuth.ClientID, cfg.OAuth.ClientSecret, &cfg.OAuth.RedirectURI)

	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*20)

		defer r.Body.Close()
		defer cancel()

		authCode := r.URL.Query().Get("code")

		user, err := client.Auth(authCode)
		if err != nil {
			panic(err)
		}

		a, err := repo.Save(ctx, *user)
		if err != nil {
			panic(err)
		}

		log.Debug("id", a.IDHex())
	}
}
