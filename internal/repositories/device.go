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

func (d *DeviceRepository) Create(ctx context.Context, payload *AddDevicePayload) (*ent.Device, error) {
	return d.dbClient.Device.Create().
		SetName(payload.Name).
		SetUsername(payload.Username).
		Save(ctx)
}

func (d *DeviceRepository) GetAll(ctx context.Context, username string) ([]*ent.Device, error) {
	return d.dbClient.Device.Query().Where(device.Username(username)).All(ctx)
}

func (d *DeviceRepository) Get(ctx context.Context, deviceId int) (*ent.Device, error) {
	return d.dbClient.Device.Query().Where(device.ID(deviceId)).First(ctx)
}

func (d *DeviceRepository) Delete(ctx context.Context, deviceId int) error {
	_, err := d.dbClient.Device.Delete().Where(device.ID(deviceId)).Exec(ctx)
	return err
}
