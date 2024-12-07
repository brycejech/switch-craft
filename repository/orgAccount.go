package repository

import (
	"context"
	"fmt"
	"switchcraft/repository/queries"
	"switchcraft/types"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewOrgAccountRepository(logger *types.Logger, db *pgxpool.Pool) *orgAccountRepo {
	return &orgAccountRepo{
		logger: logger,
		db:     db,
	}
}

type orgAccountRepo struct {
	logger *types.Logger
	db     *pgxpool.Pool
}

func (r *orgAccountRepo) Create(ctx context.Context,
	orgID int64,
	firstName string,
	lastName string,
	email string,
	username string,
	password *string,
	createdBy int64,
) (*types.Account, error) {

	var (
		account types.Account
		rows    pgx.Rows
		err     error
	)

	if rows, err = r.db.Query(
		ctx,
		queries.OrgAccountCreate,
		orgID,
		firstName,
		lastName,
		email,
		username,
		password,
		createdBy,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if account, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Account]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &account, nil
}

func (r *orgAccountRepo) GetMany(ctx context.Context, orgID int64) ([]types.Account, error) {

	var (
		accounts []types.Account
		rows     pgx.Rows
		err      error
	)

	if rows, err = r.db.Query(ctx, queries.OrgAccountGetMany, orgID); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if accounts, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Account]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return accounts, nil
}

func (r *orgAccountRepo) GetOne(ctx context.Context,
	orgID int64,
	id *int64,
	uuid *string,
	username *string,
) (*types.Account, error) {

	var (
		account types.Account
		rows    pgx.Rows
		err     error
	)

	if rows, err = r.db.Query(ctx,
		queries.OrgAccountGetOne,
		orgID,
		id,
		uuid,
		username,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if account, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Account]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &account, nil
}

func (r *orgAccountRepo) Update(ctx context.Context,
	orgID int64,
	id int64,
	firstName string,
	lastName string,
	email string,
	username string,
	modifiedBy int64,
) (*types.Account, error) {

	var (
		account types.Account
		rows    pgx.Rows
		err     error
	)

	if rows, err = r.db.Query(ctx,
		queries.OrgAccountUpdate,
		orgID,
		id,
		firstName,
		lastName,
		email,
		username,
		modifiedBy,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if account, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Account]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &account, nil
}

func (r *orgAccountRepo) Delete(ctx context.Context, orgID int64, id int64) error {
	row := r.db.QueryRow(ctx, queries.OrgAccountDelete, orgID, id)

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

func (r *orgAccountRepo) GetByUsername(ctx context.Context, username string) (*types.Account, error) {

	var (
		account types.Account
		rows    pgx.Rows
		err     error
	)

	if rows, err = r.db.Query(ctx,
		queries.AccountGetByUsername,
		username,
	); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if account, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Account]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return &account, nil
}
