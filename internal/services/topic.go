package services

import (
	"iosync/ent"
	"iosync/internal/repositories"
)

type TopicService struct {
	topicRepository *repositories.TopicRepository
}

func NewTopicService(dbClient *ent.Client) *TopicService {
	topicRepository := repositories.NewTopicRepository(dbClient)

	return &TopicService{
		topicRepository: topicRepository,
	}
}
