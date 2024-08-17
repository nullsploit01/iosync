package services

import (
	"context"
	"errors"
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

	return d.deviceRepository.Create(ctx, &addDevicePaylaod)
}

func (d *DeviceService) GetDevices(ctx context.Context, username string) ([]*ent.Device, error) {
	devices, err := d.deviceRepository.GetAll(ctx, username)

	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			return nil, errors.New("invalid username")
		}
		return nil, err
	}

	return devices, nil
}

func (d *DeviceService) GetDevice(ctx context.Context, deviceId int) (*ent.Device, error) {
	device, err := d.deviceRepository.Get(ctx, deviceId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			return nil, errors.New("device not found")
		}
		return nil, err
	}

	return device, nil
}

func (d *DeviceService) DeleteDevice(ctx context.Context, deviceId int) error {
	err := d.deviceRepository.Delete(ctx, deviceId)
	if err != nil {
		var notFoundError *ent.NotFoundError
		if errors.As(err, &notFoundError) {
			return errors.New("device not found")
		}
		return err
	}

	return nil
}
