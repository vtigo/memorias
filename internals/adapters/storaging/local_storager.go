package storaging

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type LocalStorager struct {
	dir string
}

func NewLocalStorager(dir string) *LocalStorager {
	return &LocalStorager{
		dir: dir,
	}
}

func (s *LocalStorager) Store(ctx context.Context,
			    			  key string,
							  r io.Reader) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	file, err := os.Create(filepath.Join(s.dir, key))
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	if _, err := io.Copy(file, r); err != nil {
		return fmt.Errorf("failed to write data to the created file: %w", err)
	}
	return nil
}

func (s *LocalStorager) Read(ctx context.Context,
							 key string) (io.ReadCloser, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	file, err := os.Open(filepath.Join(s.dir, key))
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	return file, nil
}

func (s *LocalStorager) Delete(ctx context.Context, key string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := os.Remove(filepath.Join(s.dir, key)); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}
