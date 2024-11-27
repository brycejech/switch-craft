package repository

import (
	"context"
	"errors"
	"fmt"
	"switchcraft/repository/queries"
	"switchcraft/types"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewAppRepository(logger *types.Logger, db *pgxpool.Pool) *appRepo {
	return &appRepo{
		logger: logger,
		db:     db,
	}
}

type appRepo struct {
	logger *types.Logger
	db     *pgxpool.Pool
}

func (r *appRepo) Create(ctx context.Context,
	orgID int64,
	name string,
	slug string,
	createdBy int64,
) (*types.Application, error) {
	var (
		application types.Application
		rows        pgx.Rows
		err         error
	)

	if rows, err = r.db.Query(ctx,
		queries.AppCreate,
		orgID,
		name,
		slug,
		createdBy,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if application, err = pgx.CollectOneRow(
		rows,
		pgx.RowToStructByName[types.Application],
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &application, nil
}

func (r *appRepo) GetMany(ctx context.Context,
	orgID int64,
) ([]types.Application, error) {
	var (
		applications []types.Application
		rows         pgx.Rows
		err          error
	)

	if rows, err = r.db.Query(ctx, queries.AppGetMany, orgID); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if applications, err = pgx.CollectRows(
		rows,
		pgx.RowToStructByName[types.Application],
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return applications, nil
}

func (r *appRepo) GetOne(ctx context.Context,
	orgID int64,
	id *int64,
	uuid *string,
	slug *string,
) (*types.Application, error) {
	var (
		application types.Application
		rows        pgx.Rows
		err         error
	)

	if rows, err = r.db.Query(ctx,
		queries.AppGetOne,
		orgID,
		id,
		uuid,
		slug,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if application, err = pgx.CollectOneRow(
		rows,
		pgx.RowToStructByName[types.Application],
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &application, nil
}

func (r *appRepo) Update(ctx context.Context,
	orgID int64,
	id int64,
	name string,
	slug string,
	modifiedBy int64,
) (*types.Application, error) {
	var (
		application types.Application
		rows        pgx.Rows
		err         error
	)

	if rows, err = r.db.Query(ctx,
		queries.AppUpdate,
		orgID,
		id,
		name,
		slug,
		modifiedBy,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if application, err = pgx.CollectOneRow(
		rows,
		pgx.RowToStructByName[types.Application],
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &application, nil
}

func (r *appRepo) Delete(ctx context.Context,
	orgID int64,
	id int64,
) error {
	row := r.db.QueryRow(ctx,
		queries.AppDelete,
		orgID,
		id,
	)

	var numDeleted int64
	if err := row.Scan(&numDeleted); err != nil {
		return handleError(ctx, r.logger, err)
	}

	if numDeleted < 1 {
		return errors.New("no rows deleted")
	}

	if numDeleted > 1 {
		return errors.New(
			fmt.Sprintf("expected to delete 1 row, deleted %v", numDeleted),
		)
	}

	return nil
}
