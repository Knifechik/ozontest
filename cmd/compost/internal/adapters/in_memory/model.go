package in_memory

import (
	"ozon_test_compost/cmd/compost/internal/app"
)

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
