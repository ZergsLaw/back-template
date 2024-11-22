package repo

import (
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sipki-tech/database"
	"github.com/sipki-tech/database/connectors"
	"github.com/sipki-tech/database/migrations"

	"github.com/Bar-Nik/back-template/cmd/twitter/internal/app"
)

var _ app.Repo = &Repo{}

type (
	Config struct {
		Cockroach  connectors.CockroachDB
		MigrateDir string
		Driver     string
	}

	Repo struct {
		sql *database.SQL
	}
)

func New(ctx context.Context, req *prometheus.Registry, namespace string, cfg Config) (*Repo, error) {
	const subsystem = "repo"

	m := database.NewMetrics(req, namespace, subsystem, new(app.Repo))
	returnErrs := []error{
		app.ErrNotFound,
		app.ErrInvalidArgument,
		app.ErrAccessDenied,
	}

	migrates, err := migrations.Parse(cfg.MigrateDir)
	if err != nil {
		return nil, fmt.Errorf("migration.Parse: %w", err)
	}

	err = migrations.Run(ctx, cfg.Driver, &cfg.Cockroach, migrations.Up, migrates)
	if err != nil {
		return nil, fmt.Errorf("migrations.Run: %w", err)
	}

	conn, err := database.NewSQL(ctx, cfg.Driver, database.SQLConfig{Metrics: m,
		ReturnErrs: returnErrs,
	}, &cfg.Cockroach)

	if err != nil {
		return nil, fmt.Errorf("librepo.NewCockroach: %w", err)
	}

	return &Repo{
		sql: conn,
	}, nil
}

func (r *Repo) Close() error {
	return r.sql.Close()
}

func (r *Repo) Save(ctx context.Context, p app.Post) (id uuid.UUID, err error) {
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		newPost := convert(p)

		const query = `
		insert into 
		posts 
		    (post, is_published, user_id) 
		values 
			($1, $2, $3)
		returning id
		`

		err := db.GetContext(ctx, &id, query, newPost.POST, newPost.IsPublished, newPost.UserId)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}

		return nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *Repo) Update(ctx context.Context, p app.Post) (upPost *app.Post, err error) {
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		updatePost := convert(p)
		const query = `
		update posts
		set
		post = $1,
		updated_at = now(),
		user_id = $2
		where id = $3
		retutning *`

		var res post
		err := db.GetContext(ctx, &res, query, updatePost.POST, updatePost.UserId, updatePost.ID)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}
		upPost = res.convert()

		return nil
	})
	if err != nil {
		return nil, err
	}
	return upPost, nil
}

func (r *Repo) Delete(ctx context.Context, id uuid.UUID) error {
	return r.sql.NoTx(func(db *sqlx.DB) error {
		const query = `
		delete
		from posts
		where id = $1 returning *`

		err := db.GetContext(ctx, &post{}, query, id)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}

		return nil
	})
}

func (r *Repo) PostById(ctx context.Context, id uuid.UUID) (p *app.Post, err error) {
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		const query = `select * from posts where id = $1`

		res := post{}
		err = db.GetContext(ctx, &res, query, id)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}
		p = res.convert()
		return nil
	})
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *Repo) PostsByUser(ctx context.Context, user_id uuid.UUID, offset int, limit int) (posts []app.Post, err error) {
	err = r.sql.NoTx(func(db *sqlx.DB) error {
		const query = `select post from posts where user_id = $1 offset $2 limit $3`

		res := make([]post, 0)
		err = db.SelectContext(ctx, &res, query, user_id, offset, limit)
		if err != nil {
			return fmt.Errorf("db.GetContext: %w", convertErr(err))
		}

		var posts []app.Post
		for i := 1; i <= len(res); i++ {
			posts[i] = *res[i].convert()
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return posts, nil
}
