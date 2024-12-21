package core

import (
	"context"
	"errors"
	"switchcraft/types"
)

type featFlagCreateArgs struct {
	orgSlug     string
	appSlug     string
	name        string
	label       string
	description string
	isEnabled   bool
}

func (a *featFlagCreateArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("featFlagCreateArgs.orgSlug cannot be empty")
	}
	if a.appSlug == "" {
		return errors.New("featFlagCreateArgs.appSlug cannot be empty")
	}
	if a.name == "" {
		return errors.New("featFlagCreateArgs.name cannot be empty")
	}
	return nil
}

func (c *Core) NewFeatFlagCreateArgs(
	orgSlug string,
	appSlug string,
	name string,
	label string,
	description string,
	isEnabled bool,
) featFlagCreateArgs {
	return featFlagCreateArgs{
		orgSlug:     orgSlug,
		appSlug:     appSlug,
		name:        name,
		label:       label,
		description: description,
		isEnabled:   isEnabled,
	}
}

func (c *Core) FeatFlagCreate(ctx context.Context, args featFlagCreateArgs) (*types.FeatureFlag, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	var (
		org *types.Organization
		app *types.Application
	)

	if org, err = c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &args.orgSlug)); err != nil {
		return nil, err
	}
	if app, err = c.AppGetOne(ctx, c.NewAppGetOneArgs(org.Slug, nil, nil, &args.appSlug)); err != nil {
		return nil, err
	}

	return c.featureFlagRepo.Create(ctx,
		org.ID,
		app.ID,
		args.name,
		args.label,
		args.description,
		args.isEnabled,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) FeatFlagGetMany(ctx context.Context,
	orgSlug string,
	appSlug string,
) ([]types.FeatureFlag, error) {
	var (
		org *types.Organization
		app *types.Application
		err error
	)

	if org, err = c.OrgGetOne(ctx,
		c.NewOrgGetOneArgs(nil, nil, &orgSlug),
	); err != nil {
		return nil, err
	}
	if app, err = c.AppGetOne(ctx,
		c.NewAppGetOneArgs(orgSlug, nil, nil, &appSlug),
	); err != nil {
		return nil, err
	}

	return c.featureFlagRepo.GetMany(ctx, org.ID, app.ID)
}

type featFlagGetOneArgs struct {
	orgSlug string
	appSlug string
	id      *int64
	uuid    *string
	name    *string
}

func (a *featFlagGetOneArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("featFlagGetOneArgs.orgSlug cannot be empty")
	}
	if a.appSlug == "" {
		return errors.New("featFlagGetOneArgs.appSlug cannot be empty")
	}
	if a.id == nil && a.uuid == nil && a.name == nil {
		return errors.New("featFlagGetOneArgs: must provide id, uuid, or name")
	}
	if a.id != nil && *a.id < 1 {
		return errors.New("featFlagGetOneArgs.id must be positive integer")
	}
	return nil
}

func (c *Core) NewFeatFlagGetOneArgs(
	orgSlug string,
	appSlug string,
	id *int64,
	uuid *string,
	name *string,
) featFlagGetOneArgs {
	return featFlagGetOneArgs{
		orgSlug: orgSlug,
		appSlug: appSlug,
		id:      id,
		uuid:    uuid,
		name:    name,
	}
}

func (c *Core) FeatFlagGetOne(ctx context.Context, args featFlagGetOneArgs) (*types.FeatureFlag, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	var (
		org *types.Organization
		app *types.Application
		err error
	)

	if org, err = c.OrgGetOne(ctx,
		c.NewOrgGetOneArgs(nil, nil, &args.orgSlug),
	); err != nil {
		return nil, err
	}
	if app, err = c.AppGetOne(ctx,
		c.NewAppGetOneArgs(org.Slug, nil, nil, &args.appSlug),
	); err != nil {
		return nil, err
	}

	return c.featureFlagRepo.GetOne(ctx,
		org.ID,
		app.ID,
		args.id,
		args.uuid,
		args.name,
	)
}

type featFlagUpdateArgs struct {
	orgSlug     string
	appSlug     string
	id          int64
	name        string
	label       string
	description string
	isEnabled   bool
}

func (a *featFlagUpdateArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("featFlagUpdateArgs.orgSlug cannot be empty")
	}
	if a.appSlug == "" {
		return errors.New("featFlagUpdateArgs.appSlug cannot be empty")
	}
	if a.id < 1 {
		return errors.New("featFlagUpdateArgs.id must be positive integer")
	}
	if a.name == "" {
		return errors.New("featFlagUpdateArgs.name cannot be empty")
	}
	return nil
}

func (c *Core) NewFeatFlagUpdateArgs(
	orgSlug string,
	appSlug string,
	id int64,
	name string,
	label string,
	description string,
	isEnabled bool,
) featFlagUpdateArgs {
	return featFlagUpdateArgs{
		orgSlug:     orgSlug,
		appSlug:     appSlug,
		id:          id,
		name:        name,
		label:       label,
		description: description,
		isEnabled:   isEnabled,
	}
}

