package services

import (
	"context"
	"iosync/ent"
	"iosync/internal/repositories"
)

type TopicService struct {
	topicRepository  *repositories.TopicRepository
	deviceRepository *repositories.DeviceRepository
}

type CreateTopicPayload struct {
	Name string `json:"name" validate:"required"`
}

func NewTopicService(dbClient *ent.Client) *TopicService {
	topicRepository := repositories.NewTopicRepository(dbClient)
	deviceRepository := repositories.NewDeviceRepository(dbClient)

	return &TopicService{
		topicRepository:  topicRepository,
		deviceRepository: deviceRepository,
	}
}

func (t *TopicService) CreateTopic(ctx context.Context, deviceId int, requestPayload *CreateTopicPayload) (*ent.Topic, error) {
	device, err := t.deviceRepository.Get(ctx, deviceId)
	if err != nil {
		return nil, err
	}

	createTopicPayload := &repositories.CreateTopicPayload{
		Name: requestPayload.Name,
	}

	return t.topicRepository.Create(ctx, createTopicPayload, device)
}
