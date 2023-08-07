package notion

type AuthBody struct {
	GrantType string `json:"grant_type"`
	Code      string `json:"code"`

	RedirectURI *string `json:"redirect_uri,omitempty"`
}

func CreateAuthBody(code string, redirectURI *string) *AuthBody {
	return &AuthBody{
		GrantType:   "authorization_code",
		Code:        code,
		RedirectURI: redirectURI,
	}
}
