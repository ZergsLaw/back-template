package app

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gofrs/uuid"

	"github.com/ZergsLaw/back-template/internal/dom"
)

// VerificationEmail check exists or not user email.
func (a *App) VerificationEmail(ctx context.Context, email string) error {
	email = strings.ToLower(email)
	_, err := a.repo.UserByEmail(ctx, email)
	switch {
	case errors.Is(err, ErrNotFound):
		return nil
	case err == nil:
		return ErrEmailExist
	default:
		return fmt.Errorf("a.repo.ByEmail: %w", err)
	}
}

// VerificationUsername check exists or not username.
func (a *App) VerificationUsername(ctx context.Context, username string) error {
	_, err := a.repo.UserByUsername(ctx, username)
	switch {
	case errors.Is(err, ErrNotFound):
		return nil
	case err == nil:
		return ErrUsernameExist
	default:
		return fmt.Errorf("a.repo.ByUsername: %w", err)
	}
}

// CreateUser create new user by params.
func (a *App) CreateUser(ctx context.Context, email, username, fullName, password string) (userID uuid.UUID, err error) {
	passHash, err := a.hash.Hashing(password)
	if err != nil {
		return uuid.Nil, fmt.Errorf("a.hash.Hashing: %w", err)
	}
	email = strings.ToLower(email)

	err = a.repo.Tx(ctx, func(repo Repo) error {
		newUser := User{
			Email:    email,
			Name:     username,
			FullName: fullName,
			PassHash: passHash,
			Status:   dom.UserStatusDefault,
		}

		userID, err = repo.UserSave(ctx, newUser)
		if err != nil {
			return fmt.Errorf("repo.UserSave: %w", err)
		}

		u, err := repo.UserByID(ctx, userID)
		if err != nil {
			return fmt.Errorf("repo.ByID: %w", err)
		}

		task := Task{
			User: *u,
			Kind: TaskKindEventUserAdd,
		}

		_, err = repo.SaveTask(ctx, task)
		if err != nil {
			return fmt.Errorf("repo.SaveTask: %w", err)
		}

		return nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

// Login make new session and returns sessions token.
func (a *App) Login(ctx context.Context, email, password string, origin Origin) (uuid.UUID, *Token, error) {
	email = strings.ToLower(email)
	user, err := a.repo.UserByEmail(ctx, email)
	if err != nil {
		return uuid.Nil, nil, fmt.Errorf("a.repo.ByEmail: %w", err)
	}

	if !a.hash.Compare(user.PassHash, []byte(password)) {
		return uuid.Nil, nil, ErrInvalidPassword
	}

	sessionID := a.id.New()
	token, err := a.auth.Token(sessionID)
	if err != nil {
		return uuid.Nil, nil, fmt.Errorf("a.auth.Token: %w", err)
	}

	err = a.sessions.SessionSave(ctx, Session{
		ID:     sessionID,
		Origin: origin,
		Token:  *token,
		UserID: user.ID,
	})
	if err != nil {
		return uuid.Nil, nil, fmt.Errorf("a.sessions.UserSave: %w", err)
	}

	return user.ID, token, nil
}

// UserByID get user by id.
func (a *App) UserByID(ctx context.Context, session Session, userID uuid.UUID) (*User, error) {
	if userID == uuid.Nil {
		userID = session.UserID
	}

	return a.repo.UserByID(ctx, userID)
}

// ListUserByFilters get users by filters.
func (a *App) ListUserByFilters(ctx context.Context, _ Session, filters SearchParams) ([]User, int, error) {
	return a.repo.SearchUsers(ctx, filters)
}

// Logout remove user's session.
func (a *App) Logout(ctx context.Context, session Session) error {
	return a.sessions.SessionDelete(ctx, session.ID)
}

// UpdatePassword update user's password.
func (a *App) UpdatePassword(ctx context.Context, session Session, oldPass, newPass string) error {
	user, err := a.repo.UserByID(ctx, session.UserID)
	if err != nil {
		return fmt.Errorf("a.repo.ByID: %w", err)
	}

	if !a.hash.Compare(user.PassHash, []byte(oldPass)) {
		return ErrInvalidPassword
	}

	if a.hash.Compare(user.PassHash, []byte(newPass)) {
		return ErrNotDifferent
	}

	passHash, err := a.hash.Hashing(newPass)
	if err != nil {
		return fmt.Errorf("a.hash.Hashing: %w", err)
	}
	user.PassHash = passHash

	_, err = a.repo.UserUpdate(ctx, *user)
	if err != nil {
		return fmt.Errorf("a.repo.UserUpdate: %w", err)
	}

	return nil
}

// UpdateUser update user's profile.
func (a *App) UpdateUser(ctx context.Context, session Session, username string, avatarID uuid.UUID) error {
	u, err := a.repo.UserByID(ctx, session.UserID)
	if err != nil {
		return fmt.Errorf("a.repo.ByID: %w", err)
	}

	if avatarID == uuid.Nil {
		avatarID = u.AvatarID
	}

	if avatarID != uuid.Nil {
		_, err = a.repo.GetAvatar(ctx, avatarID)
		if err != nil {
			return fmt.Errorf("a.repo.GetAvatarCache: %w", err)
		}
	}

	user := User{
		ID:       u.ID,
		Email:    u.Email,
		FullName: u.FullName,
		Name:     username,
		PassHash: u.PassHash,
		AvatarID: avatarID,
		Status:   u.Status,
	}

	_, err = a.repo.UserUpdate(ctx, user)
	if err != nil {
		return fmt.Errorf("a.repo.UserUpdate: %w", err)
	}

	return nil
}

func (a *App) GetUsersByIDs(ctx context.Context, _ Session, ids []uuid.UUID) ([]User, error) {
	return a.repo.UsersByIDs(ctx, ids)
}

// Auth get user session by token.
func (a *App) Auth(ctx context.Context, token string) (*Session, error) {
	subject, err := a.auth.Subject(token)
	if err != nil {
		return nil, fmt.Errorf("a.auth.Subject: %w", err)
	}

	session, err := a.sessions.SessionByID(ctx, subject)
	if err != nil {
		return nil, fmt.Errorf("a.session.ByID: %w", err)
	}

	return session, nil
}
