package core

import (
	"context"
	"errors"
	"switchcraft/types"
)

type featureFlagCreateArgs struct {
	orgID     int64
	appID     int64
	name      string
	isEnabled bool
}

func (a *featureFlagCreateArgs) Validate() error {
	if a.orgID < 1 {
		return errors.New("featureFlagCreateArgs.orgID must be positive integer")
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
	orgID int64,
	appID int64,
	name string,
	isEnabled bool,
) featureFlagCreateArgs {
	return featureFlagCreateArgs{
		orgID:     orgID,
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
		args.orgID,
		args.appID,
		args.name,
		args.isEnabled,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) FeatureFlagGetMany(ctx context.Context,
	orgID int64,
	appID int64,
) ([]types.FeatureFlag, error) {
	return c.featureFlagRepo.GetMany(ctx, orgID, appID)
}

type featureFlagGetOneArgs struct {
	orgID int64
	appID int64
	id    *int64
	uuid  *string
	name  *string
}

func (a *featureFlagGetOneArgs) Validate() error {
	if a.orgID < 1 {
		return errors.New("featureFlagGetOneArgs.orgID must be positive integer")
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
	orgID int64,
	appID int64,
	id *int64,
	uuid *string,
	name *string,
) featureFlagGetOneArgs {
	return featureFlagGetOneArgs{
		orgID: orgID,
		appID: appID,
		id:    id,
		uuid:  uuid,
		name:  name,
	}
}

func (c *Core) FeatureFlagGetOne(ctx context.Context, args featureFlagGetOneArgs) (*types.FeatureFlag, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.featureFlagRepo.GetOne(ctx,
		args.orgID,
		args.appID,
		args.id,
		args.uuid,
		args.name,
	)
}

type featureFlagUpdateArgs struct {
	orgID     int64
	appID     int64
	id        int64
	name      string
	isEnabled bool
}

func (a *featureFlagUpdateArgs) Validate() error {
	if a.orgID < 1 {
		return errors.New("featureFlagUpdateArgs.orgID must be positive integer")
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
	orgID int64,
	appID int64,
	id int64,
	name string,
	isEnabled bool,
) featureFlagUpdateArgs {
	return featureFlagUpdateArgs{
		orgID:     orgID,
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
		args.orgID,
		args.appID,
		args.id,
		args.name,
		args.isEnabled,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) FeatureFlagDelete(ctx context.Context,
	orgID int64,
	appID int64,
	id int64,
) error {
	return c.featureFlagRepo.Delete(ctx,
		orgID,
		appID,
		id,
	)
}
