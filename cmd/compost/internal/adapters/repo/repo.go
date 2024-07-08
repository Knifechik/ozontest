package repo

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sipki-tech/database/migrations"
	"ozon_test_compost/cmd/compost/internal/app"
)

var _ app.Repo = &Repo{}

type (
	Config struct {
		Postgres   Connector
		MigrateDir string
		Driver     string
	}

	Connector struct {
		ConnectionDSN string `yaml:"connection_dsn"`
	}

	Repo struct {
		sql *sqlx.DB
	}
)

func (c Connector) DSN() (string, error) {
	return c.ConnectionDSN, nil
}

func New(ctx context.Context, cfg Config) (*Repo, error) {

	migrates, err := migrations.Parse(cfg.MigrateDir)
	if err != nil {
		return nil, fmt.Errorf("migration.Parse: %w", err)
	}

	err = migrations.Run(ctx, cfg.Driver, &cfg.Postgres, migrations.Up, migrates)
	if err != nil {
		return nil, fmt.Errorf("migration.Run: %w", err)
	}

	dsn, err := cfg.Postgres.DSN()
	if err != nil {
		return nil, fmt.Errorf("connector.DSN: %w", err)
	}

	db, err := sqlx.Open(cfg.Driver, dsn)
	if err != nil {
		return nil, fmt.Errorf("sqlx.Open: %w", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("db.PingContext: %w", err)
	}

	return &Repo{
		db,
	}, nil
}

func (r *Repo) Close() error {
	return r.sql.Close()
}

func (r *Repo) SavePost(ctx context.Context, post app.Post) (*app.Post, error) {
	p := postConvert(post)
	const query = `INSERT INTO post_table
    (title, content, author_id, comments_allowed) 
    values ($1, $2, $3, $4) returning *`
	res := Post{}
	err := r.sql.GetContext(ctx, &res, query, p.Title, p.Content, p.AuthorID, p.CommentsAllowed)
	if err != nil {
		return nil, fmt.Errorf("db.GetContext: %w", err)
	}

	return res.Convert(), nil
}

func (r *Repo) PostByID(ctx context.Context, id int) (*app.Post, error) {
	const query = `SELECT * FROM post_table WHERE id = $1`
	res := Post{}
	err := r.sql.GetContext(ctx, &res, query, id)
	if err != nil {
		return nil, fmt.Errorf("db.SelectContext: %w", err)
	}

	return res.Convert(), nil
}

func (r *Repo) GetAllPosts(ctx context.Context) ([]app.Post, error) {
	const query = `SELECT * FROM post_table ORDER BY created_at DESC`
	posts := []Post{}
	err := r.sql.SelectContext(ctx, &posts, query)
	if err != nil {
		return nil, fmt.Errorf("db.SelectContext: %w", err)
	}

	res := make([]app.Post, 0, len(posts))
	for _, post := range posts {
		res = append(res, *post.Convert())
	}

	return res, nil
}

func (r *Repo) SaveComment(ctx context.Context, comment app.Comment) (*app.Comment, error) {
	c := commentConvert(comment)
	const query = `INSERT INTO comment_table
    (post_id, content,author_id, parent_comment_id) 
    values ($1, $2, $3, $4) returning *`
	res := Comment{}
	err := r.sql.GetContext(ctx, &res, query, c.PostID, c.Content, c.AuthorID, c.ParentCommentID)
	if err != nil {
		return nil, fmt.Errorf("db.GetContext: %w", err)
	}

	return res.Convert(), nil
}

func (r *Repo) CommentsByID(ctx context.Context, id int, parentID *int, limit int, offset *int) ([]app.Comment, error) {
	comments := []Comment{}
	var err error

	if parentID != nil {
		const query = `SELECT * FROM comment_table WHERE post_id = $1 and parent_comment_id = $2 ORDER BY created_at DESC LIMIT $3 OFFSET $4`
		err = r.sql.SelectContext(ctx, &comments, query, id, parentID, limit, offset)

	}
	if parentID == nil {
		const query = `SELECT * FROM comment_table WHERE post_id = $1 and parent_comment_id is null ORDER BY created_at DESC LIMIT $2 OFFSET $3`
		err = r.sql.SelectContext(ctx, &comments, query, id, limit, offset)
	}

	if err != nil {
		return nil, fmt.Errorf("db.SelectContext: %w", err)
	}

	res := make([]app.Comment, 0, len(comments))
	for _, comment := range comments {
		res = append(res, *comment.Convert())
	}

	return res, nil
}
