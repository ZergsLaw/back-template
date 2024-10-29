package files

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/minio/minio-go/v7"

	"github.com/ZergsLaw/back-template/cmd/user/internal/app"
)

// UploadAvatar implements app.FileStore.
func (c *Client) UploadAvatar(ctx context.Context, f app.Avatar) (uuid.UUID, error) {
	id := uuid.Must(uuid.NewV4())

	const partSize = 1024 * 1024 / 2
	_, err := c.store.PutObject(ctx, bucketAvatars, id.String(), f, f.Size, minio.PutObjectOptions{
		UserMetadata: map[string]string{
			headerSrcName: f.Name,
		},
		ContentType: f.ContentType,
		PartSize:    partSize,
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("c.store.PutObject: %w", err)
	}

	return id, nil
}

// DownloadAvatar implements app.FileStore.
func (c *Client) DownloadAvatar(ctx context.Context, id uuid.UUID) (*app.Avatar, error) {
	file, err := c.store.GetObject(ctx, bucketAvatars, id.String(), minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("c.store.GetObject: %w", err)
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, fmt.Errorf("file.stat: %w", err)
	}

	if stat.IsDeleteMarker {
		return nil, app.ErrNotFound
	}

	f := &app.Avatar{
		ReadSeekCloser: file,
		ID:             id,
		Name:           stat.Metadata.Get(headerSrcName),
		Size:           stat.Size,
		ModTime:        stat.LastModified,
		ContentType:    stat.ContentType,
	}

	return f, nil
}

// DeleteAvatar implements app.FileStore.
func (c *Client) DeleteAvatar(ctx context.Context, id uuid.UUID) error {
	err := c.store.RemoveObject(ctx, bucketAvatars, id.String(), minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("c.store.RemoveObject: %w", err)
	}

	return nil
}

// UploadFile implements app.FileStore.
func (c *Client) UploadFile(ctx context.Context, f app.File) (uuid.UUID, error) {
	id := uuid.Must(uuid.NewV4())

	const partSize = 1024 * 1024 / 2
	_, err := c.store.PutObject(ctx, bucketFile, id.String(), f, f.Size, minio.PutObjectOptions{
		UserMetadata: map[string]string{
			headerSrcName: f.Name,
		},
		ContentType: f.ContentType,
		PartSize:    partSize,
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("c.store.PutObject: %w", err)
	}

	return id, nil
}
