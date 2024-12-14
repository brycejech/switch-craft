package core

import (
	"context"
	"errors"
	"switchcraft/types"
)

type orgAccountCreateArgs struct {
	orgSlug   string
	firstName string
	lastName  string
	email     string
	username  string
	password  *string
}

func (a *orgAccountCreateArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("accountCreateArgs.orgSlug cannot be empty")
	}
	if a.firstName == "" {
		return errors.New("accountCreateArgs.firstName cannot be empty")
	}
	if a.lastName == "" {
		return errors.New("accountCreateArgs.lastName cannot be empty")
	}
	if a.email == "" {
		return errors.New("accountCreateArgs.email cannot be empty")
	}
	if a.username == "" {
		return errors.New("accountCreateArgs.username cannot be empty")
	}
	return nil
}

func (c *Core) NewOrgAccountCreateArgs(
	orgSlug string,
	firstName string,
	lastName string,
	email string,
	username string,
	password *string,
) orgAccountCreateArgs {
	return orgAccountCreateArgs{
		orgSlug:   orgSlug,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		username:  username,
		password:  password,
	}
}

func (c *Core) OrgAccountCreate(ctx context.Context, args orgAccountCreateArgs) (*types.Account, error) {
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

	var password *string
	if args.password != nil {
		tmpPass, err := c.AuthPasswordHash(*args.password)
		if err != nil {
			return nil, err
		}
		if tmpPass == "" {
			password = nil
		} else {
			password = &tmpPass
		}
	}

	return c.orgAccountRepo.Create(ctx,
		org.ID,
		args.firstName,
		args.lastName,
		args.email,
		args.username,
		password,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) OrgAccountGetMany(ctx context.Context, orgSlug string) ([]types.Account, error) {
	if orgSlug == "" {
		return nil, errors.New("core.OrgAccountGetMany orgSlug cannot be empty")
	}

	org, err := c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &orgSlug))
	if err != nil {
		return nil, err
	}

	return c.orgAccountRepo.GetMany(ctx, org.ID)
}

func (c *Core) OrgAccountGetManyByID(ctx context.Context,
	orgSlug string,
	accountIDs []int64,
) ([]types.Account, error) {
	if orgSlug == "" {
		return nil, errors.New("core.OrgAccountGetManyByID org slug cannot be empty")
	}

	for _, id := range accountIDs {
		if id < 1 {
			return nil, errors.New("core.OrgAccountGetManyByID accountIDs must be positive integer")
		}
	}

	org, err := c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &orgSlug))
	if err != nil {
		return nil, err
	}

	return c.orgAccountRepo.GetManyByID(ctx, org.ID, accountIDs)
}

type orgAccountGetOneArgs struct {
	orgSlug  string
	id       *int64
	uuid     *string
	username *string
}

func (a *orgAccountGetOneArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("orgAccountGetOneArgs orgSlug cannot be empty")
	}
	if a.id == nil && a.uuid == nil && a.username == nil {
		return errors.New("error: must provide id, uuid, or username")
	}
	return nil
}

func (c *Core) NewOrgAccountGetOneArgs(orgSlug string, id *int64, uuid *string, username *string) orgAccountGetOneArgs {
	return orgAccountGetOneArgs{
		orgSlug:  orgSlug,
		id:       id,
		uuid:     uuid,
		username: username,
	}
}

func (c *Core) OrgAccountGetOne(ctx context.Context, args orgAccountGetOneArgs) (*types.Account, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	org, err := c.OrgGetOne(ctx,
		c.NewOrgGetOneArgs(nil, nil, &args.orgSlug),
	)
	if err != nil {
		return nil, err
	}

	return c.orgAccountRepo.GetOne(ctx, org.ID, args.id, args.uuid, args.username)
}

type orgAccountUpdateArgs struct {
	orgSlug   string
	id        int64
	firstName string
	lastName  string
	email     string
	username  string
}

func (a *orgAccountUpdateArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("orgAccountUpdateArgs orgSlug cannot be empty")
	}
	return nil
}

func (c *Core) NewOrgAccountUpdateArgs(
	orgSlug string,
	id int64,
	firstName string,
	lastName string,
	email string,
	username string,
) orgAccountUpdateArgs {
	return orgAccountUpdateArgs{
		orgSlug:   orgSlug,
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		username:  username,
	}
}

func (c *Core) OrgAccountUpdate(ctx context.Context, args orgAccountUpdateArgs) (*types.Account, error) {
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

	return c.orgAccountRepo.Update(ctx,
		org.ID,
		args.id,
		args.firstName,
		args.lastName,
		args.email,
		args.username,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) OrgAccountDelete(ctx context.Context, orgSlug string, id int64) error {
	if orgSlug == "" {
		return errors.New("core.OrgAccountDelete orgSlug cannot be empty")
	}

	org, err := c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &orgSlug))
	if err != nil {
		return err
	}

	return c.orgAccountRepo.Delete(ctx, org.ID, id)
}