func (c *Core) FeatFlagUpdate(ctx context.Context, args featFlagUpdateArgs) (*types.FeatureFlag, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	var (
		org *types.Organization
		app *types.Application
	)
	if org, err = c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &args.orgSlug)); err != nil {
		return nil, err
	}
	if app, err = c.AppGetOne(ctx, c.NewAppGetOneArgs(org.Slug, nil, nil, &args.appSlug)); err != nil {
		return nil, err
	}

	return c.featureFlagRepo.Update(ctx,
		org.ID,
		app.ID,
		args.id,
		args.name,
		args.label,
		args.description,
		args.isEnabled,
		tracer.AuthAccount.ID,
	)
}

func (c *Core) FeatFlagDelete(ctx context.Context,
	orgSlug string,
	appSlug string,
	id int64,
) error {
	var (
		org *types.Organization
		app *types.Application
		err error
	)

	if org, err = c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &orgSlug)); err != nil {
		return err
	}
	if app, err = c.AppGetOne(ctx, c.NewAppGetOneArgs(org.Slug, nil, nil, &appSlug)); err != nil {
		return err
	}

	return c.featureFlagRepo.Delete(ctx,
		org.ID,
		app.ID,
		id,
	)
}

type groupFlagCreateArgs struct {
	orgSlug   string
	groupID   int64
	appSlug   string
	flagID    int64
	isEnabled bool
}

func (a *groupFlagCreateArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("groupFlagCreateArgs.orgSlug cannot be empty")
	}
	if a.groupID < 1 {
		return errors.New("groupFlagCreateArgs.groupID must be positive integer")
	}
	if a.appSlug == "" {
		return errors.New("groupFlagCreateArgs.appSlug cannot be empty")
	}
	if a.flagID < 1 {
		return errors.New("groupFlagCreateArgs.flagID must be positive integer")
	}
	return nil
}

func (c *Core) NewGroupFlagCreateArgs(
	orgSlug string,
	groupID int64,
	appSlug string,
	flagID int64,
	isEnabled bool,
) groupFlagCreateArgs {
	return groupFlagCreateArgs{
		orgSlug:   orgSlug,
		groupID:   groupID,
		appSlug:   appSlug,
		flagID:    flagID,
		isEnabled: isEnabled,
	}
}

func (c *Core) GroupFlagCreate(ctx context.Context, args groupFlagCreateArgs) (*types.OrgGroupFeatureFlag, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	var (
		org *types.Organization
		app *types.Application
	)
	if org, err = c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &args.orgSlug)); err != nil {
		return nil, err
	}
	if app, err = c.AppGetOne(ctx, c.NewAppGetOneArgs(org.Slug, nil, nil, &args.appSlug)); err != nil {
		return nil, err
	}

	return c.featureFlagRepo.GroupFlagCreate(ctx,
		org.ID,
		args.groupID,
		app.ID,
		args.flagID,
		args.isEnabled,
		tracer.AuthAccount.ID,
	)
}

type groupFlagsGetByFlagIDArgs struct {
	orgSlug string
	appSlug string
	flagID  int64
}

func (a *groupFlagsGetByFlagIDArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("core.GroupFlagsGetByFlagID orgSlug cannot be empty")
	}
	if a.appSlug == "" {
		return errors.New("core.GroupFlagsGetByFlagID appSlug cannot be empty")
	}
	if a.flagID < 1 {
		return errors.New("core.GroupFlagsGetByFlagID flagID must be positive integer")
	}
	return nil
}

func (c *Core) NewGroupFlagsGetByFlagIDArgs(
	orgSlug string,
	appSlug string,
	flagID int64,
) groupFlagsGetByFlagIDArgs {
	return groupFlagsGetByFlagIDArgs{
		orgSlug: orgSlug,
		appSlug: appSlug,
		flagID:  flagID,
	}
}

func (c *Core) GroupFlagsGetByFlagID(ctx context.Context, args groupFlagsGetByFlagIDArgs) ([]types.OrgGroupFeatureFlag, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	var (
		org *types.Organization
		app *types.Application
		err error
	)
	if org, err = c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &args.orgSlug)); err != nil {
		return nil, err
	}
	if app, err = c.AppGetOne(ctx, c.NewAppGetOneArgs(org.Slug, nil, nil, &args.appSlug)); err != nil {
		return nil, err
	}

	return c.featureFlagRepo.GroupFlagsGetByFlagID(ctx, org.ID, app.ID, args.flagID)
}

type groupFlagGetOneArgs struct {
	orgSlug string
	groupID int64
	appSlug string
	flagID  int64
}

func (a *groupFlagGetOneArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("core.GroupFlagGetOne orgSlug cannot be empty")
	}
	if a.groupID < 1 {
		return errors.New("core.GroupFlagGetOne groupID must be positive integer")
	}
	if a.appSlug == "" {
		return errors.New("core.GroupFlagGetOne appSlug cannot be empty")
	}
	if a.flagID < 1 {
		return errors.New("core.GroupFlagGetOne flagID must be positive integer")
	}
	return nil
}

