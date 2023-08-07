package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gettmure/go-ntn-bot/pkg/lib"
)

type Client interface {
	Auth(authCode string) (*UserAuth, error)
}

type client struct {
	http         http.Client
	clientID     string
	clientSecret string
	redirectURI  *string
}

func InitClient(clientID string, clientSecret string, redirectURI *string) Client {
	return &client{
		http:         http.Client{},
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURI:  redirectURI,
	}
}

func (c *client) Auth(authCode string) (*UserAuth, error) {
	url := BuildAuthURL()
	credentials := fmt.Sprintf("%s:%s", c.clientID, c.clientSecret)

	data := CreateAuthBody(authCode, c.redirectURI)
	bts, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bts))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf(`Basic "%s"`, lib.StringToBase64(credentials)))

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	userAuth := &UserAuth{}
	if err = json.Unmarshal(body, userAuth); err != nil {
		return nil, err
	}

	if userAuth.Error != nil {
		return nil, fmt.Errorf("failed to auth notion: %s", *userAuth.Error)
	}

	return userAuth, nil
}
