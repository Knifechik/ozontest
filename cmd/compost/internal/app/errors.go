package app

import "errors"

var (
	ErrNotAllowed        = errors.New("comments not allowed")
	ErrOverCommentLength = errors.New("comments over maximum length")
)
