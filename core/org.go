package core

import (
	"context"
	"errors"
	"switchcraft/types"
)

type orgCreateArgs struct {
	name  string
	slug  string
	owner int64
}

func (a *orgCreateArgs) Validate() error {
	if err := validateSlug(a.slug); err != nil {
		return err
	}
	return nil
}

func (c *Core) NewOrgCreateArgs(
	name string,
	slug string,
	owner int64,
) orgCreateArgs {
	return orgCreateArgs{
		name:  name,
		slug:  slug,
		owner: owner,
	}
}

func (c *Core) OrgCreate(ctx context.Context, args orgCreateArgs) (*types.Org, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.orgRepo.Create(ctx, args.name, args.slug, args.owner, tracer.AuthAccount.ID)
}

func (c *Core) OrgGetMany(ctx context.Context) ([]types.Org, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if tracer.AuthAccount.ID == 0 {
		return nil, errors.New("error: invalid operation context authAccount")
	}
	if tracer.AuthAccount.OrgID == nil {
		c.logger.Warn(tracer, "core.OrgGetMany caught nil authAccount.OrgID", nil)
	}

	return c.orgRepo.GetMany(ctx)
}

type orgGetOneArgs struct {
	id   *int64
	uuid *string
	slug *string
}

func (a *orgGetOneArgs) Validate() error {
	if a.id == nil && a.uuid == nil && a.slug == nil {
		return errors.New("orgGetOneArgs: must provide one of id, uuid, or slug")
	}

	if a.id != nil && *a.id < 1 {
		return errors.New("orgGetOneArgs: id must be positive integer")
	}

	return nil
}

func (c *Core) NewOrgGetOneArgs(id *int64, uuid *string, slug *string) orgGetOneArgs {
	return orgGetOneArgs{
		id:   id,
		uuid: uuid,
		slug: slug,
	}
}

func (c *Core) OrgGetOne(ctx context.Context, args orgGetOneArgs) (*types.Org, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if tracer.AuthAccount.ID == 0 {
		return nil, errors.New("error: invalid operation context authAccount")
	}

	return c.orgRepo.GetOne(ctx, args.id, args.uuid, args.slug)
}

type orgUpdateArgs struct {
	id    int64
	name  string
	slug  string
	owner int64
}

func (a *orgUpdateArgs) Validate() error {
	if err := validateSlug(a.slug); err != nil {
		return err
	}
	return nil
}

func (c *Core) NewOrgUpdateArgs(
	id int64,
	name string,
	slug string,
	owner int64,
) orgUpdateArgs {
	return orgUpdateArgs{
		id:    id,
		name:  name,
		slug:  slug,
		owner: owner,
	}
}

func (c *Core) OrgUpdate(ctx context.Context, args orgUpdateArgs) (*types.Org, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.orgRepo.Update(ctx, args.id, args.name, args.slug, args.owner, tracer.AuthAccount.ID)
}

func (c *Core) OrgDelete(ctx context.Context, id int64) error {
	return c.orgRepo.Delete(ctx, id)
}
