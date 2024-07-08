package graph_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"ozon_test_compost/cmd/compost/internal/api/graph"
	"testing"
)

var (
	errAny = errors.New("any err")
)

func start(t *testing.T) (*Mockapplication, *graph.Resolver, *require.Assertions) {
	ctrl := gomock.NewController(t)
	mockApp := NewMockapplication(ctrl)
	assert := require.New(t)

	return mockApp, graph.NewResolver(mockApp), assert
}
