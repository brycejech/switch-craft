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
	globalAccountRepo GlobalAccountRepo,
	orgAccountRepo OrgAccountRepo,
	orgGroupRepo OrgGroupRepo,
	orgRepo OrgRepo,
	appRepo AppRepo,
	featureFlagRepo FeatureFlagRepo,
	jwtSigningKey []byte,
) *Core {
	return &Core{
		logger:            logger,
		repository:        repo,
		globalAccountRepo: globalAccountRepo,
		orgAccountRepo:    orgAccountRepo,
		orgGroupRepo:      orgGroupRepo,
		orgRepo:           orgRepo,
		appRepo:           appRepo,
		featureFlagRepo:   featureFlagRepo,
		jwtSigningKey:     jwtSigningKey,
	}
}

type Core struct {
	logger            *types.Logger
	repository        Repo
	globalAccountRepo GlobalAccountRepo
	orgAccountRepo    OrgAccountRepo
	orgGroupRepo      OrgGroupRepo
	orgRepo           OrgRepo
	appRepo           AppRepo
	featureFlagRepo   FeatureFlagRepo
	jwtSigningKey     []byte
}

func (c *Core) getOperationTracer(ctx context.Context) (types.OperationTracer, error) {
	tracer, ok := ctx.Value(types.CtxOperationTracer).(types.OperationTracer)
	if !ok {
		return tracer, errors.New("invalid operation context")
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

type GlobalAccountRepo interface {
	Create(ctx context.Context,
		isInstanceAdmin bool,
		firstName string,
		lastName string,
		email string,
		username string,
		password string,
		createdBy int64,
	) (*types.Account, error)
	GetMany(ctx context.Context) ([]types.Account, error)
	GetOne(ctx context.Context,
		id *int64,
		uuid *string,
		username *string,
	) (*types.Account, error)
	Update(ctx context.Context,
		id int64,
		isInstanceAdmin bool,
		firstName string,
		lastName string,
		email string,
		username string,
		modifiedBy int64,
	) (*types.Account, error)
	Delete(ctx context.Context, id int64) error
	GetByUsername(ctx context.Context, username string) (*types.Account, error)
}

type OrgAccountRepo interface {
	Create(ctx context.Context,
		orgID int64,
		firstName string,
		lastName string,
		email string,
		username string,
		password *string,
		createdBy int64,
	) (*types.Account, error)
	GetMany(ctx context.Context, orgID int64) ([]types.Account, error)
	GetOne(ctx context.Context,
		orgID int64,
		id *int64,
		uuid *string,
		username *string,
	) (*types.Account, error)
	Update(ctx context.Context,
		orgID int64,
		id int64,
		firstName string,
		lastName string,
		email string,
		username string,
		modifiedBy int64,
	) (*types.Account, error)
	Delete(ctx context.Context, orgID int64, id int64) error
}

type OrgGroupRepo interface {
	Create(ctx context.Context,
		orgID int64,
		name string,
		description string,
		createdBy int64,
	) (*types.OrgGroup, error)
	GetMany(ctx context.Context, orgID int64) ([]types.OrgGroup, error)
	GetOne(ctx context.Context,
		orgID int64,
		id *int64,
		uuid *string,
	) (*types.OrgGroup, error)
	Update(ctx context.Context,
		orgID int64,
		id int64,
		name string,
		description string,
		modifiedBy int64,
	) (*types.OrgGroup, error)
	Delete(ctx context.Context,
		orgID int64,
		id int64,
	) error
}

type OrgRepo interface {
	Create(ctx context.Context,
		name string,
		slug string,
		owner int64,
		createdBy int64,
	) (*types.Organization, error)
	GetMany(ctx context.Context) ([]types.Organization, error)
	GetOne(ctx context.Context,
		id *int64,
		uuid *string,
		slug *string,
	) (*types.Organization, error)
	Update(ctx context.Context,
		id int64,
		name string,
		slug string,
		owner int64,
		modifiedBy int64,
	) (*types.Organization, error)
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
		label string,
		description string,
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
		name string,
		label string,
		description string,
		isEnabled bool,
		modifiedBy int64,
	) (*types.FeatureFlag, error)
	Delete(ctx context.Context,
		orgID int64,
		applicationID int64,
		id int64,
	) error
}
