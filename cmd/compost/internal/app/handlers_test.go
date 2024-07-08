package app_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"ozon_test_compost/cmd/compost/internal/app"
	"testing"
)

func TestResolvers_CreatePost(t *testing.T) {
	t.Parallel()

	var (
		newPost = app.Post{
			Title:           "Test",
			Content:         "some content",
			AuthorID:        1,
			CommentsAllowed: true,
		}
		repoRes = &app.Post{
			ID:              1,
			Title:           "Test",
			Content:         "some content",
			AuthorID:        1,
			CommentsAllowed: true,
		}
		wantRes = &app.Post{
			ID:              1,
			Title:           "Test",
			Content:         "some content",
			AuthorID:        1,
			CommentsAllowed: true,
		}
	)

	testcases := map[string]struct {
		repoRes *app.Post
		wantRes *app.Post
		repoErr error
		wantErr error
	}{
		"success":         {repoRes: repoRes, wantRes: wantRes, repoErr: nil, wantErr: nil},
		"r.repo.SavePost": {repoRes: nil, wantRes: nil, repoErr: errAny, wantErr: errAny},
	}

	for name, tc := range testcases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockApp, c, assert := start(t)

			mockApp.EXPECT().SavePost(gomock.Any(), newPost).Return(tc.repoRes, tc.repoErr)

			res, err := c.CreatePost(context.Background(), "Test", "some content",
				1, true)

			assert.ErrorIs(err, tc.wantErr)
			assert.Equal(tc.wantRes, res)
		})
	}
}

func TestResolvers_CreateComment(t *testing.T) {
	t.Parallel()

	var (
		comment = app.Comment{
			ID:              0,
			PostID:          1,
			Content:         "texttext",
			AuthorID:        1,
			ParentCommentID: nil,
		}
		commentAllowed = &app.Post{
			ID:              1,
			Title:           "Text",
			Content:         "texttext",
			AuthorID:        1,
			CommentsAllowed: true,
		}

		repoRes = &app.Comment{
			ID:              1,
			PostID:          1,
			Content:         "texttext",
			AuthorID:        1,
			ParentCommentID: nil,
		}

		wantRes = &app.Comment{
			ID:              1,
			PostID:          1,
			Content:         "texttext",
			AuthorID:        1,
			ParentCommentID: nil,
		}
	)

	testcases := map[string]struct {
		repoRes     *app.Comment
		wantRes     *app.Comment
		repoByIDErr error
		repoErr     error
		wantErr     error
	}{
		"success":            {repoRes: repoRes, wantRes: wantRes, repoByIDErr: nil, repoErr: nil, wantErr: nil},
		"r.repo.SaveComment": {repoRes: nil, wantRes: nil, repoByIDErr: nil, repoErr: errAny, wantErr: errAny},
	}

	for name, tc := range testcases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockApp, c, assert := start(t)

			mockApp.EXPECT().PostByID(gomock.Any(), 1).Return(commentAllowed, tc.repoByIDErr)
			mockApp.EXPECT().SaveComment(gomock.Any(), comment).Return(tc.repoRes, tc.repoErr)

			res, err := c.CreateComment(context.Background(), 1,
				"texttext", 1, nil)

			assert.ErrorIs(err, tc.wantErr)
			assert.Equal(tc.wantRes, res)
		})
	}
}

func TestResolvers_GetPost(t *testing.T) {
	t.Parallel()

	var (
		appRes = &app.Post{
			ID:              1,
			Title:           "Test",
			Content:         "some content",
			AuthorID:        1,
			CommentsAllowed: true,
		}
		wantRes = &app.Post{
			ID:              1,
			Title:           "Test",
			Content:         "some content",
			AuthorID:        1,
			CommentsAllowed: true,
		}
	)

	testcases := map[string]struct {
		appRes  *app.Post
		wantRes *app.Post
		appErr  error
		wantErr error
	}{
		"success":         {appRes: appRes, wantRes: wantRes, appErr: nil, wantErr: nil},
		"a.repo.PostByID": {appRes: nil, wantRes: nil, appErr: errAny, wantErr: errAny},
	}

	for name, tc := range testcases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockApp, c, assert := start(t)

			mockApp.EXPECT().PostByID(gomock.Any(), 1).Return(appRes, tc.appErr)

			res, err := c.GetPost(context.Background(), 1)

			assert.ErrorIs(err, tc.wantErr)
			assert.Equal(tc.wantRes, res)
		})
	}
}

