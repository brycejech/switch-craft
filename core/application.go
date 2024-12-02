package core

import (
	"context"
	"errors"
	"switchcraft/types"
)

type appCreateArgs struct {
	orgSlug string
	name    string
	slug    string
}

func (a *appCreateArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("appCreateArgs.orgSlug cannot be empty")
	}
	if a.name == "" {
		return errors.New("appCreateArgs.name cannot be empty")
	}
	if err := validateSlug(a.slug); err != nil {
		return err
	}
	return nil
}

func (c *Core) NewAppCreateArgs(
	orgSlug string,
	name string,
	slug string,
) appCreateArgs {
	return appCreateArgs{
		orgSlug: orgSlug,
		name:    name,
		slug:    slug,
	}
}

func (c *Core) AppCreate(ctx context.Context, args appCreateArgs) (*types.Application, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	org, err := c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &args.orgSlug))
	if err != nil {
		return nil, err
	}

	return c.appRepo.Create(ctx,
		org.ID,
		args.name,
		args.slug,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) AppGetMany(ctx context.Context, orgSlug string) ([]types.Application, error) {
	var (
		org *types.Org
		err error
	)
	if org, err = c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &orgSlug)); err != nil {
		return nil, err
	}

	return c.appRepo.GetMany(ctx, org.ID)
}

type appGetOneArgs struct {
	orgSlug string
	id      *int64
	uuid    *string
	slug    *string
}

func (a *appGetOneArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("appGetOneArgs.orgSlug cannot be empty")
	}
	if a.id == nil && a.uuid == nil && a.slug == nil {
		return errors.New("appGetOneArgs: must provide id, uuid, or slug")
	}
	if a.slug != nil {
		if err := validateSlug(*a.slug); err != nil {
			return err
		}
	}
	return nil
}

func (c *Core) NewAppGetOneArgs(
	orgSlug string,
	id *int64,
	uuid *string,
	slug *string,
) appGetOneArgs {
	return appGetOneArgs{
		orgSlug: orgSlug,
		id:      id,
		uuid:    uuid,
		slug:    slug,
	}
}

func (c *Core) AppGetOne(ctx context.Context, args appGetOneArgs) (*types.Application, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	org, err := c.OrgGetOne(ctx,
		c.NewOrgGetOneArgs(nil, nil, &args.orgSlug),
	)
	if err != nil {
		return nil, err
	}

	return c.appRepo.GetOne(ctx, org.ID, args.id, args.uuid, args.slug)
}

type appUpdateArgs struct {
	orgSlug string
	id      int64
	name    string
	slug    string
}

func (a *appUpdateArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("appUpdateArgs.orgSlug cannot be empty")
	}
	if a.id < 1 {
		return errors.New("appUpdateArgs: id must be positive integer")
	}
	if a.name == "" {
		return errors.New("appUpdateArgs: must provide application name")
	}
	if err := validateSlug(a.slug); err != nil {
		return err
	}
	return nil
}

func (c *Core) NewAppUpdateArgs(
	orgSlug string,
	id int64,
	name string,
	slug string,
) appUpdateArgs {
	return appUpdateArgs{
		orgSlug: orgSlug,
		id:      id,
		name:    name,
		slug:    slug,
	}
}

func (c *Core) AppUpdate(ctx context.Context, args appUpdateArgs) (*types.Application, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	org, err := c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &args.orgSlug))
	if err != nil {
		return nil, err
	}

	return c.appRepo.Update(ctx,
		org.ID,
		args.id,
		args.name,
		args.slug,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) AppDelete(ctx context.Context, orgSlug string, appSlug string) error {
	var (
		org *types.Org
		app *types.Application
		err error
	)

	if org, err = c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &orgSlug)); err != nil {
		return err
	}
	if app, err = c.AppGetOne(ctx, c.NewAppGetOneArgs(orgSlug, nil, nil, &appSlug)); err != nil {
		return err
	}

	return c.appRepo.Delete(ctx, org.ID, app.ID)
}
