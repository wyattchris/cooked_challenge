package postgres

import (
	"log/slog"

	"github.com/GenerateNU/cooked/backend/internal/settings"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/GenerateNU/cooked/backend/internal/types"

	
	"context"
)

type DB struct {
	db *sqlx.DB
}

// TODO: implement connecting
func New(settings settings.Postgres) *DB {
	slog.Info("creating new postgres db", "settings", settings, "connection", settings.Connection())
	return &DB{db: sqlx.MustConnect("postgres", settings.Connection())}
}

func (db *DB) Ping() error {
	return db.db.Ping()
}

func (db *DB) CreateRecipe(ctx context.Context, recipe types.Recipe) (types.Recipe, error) {
	if _, err := db.db.ExecContext(ctx, "INSERT INTO recipes (id, name, cook_duration, instructions, image_url, meal) VALUES ($1, $2, $3, $4, $5, $6)",
	 recipe.ID, recipe.Name, recipe.Cook, recipe.Instructions, recipe.ImageURL, recipe.Meal); err != nil {
		return types.Recipe{}, err
	}

	return recipe, nil
}

// TODO: implement the necessary queries to satisfy the storage.Storage interface
