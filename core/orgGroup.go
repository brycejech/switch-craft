package core

import (
	"context"
	"errors"
	"switchcraft/types"
)

type orgGroupCreateArgs struct {
	orgSlug     string
	name        string
	description string
}

func (a *orgGroupCreateArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("orgGroupCreateArgs.orgSlug cannot be empty")
	}
	if a.name == "" {
		return errors.New("orgGroupCreateArgs.name cannot be empty")
	}
	return nil
}

func (c *Core) NewOrgGroupCreateArgs(
	orgSlug string,
	name string,
	description string,
) orgGroupCreateArgs {
	return orgGroupCreateArgs{
		orgSlug:     orgSlug,
		name:        name,
		description: description,
	}
}

func (c *Core) OrgGroupCreate(ctx context.Context, args orgGroupCreateArgs) (*types.OrgGroup, error) {
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

	return c.orgGroupRepo.Create(ctx,
		org.ID,
		args.name,
		args.description,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) OrgGroupGetMany(ctx context.Context, orgSlug string) ([]types.OrgGroup, error) {
	if orgSlug == "" {
		return nil, errors.New("core.OrgGroupGetMany orgSlug cannot be empty")
	}

	org, err := c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &orgSlug))
	if err != nil {
		return nil, err
	}

	return c.orgGroupRepo.GetMany(ctx, org.ID)
}

type orgGroupGetOneArgs struct {
	orgSlug string
	id      *int64
	uuid    *string
}

func (a *orgGroupGetOneArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("orgGroupGetOneArgs orgSlug cannot be empty")
	}
	if a.id == nil && a.uuid == nil {
		return errors.New("orgGroupGetOneArgs must provide id or uuid")
	}
	return nil
}

func (c *Core) NewOrgGroupGetOneArgs(orgSlug string, id *int64, uuid *string) orgGroupGetOneArgs {
	return orgGroupGetOneArgs{
		orgSlug: orgSlug,
		id:      id,
		uuid:    uuid,
	}
}

func (c *Core) OrgGroupGetOne(ctx context.Context, args orgGroupGetOneArgs) (*types.OrgGroup, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	org, err := c.OrgGetOne(ctx,
		c.NewOrgGetOneArgs(nil, nil, &args.orgSlug),
	)
	if err != nil {
		return nil, err
	}

	return c.orgGroupRepo.GetOne(ctx, org.ID, args.id, args.uuid)
}

type orgGroupUpdateArgs struct {
	orgSlug     string
	id          int64
	name        string
	description string
}

func (a *orgGroupUpdateArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("orgGroupUpdateArgs orgSlug cannot be empty")
	}
	if a.id < 1 {
		return errors.New("orgGroupUpdateArgs id must be positive integer")
	}
	if a.name == "" {
		return errors.New("orgGroupUpdateArgs name cannot be empty")
	}

	return nil
}

func (c *Core) NewOrgGroupUpdateArgs(
	orgSlug string,
	id int64,
	name string,
	description string,
) orgGroupUpdateArgs {
	return orgGroupUpdateArgs{
		orgSlug:     orgSlug,
		id:          id,
		name:        name,
		description: description,
	}
}

func (c *Core) OrgGroupUpdate(ctx context.Context, args orgGroupUpdateArgs) (*types.OrgGroup, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	org, err := c.OrgGetOne(ctx,
		c.NewOrgGetOneArgs(nil, nil, &args.orgSlug),
	)
	if err != nil {
		return nil, err
	}

	return c.orgGroupRepo.Update(ctx,
		org.ID,
		args.id,
		args.name,
		args.description,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) OrgGroupDelete(ctx context.Context, orgSlug string, id int64) error {
	if orgSlug == "" {
		return errors.New("core.OrgGroupDelete orgSlug cannot be empty")
	}
	if id < 1 {
		return errors.New("core.OrgGroupDelete id must be positive integer")
	}

	org, err := c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &orgSlug))
	if err != nil {
		return err
	}

	return c.orgGroupRepo.Delete(ctx, org.ID, id)
}
