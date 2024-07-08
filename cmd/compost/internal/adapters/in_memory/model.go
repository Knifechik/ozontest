package in_memory

import (
	"ozon_test_compost/cmd/compost/internal/app"
	"sort"
)

func postConvert(post app.Post) *Post {
	return &Post{
		ID:              post.ID,
		Title:           post.Title,
		Content:         post.Content,
		AuthorID:        post.AuthorID,
		CommentsAllowed: post.CommentsAllowed,
		CreatedAt:       post.CreatedAt,
	}
}

func (p Post) Convert() *app.Post {
	return &app.Post{
		ID:              p.ID,
		Title:           p.Title,
		Content:         p.Content,
		AuthorID:        p.AuthorID,
		CommentsAllowed: p.CommentsAllowed,
		CreatedAt:       p.CreatedAt,
	}
}

func commentConvert(comment app.Comment) *Comment {
	return &Comment{
		ID:              comment.ID,
		PostID:          comment.PostID,
		Content:         comment.Content,
		AuthorID:        comment.AuthorID,
		ParentCommentID: comment.ParentCommentID,
		CreatedAt:       comment.CreatedAt,
	}
}

func (c Comment) Convert() *app.Comment {
	return &app.Comment{
		ID:              c.ID,
		PostID:          c.PostID,
		Content:         c.Content,
		AuthorID:        c.AuthorID,
		ParentCommentID: c.ParentCommentID,
		CreatedAt:       c.CreatedAt,
	}
}

func sortComments(comments []app.Comment) []app.Comment {
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].CreatedAt.Before(comments[j].CreatedAt)
	})
	return comments
}

func paginationComments(comments []app.Comment, limit int, off *int) []app.Comment {
	var offset int
	switch {
	case off == nil:
		offset = 0
	case off != nil:
		offset = *off
	}
	comments = sortComments(comments)
	if limit > len(comments) && offset == 0 {
		return comments
	}
	if offset >= len(comments) {
		return nil
	}
	end := offset + limit
	if end > len(comments) {
		end = len(comments)
	}
	result := comments[offset:end]
	return result
}

func sortPosts(Posts []app.Post) []app.Post {
	sortedPosts := make([]Post, len(Posts))
	sort.Slice(sortedPosts, func(i, j int) bool {
		return Posts[i].CreatedAt.Before(Posts[j].CreatedAt)
	})
	return Posts
}
