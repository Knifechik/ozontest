package graph_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"ozon_test_compost/cmd/compost/internal/api/graph/model"
	"ozon_test_compost/cmd/compost/internal/app"
	"testing"
)

func TestResolvers_CreatePost(t *testing.T) {
	t.Parallel()

	var (
		appRes = &app.Post{
			ID:              1,
			Title:           "Test",
			Content:         "some content",
			AuthorID:        1,
			CommentsAllowed: true,
		}
		wantRes = &model.Post{
			ID:              1,
			Title:           "Test",
			Content:         "some content",
			AuthorID:        1,
			CommentsAllowed: true,
		}
	)

	testcases := map[string]struct {
		appRes  *app.Post
		wantRes *model.Post
		appErr  error
		wantErr error
	}{
		"success":          {appRes: appRes, wantRes: wantRes, appErr: nil, wantErr: nil},
		"a.app.CreatePost": {appRes: nil, wantRes: nil, appErr: errAny, wantErr: errAny},
	}

	for name, tc := range testcases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockApp, c, assert := start(t)

			mockApp.EXPECT().CreatePost(gomock.Any(), "Test",
				"some content", 1, true).Return(tc.appRes, tc.appErr)

			res, err := c.Mutation().CreatePost(context.Background(), "Test",
				"some content", 1, true)

			assert.ErrorIs(err, tc.wantErr)
			assert.Equal(tc.wantRes, res)
		})
	}
}

func TestResolvers_Post(t *testing.T) {
	t.Parallel()

	var (
		appRes = &app.Post{
			ID:              1,
			Title:           "Test",
			Content:         "some content",
			AuthorID:        1,
			CommentsAllowed: true,
		}
		wantRes = &model.Post{
			ID:              1,
			Title:           "Test",
			Content:         "some content",
			AuthorID:        1,
			CommentsAllowed: true,
		}
	)

	testcases := map[string]struct {
		appRes  *app.Post
		wantRes *model.Post
		appErr  error
		wantErr error
	}{
		"success":       {appRes: appRes, wantRes: wantRes, appErr: nil, wantErr: nil},
		"a.app.GetPost": {appRes: nil, wantRes: nil, appErr: errAny, wantErr: errAny},
	}

	for name, tc := range testcases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockApp, c, assert := start(t)

			mockApp.EXPECT().GetPost(gomock.Any(), 1).Return(tc.appRes, tc.appErr)

			res, err := c.Query().Post(context.Background(), 1)

			assert.ErrorIs(err, tc.wantErr)
			assert.Equal(tc.wantRes, res)
		})
	}
}

func TestResolvers_Posts(t *testing.T) {
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
		wantRes = []*model.Post{{
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
		wantRes []*model.Post
		appErr  error
		wantErr error
	}{
		"success":        {appRes: appRes, wantRes: wantRes, appErr: nil, wantErr: nil},
		"a.app.GetPosts": {appRes: nil, wantRes: nil, appErr: errAny, wantErr: errAny},
	}

	for name, tc := range testcases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockApp, c, assert := start(t)

			mockApp.EXPECT().GetPosts(gomock.Any()).Return(tc.appRes, tc.appErr)

			res, err := c.Query().Posts(context.Background())

			assert.ErrorIs(err, tc.wantErr)
			assert.Equal(tc.wantRes, res)
		})
	}
}

