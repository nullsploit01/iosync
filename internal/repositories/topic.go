package repositories

import (
	"context"
	"iosync/ent"
)

type TopicRepository struct {
	dbClient *ent.Client
}

type CreateTopicPayload struct {
	Name string
}

func NewTopicRepository(dbClient *ent.Client) *TopicRepository {
	return &TopicRepository{
		dbClient: dbClient,
	}
}

func (t *TopicRepository) Create(ctx context.Context, payload *CreateTopicPayload, device *ent.Device) (*ent.Topic, error) {
	return t.dbClient.Topic.Create().
		SetName(payload.Name).
		SetDevice(device).
		Save(ctx)
}
