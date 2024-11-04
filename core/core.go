package core

import (
	"context"
	"switchcraft/types"
)

func NewCore(repo Repo, accountRepo AccountRepo, tenantRepo TenantRepo) *Core {
	return &Core{
		repository:  repo,
		accountRepo: accountRepo,
		tenantRepo:  tenantRepo,
	}
}

type Core struct {
	repository  Repo
	accountRepo AccountRepo
	tenantRepo  TenantRepo
}

func (c *Core) MigrateUp() error {
	return c.repository.MigrateUp()
}

func (c *Core) MigrateDown() error {
	return c.repository.MigrateDown()
}

/* === PORTS === */

type Repo interface {
	MigrateUp() error
	MigrateDown() error
}

type AccountRepo interface {
	Create(ctx context.Context,
		tenantID *int64,
		firstName string,
		lastName string,
		email string,
		username string,
		password *string,
		createdBy int64,
	) (*types.Account, error)
	GetMany(ctx context.Context, tenantID *int64) ([]types.Account, error)
	GetOne(ctx context.Context,
		tenantID *int64,
		id *int64,
		uuid *string,
		username *string,
	) (*types.Account, error)
	Update(ctx context.Context,
		tenantID *int64,
		id int64,
		firstName string,
		lastName string,
		email string,
		username string,
		modifiedBy int64,
	) (*types.Account, error)
	Delete(ctx context.Context, tenantID *int64, id int64) error
}

type TenantRepo interface {
	Create(ctx context.Context,
		name string,
		slug string,
		owner int64,
		createdBy int64,
	) (*types.Tenant, error)
	GetMany(ctx context.Context) ([]types.Tenant, error)
	GetOne(ctx context.Context, ID int64) (*types.Tenant, error)
	Update(ctx context.Context,
		id int64,
		name string,
		slug string,
		owner int64,
		modifiedBy int64,
	) (*types.Tenant, error)
	Delete(ctx context.Context, id int64) error
}
