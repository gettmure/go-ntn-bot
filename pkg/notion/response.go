package notion

// https://developers.notion.com/docs/authorization#step-4-notion-responds-with-an-access_token-and-additional-information
type UserAuth struct {
	AccessToken          string `json:"access_token" bson:"accesstoken"`
	BotID                string `json:"bot_id" bson:"botid"`
	DuplicatedTemplateID string `json:"duplicated_template_id" bson:"duplicatedtemplateid"`
	Owner                Owner  `json:"owner" bson:"owner"`

	WorkspaceID   string `json:"workspace_id" bson:"workspaceid"`
	WorkspaceIcon string `json:"workspace_icon" bson:"workspaceicon"`
	WorkspaceName string `json:"workspace_name" bson:"workspacename"`

	Error *string `json:"error"`
}

type Owner struct {
	Type string `json:"type" bson:"type"`
	User struct {
		ID        string `json:"id" bson:"id"`
		Object    string `json:"object" bson:"object"`
		Name      string `json:"name" bson:"name"`
		AvatarURL string `json:"avatar_url" bson:"avatarurl"`
		Type      string `json:"type" bson:"type"`
		Person    struct {
			Email string `json:"email" bson:"email"`
		} `json:"person" bson:"person"`
	} `json:"user" bson:"user"`
}
