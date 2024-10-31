package repository

import (
	"context"
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
	firstName string,
	lastName string,
	email string,
	username string,
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
		firstName,
		lastName,
		email,
		username,
		createdBy,
	); err != nil {
		return nil, handleError(err)
	}

	if account, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Account]); err != nil {
		return nil, handleError(err)
	}

	return &account, nil
}

func (r *accountRepository) GetMany(ctx context.Context) ([]types.Account, error) {

	var (
		accounts []types.Account
		rows     pgx.Rows
		err      error
	)

	if rows, err = r.db.Query(ctx, queries.AccountGetMany); err != nil {
		return nil, handleError(err)
	}

	if accounts, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Account]); err != nil {
		return nil, handleError(err)
	}

	return accounts, nil
}

func (r *accountRepository) GetOne(ctx context.Context,
	id *int64,
	uuid *string,
	username *string,
) (*types.Account, error) {

	var (
		account types.Account
		rows    pgx.Rows
		err     error
	)

	if id == nil {
		fmt.Println("id is null")
	} else {
		fmt.Println("id is", *id)
	}

	if rows, err = r.db.Query(ctx,
		queries.AccountGetOne,
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

func (r *accountRepository) Delete(ctx context.Context, id int64) error {
	if _, err := r.db.Query(ctx, queries.AccountDelete, id); err != nil {
		return handleError(err)
	}
	return nil
}