func TestResolvers_GetPosts(t *testing.T) {
	t.Parallel()

	var (
		appRes = []app.Post{{
			ID:              1,
			Title:           "Test",
			Content:         "some content",
			AuthorID:        1,
			CommentsAllowed: true,
		}, {
			ID:              2,
			Title:           "TestNewContent",
			Content:         "new content",
			AuthorID:        1,
			CommentsAllowed: false,
		},
		}
		wantRes = []app.Post{{
			ID:              1,
			Title:           "Test",
			Content:         "some content",
			AuthorID:        1,
			CommentsAllowed: true,
		}, {
			ID:              2,
			Title:           "TestNewContent",
			Content:         "new content",
			AuthorID:        1,
			CommentsAllowed: false,
		},
		}
	)

	testcases := map[string]struct {
		appRes  []app.Post
		wantRes []app.Post
		appErr  error
		wantErr error
	}{
		"success":           {appRes: appRes, wantRes: wantRes, appErr: nil, wantErr: nil},
		"a.app.GetAllPosts": {appRes: nil, wantRes: nil, appErr: errAny, wantErr: errAny},
	}

	for name, tc := range testcases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockApp, c, assert := start(t)

			mockApp.EXPECT().GetAllPosts(gomock.Any()).Return(tc.appRes, tc.appErr)

			res, err := c.GetPosts(context.Background())

			assert.ErrorIs(err, tc.wantErr)
			assert.Equal(tc.wantRes, res)
		})
	}
}

func TestResolvers_Comments(t *testing.T) {
	t.Parallel()

	var (
		appRes = []app.Comment{
			{
				ID:              1,
				PostID:          1,
				Content:         "some content",
				AuthorID:        1,
				ParentCommentID: nil,
			}, {
				ID:              2,
				PostID:          1,
				Content:         "new content",
				AuthorID:        2,
				ParentCommentID: nil,
			},
		}
		wantRes = []app.Comment{{
			ID:              1,
			PostID:          1,
			Content:         "some content",
			AuthorID:        1,
			ParentCommentID: nil,
		}, {
			ID:              2,
			PostID:          1,
			Content:         "new content",
			AuthorID:        2,
			ParentCommentID: nil,
		},
		}
	)

	testcases := map[string]struct {
		appRes  []app.Comment
		wantRes []app.Comment
		appErr  error
		wantErr error
	}{
		"success":            {appRes: appRes, wantRes: wantRes, appErr: nil, wantErr: nil},
		"a.app.CommentsByID": {appRes: nil, wantRes: nil, appErr: errAny, wantErr: errAny},
	}

	for name, tc := range testcases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockApp, c, assert := start(t)

			mockApp.EXPECT().CommentsByID(gomock.Any(), 1, nil, 10, nil).Return(tc.appRes, tc.appErr)

			res, err := c.Comments(context.Background(), 1, 10, nil)

			assert.ErrorIs(err, tc.wantErr)
			assert.Equal(tc.wantRes, res)
		})
	}
}

func TestResolvers_ChildComments(t *testing.T) {
	t.Parallel()

	var (
		parent          = 1
		parentCommentID = &parent
		appRes          = []app.Comment{
			{
				ID:              1,
				PostID:          1,
				Content:         "some content",
				AuthorID:        1,
				ParentCommentID: parentCommentID,
			}, {
				ID:              2,
				PostID:          1,
				Content:         "new content",
				AuthorID:        2,
				ParentCommentID: parentCommentID,
			},
		}
		wantRes = []app.Comment{{
			ID:              1,
			PostID:          1,
			Content:         "some content",
			AuthorID:        1,
			ParentCommentID: parentCommentID,
		}, {
			ID:              2,
			PostID:          1,
			Content:         "new content",
			AuthorID:        2,
			ParentCommentID: parentCommentID,
		},
		}
	)

	testcases := map[string]struct {
		appRes  []app.Comment
		wantRes []app.Comment
		appErr  error
		wantErr error
	}{
		"success":             {appRes: appRes, wantRes: wantRes, appErr: nil, wantErr: nil},
		"a.repo.CommentsByID": {appRes: nil, wantRes: nil, appErr: errAny, wantErr: errAny},
	}

	for name, tc := range testcases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockApp, c, assert := start(t)

			mockApp.EXPECT().CommentsByID(gomock.Any(), 1, parentCommentID, 10, nil).Return(tc.appRes, tc.appErr)

			res, err := c.ChildComments(context.Background(), 1, parentCommentID, 10, nil)

			assert.ErrorIs(err, tc.wantErr)
			assert.Equal(tc.wantRes, res)
		})
	}
}
