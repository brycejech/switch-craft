package core

import (
	"context"
	"errors"
	"switchcraft/types"
)

type accountCreateArgs struct {
	tenantID  int64
	firstName string
	lastName  string
	email     string
	username  string
	password  *string
}

func (a *accountCreateArgs) Validate() error {
	if a.tenantID < 1 {
		return errors.New("accountCreateArgs.tenantID must be positive integer")
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

func NewAccountCreateArgs(
	tenantID int64,
	firstName string,
	lastName string,
	email string,
	username string,
	password *string,
) accountCreateArgs {
	return accountCreateArgs{
		tenantID:  tenantID,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		username:  username,
		password:  password,
	}
}

func (c *Core) AccountCreate(ctx context.Context, args accountCreateArgs) (*types.Account, error) {
	opCtx, ok := ctx.Value(types.CtxOperationTracker).(types.OperationTracker)
	if !ok {
		return nil, errors.New("error: invalid operation context")
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
		&args.tenantID,
		args.firstName,
		args.lastName,
		args.email,
		args.username,
		password,
		opCtx.AuthAccount.ID,
	)
}

type accountCreateGlobalArgs struct {
	firstName string
	lastName  string
	email     string
	username  string
	password  string
}

func (a *accountCreateGlobalArgs) Validate() error {
	if len(a.password) < 12 {
		return errors.New("accountCreateGlobalArgs.password must be at least 12 characters")
	}
	return nil
}

func NewAccountCreateGlobalArgs(
	firstName string,
	lastName string,
	email string,
	username string,
	password string,
) accountCreateGlobalArgs {
	return accountCreateGlobalArgs{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		username:  username,
		password:  password,
	}
}

func (c *Core) AccountCreateGlobal(ctx context.Context, args accountCreateGlobalArgs) (*types.Account, error) {
	opCtx, ok := ctx.Value(types.CtxOperationTracker).(types.OperationTracker)
	if !ok {
		return nil, errors.New("error casting operation context")
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
		args.firstName,
		args.lastName,
		args.email,
		args.username,
		&hashedPassword,
		opCtx.AuthAccount.ID,
	)
}

func (c *Core) AccountGetMany(ctx context.Context, tenantID *int64) ([]types.Account, error) {
	return c.accountRepo.GetMany(ctx, tenantID)
}

type accountGetOneArgs struct {
	tenantID *int64
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

func NewAccountGetOneArgs(tenantID *int64, id *int64, uuid *string, username *string) accountGetOneArgs {
	return accountGetOneArgs{
		tenantID: tenantID,
		id:       id,
		uuid:     uuid,
		username: username,
	}
}

func (c *Core) AccountGetOne(ctx context.Context, args accountGetOneArgs) (*types.Account, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.accountRepo.GetOne(ctx, args.tenantID, args.id, args.uuid, args.username)
}

type accountUpdateArgs struct {
	tenantID  *int64
	id        int64
	firstName string
	lastName  string
	email     string
	username  string
}

func (a *accountUpdateArgs) Validate() error {
	return nil
}

func NewAccountUpdateArgs(
	tenantID *int64,
	id int64,
	firstName string,
	lastName string,
	email string,
	username string,
) accountUpdateArgs {
	return accountUpdateArgs{
		tenantID:  tenantID,
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		username:  username,
	}
}

func (c *Core) AccountUpdate(ctx context.Context, args accountUpdateArgs) (*types.Account, error) {
	opCtx, ok := ctx.Value(types.CtxOperationTracker).(types.OperationTracker)
	if !ok {
		return nil, errors.New("error casting operation context")
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	return c.accountRepo.Update(ctx,
		args.tenantID,
		args.id,
		args.firstName,
		args.lastName,
		args.email,
		args.username,
		opCtx.AuthAccount.ID,
	)
}

func (c *Core) AccountDelete(ctx context.Context, tenantID *int64, id int64) error {
	return c.accountRepo.Delete(ctx, tenantID, id)
}
