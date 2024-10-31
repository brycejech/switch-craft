package core

import (
	"context"
	"errors"
	"switchcraft/types"
)

type createAccountArgs struct {
	firstName string
	lastName  string
	email     string
	username  string
	createdBy int64
}

func (a *createAccountArgs) Validate() error {
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

func NewCreateAccountArgs(
	firstName string,
	lastName string,
	email string,
	username string,
	createdBy int64,
) createAccountArgs {
	return createAccountArgs{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		username:  username,
		createdBy: createdBy,
	}
}

func (c *Core) CreateAccount(ctx context.Context, args createAccountArgs) (*types.Account, error) {
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

func (c *Core) GetAccounts(ctx context.Context) ([]types.Account, error) {
	return c.accountRepo.GetMany(ctx)
}

type getAccountArgs struct {
	id       *int64
	uuid     *string
	username *string
}

func (a *getAccountArgs) Validate() error {
	if a.id == nil && a.uuid == nil && a.username == nil {
		return errors.New("error: must provide id, uuid, or username")
	}
	return nil
}

func NewGetAccountArgs(id *int64, uuid *string, username *string) getAccountArgs {
	return getAccountArgs{
		id:       id,
		uuid:     uuid,
		username: username,
	}
}

func (c *Core) GetAccount(ctx context.Context, args getAccountArgs) (*types.Account, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.accountRepo.GetOne(ctx, args.id, args.uuid, args.username)
}

type updateAccountArgs struct {
	id         int64
	firstName  string
	lastName   string
	email      string
	username   string
	modifiedBy int64
}

func (a *updateAccountArgs) Validate() error {
	return nil
}

func NewUpdateAccountArgs(
	id int64,
	firstName string,
	lastName string,
	email string,
	username string,
	modifiedBy int64,
) updateAccountArgs {
	return updateAccountArgs{
		id:         id,
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		username:   username,
		modifiedBy: modifiedBy,
	}
}

func (c *Core) UpdateAccount(ctx context.Context, args updateAccountArgs) (*types.Account, error) {
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

func (c *Core) DeleteAccount(ctx context.Context, id int64) error {
	return c.accountRepo.Delete(ctx, id)
}
