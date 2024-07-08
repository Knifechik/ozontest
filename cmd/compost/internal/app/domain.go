package app

type Post struct {
	ID              int
	Title           string
	Content         string
	AuthorID        int
	CommentsAllowed bool
}

type Comment struct {
	ID              int
	PostID          int
	Content         string
	AuthorID        int
	ParentCommentID *int
}
