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

func NewOrgRepository(logger *types.Logger, db *pgxpool.Pool) *orgRepo {
	return &orgRepo{
		logger: logger,
		db:     db,
	}
}

type orgRepo struct {
	logger *types.Logger
	db     *pgxpool.Pool
}

func (r *orgRepo) Create(ctx context.Context,
	name string,
	slug string,
	owner int64,
	createdBy int64,
) (*types.Org, error) {
	var (
		org  types.Org
		rows pgx.Rows
		err  error
	)

	if rows, err = r.db.Query(ctx,
		queries.OrgCreate,
		name,
		slug,
		owner,
		createdBy,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if org, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Org]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &org, nil
}

func (r *orgRepo) GetMany(ctx context.Context) ([]types.Org, error) {
	var (
		orgs []types.Org
		rows pgx.Rows
		err  error
	)

	if rows, err = r.db.Query(ctx, queries.OrgGetMany); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if orgs, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Org]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return orgs, nil
}

func (r *orgRepo) GetOne(ctx context.Context,
	id *int64,
	uuid *string,
	slug *string,
) (*types.Org, error) {
	if id == nil && uuid == nil && slug == nil {
		return nil, errors.New("orgRepository.GetOne must provide at least one of id, uuid, or slug")
	}

	var (
		org  types.Org
		rows pgx.Rows
		err  error
	)

	if rows, err = r.db.Query(ctx,
		queries.OrgGetOne,
		id,
		uuid,
		slug,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if org, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Org]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &org, nil
}

func (r *orgRepo) Update(ctx context.Context,
	id int64,
	name string,
	slug string,
	owner int64,
	modifiedBy int64,
) (*types.Org, error) {
	var (
		org  types.Org
		rows pgx.Rows
		err  error
	)

	if rows, err = r.db.Query(ctx,
		queries.OrgUpdate,
		id,
		name,
		slug,
		owner,
		modifiedBy,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if org, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Org]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &org, nil
}

func (r *orgRepo) Delete(ctx context.Context, id int64) error {
	row := r.db.QueryRow(ctx, queries.OrgDelete, id)

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
