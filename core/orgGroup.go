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

type orgGroupAccountAddArgs struct {
	orgSlug   string
	groupID   int64
	accountID int64
}

func (a *orgGroupAccountAddArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("orgGroupAccountAddArgs orgSlug cannot be empty")
	}
	if a.groupID < 1 {
		return errors.New("orgGroupAccountAddArgs groupID must be a positive integer")
	}
	if a.accountID < 1 {
		return errors.New("orgGroupAccountAddArgs accountID must be a positive integer")
	}

	return nil
}

func (c *Core) NewOrgGroupAccountAddArgs(
	orgSlug string,
	groupID int64,
	accountID int64,
) orgGroupAccountAddArgs {
	return orgGroupAccountAddArgs{
		orgSlug:   orgSlug,
		groupID:   groupID,
		accountID: accountID,
	}
}

func (c *Core) OrgGroupAccountAdd(ctx context.Context, args orgGroupAccountAddArgs) (*types.OrgGroupAccount, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	var (
		org     *types.Organization
		group   *types.OrgGroup
		account *types.Account
		err     error
	)

	if org, err = c.OrgGetOne(ctx,
		c.NewOrgGetOneArgs(nil, nil, &args.orgSlug),
	); err != nil {
		return nil, err
	}
	if group, err = c.OrgGroupGetOne(ctx,
		c.NewOrgGroupGetOneArgs(args.orgSlug, &args.groupID, nil),
	); err != nil {
		return nil, err
	}
	if group.OrgID != org.ID {
		return nil, types.ErrNotFound
	}

	if account, err = c.OrgAccountGetOne(ctx,
		c.NewOrgAccountGetOneArgs(args.orgSlug, &args.accountID, nil, nil),
	); err != nil {
		return nil, err
	}

	orgMismatchErr := errors.New("core.OrgGroupAccountAdd account must be org member")
	if account.OrgID == nil {
		return nil, orgMismatchErr
	}
	if *account.OrgID != org.ID {
		return nil, orgMismatchErr
	}

	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	return c.orgGroupRepo.AddAccount(ctx,
		org.ID,
		args.groupID,
		args.accountID,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) OrgGroupAccountGetAll(ctx context.Context,
	orgSlug string,
	groupID int64,
) ([]types.Account, error) {
	org, err := c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &orgSlug))
	if err != nil {
		return nil, err
	}

	return c.orgGroupRepo.GetAccounts(ctx, org.ID, groupID)
}

type orgGroupAccountsSetArgs struct {
	orgSlug    string
	groupID    int64
	accountIDs []int64
}

func (a *orgGroupAccountsSetArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("orgGroupAccountsSetArgs orgSlug cannot be empty")
	}
	if a.groupID < 1 {
		return errors.New("orgGroupAccountsSetArgs groupID must be a positive integer")
	}

	for _, id := range a.accountIDs {
		if id < 1 {
			return errors.New("orgGroupAccountsSetArgs accountIDs must be positive integers")
		}
	}

	return nil
}

func (c *Core) NewOrgGroupAccountsSetArgs(
	orgSlug string,
	groupID int64,
	accountIDs []int64,
) orgGroupAccountsSetArgs {
	return orgGroupAccountsSetArgs{
		orgSlug:    orgSlug,
		groupID:    groupID,
		accountIDs: accountIDs,
	}
}

func (c *Core) OrgGroupAccountsSet(ctx context.Context, args orgGroupAccountsSetArgs) ([]types.Account, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	var (
		org   *types.Organization
		group *types.OrgGroup
		err   error
	)
	if org, err = c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &args.orgSlug)); err != nil {
		return nil, err
	}
	if group, err = c.OrgGroupGetOne(ctx, c.NewOrgGroupGetOneArgs(args.orgSlug, &args.groupID, nil)); err != nil {
		return nil, err
	}
	if group.OrgID != org.ID {
		return nil, types.ErrNotFound
	}

	existingAccounts, err := c.OrgAccountGetManyByID(ctx, args.orgSlug, args.accountIDs)
	if err != nil {
		return nil, err
	}
	// Caller tried to add accounts that don't exist or don't belong to org
	if len(existingAccounts) != len(args.accountIDs) {
		return nil, types.ErrNotFound
	}

	orgMismatchErr := errors.New("core.OrgGroupAccountsSet account must be org member")
	for _, a := range existingAccounts {
		if a.OrgID == nil {
			return nil, orgMismatchErr
		}
		if *a.OrgID != org.ID {
			return nil, orgMismatchErr
		}
	}

	tracer, _ := c.getOperationTracer(ctx)
	return c.orgGroupRepo.UpdateAccounts(ctx, org.ID, args.groupID, args.accountIDs, tracer.AuthAccount.ID)
}

func (c *Core) OrgGroupAccountRemove(ctx context.Context,
	orgSlug string,
	groupID int64,
	accountID int64,
) error {
	org, err := c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &orgSlug))
	if err != nil {
		return err
	}

	return c.orgGroupRepo.RemoveAccount(ctx, org.ID, groupID, accountID)
}
