// Code generated by MockGen. DO NOT EDIT.
// Source: resolver.go

// Package graph_test is a generated GoMock package.
package graph_test

import (
	context "context"
	app "ozon_test_compost/cmd/compost/internal/app"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockapplication is a mock of application interface.
type Mockapplication struct {
	ctrl     *gomock.Controller
	recorder *MockapplicationMockRecorder
}

// MockapplicationMockRecorder is the mock recorder for Mockapplication.
type MockapplicationMockRecorder struct {
	mock *Mockapplication
}

// NewMockapplication creates a new mock instance.
func NewMockapplication(ctrl *gomock.Controller) *Mockapplication {
	mock := &Mockapplication{ctrl: ctrl}
	mock.recorder = &MockapplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockapplication) EXPECT() *MockapplicationMockRecorder {
	return m.recorder
}

// ChildComments mocks base method.
func (m *Mockapplication) ChildComments(ctx context.Context, postID int, parentID *int, limit int, offset *int) ([]app.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChildComments", ctx, postID, parentID, limit, offset)
	ret0, _ := ret[0].([]app.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChildComments indicates an expected call of ChildComments.
func (mr *MockapplicationMockRecorder) ChildComments(ctx, postID, parentID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChildComments", reflect.TypeOf((*Mockapplication)(nil).ChildComments), ctx, postID, parentID, limit, offset)
}

// Comments mocks base method.
func (m *Mockapplication) Comments(ctx context.Context, ID, limit int, offset *int) ([]app.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Comments", ctx, ID, limit, offset)
	ret0, _ := ret[0].([]app.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Comments indicates an expected call of Comments.
func (mr *MockapplicationMockRecorder) Comments(ctx, ID, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Comments", reflect.TypeOf((*Mockapplication)(nil).Comments), ctx, ID, limit, offset)
}

// CreateComment mocks base method.
func (m *Mockapplication) CreateComment(ctx context.Context, postID int, content string, authorID int, parentCommentID *int) (*app.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", ctx, postID, content, authorID, parentCommentID)
	ret0, _ := ret[0].(*app.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockapplicationMockRecorder) CreateComment(ctx, postID, content, authorID, parentCommentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*Mockapplication)(nil).CreateComment), ctx, postID, content, authorID, parentCommentID)
}

// CreatePost mocks base method.
func (m *Mockapplication) CreatePost(ctx context.Context, title, content string, authorID int, commentAllowed bool) (*app.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", ctx, title, content, authorID, commentAllowed)
	ret0, _ := ret[0].(*app.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockapplicationMockRecorder) CreatePost(ctx, title, content, authorID, commentAllowed interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*Mockapplication)(nil).CreatePost), ctx, title, content, authorID, commentAllowed)
}

// GetPost mocks base method.
func (m *Mockapplication) GetPost(ctx context.Context, id int) (*app.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", ctx, id)
	ret0, _ := ret[0].(*app.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockapplicationMockRecorder) GetPost(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*Mockapplication)(nil).GetPost), ctx, id)
}

// GetPosts mocks base method.
func (m *Mockapplication) GetPosts(ctx context.Context) ([]app.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPosts", ctx)
	ret0, _ := ret[0].([]app.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPosts indicates an expected call of GetPosts.
func (mr *MockapplicationMockRecorder) GetPosts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosts", reflect.TypeOf((*Mockapplication)(nil).GetPosts), ctx)
}

// Subscriptions mocks base method.
func (m *Mockapplication) Subscriptions(ctx context.Context, postID, userID int) (<-chan app.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscriptions", ctx, postID, userID)
	ret0, _ := ret[0].(<-chan app.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscriptions indicates an expected call of Subscriptions.
func (mr *MockapplicationMockRecorder) Subscriptions(ctx, postID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscriptions", reflect.TypeOf((*Mockapplication)(nil).Subscriptions), ctx, postID, userID)
}
