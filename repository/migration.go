package repository

import (
	"errors"
	"fmt"
	"switchcraft/repository/queries"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/stdlib"
)

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
