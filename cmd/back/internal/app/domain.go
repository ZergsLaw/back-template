package app

import (
	"io"
	"net"
	"strings"
	"time"

	"github.com/gofrs/uuid"

	"github.com/ZergsLaw/back-template/internal/dom"
)

type (
	// SearchParams params for search users.
	SearchParams struct {
		OwnerID        uuid.UUID
		Username       string
		FullName       string
		Statuses       []dom.UserStatus
		Email          string
		StartCreatedAt time.Time
		EndCreatedAt   time.Time
		Limit          uint64
		Offset         uint64
	}

	// User contains user information.
	User struct {
		ID        uuid.UUID
		Email     string
		FullName  string
		Name      string
		PassHash  []byte
		AvatarID  uuid.UUID
		Status    dom.UserStatus
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	// Avatar contains file information.
	Avatar struct {
		ID          uuid.UUID
		Name        string
		ContentType string
		Size        int64
		ModTime     time.Time
		io.ReadSeekCloser
	}

	// File contains file information.
	File struct {
		ID          uuid.UUID
		UserID      uuid.UUID
		Name        string
		ContentType string
		Size        int64
		ModTime     time.Time
		io.ReadSeekCloser
	}

	// AvatarInfo struct for caching info for finding file.
	AvatarInfo struct {
		FileID    uuid.UUID
		OwnerID   uuid.UUID
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	// TaskKind represents kind of task.
	TaskKind uint8

	// Task contains information for executing any deferred logic.
	Task struct {
		ID         uuid.UUID
		User       User
		Kind       TaskKind
		CreatedAt  time.Time
		UpdatedAt  time.Time
		FinishedAt time.Time
	}

	// FileFormat represents format of file.
	FileFormat uint8

	// SolutionStatus decision made at the time of the update.
	SolutionStatus uint8

	// StatusUpdateRequest user status update request.
	StatusUpdateRequest struct {
		ID             uuid.UUID
		UserID         uuid.UUID
		SolutionStatus SolutionStatus
		CreatedAt      time.Time
		UpdatedAt      time.Time
	}

	// SearchStatusUpdateRequest params for search request for update.
	SearchStatusUpdateRequest struct {
		SolutionStatus SolutionStatus
		Limit          uint
		Offset         uint
	}

	// Token contains auth token.
	Token struct {
		// Generate by Auth contract.
		Value string
	}

	// Origin information about req user.
	Origin struct {
		IP        net.IP
		UserAgent string
	}

	// Session contains session info for identify a user.
	Session struct {
		ID        uuid.UUID // Generate by repository layer.
		Origin    Origin
		Token     Token
		UserID    uuid.UUID
		CreatedAt time.Time // Generate by repository layer.
		UpdatedAt time.Time // Generate by repository layer.
	}
)

//go:generate stringer -output=stringer.TaskKind.go -type=TaskKind -trimprefix=TaskKind
const (
	_ TaskKind = iota
	TaskKindEventUserAdd
	TaskKindEventUserDel
	TaskKindEventUserUpdate
)

//go:generate stringer -output=stringer.FileFormat.go -type=FileFormat -trimprefix=FileFormat
const (
	_ FileFormat = iota
	FileFormatWebp
	FileFormatPng
	FileFormatJpeg
	FileFormatGif
	FileFormatRaw
	FileFormatSvg
)

//go:generate stringer -output=stringer.SolutionStatus.go -type=SolutionStatus -trimprefix=SolutionStatus
const (
	_ SolutionStatus = iota
	SolutionStatusNew
	SolutionStatusApprove
	SolutionStatusCancel
)

func validateFileFormat(format string) error {
	switch format {
	case strings.ToLower(FileFormatWebp.String()), strings.ToLower(FileFormatPng.String()),
		strings.ToLower(FileFormatJpeg.String()), strings.ToLower(FileFormatGif.String()),
		strings.ToLower(FileFormatRaw.String()), strings.ToLower(FileFormatSvg.String()):
		return nil
	default:
		return ErrInvalidImageFormat
	}
}
