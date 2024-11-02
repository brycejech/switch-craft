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

func NewAccountRepository(pool *pgxpool.Pool) *accountRepository {
	return &accountRepository{
		db: pool,
	}
}

type accountRepository struct {
	db *pgxpool.Pool
}

func (r *accountRepository) Create(ctx context.Context,
	tenantID *int64,
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
		tenantID,
		firstName,
		lastName,
		email,
		username,
		password,
		createdBy,
	); err != nil {
		return nil, handleError(err)
	}

	if account, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Account]); err != nil {
		return nil, handleError(err)
	}

	return &account, nil
}

func (r *accountRepository) GetMany(ctx context.Context, tenantID *int64) ([]types.Account, error) {

	var (
		accounts []types.Account
		rows     pgx.Rows
		err      error
	)

	if rows, err = r.db.Query(ctx, queries.AccountGetMany, tenantID); err != nil {
		return nil, handleError(err)
	}

	if accounts, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Account]); err != nil {
		return nil, handleError(err)
	}

	return accounts, nil
}

func (r *accountRepository) GetOne(ctx context.Context,
	tenantID *int64,
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
		tenantID,
		id,
		uuid,
		username,
	); err != nil {
		return nil, handleError(err)
	}

	if account, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Account]); err != nil {
		return nil, handleError(err)
	}

	return &account, nil
}

func (r *accountRepository) Update(ctx context.Context,
	tenantID *int64,
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
		queries.AccountUpdate,
		tenantID,
		id,
		firstName,
		lastName,
		email,
		username,
		modifiedBy,
	); err != nil {
		return nil, handleError(err)
	}

	if account, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Account]); err != nil {
		return nil, handleError(err)
	}

	return &account, nil
}

func (r *accountRepository) Delete(ctx context.Context, tenantID *int64, id int64) error {
	row := r.db.QueryRow(ctx, queries.AccountDelete, tenantID, id)

	var numDeleted int64
	if err := row.Scan(&numDeleted); err != nil {
		return handleError(err)
	}

	if numDeleted < 1 {
		return errors.New("no rows deleted")
	}

	if numDeleted > 1 {
		return errors.New(fmt.Sprintf("expected to delete 1 row, deleted %v", numDeleted))
	}

	return nil
}
