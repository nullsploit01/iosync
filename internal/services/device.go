package services

import (
	"context"
	"iosync/ent"
	"iosync/internal/repositories"
)

type DeviceService struct {
	deviceRepository *repositories.DeviceRepository
}

type AddDeviceRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string
}

func NewDeviceService(dbClient *ent.Client) *DeviceService {
	deviceRepository := repositories.NewDeviceRepository(dbClient)

	return &DeviceService{
		deviceRepository: deviceRepository,
	}
}

func (d *DeviceService) AddDevice(ctx context.Context, request AddDeviceRequest) (*ent.Device, error) {
	addDevicePaylaod := repositories.AddDevicePayload{
		Name:     request.Name,
		Username: request.Username,
	}

	return d.deviceRepository.AddDevice(ctx, &addDevicePaylaod)
}
