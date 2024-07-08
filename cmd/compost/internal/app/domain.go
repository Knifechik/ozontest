package app

import "time"

type Post struct {
	ID              int
	Title           string
	Content         string
	AuthorID        int
	CommentsAllowed bool
	CreatedAt       time.Time
}

type Comment struct {
	ID              int
	PostID          int
	Content         string
	AuthorID        int
	ParentCommentID *int
	CreatedAt       time.Time
}
