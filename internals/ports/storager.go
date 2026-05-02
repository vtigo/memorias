package ports

import (
	"context"
	"io"
)

type FileStorager interface {
	Store(ctx context.Context, key string, r io.Reader, contentType string) error
	Read(ctx context.Context, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, key string) error
}
