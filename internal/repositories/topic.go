package repositories

import "iosync/ent"

type TopicRepository struct {
	dbClient *ent.Client
}

func NewTopicRepository(dbClient *ent.Client) *TopicRepository {
	return &TopicRepository{
		dbClient: dbClient,
	}
}
