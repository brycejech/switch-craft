package core

import (
	"context"
	"errors"
	"switchcraft/types"
)

type featFlagCreateArgs struct {
	orgSlug   string
	appSlug   string
	name      string
	isEnabled bool
}

func (a *featFlagCreateArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("featFlagCreateArgs.orgSlug cannot be empty")
	}
	if a.appSlug == "" {
		return errors.New("featFlagCreateArgs.appSlug cannot be empty")
	}
	if a.name == "" {
		return errors.New("featFlagCreateArgs.name cannot be empty")
	}
	return nil
}

func (c *Core) NewFeatFlagCreateArgs(
	orgSlug string,
	appSlug string,
	name string,
	isEnabled bool,
) featFlagCreateArgs {
	return featFlagCreateArgs{
		orgSlug:   orgSlug,
		appSlug:   appSlug,
		name:      name,
		isEnabled: isEnabled,
	}
}

func (c *Core) FeatFlagCreate(ctx context.Context, args featFlagCreateArgs) (*types.FeatureFlag, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	var (
		org *types.Organization
		app *types.Application
	)

	if org, err = c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &args.orgSlug)); err != nil {
		return nil, err
	}
	if app, err = c.AppGetOne(ctx, c.NewAppGetOneArgs(org.Slug, nil, nil, &args.appSlug)); err != nil {
		return nil, err
	}

	return c.featureFlagRepo.Create(ctx,
		org.ID,
		app.ID,
		args.name,
		args.isEnabled,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) FeatFlagGetMany(ctx context.Context,
	orgSlug string,
	appSlug string,
) ([]types.FeatureFlag, error) {
	var (
		org *types.Organization
		app *types.Application
		err error
	)

	if org, err = c.OrgGetOne(ctx,
		c.NewOrgGetOneArgs(nil, nil, &orgSlug),
	); err != nil {
		return nil, err
	}
	if app, err = c.AppGetOne(ctx,
		c.NewAppGetOneArgs(orgSlug, nil, nil, &appSlug),
	); err != nil {
		return nil, err
	}

	return c.featureFlagRepo.GetMany(ctx, org.ID, app.ID)
}

type featFlagGetOneArgs struct {
	orgSlug string
	appSlug string
	id      *int64
	uuid    *string
	name    *string
}

func (a *featFlagGetOneArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("featFlagGetOneArgs.orgSlug cannot be empty")
	}
	if a.appSlug == "" {
		return errors.New("featFlagGetOneArgs.appSlug cannot be empty")
	}
	if a.id == nil && a.uuid == nil && a.name == nil {
		return errors.New("featFlagGetOneArgs: must provide id, uuid, or name")
	}
	if a.id != nil && *a.id < 1 {
		return errors.New("featFlagGetOneArgs.id must be positive integer")
	}
	return nil
}

func (c *Core) NewFeatFlagGetOneArgs(
	orgSlug string,
	appSlug string,
	id *int64,
	uuid *string,
	name *string,
) featFlagGetOneArgs {
	return featFlagGetOneArgs{
		orgSlug: orgSlug,
		appSlug: appSlug,
		id:      id,
		uuid:    uuid,
		name:    name,
	}
}

func (c *Core) FeatFlagGetOne(ctx context.Context, args featFlagGetOneArgs) (*types.FeatureFlag, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	var (
		org *types.Organization
		app *types.Application
		err error
	)

	if org, err = c.OrgGetOne(ctx,
		c.NewOrgGetOneArgs(nil, nil, &args.orgSlug),
	); err != nil {
		return nil, err
	}
	if app, err = c.AppGetOne(ctx,
		c.NewAppGetOneArgs(org.Slug, nil, nil, &args.appSlug),
	); err != nil {
		return nil, err
	}

	return c.featureFlagRepo.GetOne(ctx,
		org.ID,
		app.ID,
		args.id,
		args.uuid,
		args.name,
	)
}

type featFlagUpdateArgs struct {
	orgSlug   string
	appSlug   string
	id        int64
	name      string
	isEnabled bool
}

func (a *featFlagUpdateArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("featFlagUpdateArgs.orgSlug cannot be empty")
	}
	if a.appSlug == "" {
		return errors.New("featFlagUpdateArgs.appSlug cannot be empty")
	}
	if a.id < 1 {
		return errors.New("featFlagUpdateArgs.id must be positive integer")
	}
	if a.name == "" {
		return errors.New("featFlagUpdateArgs.name cannot be empty")
	}
	return nil
}

func (c *Core) NewFeatFlagUpdateArgs(
	orgSlug string,
	appSlug string,
	id int64,
	name string,
	isEnabled bool,
) featFlagUpdateArgs {
	return featFlagUpdateArgs{
		orgSlug:   orgSlug,
		appSlug:   appSlug,
		id:        id,
		name:      name,
		isEnabled: isEnabled,
	}
}

func (c *Core) FeatFlagUpdate(ctx context.Context, args featFlagUpdateArgs) (*types.FeatureFlag, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	var (
		org *types.Organization
		app *types.Application
	)
	if org, err = c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &args.orgSlug)); err != nil {
		return nil, err
	}
	if app, err = c.AppGetOne(ctx, c.NewAppGetOneArgs(org.Slug, nil, nil, &args.appSlug)); err != nil {
		return nil, err
	}

	return c.featureFlagRepo.Update(ctx,
		org.ID,
		app.ID,
		args.id,
		args.name,
		args.isEnabled,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) FeatFlagDelete(ctx context.Context,
	orgSlug string,
	appSlug string,
	id int64,
) error {
	var (
		org *types.Organization
		app *types.Application
		err error
	)

	if org, err = c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &orgSlug)); err != nil {
		return err
	}
	if app, err = c.AppGetOne(ctx, c.NewAppGetOneArgs(org.Slug, nil, nil, &appSlug)); err != nil {
		return err
	}

	return c.featureFlagRepo.Delete(ctx,
		org.ID,
		app.ID,
		id,
	)
}