func TestResolvers_CreateComment(t *testing.T) {
	t.Parallel()

	var (
		appRes = &app.Comment{
			ID:              5,
			PostID:          1,
			Content:         "some content",
			AuthorID:        3,
			ParentCommentID: nil,
		}
		wantRes = &model.Comment{
			ID:              5,
			PostID:          1,
			Content:         "some content",
			AuthorID:        3,
			ParentCommentID: nil,
		}
	)

	testcases := map[string]struct {
		appRes  *app.Comment
		wantRes *model.Comment
		appErr  error
		wantErr error
	}{
		"success":     {appRes: appRes, wantRes: wantRes, appErr: nil, wantErr: nil},
		"a.app.Posts": {appRes: nil, wantRes: nil, appErr: errAny, wantErr: errAny},
	}

	for name, tc := range testcases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockApp, c, assert := start(t)

			mockApp.EXPECT().CreateComment(gomock.Any(), 1,
				"some content", 3, nil).Return(tc.appRes, tc.appErr)

			res, err := c.Mutation().CreateComment(context.Background(), 1,
				"some content", 3, nil)

			assert.ErrorIs(err, tc.wantErr)
			assert.Equal(tc.wantRes, res)
		})
	}
}

func TestResolvers_Comments(t *testing.T) {
	t.Parallel()

	var (
		obj = &model.Post{
			ID:              1,
			Title:           "Test",
			Content:         "some content",
			AuthorID:        1,
			CommentsAllowed: true,
		}
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
		wantRes = []*model.Comment{{
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
		wantRes []*model.Comment
		appErr  error
		wantErr error
	}{
		"success":        {appRes: appRes, wantRes: wantRes, appErr: nil, wantErr: nil},
		"a.app.Comments": {appRes: nil, wantRes: nil, appErr: errAny, wantErr: errAny},
	}

	for name, tc := range testcases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockApp, c, assert := start(t)

			mockApp.EXPECT().Comments(gomock.Any(), obj.ID, 10, nil).Return(tc.appRes, tc.appErr)

			res, err := c.Post().Comments(context.Background(), obj, 10, nil)

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
		obj             = &model.Comment{
			ID:              1,
			PostID:          1,
			Content:         "some content",
			AuthorID:        1,
			ParentCommentID: nil,
		}
		appRes = []app.Comment{
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
		wantRes = []*model.Comment{{
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
		wantRes []*model.Comment
		appErr  error
		wantErr error
	}{
		"success":             {appRes: appRes, wantRes: wantRes, appErr: nil, wantErr: nil},
		"a.app.ChildComments": {appRes: nil, wantRes: nil, appErr: errAny, wantErr: errAny},
	}

	for name, tc := range testcases {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mockApp, c, assert := start(t)

			mockApp.EXPECT().ChildComments(gomock.Any(), obj.PostID, &obj.ID, 10, nil).Return(tc.appRes, tc.appErr)

			res, err := c.Comment().ChildComments(context.Background(), obj, 10, nil)

			assert.ErrorIs(err, tc.wantErr)
			assert.Equal(tc.wantRes, res)
		})
	}
}

//func TestResolvers_CommentAdded(t *testing.T) {
//	t.Parallel()
//
//	var (
//		appRes = &app.Comment{
//			ID:              1,
//			PostID:          1,
//			Content:         "some content",
//			AuthorID:        1,
//			ParentCommentID: nil,
//		}
//
//		resres = make(chan app.Comment)
//
//
//		wantRes = &model.Comment{
//			ID:              1,
//			PostID:          1,
//			Content:         "some content",
//			AuthorID:        1,
//			ParentCommentID: nil,
//		}
//	)
//
//	testcases := map[string]struct {
//		appRes  app.Comment
//		wantRes *model.Comment
//		appErr  error
//		wantErr error
//	}{
//		"success":             {appRes: appRes, wantRes: wantRes, appErr: nil, wantErr: nil},
//		"a.app.Subscriptions": {appRes: nil, wantRes: nil, appErr: errAny, wantErr: errAny},
//	}
//
//	for name, tc := range testcases {
//		name, tc := name, tc
//		t.Run(name, func(t *testing.T) {
//			t.Parallel()
//
//			mockApp, c, assert := start(t)
//
//			mockApp.EXPECT().Subscriptions(gomock.Any(), 1, 1)
//
//			res, err := c.Subscription().CommentAdded(context.Background(), 1, 1)
//
//			assert.ErrorIs(tc.wantErr, err)
//			assert.Equal(tc.wantRes, res)
//		})
//	}
//}
