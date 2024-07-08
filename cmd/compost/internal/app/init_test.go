package app_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"ozon_test_compost/cmd/compost/internal/app"
	"testing"
)

var (
	errAny = errors.New("any err")
)

func start(t *testing.T) (*MockRepo, *app.App, *require.Assertions) {
	ctrl := gomock.NewController(t)
	mockRepo := NewMockRepo(ctrl)
	assert := require.New(t)

	module := app.New(mockRepo)

	return mockRepo, module, assert
}
