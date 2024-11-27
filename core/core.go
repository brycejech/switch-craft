package core

import (
	"context"
	"errors"
	"regexp"
	"switchcraft/types"
)

func NewCore(
	logger *types.Logger,
	repo Repo,
	accountRepo AccountRepo,
	orgRepo OrgRepo,
	appRepo AppRepo,
	featureFlagRepo FeatureFlagRepo,
	jwtSigningKey []byte,
) *Core {
	return &Core{
		logger:          logger,
		repository:      repo,
		accountRepo:     accountRepo,
		orgRepo:         orgRepo,
		appRepo:         appRepo,
		featureFlagRepo: featureFlagRepo,
		jwtSigningKey:   jwtSigningKey,
	}
}

type Core struct {
	logger          *types.Logger
	repository      Repo
	accountRepo     AccountRepo
	orgRepo         OrgRepo
	appRepo         AppRepo
	featureFlagRepo FeatureFlagRepo
	jwtSigningKey   []byte
}

func (c *Core) getOperationTracer(ctx context.Context) (types.OperationTracer, error) {
	tracer, ok := ctx.Value(types.CtxOperationTracer).(types.OperationTracer)
	if !ok {
		return tracer, errors.New("Invalid operation context")
	}
	return tracer, nil
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
		orgID *int64,
		firstName string,
		lastName string,
		email string,
		username string,
		password *string,
		createdBy int64,
	) (*types.Account, error)
	GetMany(ctx context.Context, orgID *int64) ([]types.Account, error)
	GetOne(ctx context.Context,
		orgID *int64,
		id *int64,
		uuid *string,
		username *string,
	) (*types.Account, error)
	Update(ctx context.Context,
		orgID *int64,
		id int64,
		firstName string,
		lastName string,
		email string,
		username string,
		modifiedBy int64,
	) (*types.Account, error)
	Delete(ctx context.Context, orgID *int64, id int64) error
}

type OrgRepo interface {
	Create(ctx context.Context,
		name string,
		slug string,
		owner int64,
		createdBy int64,
	) (*types.Org, error)
	GetMany(ctx context.Context) ([]types.Org, error)
	GetOne(ctx context.Context,
		id *int64,
		uuid *string,
		slug *string,
	) (*types.Org, error)
	Update(ctx context.Context,
		id int64,
		name string,
		slug string,
		owner int64,
		modifiedBy int64,
	) (*types.Org, error)
	Delete(ctx context.Context, id int64) error
}

type AppRepo interface {
	Create(ctx context.Context,
		orgID int64,
		name string,
		slug string,
		createdBy int64,
	) (*types.Application, error)
	GetMany(ctx context.Context,
		orgID int64,
	) ([]types.Application, error)
	GetOne(ctx context.Context,
		orgID int64,
		id *int64,
		uuid *string,
		slug *string,
	) (*types.Application, error)
	Update(ctx context.Context,
		orgID int64,
		id int64,
		name string,
		slug string,
		modifiedBy int64,
	) (*types.Application, error)
	Delete(ctx context.Context,
		orgID int64,
		id int64,
	) error
}

type FeatureFlagRepo interface {
	Create(ctx context.Context,
		orgID int64,
		applicationID int64,
		name string,
		isEnabled bool,
		createdBy int64,
	) (*types.FeatureFlag, error)
	GetMany(ctx context.Context,
		orgID int64,
		applicationID int64,
	) ([]types.FeatureFlag, error)
	GetOne(ctx context.Context,
		orgID int64,
		applicationID int64,
		id *int64,
		uuid *string,
		name *string,
	) (*types.FeatureFlag, error)
	Update(ctx context.Context,
		orgID int64,
		applicationID int64,
		id int64,
		slug string,
		isEnabled bool,
		modifiedBy int64,
	) (*types.FeatureFlag, error)
	Delete(ctx context.Context,
		orgID int64,
		applicationID int64,
		id int64,
	) error
}
