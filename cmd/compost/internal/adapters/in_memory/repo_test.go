package in_memory_test

import (
	"ozon_test_compost/cmd/compost/internal/app"
	"testing"
)

func TestRepo_Smoke(t *testing.T) {
	t.Parallel()

	ctx, r, assert := start(t)

	newPost := app.Post{
		Title:           "Test",
		Content:         "TestTest",
		AuthorID:        1,
		CommentsAllowed: true,
	}

	post, err := r.SavePost(ctx, newPost)

	assert.NoError(err)
	assert.NotZero(post)

	post2, err := r.PostByID(ctx, 1)

	assert.NoError(err)
	assert.NotZero(post2)
	assert.Equal(post, post2)

	newPost.Title = "TestNew"
	newPost.Content = "TestingNew"
	newPost.AuthorID = 2
	newPost.CommentsAllowed = true

	_, err = r.SavePost(ctx, newPost)
	assert.NoError(err)

	posts, err := r.GetAllPosts(ctx)
	assert.Len(posts, 2)

	newComment := app.Comment{
		PostID:          1,
		Content:         "TestTest",
		AuthorID:        2,
		ParentCommentID: nil,
	}

	comment, err := r.SaveComment(ctx, newComment)
	assert.NoError(err)
	assert.NotZero(comment)

	parent := 1
	parentId := &parent

	newComment.PostID = 1
	newComment.Content = "TestingNew"
	newComment.AuthorID = 3
	newComment.ParentCommentID = parentId

	comment2, err := r.SaveComment(ctx, newComment)
	assert.NoError(err)
	assert.NotZero(comment2)

	commentsWithoutParent, err := r.CommentsByID(ctx, 1, nil, 10, nil)

	assert.NoError(err)
	assert.Len(commentsWithoutParent, 1)

	commentWithParent, err := r.CommentsByID(ctx, 1, parentId, 10, nil)
	assert.NoError(err)
	assert.Len(commentWithParent, 1)

	assert.Contains(commentsWithoutParent, *comment)
	assert.Contains(commentWithParent, *comment2)

}
