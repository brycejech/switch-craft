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

func (r *orgGroupRepo) AddAccount(ctx context.Context,
	orgID int64,
	groupID int64,
	accountID int64,
	createdBy int64,
) (*types.OrgGroupAccount, error) {
	var (
		groupAcct types.OrgGroupAccount
		rows      pgx.Rows
		err       error
	)

	if rows, err = r.db.Query(ctx,
		queries.OrgGroupAccountCreate,
		orgID,
		groupID,
		accountID,
		createdBy,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if groupAcct, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.OrgGroupAccount]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &groupAcct, nil
}

func (r *orgGroupRepo) GetAccounts(ctx context.Context,
	orgID int64,
	groupID int64,
) ([]types.Account, error) {

	var (
		accounts []types.Account
		rows     pgx.Rows
		err      error
	)

	if rows, err = r.db.Query(ctx,
		queries.OrgGroupAccountGetMany,
		orgID,
		groupID,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if accounts, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Account]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return accounts, nil
}

func (r *orgGroupRepo) UpdateAccounts(ctx context.Context,
	orgID int64,
	groupID int64,
	accountIDs []int64,
	createdBy int64,
) ([]types.Account, error) {
	if err := r.RemoveAllAccounts(ctx, orgID, groupID); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	var (
		numAccounts = len(accountIDs)
		numInserted int64
		err         error
	)
	numInserted, err = r.db.CopyFrom(ctx,
		pgx.Identifier{"account", "org_group_account"},
		[]string{"org_id", "group_id", "account_id", "created_by"},
		pgx.CopyFromSlice(numAccounts, func(i int) ([]interface{}, error) {
			return []interface{}{orgID, groupID, accountIDs[i], createdBy}, nil
		}),
	)
	if err != nil {
		return nil, handleError(ctx, r.logger, err)
	}
	if numInserted != int64(numAccounts) {
		return nil, handleError(ctx, r.logger, errors.New("orgGroup.UpdateAccounts mismatched insert and accountIDs len"))
	}

	return r.GetAccounts(ctx, orgID, groupID)
}

func (r *orgGroupRepo) RemoveAccount(ctx context.Context,
	orgID int64,
	groupID int64,
	accountID int64,
) error {

	row := r.db.QueryRow(ctx,
		queries.OrgGroupAccountDeleteOne,
		orgID,
		groupID,
		accountID,
	)

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

func (r *orgGroupRepo) RemoveAllAccounts(ctx context.Context,
	orgID int64,
	groupID int64,
) error {
	if _, err := r.db.Exec(ctx,
		queries.OrgGroupAccountDeleteAll,
		orgID,
		groupID,
	); err != nil {
		return handleError(ctx, r.logger, err)
	}

	return nil
}
