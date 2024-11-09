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

func NewAppRepository(db *pgxpool.Pool) *appRepo {
	return &appRepo{
		db: db,
	}
}

type appRepo struct {
	db *pgxpool.Pool
}

func (r *appRepo) Create(ctx context.Context,
	tenantID int64,
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
		tenantID,
		name,
		slug,
		createdBy,
	); err != nil {
		return nil, handleError(err)
	}

	if application, err = pgx.CollectOneRow(
		rows,
		pgx.RowToStructByName[types.Application],
	); err != nil {
		return nil, handleError(err)
	}

	return &application, nil
}

func (r *appRepo) GetMany(ctx context.Context,
	tenantID int64,
) ([]types.Application, error) {
	var (
		applications []types.Application
		rows         pgx.Rows
		err          error
	)

	if rows, err = r.db.Query(ctx, queries.AppGetMany, tenantID); err != nil {
		return nil, handleError(err)
	}

	if applications, err = pgx.CollectRows(
		rows,
		pgx.RowToStructByName[types.Application],
	); err != nil {
		return nil, handleError(err)
	}

	return applications, nil
}

func (r *appRepo) GetOne(ctx context.Context,
	tenantID int64,
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
		tenantID,
		id,
		uuid,
		slug,
	); err != nil {
		return nil, handleError(err)
	}

	if application, err = pgx.CollectOneRow(
		rows,
		pgx.RowToStructByName[types.Application],
	); err != nil {
		return nil, handleError(err)
	}

	return &application, nil
}

func (r *appRepo) Update(ctx context.Context,
	tenantID int64,
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
		tenantID,
		id,
		name,
		slug,
		modifiedBy,
	); err != nil {
		return nil, handleError(err)
	}

	if application, err = pgx.CollectOneRow(
		rows,
		pgx.RowToStructByName[types.Application],
	); err != nil {
		return nil, handleError(err)
	}

	return &application, nil
}

func (r *appRepo) Delete(ctx context.Context,
	tenantID int64,
	id int64,
) error {
	row := r.db.QueryRow(ctx,
		queries.AppDelete,
		tenantID,
		id,
	)

	var numDeleted int64
	if err := row.Scan(&numDeleted); err != nil {
		return handleError(err)
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
