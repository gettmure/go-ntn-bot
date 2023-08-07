package repo

import (
	"context"

	"github.com/gettmure/go-ntn-bot/internal/domain/model"
	"github.com/gettmure/go-ntn-bot/internal/storage/mongodb"
	"github.com/gettmure/go-ntn-bot/pkg/notion"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const tableName = "notionUsers"

type NotionUserRepository interface {
	Save(ctx context.Context, user notion.UserAuth) (*model.NotionUser, error)
	FindOne(ctx context.Context, filter bson.M) (*model.NotionUser, error)
	FindOneByToken(ctx context.Context, token string) (*model.NotionUser, error)
}

type notionUserRepository struct {
	db *mongodb.Client
}

func NewNotionUserRepository(db *mongodb.Client) NotionUserRepository {
	return &notionUserRepository{db: db}
}

func (r *notionUserRepository) Save(ctx context.Context, user notion.UserAuth) (*model.NotionUser, error) {
	if record, _ := r.FindOneByToken(ctx, user.AccessToken); record != nil {
		return record, nil
	}

	res, err := r.db.Database("ntn").Collection(tableName).InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	oid := res.InsertedID.(primitive.ObjectID)
	notionUser := model.NewNotionUser(oid, user)

	return notionUser, nil
}

func (r *notionUserRepository) FindOne(ctx context.Context, filter bson.M) (*model.NotionUser, error) {
	user := &model.NotionUser{}

	res := r.db.Database("ntn").Collection(tableName).FindOne(ctx, filter)
	if err := res.Err(); err != nil {
		return nil, err
	}

	if err := res.Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *notionUserRepository) FindOneByToken(ctx context.Context, token string) (*model.NotionUser, error) {
	filter := bson.M{"accesstoken": token}

	user, err := r.FindOne(ctx, filter)
	if err != nil {
		return nil, err
	}

	return user, nil
}
