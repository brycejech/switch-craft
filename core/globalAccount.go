package core

import (
	"context"
	"errors"
	"switchcraft/types"
)

type globalAccountCreateArgs struct {
	isInstanceAdmin bool
	firstName       string
	lastName        string
	email           string
	username        string
	password        string
}

func (a *globalAccountCreateArgs) Validate() error {
	if a.firstName == "" {
		return errors.New("globalAccountCreateArgs.firstName cannot be empty")
	}
	if a.lastName == "" {
		return errors.New("globalAccountCreateArgs.lastName cannot be empty")
	}
	if a.email == "" {
		return errors.New("globalAccountCreateArgs.email cannot be empty")
	}
	if a.username == "" {
		return errors.New("globalAccountCreateArgs.username cannot be empty")
	}
	return nil
}

func (c *Core) NewGlobalAccountCreateArgs(
	isInstanceAdmin bool,
	firstName string,
	lastName string,
	email string,
	username string,
	password string,
) globalAccountCreateArgs {
	return globalAccountCreateArgs{
		isInstanceAdmin: isInstanceAdmin,
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		username:        username,
		password:        password,
	}
}

func (c *Core) GlobalAccountCreate(ctx context.Context, args globalAccountCreateArgs) (*types.Account, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	password, err := c.AuthPasswordHash(args.password)
	if err != nil {
		return nil, err
	}

	return c.globalAccountRepo.Create(ctx,
		args.isInstanceAdmin,
		args.firstName,
		args.lastName,
		args.email,
		args.username,
		password,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) GlobalAccountGetMany(ctx context.Context) ([]types.Account, error) {
	return c.globalAccountRepo.GetMany(ctx)
}

type globalAccountGetOneArgs struct {
	id       *int64
	uuid     *string
	username *string
}

func (a *globalAccountGetOneArgs) Validate() error {
	if a.id == nil && a.uuid == nil && a.username == nil {
		return errors.New("globalAccountGetOneArgs must provide id, uuid, or username")
	}
	return nil
}

func (c *Core) NewGlobalAccountGetOneArgs(id *int64, uuid *string, username *string) globalAccountGetOneArgs {
	return globalAccountGetOneArgs{
		id:       id,
		uuid:     uuid,
		username: username,
	}
}

func (c *Core) GlobalAccountGetOne(ctx context.Context, args globalAccountGetOneArgs) (*types.Account, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.globalAccountRepo.GetOne(ctx, args.id, args.uuid, args.username)

}

type globalAccountUpdateArgs struct {
	id              int64
	isInstanceAdmin bool
	firstName       string
	lastName        string
	email           string
	username        string
}

func (a *globalAccountUpdateArgs) Validate() error {
	return nil
}

func (c *Core) NewGlobalAccountUpdateArgs(
	id int64,
	isInstanceAdmin bool,
	firstName string,
	lastName string,
	email string,
	username string,
) globalAccountUpdateArgs {
	return globalAccountUpdateArgs{
		id:              id,
		isInstanceAdmin: isInstanceAdmin,
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		username:        username,
	}
}

func (c *Core) GlobalAccountUpdate(ctx context.Context, args globalAccountUpdateArgs) (*types.Account, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.globalAccountRepo.Update(ctx,
		args.id,
		args.isInstanceAdmin,
		args.firstName,
		args.lastName,
		args.email,
		args.username,
		tracer.AuthAccount.ID,
	)

}

func (c *Core) GlobalAccountDelete(ctx context.Context, id int64) error {
	return c.globalAccountRepo.Delete(ctx, id)
}
