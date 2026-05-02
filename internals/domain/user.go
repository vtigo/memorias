package domain

import "context"

type User struct {
	ID           int
	Nickname     string
	Username     string
	PasswordHash string
}

type UserRepository interface {
	Create(ctx context.Context, u User) (User, error)
	Update(ctx context.Context, u User) (User, error)
	FindByID(ctx context.Context, id int) (User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	Delete(ctx context.Context, id int) error
}
