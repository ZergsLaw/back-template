package files

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/minio/minio-go/v7"

	"github.com/ZergsLaw/back-template/cmd/back/internal/app"
)

// DeleteFile implements app.FileStore.
func (c *Client) DeleteFile(ctx context.Context, id uuid.UUID) error {
	err := c.store.RemoveObject(ctx, bucketFile, id.String(), minio.RemoveObjectOptions{})
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
			headerUserID:  f.UserID.String(),
		},
		ContentType: f.ContentType,
		PartSize:    partSize,
	})
	if err != nil {
		return uuid.Nil, fmt.Errorf("c.store.PutObject: %w", err)
	}

	return id, nil
}

// DownloadFile implements app.FileStore.
func (c *Client) DownloadFile(ctx context.Context, id uuid.UUID) (*app.File, error) {
	file, err := c.store.GetObject(ctx, bucketFile, id.String(), minio.GetObjectOptions{})
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

	f := &app.File{
		ReadSeekCloser: file,
		ID:             id,
		UserID:         uuid.FromStringOrNil(stat.UserMetadata["User_id"]),
		Name:           stat.Metadata.Get(headerSrcName),
		Size:           stat.Size,
		ModTime:        stat.LastModified,
		ContentType:    stat.ContentType,
	}

	return f, nil
}
