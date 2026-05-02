package domain

import (
	"context"
	"time"
)

type PhotoTag struct {
	ID    int
	Label string
}

type PhotoTagRepository interface {
	Create(ctx context.Context, t PhotoTag) (PhotoTag, error)
	Update(ctx context.Context, t PhotoTag) (PhotoTag, error)
	FindByID(ctx context.Context, id int) (PhotoTag, error)
	FindAll(ctx context.Context) ([]PhotoTag, error)
	Delete(ctx context.Context, id int) error
}

type Photo struct {
	ID          int
	CameraID    int
	FilmID      *int
	TakenAt     time.Time
	UserID      *int
	Description string
	TagIDs      []int
	StorageKey  string
}

type PhotoRepository interface {
	Create(ctx context.Context, p Photo) (Photo, error)
	Update(ctx context.Context, p Photo) (Photo, error)
	FindByID(ctx context.Context, id int) (Photo, error)
	FindByUser(ctx context.Context, userID int) ([]Photo, error)
	FindByCamera(ctx context.Context, cameraID int) ([]Photo, error)
	FindByFilm(ctx context.Context, filmID int) ([]Photo, error)
	FindByTags(ctx context.Context, tagIDs []int) ([]Photo, error)
	Delete(ctx context.Context, id int) error
}
