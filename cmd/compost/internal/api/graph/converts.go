package graph

import (
	"ozon_test_compost/cmd/compost/internal/api/graph/model"
	"ozon_test_compost/cmd/compost/internal/app"
)

func convertPost(post app.Post) *model.Post {
	return &model.Post{
		ID:              post.ID,
		Title:           post.Title,
		Content:         post.Content,
		AuthorID:        post.AuthorID,
		CommentsAllowed: post.CommentsAllowed,
	}
}

func convertComment(comment app.Comment) *model.Comment {
	return &model.Comment{
		ID:              comment.ID,
		PostID:          comment.PostID,
		AuthorID:        comment.AuthorID,
		Content:         comment.Content,
		ParentCommentID: comment.ParentCommentID,
	}
}
