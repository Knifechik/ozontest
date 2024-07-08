package repo

import (
	"ozon_test_compost/cmd/compost/internal/app"
	"time"
)

type Post struct {
	ID              int       `db:"id" json:"id"`
	Title           string    `db:"title" json:"title"`
	Content         string    `db:"content" json:"content"`
	AuthorID        int       `db:"author_id" json:"authorId"`
	CommentsAllowed bool      `db:"comments_allowed" json:"commentsAllowed"`
	CreatedAt       time.Time `db:"created_at" json:"createdAt"`
}

type Comment struct {
	ID              int       `db:"id" json:"id"`
	PostID          int       `db:"post_id" json:"postId"`
	Content         string    `db:"content" json:"content"`
	AuthorID        int       `db:"author_id" json:"authorId"`
	ParentCommentID *int      `db:"parent_comment_id" json:"parentCommentId,omitempty"`
	CreatedAt       time.Time `db:"created_at" json:"createdAt"`
}

func postConvert(post app.Post) *Post {
	return &Post{
		ID:              post.ID,
		Title:           post.Title,
		Content:         post.Content,
		AuthorID:        post.AuthorID,
		CommentsAllowed: post.CommentsAllowed,
	}
}

func (p Post) Convert() *app.Post {
	return &app.Post{
		ID:              p.ID,
		Title:           p.Title,
		Content:         p.Content,
		AuthorID:        p.AuthorID,
		CommentsAllowed: p.CommentsAllowed,
	}
}

func commentConvert(comment app.Comment) *Comment {
	return &Comment{
		ID:              comment.ID,
		PostID:          comment.PostID,
		Content:         comment.Content,
		AuthorID:        comment.AuthorID,
		ParentCommentID: comment.ParentCommentID,
	}
}

func (c Comment) Convert() *app.Comment {
	return &app.Comment{
		ID:              c.ID,
		PostID:          c.PostID,
		Content:         c.Content,
		AuthorID:        c.AuthorID,
		ParentCommentID: c.ParentCommentID,
	}
}
