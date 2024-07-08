package in_memory_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"ozon_test_compost/cmd/compost/internal/adapters/in_memory"
	"testing"
)

func start(t *testing.T) (context.Context, *in_memory.Repo, *require.Assertions) {
	repo := in_memory.New()

	assert := require.New(t)
	return context.Background(), repo, assert
}
