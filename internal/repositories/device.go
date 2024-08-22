package repositories

import (
	"context"
	"iosync/ent"
	"iosync/ent/device"
	"iosync/ent/user"
)

type DeviceRepository struct {
	dbClient *ent.Client
}

func NewDeviceRepository(dbClient *ent.Client) *DeviceRepository {
	return &DeviceRepository{dbClient: dbClient}
}

func (d *DeviceRepository) Create(ctx context.Context, deviceName string, user *ent.User) (*ent.Device, error) {
	return d.dbClient.Device.Create().
		SetName(deviceName).
		SetUser(user).
		Save(ctx)
}

func (d *DeviceRepository) GetAll(ctx context.Context, username string) ([]*ent.Device, error) {
	user, err := d.dbClient.User.
		Query().
		Where(user.Username(username)).
		WithDevices().
		Only(ctx)

	return user.Edges.Devices, err
}

func (d *DeviceRepository) Get(ctx context.Context, deviceId int) (*ent.Device, error) {
	return d.dbClient.Device.Query().Where(device.ID(deviceId)).WithAPIKey().First(ctx)
}

func (d *DeviceRepository) Delete(ctx context.Context, deviceId int) error {
	_, err := d.dbClient.Device.Delete().Where(device.ID(deviceId)).Exec(ctx)
	return err
}
