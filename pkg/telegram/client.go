package telegram

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gettmure/go-ntn-bot/pkg/lib"
)

type Client interface {
	// https://core.telegram.org/bots/api#sendmessage
	SendMessage(token string, chatId int64, msg string) (*Message, error)
}

type client struct {
	http http.Client
}

func InitClient() Client {
	return &client{http: http.Client{}}
}

func (c *client) SendMessage(token string, chatId int64, msg string) (*Message, error) {
	url, err := lib.BuildTelegramURL(token, lib.TelegramSendMessage, nil)
	if err != nil {
		return nil, err
	}

	sendMessage := CreateSendMessage(chatId, msg)
	bts, err := json.Marshal(sendMessage)
	if err != nil {
		return nil, err
	}

	resp, err := c.http.Post(url.String(), "application/json", bytes.NewBuffer(bts))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var message *Message
	if err = json.Unmarshal(body, message); err != nil {
		return nil, err
	}

	return message, nil
}
