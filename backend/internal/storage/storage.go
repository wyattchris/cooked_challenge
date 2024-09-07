package storage

import (
	"context"

	"github.com/GenerateNU/cooked/backend/internal/types"
)

// TODO: add the necessary queries
type Storage interface {
	Ping() error
	CreateRecipe(ctx context.Context, recipe types.Recipe) (types.Recipe, error)
}
