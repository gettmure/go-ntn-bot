package notion

import "fmt"

// TODO: move to config
const BaseURL = "https://api.notion.com/v1"

func BuildAuthURL() string {
	return fmt.Sprintf("%s%s", BaseURL, "/oauth/token")
}
