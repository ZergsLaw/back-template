package app

import (
	"context"

	"github.com/gofrs/uuid"
)

type (
	// Repo interface for user data repository.
	Repo interface {
		FileInfoRepo
		TaskRepo
		// Tx starts transaction in database.
		// Errors: unknown.
		Tx(ctx context.Context, f func(Repo) error) error
		// UserSave adds to the new user to repository.
		// Errors: ErrEmailExist, ErrUsernameExist, unknown.
		UserSave(context.Context, User) (uuid.UUID, error)
		// UserUpdate update user info.
		// Errors: ErrUsernameExist, ErrEmailExist, unknown.
		UserUpdate(context.Context, User) (*User, error)
		// UserByID returns user info by id.
		// Errors: ErrNotFound, unknown.
		UserByID(context.Context, uuid.UUID) (*User, error)
		// UserByEmail returns user info by email.
		// Errors: ErrNotFound, unknown.
		UserByEmail(context.Context, string) (*User, error)
		// UserByUsername returns user info by username.
		// Errors: ErrNotFound, unknown.
		UserByUsername(context.Context, string) (*User, error)
		// SearchUsers returns list user info.
		// Errors: unknown.
		SearchUsers(context.Context, SearchParams) ([]User, int, error)
		// UsersByIDs returns list of users.
		// Errors: ErrNotFound, unknown.
		UsersByIDs(ctx context.Context, ids []uuid.UUID) (users []User, err error)
	}
	// FileInfoRepo provides to file info repository
	FileInfoRepo interface {
		// SaveAvatar adds to the new cache about user avatar to repository.
		// Errors: ErrUserIDAndFileIDExist, ErrMaximumNumberOfStoredFilesReached, unknown.
		SaveAvatar(ctx context.Context, fileCache AvatarInfo) error
		// DeleteAvatar delete cache about user avatar info.
		// Errors: unknown.
		DeleteAvatar(ctx context.Context, userID, fileID uuid.UUID) error
		// GetAvatar returns cache about user avatar by id.
		// Errors: ErrNotFound, unknown.
		GetAvatar(ctx context.Context, fileID uuid.UUID) (*AvatarInfo, error)
		// ListAvatarByUserID returns list cache user file.
		// Errors: unknown.
		ListAvatarByUserID(ctx context.Context, userID uuid.UUID) ([]AvatarInfo, error)
		// GetCountAvatars returns count user avatars.
		// Errors: ErrNotFound, unknown.
		GetCountAvatars(ctx context.Context, ownerID uuid.UUID) (total int, err error)
	}

	// FileStore interface for saving and getting files.
	FileStore interface {
		// UploadAvatar save new file in database.
		// Errors: unknown.
		UploadAvatar(ctx context.Context, f Avatar) (uuid.UUID, error)
		// DownloadAvatar get file by id.
		// Errors: unknown.
		DownloadAvatar(ctx context.Context, id uuid.UUID) (*Avatar, error)
		// DeleteAvatar delete file by id.
		// Errors: unknown.
		DeleteAvatar(ctx context.Context, id uuid.UUID) error
		// UploadFile save new file in database.
		// Errors: unknown.
		UploadFile(ctx context.Context, f File) (uuid.UUID, error)
	}

	// PasswordHash module responsible for hashing password.
	PasswordHash interface {
		// Hashing returns the hashed version of the password.
		// Errors: unknown.
		Hashing(password string) ([]byte, error)
		// Compare compares two passwords for matches.
		Compare(hashedPassword []byte, password []byte) bool
	}

	// TaskRepo interface for saving tasks.
	TaskRepo interface {
		// SaveTask adds new task to repository.
		// Errors: unknown.
		SaveTask(context.Context, Task) (uuid.UUID, error)
		// FinishTask set column Task.FinishedAt task.
		// Errors: unknown.
		FinishTask(context.Context, uuid.UUID) error
		// ListActualTask returns list task by limit and ordered by created_at (ask).
		// Return tasks without Task.FinishedAt.
		// Errors: unknown.
		ListActualTask(context.Context, int) ([]Task, error)
	}

	// Sessions module for manager user's session.
	Sessions interface {
		// SessionSave saves the new user session in a database.
		// Errors: unknown.
		SessionSave(context.Context, Session) error
		// SessionByID returns user session by session id.
		// Errors: ErrNotFound, unknown.
		SessionByID(context.Context, uuid.UUID) (*Session, error)
		// SessionDelete removes user session.
		// Errors: ErrNotFound, unknown.
		SessionDelete(context.Context, uuid.UUID) error
	}

	// Auth interface for generate access and refresh token by subject.
	Auth interface {
		// Token generate tokens by subject with expire time.
		// Errors: unknown.
		Token(uuid.UUID) (*Token, error)
		// Subject unwrap Subject info from token.
		// Errors: ErrInvalidToken, ErrExpiredToken, unknown.
		Subject(token string) (uuid.UUID, error)
	}

	// ID generator for session.
	ID interface {
		// New generate new ID for session.
		New() uuid.UUID
	}
)
