package model

type Comment struct {
	ID              int    `json:"id"`
	PostID          int    `json:"postId"`
	Content         string `json:"content"`
	AuthorID        int    `json:"authorId"`
	ParentCommentID *int   `json:"parentCommentId,omitempty"`
}
