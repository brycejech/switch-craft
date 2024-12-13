package cli

import (
	"fmt"
	"log"
	"switchcraft/core"
	"switchcraft/types"
	"time"

	"github.com/spf13/cobra"
)

func registerOrgGroupModule(core *core.Core) {
	var orgGroupCmd = &cobra.Command{
		Use:   "orgGroup",
		Short: "SwitchCraft CLI organization group module",
	}
	orgGroupGetManyCmd(core, orgGroupCmd)
	orgGroupCreateCmd(core, orgGroupCmd)
	orgGroupGetOneCmd(core, orgGroupCmd)
	orgGroupUpdateCmd(core, orgGroupCmd)
	orgGroupDeleteCmd(core, orgGroupCmd)

	rootCmd.AddCommand(orgGroupCmd)
}

func orgGroupGetManyCmd(core *core.Core, parentCmd *cobra.Command) {
	var orgSlug string
	getManyCmd := &cobra.Command{
		Use:   "getMany",
		Short: "Get multiple organization groups",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			groups, err := core.OrgGroupGetMany(opCtx, orgSlug)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(groups)
		},
	}
	getManyCmd.Flags().StringVar(&orgSlug, "orgSlug", "", "Organization slug")
	getManyCmd.MarkFlagRequired("orgSlug")

	parentCmd.AddCommand(getManyCmd)
}

func orgGroupCreateCmd(core *core.Core, parentCmd *cobra.Command) {
	args := struct {
		orgSlug     string
		name        string
		description string
	}{}
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create new organization group",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewOrgGroupCreateArgs(
				args.orgSlug,
				args.name,
				args.description,
			)

			group, err := core.OrgGroupCreate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(group)
		},
	}
	createCmd.Flags().StringVar(&args.orgSlug, "orgSlug", "", "Organization slug")
	createCmd.MarkFlagRequired("orgSlug")
	createCmd.Flags().StringVar(&args.name, "name", "", "group.name")
	createCmd.MarkFlagRequired("name")
	createCmd.Flags().StringVar(&args.description, "description", "", "group.description")

	parentCmd.AddCommand(createCmd)
}

func orgGroupGetOneCmd(core *core.Core, parentCmd *cobra.Command) {
	var orgSlug string
	var groupID int64
	var groupUUID string
	getOneCmd := &cobra.Command{
		Use:   "getOne",
		Short: "Get an organization group by id or uuid",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			var (
				id   *int64
				uuid *string
			)
			if cmd.Flags().Changed("id") {
				id = &groupID
			}
			if cmd.Flags().Changed("uuid") {
				uuid = &groupUUID
			}

			args := core.NewOrgGroupGetOneArgs(
				orgSlug,
				id,
				uuid,
			)

			group, err := core.OrgGroupGetOne(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(group)
		},
	}
	getOneCmd.Flags().StringVar(&orgSlug, "orgSlug", "", "Organization slug")
	getOneCmd.MarkFlagRequired("orgSlug")
	getOneCmd.Flags().Int64Var(&groupID, "id", 0, "group.id")
	getOneCmd.Flags().StringVar(&groupUUID, "uuid", "", "group.uuid")

	parentCmd.AddCommand(getOneCmd)
}

func orgGroupUpdateCmd(core *core.Core, parentCmd *cobra.Command) {
	args := struct {
		orgSlug     string
		id          int64
		name        string
		description string
	}{}
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing organization group",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewOrgGroupUpdateArgs(
				args.orgSlug,
				args.id,
				args.name,
				args.description,
			)

			group, err := core.OrgGroupUpdate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(group)
		},
	}
	updateCmd.Flags().StringVar(&args.orgSlug, "orgSlug", "", "Organization slug")
	updateCmd.MarkFlagRequired("orgSlug")
	updateCmd.Flags().Int64Var(&args.id, "id", 0, "group.id")
	updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().StringVar(&args.name, "name", "", "group.name")
	updateCmd.MarkFlagRequired("name")
	updateCmd.Flags().StringVar(&args.description, "description", "", "group.description")
	updateCmd.MarkFlagRequired("description")

	parentCmd.AddCommand(updateCmd)
}

func orgGroupDeleteCmd(core *core.Core, parentCmd *cobra.Command) {
	var orgSlug string
	var groupID int64
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an organization group",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			if err := core.OrgGroupDelete(opCtx, orgSlug, groupID); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Organization group '%v' deleted successfully\n", groupID)
		},
	}
	deleteCmd.Flags().StringVar(&orgSlug, "orgSlug", "", "Organization slug")
	deleteCmd.MarkFlagRequired("orgSlug")
	deleteCmd.Flags().Int64Var(&groupID, "id", 0, "group.id")
	deleteCmd.MarkFlagRequired("id")

	parentCmd.AddCommand(deleteCmd)
}
