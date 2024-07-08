package app

import "context"

type Repo interface {
	SavePost(context.Context, Post) (*Post, error)
	PostByID(context.Context, int) (*Post, error)
	GetAllPosts(context.Context) ([]Post, error)
	SaveComment(context.Context, Comment) (*Comment, error)
	CommentsByID(context.Context, int, *int, int, *int) ([]Comment, error)
}