func (c *Core) NewGroupFlagGetOneArgs(
	orgSlug string,
	groupID int64,
	appSlug string,
	flagID int64,
) groupFlagGetOneArgs {
	return groupFlagGetOneArgs{
		orgSlug: orgSlug,
		groupID: groupID,
		appSlug: appSlug,
		flagID:  flagID,
	}
}

func (c *Core) GroupFlagGetOne(ctx context.Context, args groupFlagGetOneArgs) (*types.OrgGroupFeatureFlag, error) {
	if err := args.Validate(); err != nil {
		return nil, err
	}

	var (
		org *types.Organization
		app *types.Application
		err error
	)
	if org, err = c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &args.orgSlug)); err != nil {
		return nil, err
	}
	if app, err = c.AppGetOne(ctx, c.NewAppGetOneArgs(org.Slug, nil, nil, &args.appSlug)); err != nil {
		return nil, err
	}

	return c.featureFlagRepo.GroupFlagGetOne(ctx, org.ID, args.groupID, app.ID, args.flagID)
}

type groupFlagUpdateArgs struct {
	orgSlug   string
	groupID   int64
	appSlug   string
	flagID    int64
	isEnabled bool
}

func (a *groupFlagUpdateArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("groupFlagUpdateArgs.orgSlug cannot be empty")
	}
	if a.groupID < 1 {
		return errors.New("groupFlagUpdateArgs.groupID must be positive integer")
	}
	if a.appSlug == "" {
		return errors.New("groupFlagUpdateArgs.appSlug cannot be empty")
	}
	if a.flagID < 1 {
		return errors.New("groupFlagUpdateArgs.flagID must be positive integer")
	}
	return nil
}

func (c *Core) NewGroupFlagUpdateArgs(
	orgSlug string,
	groupID int64,
	appSlug string,
	flagID int64,
	isEnabled bool,
) groupFlagUpdateArgs {
	return groupFlagUpdateArgs{
		orgSlug:   orgSlug,
		groupID:   groupID,
		appSlug:   appSlug,
		flagID:    flagID,
		isEnabled: isEnabled,
	}
}

func (c *Core) GroupFlagUpdate(ctx context.Context, args groupFlagUpdateArgs) (*types.OrgGroupFeatureFlag, error) {
	tracer, err := c.getOperationTracer(ctx)
	if err != nil {
		return nil, err
	}

	if err := args.Validate(); err != nil {
		return nil, err
	}

	var (
		org *types.Organization
		app *types.Application
	)
	if org, err = c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &args.orgSlug)); err != nil {
		return nil, err
	}
	if app, err = c.AppGetOne(ctx, c.NewAppGetOneArgs(org.Slug, nil, nil, &args.appSlug)); err != nil {
		return nil, err
	}

	return c.featureFlagRepo.GroupFlagUpdate(ctx,
		org.ID,
		args.groupID,
		app.ID,
		args.flagID,
		args.isEnabled,
		tracer.AuthAccount.ID,
	)
}

type groupFlagDeleteArgs struct {
	orgSlug string
	groupID int64
	appSlug string
	flagID  int64
}

func (a *groupFlagDeleteArgs) Validate() error {
	if a.orgSlug == "" {
		return errors.New("core.GroupFlagDelete orgSlug cannot be empty")
	}
	if a.groupID < 1 {
		return errors.New("core.GroupFlagDelete groupID must be positive integer")
	}
	if a.appSlug == "" {
		return errors.New("core.GroupFlagDelete appSlug cannot be empty")
	}
	if a.flagID < 1 {
		return errors.New("core.GroupFlagDelete flagID must be positive integer")
	}
	return nil
}

func (c *Core) NewGroupFlagDeleteArgs(
	orgSlug string,
	groupID int64,
	appSlug string,
	flagID int64,
) groupFlagDeleteArgs {
	return groupFlagDeleteArgs{
		orgSlug: orgSlug,
		groupID: groupID,
		appSlug: appSlug,
		flagID:  flagID,
	}
}

func (c *Core) GroupFlagDelete(ctx context.Context, args groupFlagDeleteArgs) error {
	if err := args.Validate(); err != nil {
		return err
	}

	var (
		org *types.Organization
		app *types.Application
		err error
	)
	if org, err = c.OrgGetOne(ctx, c.NewOrgGetOneArgs(nil, nil, &args.orgSlug)); err != nil {
		return err
	}
	if app, err = c.AppGetOne(ctx, c.NewAppGetOneArgs(org.Slug, nil, nil, &args.appSlug)); err != nil {
		return err
	}

	return c.featureFlagRepo.GroupFlagDelete(ctx,
		org.ID,
		args.groupID,
		app.ID,
		args.flagID,
	)
}
