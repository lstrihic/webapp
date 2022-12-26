package db

import (
	"context"
	"fmt"
	"github.com/lstrihic/webapp/adapter/db/entity"
	"github.com/lstrihic/webapp/pkg/config"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"go.uber.org/fx"
)

var Provider = fx.Provide(InitDB)

type DB interface {
	Client() *entity.Client
	Migrate(ctx context.Context) error
}

type db struct {
	client    *entity.Client
	lifecycle fx.Lifecycle
}

// InitDB initialize db.
func InitDB(cfg *config.Config, lifecycle fx.Lifecycle) (DB, error) {
	client, err := entity.Open("sqlite3", fmt.Sprintf("file:%s?cache=shared&_fk=1", cfg.DB.File))
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database connection")
	}
	db := &db{
		client:    client,
		lifecycle: lifecycle,
	}
	db.registerShutdownHook()
	return db, nil
}

// Client return ent client.
func (database *db) Client() *entity.Client {
	return database.client
}

// Migrate create db schema.
func (database *db) Migrate(ctx context.Context) error {
	if err := database.client.Schema.Create(ctx); err != nil {
		return errors.Wrap(err, "failed to create db schema")
	}
	return nil
}

// registerShutdownHook add shutdown hook.
func (database *db) registerShutdownHook() {
	database.lifecycle.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			return database.client.Close()
		},
	})
}
