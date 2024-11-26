package core

import (
	"context"
	"errors"
	"switchcraft/types"
)

type appCreateArgs struct {
	tenantID int64
	name     string
	slug     string
}

func (a *appCreateArgs) Validate() error {
	if a.tenantID < 1 {
		return errors.New("appCreateArgs.tenantID must be positive integer")
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
	tenantID int64,
	name string,
	slug string,
) appCreateArgs {
	return appCreateArgs{
		tenantID: tenantID,
		name:     name,
		slug:     slug,
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

	return c.appRepo.Create(ctx,
		args.tenantID,
		args.name,
		args.slug,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) AppGetMany(ctx context.Context, tenantID int64) ([]types.Application, error) {
	return c.appRepo.GetMany(ctx, tenantID)
}

type appGetOneArgs struct {
	tenantID int64
	id       *int64
	uuid     *string
	slug     *string
}

func (a *appGetOneArgs) Validate() error {
	if a.tenantID < 1 {
		return errors.New("appGetOneArgs.tenantID must be positive integer")
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
	tenantID int64,
	id *int64,
	uuid *string,
	slug *string,
) appGetOneArgs {
	return appGetOneArgs{
		tenantID: tenantID,
		id:       id,
		uuid:     uuid,
		slug:     slug,
	}
}

func (c *Core) AppGetOne(ctx context.Context, args appGetOneArgs) (*types.Application, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.appRepo.GetOne(ctx, args.tenantID, args.id, args.uuid, args.slug)
}

type appUpdateArgs struct {
	tenantID int64
	id       int64
	name     string
	slug     string
}

func (a *appUpdateArgs) Validate() error {
	if a.tenantID < 1 {
		return errors.New("appUpdateArgs: tenantID must be positive integer")
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
	tenantID int64,
	id int64,
	name string,
	slug string,
) appUpdateArgs {
	return appUpdateArgs{
		id:       id,
		tenantID: tenantID,
		name:     name,
		slug:     slug,
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

	return c.appRepo.Update(ctx,
		args.tenantID,
		args.id,
		args.name,
		args.slug,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) AppDelete(ctx context.Context, tenantID int64, id int64) error {
	return c.appRepo.Delete(ctx, tenantID, id)
}
