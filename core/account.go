package core

import (
	"context"
	"errors"
	"switchcraft/types"
)

type accountCreateArgs struct {
	firstName string
	lastName  string
	email     string
	username  string
	createdBy int64
}

func (a *accountCreateArgs) Validate() error {
	if a.firstName == "" {
		return errors.New("createAccountArgs.firstName cannot be empty")
	}
	if a.lastName == "" {
		return errors.New("createAccountArgs.lastName cannot be empty")
	}
	if a.email == "" {
		return errors.New("createAccountArgs.email cannot be empty")
	}
	if a.username == "" {
		return errors.New("createAccountArgs.username cannot be empty")
	}
	if a.createdBy == 0 {
		return errors.New("createAccountArgs.createdBy cannot be empty")
	}
	return nil
}

func NewAccountCreateArgs(
	firstName string,
	lastName string,
	email string,
	username string,
	createdBy int64,
) accountCreateArgs {
	return accountCreateArgs{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		username:  username,
		createdBy: createdBy,
	}
}

func (c *Core) AccountCreate(ctx context.Context, args accountCreateArgs) (*types.Account, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.accountRepo.Create(ctx,
		args.firstName,
		args.lastName,
		args.email,
		args.username,
		args.createdBy,
	)
}

func (c *Core) AccountGetMany(ctx context.Context) ([]types.Account, error) {
	return c.accountRepo.GetMany(ctx)
}

type accountGetOneArgs struct {
	id       *int64
	uuid     *string
	username *string
}

func (a *accountGetOneArgs) Validate() error {
	if a.id == nil && a.uuid == nil && a.username == nil {
		return errors.New("error: must provide id, uuid, or username")
	}
	return nil
}

func NewAccountGetOneArgs(id *int64, uuid *string, username *string) accountGetOneArgs {
	return accountGetOneArgs{
		id:       id,
		uuid:     uuid,
		username: username,
	}
}

func (c *Core) AccountGetOne(ctx context.Context, args accountGetOneArgs) (*types.Account, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.accountRepo.GetOne(ctx, args.id, args.uuid, args.username)
}

type accountUpdateArgs struct {
	id         int64
	firstName  string
	lastName   string
	email      string
	username   string
	modifiedBy int64
}

func (a *accountUpdateArgs) Validate() error {
	return nil
}

func NewAccountUpdateArgs(
	id int64,
	firstName string,
	lastName string,
	email string,
	username string,
	modifiedBy int64,
) accountUpdateArgs {
	return accountUpdateArgs{
		id:         id,
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		username:   username,
		modifiedBy: modifiedBy,
	}
}

func (c *Core) AccountUpdate(ctx context.Context, args accountUpdateArgs) (*types.Account, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.accountRepo.Update(ctx,
		args.id,
		args.firstName,
		args.lastName,
		args.email,
		args.username,
		args.modifiedBy,
	)
}

func (c *Core) AccountDelete(ctx context.Context, id int64) error {
	return c.accountRepo.Delete(ctx, id)
}
