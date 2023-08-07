package webhook

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gettmure/go-ntn-bot/pkg/lib"
)

func RegisterTelegramWebhook(webhookURL, botToken string) ([]byte, error) {
	params := map[string]string{
		"url":                  webhookURL,
		"drop_pending_updates": "true",
	}

	url, err := lib.BuildTelegramURL(botToken, "setWebhook", &params)
	if err != nil {
		return []byte{}, err
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
