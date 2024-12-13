package repository

import (
	"context"
	"fmt"
	"switchcraft/repository/queries"
	"switchcraft/types"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewOrgGroupRepository(logger *types.Logger, db *pgxpool.Pool) *orgGroupRepo {
	return &orgGroupRepo{
		logger: logger,
		db:     db,
	}
}

type orgGroupRepo struct {
	logger *types.Logger
	db     *pgxpool.Pool
}

func (r *orgGroupRepo) Create(ctx context.Context,
	orgID int64,
	name string,
	description string,
	createdBy int64,
) (*types.OrgGroup, error) {

	var (
		group types.OrgGroup
		rows  pgx.Rows
		err   error
	)

	if rows, err = r.db.Query(
		ctx,
		queries.OrgGroupCreate,
		orgID,
		name,
		description,
		createdBy,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if group, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.OrgGroup]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &group, nil
}

func (r *orgGroupRepo) GetMany(ctx context.Context, orgID int64) ([]types.OrgGroup, error) {

	var (
		groups []types.OrgGroup
		rows   pgx.Rows
		err    error
	)

	if rows, err = r.db.Query(ctx, queries.OrgGroupGetMany, orgID); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if groups, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.OrgGroup]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return groups, nil
}

func (r *orgGroupRepo) GetOne(ctx context.Context,
	orgID int64,
	id *int64,
	uuid *string,
) (*types.OrgGroup, error) {

	var (
		group types.OrgGroup
		rows  pgx.Rows
		err   error
	)

	if rows, err = r.db.Query(ctx,
		queries.OrgGroupGetOne,
		orgID,
		id,
		uuid,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if group, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.OrgGroup]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &group, nil
}

func (r *orgGroupRepo) Update(ctx context.Context,
	orgID int64,
	id int64,
	name string,
	description string,
	modifiedBy int64,
) (*types.OrgGroup, error) {

	var (
		group types.OrgGroup
		rows  pgx.Rows
		err   error
	)

	if rows, err = r.db.Query(ctx,
		queries.OrgGroupUpdate,
		orgID,
		id,
		name,
		description,
		modifiedBy,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if group, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.OrgGroup]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &group, nil
}

func (r *orgGroupRepo) Delete(ctx context.Context,
	orgID int64,
	id int64,
) error {
	row := r.db.QueryRow(ctx, queries.OrgGroupDelete, orgID, id)

	var numDeleted int64
	if err := row.Scan(&numDeleted); err != nil {
		return handleError(ctx, r.logger, err)
	}

	if numDeleted < 1 {
		return types.ErrNotFound
	}

	if numDeleted > 1 {
		return fmt.Errorf("expected to delete 1 row, deleted %v", numDeleted)
	}

	return nil
}
