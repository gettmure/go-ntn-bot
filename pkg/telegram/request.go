package telegram

// https://core.telegram.org/bots/api#sendmessage
type SendMessage struct {
	ChatId int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func CreateSendMessage(chatId int64, text string) *SendMessage {
	return &SendMessage{
		ChatId: chatId,
		Text:   text,
	}
}
