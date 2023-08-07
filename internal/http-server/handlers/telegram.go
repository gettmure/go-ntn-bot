package handlers

import (
	"io"
	"net/http"

	"github.com/gettmure/go-ntn-bot/internal/slogger"
	"github.com/gettmure/go-ntn-bot/pkg/telegram"
)

func TelegramWebhook(log slogger.Logger, botToken string) func(w http.ResponseWriter, r *http.Request) {
	const op = "internal.http-server.handlers.telegram.TelegramWebhook"
	log.With("op", op)

	client := telegram.InitClient()
	bot := telegram.InitBot(botToken, client)

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error("failed to parse telegram request", slogger.FormatError(err))
		}

		update, err := bot.ParseUpdate(body)
		if err != nil {
			panic(err)
		}

		bot.SendMessage(update.Message.Chat.ID, "I AM ALIVE")

		w.WriteHeader(http.StatusOK)
	}
}
