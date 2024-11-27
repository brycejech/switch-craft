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

func NewTenantRepository(db *pgxpool.Pool) *tenantRepo {
	return &tenantRepo{
		db: db,
	}
}

type tenantRepo struct {
	db *pgxpool.Pool
}

func (r *tenantRepo) Create(ctx context.Context,
	name string,
	slug string,
	owner int64,
	createdBy int64,
) (*types.Tenant, error) {
	var (
		tenant types.Tenant
		rows   pgx.Rows
		err    error
	)

	if rows, err = r.db.Query(ctx,
		queries.TenantCreate,
		name,
		slug,
		owner,
		createdBy,
	); err != nil {
		return nil, handleError(err)
	}

	if tenant, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Tenant]); err != nil {
		return nil, handleError(err)
	}

	return &tenant, nil
}

func (r *tenantRepo) GetMany(ctx context.Context) ([]types.Tenant, error) {
	var (
		tenants []types.Tenant
		rows    pgx.Rows
		err     error
	)

	if rows, err = r.db.Query(ctx, queries.TenantGetMany); err != nil {
		return nil, handleError(err)
	}

	if tenants, err = pgx.CollectRows(rows, pgx.RowToStructByName[types.Tenant]); err != nil {
		return nil, handleError(err)
	}

	return tenants, nil
}

func (r *tenantRepo) GetOne(ctx context.Context,
	id *int64,
	uuid *string,
	slug *string,
) (*types.Tenant, error) {
	if id == nil && uuid == nil && slug == nil {
		return nil, errors.New("tenantRepository.GetOne must provide at least one of id, uuid, or slug")
	}

	var (
		tenant types.Tenant
		rows   pgx.Rows
		err    error
	)

	if rows, err = r.db.Query(ctx,
		queries.TenantGetOne,
		id,
		uuid,
		slug,
	); err != nil {
		return nil, handleError(err)
	}

	if tenant, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Tenant]); err != nil {
		return nil, handleError(err)
	}

	return &tenant, nil
}

func (r *tenantRepo) Update(ctx context.Context,
	id int64,
	name string,
	slug string,
	owner int64,
	modifiedBy int64,
) (*types.Tenant, error) {
	var (
		tenant types.Tenant
		rows   pgx.Rows
		err    error
	)

	if rows, err = r.db.Query(ctx,
		queries.TenantUpdate,
		id,
		name,
		slug,
		owner,
		modifiedBy,
	); err != nil {
		return nil, handleError(err)
	}

	if tenant, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[types.Tenant]); err != nil {
		return nil, handleError(err)
	}

	return &tenant, nil
}

func (r *tenantRepo) Delete(ctx context.Context, id int64) error {
	row := r.db.QueryRow(ctx, queries.TenantDelete, id)

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
