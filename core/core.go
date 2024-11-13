package core

import (
	"context"
	"errors"
	"regexp"
	"switchcraft/types"
)

func NewCore(
	repo Repo,
	accountRepo AccountRepo,
	tenantRepo TenantRepo,
	appRepo AppRepo,
	featureFlagRepo FeatureFlagRepo,
	jwtSigningKey []byte,
) *Core {
	return &Core{
		repository:      repo,
		accountRepo:     accountRepo,
		tenantRepo:      tenantRepo,
		appRepo:         appRepo,
		featureFlagRepo: featureFlagRepo,
		jwtSigningKey:   jwtSigningKey,
	}
}

type Core struct {
	repository      Repo
	accountRepo     AccountRepo
	tenantRepo      TenantRepo
	appRepo         AppRepo
	featureFlagRepo FeatureFlagRepo
	jwtSigningKey   []byte
}

func (c *Core) MigrateUp() error {
	return c.repository.MigrateUp()
}

func (c *Core) MigrateDown() error {
	return c.repository.MigrateDown()
}

func validateSlug(slug string) error {
	exp := regexp.MustCompile(`^[a-zA-Z0-9\-_]+$`)
	if !exp.MatchString(slug) {
		return errors.New("slug must be alphanumeric with '-' or '_' only")
	}
	return nil
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

type AppRepo interface {
	Create(ctx context.Context,
		tenantID int64,
		name string,
		slug string,
		createdBy int64,
	) (*types.Application, error)
	GetMany(ctx context.Context,
		tenantID int64,
	) ([]types.Application, error)
	GetOne(ctx context.Context,
		tenantID int64,
		id *int64,
		uuid *string,
		slug *string,
	) (*types.Application, error)
	Update(ctx context.Context,
		tenantID int64,
		id int64,
		name string,
		slug string,
		modifiedBy int64,
	) (*types.Application, error)
	Delete(ctx context.Context,
		tenantID int64,
		id int64,
	) error
}

type FeatureFlagRepo interface {
	Create(ctx context.Context,
		tenantID int64,
		applicationID int64,
		name string,
		isEnabled bool,
		createdBy int64,
	) (*types.FeatureFlag, error)
	GetMany(ctx context.Context,
		tenantID int64,
		applicationID int64,
	) ([]types.FeatureFlag, error)
	GetOne(ctx context.Context,
		tenantID int64,
		applicationID int64,
		id *int64,
		uuid *string,
		name *string,
	) (*types.FeatureFlag, error)
	Update(ctx context.Context,
		tenantID int64,
		applicationID int64,
		id int64,
		slug string,
		isEnabled bool,
		modifiedBy int64,
	) (*types.FeatureFlag, error)
	Delete(ctx context.Context,
		tenantID int64,
		applicationID int64,
		id int64,
	) error
}
