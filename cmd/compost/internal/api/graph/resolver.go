package graph

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"net/http"
	"ozon_test_compost/cmd/compost/internal/app"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen generate

type application interface {
	CreatePost(ctx context.Context, title, content string, authorID int, commentAllowed bool) (*app.Post, error)
	GetPost(ctx context.Context, id int) (*app.Post, error)
	GetPosts(ctx context.Context) ([]app.Post, error)
	CreateComment(ctx context.Context, postID int, content string, authorID int, parentCommentID *int) (*app.Comment, error)
	Comments(ctx context.Context, ID, limit int, offset *int) ([]app.Comment, error)
	ChildComments(ctx context.Context, postID int, parentID *int, limit int, offset *int) ([]app.Comment, error)
	Subscriptions(ctx context.Context, postID, userID int) (<-chan app.Comment, error)
}

type Resolver struct {
	app application
}

func New(a application) *http.ServeMux {
	newResolver := NewResolver(a)

	mux := http.NewServeMux()

	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: newResolver}))

	srv.AddTransport(&transport.Websocket{})

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	return mux
}
func NewResolver(app application) *Resolver {
	return &Resolver{app: app}
}
