package telegram

import (
	"encoding/json"
)

type Bot interface {
	SendMessage(chatId int64, msg string) (*Message, error)
	ParseUpdate(body []byte) (*Update, error)
}

// https://core.telegram.org/bots/api#authorizing-your-bot
type bot struct {
	token  string
	client Client
}

func InitBot(token string, c Client) Bot {
	bot := &bot{token: token, client: c}

	return bot
}

func (b *bot) SendMessage(chatId int64, msg string) (*Message, error) {
	return b.client.SendMessage(b.token, chatId, msg)
}

func (b *bot) ParseUpdate(body []byte) (*Update, error) {
	var update *Update

	err := json.Unmarshal(body, &update)
	if err != nil {
		return nil, err
	}

	return update, nil
}
