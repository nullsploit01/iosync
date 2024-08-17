package repositories

import (
	"context"
	"iosync/ent"
	"iosync/ent/device"
)

type DeviceRepository struct {
	dbClient *ent.Client
}

type AddDevicePayload struct {
	Username string
	Name     string
}

func NewDeviceRepository(dbClient *ent.Client) *DeviceRepository {
	return &DeviceRepository{dbClient: dbClient}
}

func (d *DeviceRepository) AddDevice(ctx context.Context, payload *AddDevicePayload) (*ent.Device, error) {
	return d.dbClient.Device.Create().
		SetName(payload.Name).
		SetUsername(payload.Username).
		Save(ctx)
}

func (d *DeviceRepository) GetDevices(ctx context.Context, username string) ([]*ent.Device, error) {
	return d.dbClient.Device.Query().Where(device.Username(username)).All(ctx)
}

func (d *DeviceRepository) GetDevice(ctx context.Context, deviceId int) (*ent.Device, error) {
	return d.dbClient.Device.Query().Where(device.ID(deviceId)).First(ctx)
}

func (d *DeviceRepository) DeleteDevice(ctx context.Context, deviceId int) error {
	_, err := d.dbClient.Device.Delete().Where(device.ID(deviceId)).Exec(ctx)
	return err
}
