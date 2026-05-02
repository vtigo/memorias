package domain

import "context"

type FilmFormat int

const (
	F35mm FilmFormat = iota
	F120mm
)

type FilmBrand struct {
	ID      int
	Name    string
	Country string
}

type FilmBrandRepository interface {
	Create(ctx context.Context, b FilmBrand) (FilmBrand, error)
	Update(ctx context.Context, b FilmBrand) (FilmBrand, error)
	FindByID(ctx context.Context, id int) (FilmBrand, error)
	FindAll(ctx context.Context) ([]FilmBrand, error)
	Delete(ctx context.Context, id int) error
}

type Film struct {
	ID         int
	BrandID    int
	Stock      string
	Format     FilmFormat
	ISO        int
	ExpiryYear int
}

type FilmRepository interface {
	Create(ctx context.Context, f Film) (Film, error)
	Update(ctx context.Context, f Film) (Film, error)
	FindByID(ctx context.Context, id int) (Film, error)
	FindAll(ctx context.Context) ([]Film, error)
	Delete(ctx context.Context, id int) error
}
