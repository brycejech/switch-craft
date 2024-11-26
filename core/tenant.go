package core

import (
	"context"
	"errors"
	"switchcraft/types"
)

type tenantCreateArgs struct {
	name  string
	slug  string
	owner int64
}

func (a *tenantCreateArgs) Validate() error {
	if err := validateSlug(a.slug); err != nil {
		return err
	}
	return nil
}

func (c *Core) NewTenantCreateArgs(
	name string,
	slug string,
	owner int64,
) tenantCreateArgs {
	return tenantCreateArgs{
		name:  name,
		slug:  slug,
		owner: owner,
	}
}

func (c *Core) TenantCreate(ctx context.Context, args tenantCreateArgs) (*types.Tenant, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.tenantRepo.Create(ctx, args.name, args.slug, args.owner, tracer.AuthAccount.ID)
}

func (c *Core) TenantGetMany(ctx context.Context) ([]types.Tenant, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if tracer.AuthAccount.ID == 0 {
		return nil, errors.New("error: invalid operation context authAccount")
	}
	if tracer.AuthAccount.TenantID == nil {
		c.logger.Warn(tracer, "core.TenantGetMany caught nil authAccount.TenantID", nil)
	}

	return c.tenantRepo.GetMany(ctx)
}

func (c *Core) TenantGetOne(ctx context.Context, id int64) (*types.Tenant, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if tracer.AuthAccount.ID == 0 {
		return nil, errors.New("error: invalid operation context authAccount")
	}

	return c.tenantRepo.GetOne(ctx, id)
}

type tenantUpdateArgs struct {
	id    int64
	name  string
	slug  string
	owner int64
}

func (a *tenantUpdateArgs) Validate() error {
	if err := validateSlug(a.slug); err != nil {
		return err
	}
	return nil
}

func (c *Core) NewTenantUpdateArgs(
	id int64,
	name string,
	slug string,
	owner int64,
) tenantUpdateArgs {
	return tenantUpdateArgs{
		id:    id,
		name:  name,
		slug:  slug,
		owner: owner,
	}
}

func (c *Core) TenantUpdate(ctx context.Context, args tenantUpdateArgs) (*types.Tenant, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.tenantRepo.Update(ctx, args.id, args.name, args.slug, args.owner, tracer.AuthAccount.ID)
}

func (c *Core) TenantDelete(ctx context.Context, id int64) error {
	return c.tenantRepo.Delete(ctx, id)
}
