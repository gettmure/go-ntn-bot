package model

import (
	"github.com/gettmure/go-ntn-bot/pkg/notion"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NotionUser struct {
	ID              primitive.ObjectID `bson:"_id"`
	notion.UserAuth `bson:"inline"`
}

func NewNotionUser(oid primitive.ObjectID, user notion.UserAuth) *NotionUser {
	return &NotionUser{ID: oid, UserAuth: user}
}

func (nu *NotionUser) IDHex() string {
	return nu.ID.Hex()
}
