package core

import (
	"context"
	"errors"
	"switchcraft/types"
)

type accountCreateArgs struct {
	orgID     int64
	firstName string
	lastName  string
	email     string
	username  string
	password  *string
}

func (a *accountCreateArgs) Validate() error {
	if a.orgID < 1 {
		return errors.New("accountCreateArgs.orgID must be positive integer")
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

func (c *Core) NewAccountCreateArgs(
	orgID int64,
	firstName string,
	lastName string,
	email string,
	username string,
	password *string,
) accountCreateArgs {
	return accountCreateArgs{
		orgID:     orgID,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		username:  username,
		password:  password,
	}
}

func (c *Core) AccountCreate(ctx context.Context, args accountCreateArgs) (*types.Account, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
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

	return c.accountRepo.Create(ctx,
		&args.orgID,
		false,
		args.firstName,
		args.lastName,
		args.email,
		args.username,
		password,
		tracer.AuthAccount.ID,
	)
}

type accountCreateGlobalArgs struct {
	isInstanceAdmin bool
	firstName       string
	lastName        string
	email           string
	username        string
	password        string
}

func (a *accountCreateGlobalArgs) Validate() error {
	if len(a.password) < 12 {
		return errors.New("accountCreateGlobalArgs.password must be at least 12 characters")
	}
	return nil
}

func (c *Core) NewAccountCreateGlobalArgs(
	isInstanceAdmin bool,
	firstName string,
	lastName string,
	email string,
	username string,
	password string,
) accountCreateGlobalArgs {
	return accountCreateGlobalArgs{
		isInstanceAdmin: isInstanceAdmin,
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		username:        username,
		password:        password,
	}
}

func (c *Core) AccountCreateGlobal(ctx context.Context, args accountCreateGlobalArgs) (*types.Account, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	hashedPassword, err := c.AuthPasswordHash(args.password)
	if err != nil {
		return nil, err
	}

	return c.accountRepo.Create(ctx,
		nil,
		args.isInstanceAdmin,
		args.firstName,
		args.lastName,
		args.email,
		args.username,
		&hashedPassword,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) AccountGetMany(ctx context.Context, orgID *int64) ([]types.Account, error) {
	return c.accountRepo.GetMany(ctx, orgID)
}

type accountGetOneArgs struct {
	orgID    *int64
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

func (c *Core) NewAccountGetOneArgs(orgID *int64, id *int64, uuid *string, username *string) accountGetOneArgs {
	return accountGetOneArgs{
		orgID:    orgID,
		id:       id,
		uuid:     uuid,
		username: username,
	}
}

func (c *Core) AccountGetOne(ctx context.Context, args accountGetOneArgs) (*types.Account, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.accountRepo.GetOne(ctx, args.orgID, args.id, args.uuid, args.username)
}

type accountUpdateArgs struct {
	orgID           *int64
	id              int64
	isInstanceAdmin bool
	firstName       string
	lastName        string
	email           string
	username        string
}

func (a *accountUpdateArgs) Validate() error {
	return nil
}

func (c *Core) NewAccountUpdateArgs(
	orgID *int64,
	id int64,
	isInstanceAdmin bool,
	firstName string,
	lastName string,
	email string,
	username string,
) accountUpdateArgs {
	return accountUpdateArgs{
		orgID:           orgID,
		id:              id,
		isInstanceAdmin: isInstanceAdmin,
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		username:        username,
	}
}

func (c *Core) AccountUpdate(ctx context.Context, args accountUpdateArgs) (*types.Account, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	existingAccount, err := c.AccountGetOne(ctx, c.NewAccountGetOneArgs(args.orgID, &args.id, nil, nil))
	if err != nil {
		return nil, err
	}

	// Only allow instance admins to modify Account.isInstanceAdmin
	isInstanceAdmin := existingAccount.IsInstanceAdmin
	if existingAccount.IsInstanceAdmin != args.isInstanceAdmin {
		if tracer.AuthAccount.IsInstanceAdmin {
			isInstanceAdmin = args.isInstanceAdmin
		} else {
			return nil, types.ErrOperationNotPermitted
		}
	}

	return c.accountRepo.Update(ctx,
		args.orgID,
		args.id,
		isInstanceAdmin,
		args.firstName,
		args.lastName,
		args.email,
		args.username,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) AccountDelete(ctx context.Context, orgID *int64, id int64) error {
	return c.accountRepo.Delete(ctx, orgID, id)
}
