package domain

import "context"

type Album struct {
	ID          int
	Name        string
	Description string
	PhotoIDs    []int
}

type AlbumRepository interface {
	Create(ctx context.Context, a Album) (Album, error)
	Update(ctx context.Context, a Album) (Album, error)
	FindByID(ctx context.Context, id int) (Album, error)
	FindAll(ctx context.Context) ([]Album, error)
	Delete(ctx context.Context, id int) error
}
