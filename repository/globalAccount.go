package repository

import (
	"context"
	"fmt"
	"switchcraft/repository/queries"
	"switchcraft/types"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type globalAccountRepo struct {
	logger *types.Logger
	db     *pgxpool.Pool
}

func NewGlobalAccountRepository(logger *types.Logger, db *pgxpool.Pool) *globalAccountRepo {
	return &globalAccountRepo{
		logger: logger,
		db:     db,
	}
}

func (r *globalAccountRepo) Create(ctx context.Context,
	isInstanceAdmin bool,
	firstName string,
	lastName string,
	email string,
	username string,
	password string,
	createdBy int64,
) (*types.Account, error) {

	var (
		account types.Account
		rows    pgx.Rows
		err     error
	)

	if rows, err = r.db.Query(
		ctx,
		queries.GlobalAccountCreate,
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

func (r *globalAccountRepo) GetMany(ctx context.Context) ([]types.Account, error) {

	var (
		accounts []types.Account
		rows     pgx.Rows
		err      error
	)

	if rows, err = r.db.Query(ctx, queries.GlobalAccountGetMany); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	if accounts, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Account]); err != nil {
		return nil, handleError(ctx, r.logger, err)
	}

	return accounts, nil
}

func (r *globalAccountRepo) GetOne(ctx context.Context,
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
		queries.GlobalAccountGetOne,
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

func (r *globalAccountRepo) Update(ctx context.Context,
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
		queries.GlobalAccountUpdate,
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

func (r *globalAccountRepo) Delete(ctx context.Context, id int64) error {
	row := r.db.QueryRow(ctx, queries.GlobalAccountDelete, id)

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

func (r *globalAccountRepo) GetByUsername(ctx context.Context, username string) (*types.Account, error) {

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
