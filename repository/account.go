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

func NewAccountRepository(logger *types.Logger, db *pgxpool.Pool) *accountRepo {
	return &accountRepo{
		logger: logger,
		db:     db,
	}
}

type accountRepo struct {
	logger *types.Logger
	db     *pgxpool.Pool
}

func (r *accountRepo) Create(ctx context.Context,
	orgID *int64,
	isInstanceAdmin bool,
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
		queries.AccountCreate,
		orgID,
		isInstanceAdmin,
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

func (r *accountRepo) GetMany(ctx context.Context, orgID *int64) ([]types.Account, error) {

	var (
		accounts []types.Account
		rows     pgx.Rows
		err      error
	)

	if rows, err = r.db.Query(ctx, queries.AccountGetMany, orgID); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if accounts, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Account]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return accounts, nil
}

func (r *accountRepo) GetOne(ctx context.Context,
	orgID *int64,
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
		queries.AccountGetOne,
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

func (r *accountRepo) Update(ctx context.Context,
	orgID *int64,
	id int64,
	isInstanceAdmin bool,
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
		queries.AccountUpdate,
		orgID,
		id,
		isInstanceAdmin,
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

func (r *accountRepo) Delete(ctx context.Context, orgID *int64, id int64) error {
	row := r.db.QueryRow(ctx, queries.AccountDelete, orgID, id)

	var numDeleted int64
	if err := row.Scan(&numDeleted); err != nil {
		return handleError(ctx, r.logger, err)
	}

	if numDeleted < 1 {
		return errors.New("no rows deleted")
	}

	if numDeleted > 1 {
		return errors.New(fmt.Sprintf("expected to delete 1 row, deleted %v", numDeleted))
	}

	return nil
}
