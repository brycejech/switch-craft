package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"switchcraft/repository/queries"
	"switchcraft/types"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

type Repository struct {
	logger *types.Logger
	db     *pgxpool.Pool
}

func NewRepository(logger *types.Logger, db *pgxpool.Pool) *Repository {
	return &Repository{
		logger: logger,
		db:     db,
	}
}

func (r *Repository) MigrateUp() error {
	migration, err := r.getMigration()
	if err != nil {
		return err
	}
	defer migration.Close()

	err = migration.Up()
	if err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("warning: database is fully migrated, cannot migrate up")
			return nil
		}
		return err
	}

	return nil
}

func (r *Repository) MigrateDown() error {
	migration, err := r.getMigration()
	if err != nil {
		return err
	}
	defer migration.Close()

	if _, _, err = migration.Version(); errors.Is(err, migrate.ErrNilVersion) {
		fmt.Println("warning: cannot migrate down, no migrations have been applied")
		return nil
	}

	err = migration.Steps(-1)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) getMigration() (*migrate.Migrate, error) {
	databaseDriver, err := postgres.WithInstance(stdlib.OpenDBFromPool(r.db), &postgres.Config{})
	if err != nil {
		return nil, err
	}

	sourceDriver, err := iofs.New(queries.Migrations, "migrations")
	if err != nil {
		return nil, err
	}

	migration, err := migrate.NewWithInstance("migrations", sourceDriver, "postgres", databaseDriver)
	if err != nil {
		return nil, err
	}

	return migration, nil
}

func handleError(ctx context.Context, logger *types.Logger, err error) error {
	tracer, _ := ctx.Value(types.CtxOperationTracer).(types.OperationTracer)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == "23505" {
			logger.Error(tracer, types.ErrItemExists.Error(), nil)
			return types.ErrItemExists
		} else if pgErr.Code == "23503" {
			logger.Error(tracer, types.ErrLinkedItemNotFound.Error(), nil)
			return types.ErrLinkedItemNotFound
		}
		bytes, _ := json.MarshalIndent(pgErr, "", "  ")
		fmt.Println(string(bytes))

		return err
	} else {
		if err.Error() == "no rows in result set" {
			logger.Error(tracer, types.ErrNotFound.Error(), nil)
			return types.ErrNotFound
		}

		err = fmt.Errorf("unkonwn postgres error - %w", err)
		logger.Error(tracer, err.Error(), nil)

		return err
	}
}
