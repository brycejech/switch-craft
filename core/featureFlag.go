package core

import (
	"context"
	"errors"
	"switchcraft/types"
)

type featureFlagCreateArgs struct {
	tenantID  int64
	appID     int64
	name      string
	isEnabled bool
}

func (a *featureFlagCreateArgs) Validate() error {
	if a.tenantID < 1 {
		return errors.New("featureFlagCreateArgs.tenantID must be positive integer")
	}
	if a.appID < 1 {
		return errors.New("featureFlagCreateArgs.appID must be positive integer")
	}
	if a.name == "" {
		return errors.New("featureFlagCreateArgs.name cannot be empty")
	}
	return nil
}

func (c *Core) NewFeatureFlagCreateArgs(
	tenantID int64,
	appID int64,
	name string,
	isEnabled bool,
) featureFlagCreateArgs {
	return featureFlagCreateArgs{
		tenantID:  tenantID,
		appID:     appID,
		name:      name,
		isEnabled: isEnabled,
	}
}

func (c *Core) FeatureFlagCreate(ctx context.Context, args featureFlagCreateArgs) (*types.FeatureFlag, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.featureFlagRepo.Create(ctx,
		args.tenantID,
		args.appID,
		args.name,
		args.isEnabled,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) FeatureFlagGetMany(ctx context.Context,
	tenantID int64,
	appID int64,
) ([]types.FeatureFlag, error) {
	return c.featureFlagRepo.GetMany(ctx, tenantID, appID)
}

type featureFlagGetOneArgs struct {
	tenantID int64
	appID    int64
	id       *int64
	uuid     *string
	name     *string
}

func (a *featureFlagGetOneArgs) Validate() error {
	if a.tenantID < 1 {
		return errors.New("featureFlagGetOneArgs.tenantID must be positive integer")
	}
	if a.appID < 1 {
		return errors.New("featureFlagGetOneArgs.appID must be positive integer")
	}
	if a.id == nil && a.uuid == nil && a.name == nil {
		return errors.New("featureFlagGetOneArgs: must provide id, uuid, or name")
	}
	if a.id != nil && *a.id < 1 {
		return errors.New("featureFlagGetOneArgs.id must be positive integer")
	}
	return nil
}

func (c *Core) NewFeatureFlagGetOneArgs(
	tenantID int64,
	appID int64,
	id *int64,
	uuid *string,
	name *string,
) featureFlagGetOneArgs {
	return featureFlagGetOneArgs{
		tenantID: tenantID,
		appID:    appID,
		id:       id,
		uuid:     uuid,
		name:     name,
	}
}

func (c *Core) FeatureFlagGetOne(ctx context.Context, args featureFlagGetOneArgs) (*types.FeatureFlag, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.featureFlagRepo.GetOne(ctx,
		args.tenantID,
		args.appID,
		args.id,
		args.uuid,
		args.name,
	)
}

type featureFlagUpdateArgs struct {
	tenantID  int64
	appID     int64
	id        int64
	name      string
	isEnabled bool
}

func (a *featureFlagUpdateArgs) Validate() error {
	if a.tenantID < 1 {
		return errors.New("featureFlagUpdateArgs.tenantID must be positive integer")
	}
	if a.appID < 1 {
		return errors.New("featureFlagUpdateArgs.appID must be positive integer")
	}
	if a.id < 1 {
		return errors.New("featureFlagUpdateArgs.id must be positive integer")
	}
	if a.name == "" {
		return errors.New("featureFlagUpdateArgs.name cannot be empty")
	}
	return nil
}

func (c *Core) NewFeatureFlagUpdateArgs(
	tenantID int64,
	appID int64,
	id int64,
	name string,
	isEnabled bool,
) featureFlagUpdateArgs {
	return featureFlagUpdateArgs{
		tenantID:  tenantID,
		appID:     appID,
		id:        id,
		name:      name,
		isEnabled: isEnabled,
	}
}

func (c *Core) FeatureFlagUpdate(ctx context.Context, args featureFlagUpdateArgs) (*types.FeatureFlag, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.featureFlagRepo.Update(ctx,
		args.tenantID,
		args.appID,
		args.id,
		args.name,
		args.isEnabled,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) FeatureFlagDelete(ctx context.Context,
	tenantID int64,
	appID int64,
	id int64,
) error {
	return c.featureFlagRepo.Delete(ctx,
		tenantID,
		appID,
		id,
	)
}
