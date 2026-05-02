package domain

import "context"

type CameraType int
const (
	Analog CameraType = iota
	Digital
	Cellphone
)

type CameraBrand struct {
	ID      int
	Name    string
	Country string
}

type CameraBrandRepository interface {
	Create(ctx context.Context, b CameraBrand) (CameraBrand, error)
	Update(ctx context.Context, b CameraBrand) (CameraBrand, error)
	FindByID(ctx context.Context, id int) (CameraBrand, error)
	FindAll(ctx context.Context) ([]CameraBrand, error)
	Delete(ctx context.Context, id int) error
}

type CameraModel struct {
	ID      int
	Name    string
	BrandID int
	Type    CameraType
}

type CameraModelRepository interface {
	Create(ctx context.Context, m CameraModel) (CameraModel, error)
	Update(ctx context.Context, m CameraModel) (CameraModel, error)
	FindByID(ctx context.Context, id int) (CameraModel, error)
	FindByBrand(ctx context.Context, brandID int) ([]CameraModel, error)
	Delete(ctx context.Context, id int) error
}

type Camera struct {
	ID       int
	Nickname string
	ModelID  int
	OwnerID  int
}

type CameraRepository interface {
	Create(ctx context.Context, c Camera) (Camera, error)
	Update(ctx context.Context, c Camera) (Camera, error)
	FindByID(ctx context.Context, id int) (Camera, error)
	FindByOwner(ctx context.Context, ownerID int) ([]Camera, error)
	Delete(ctx context.Context, id int) error
}
