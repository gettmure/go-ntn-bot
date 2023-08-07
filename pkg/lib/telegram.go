package lib

import (
	"fmt"
	"net/url"
)

// https://core.telegram.org/bots/api#available-methods
const (
	BaseURL string = `https://api.telegram.org/bot`

	TelegramSetWebhook  string = "setWebhook"
	TelegramSendMessage string = "sendMessage"
)

func BuildTelegramURL(token, method string, params *map[string]string) (*url.URL, error) {
	endpoint := fmt.Sprintf("%s%s/%s", BaseURL, token, method)

	url, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	if params == nil {
		return url, nil
	}

	query := url.Query()
	for key, value := range *params {
		query.Set(key, value)
	}

	url.RawQuery = query.Encode()

	return url, nil
}
