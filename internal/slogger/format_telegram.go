package slogger

import (
	"github.com/gettmure/go-ntn-bot/pkg/telegram"
	"golang.org/x/exp/slog"
)

func FormatTelegramUpdate(update *telegram.Update) slog.Attr {
	var updateID int64

	if update != nil {
		updateID = update.UpdateID
	}

	return slog.Group("update",
		"id", updateID,
	)
}

func FormatTelegramMessage(message *telegram.Message) slog.Attr {
	var msgID int64
	var msgText string

	if message != nil {
		msgID = message.MessageID
		msgText = message.Text
	}

	return slog.Group("message",
		"id", msgID,
		"text", msgText,
	)
}
