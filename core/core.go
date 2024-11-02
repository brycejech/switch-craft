package core

import (
	"context"
	"switchcraft/types"
)

func NewCore(repo Repo, accountRepo AccountRepo) *Core {
	return &Core{
		repository:  repo,
		accountRepo: accountRepo,
	}
}

type Core struct {
	repository  Repo
	accountRepo AccountRepo
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
