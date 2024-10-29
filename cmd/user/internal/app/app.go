// Package app contains business logic.
package app

// App manages business logic methods.
type App struct {
	repo     Repo
	hash     PasswordHash
	auth     Auth
	id       ID
	sessions Sessions
	file     FileStore
}

// New build and returns new App.
func New(r Repo, ph PasswordHash, a Auth, id ID, s Sessions, f FileStore) *App {
	return &App{
		repo:     r,
		auth:     a,
		id:       id,
		hash:     ph,
		sessions: s,
		file:     f,
	}
}
