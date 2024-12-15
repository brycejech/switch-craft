package cli

import (
	"fmt"
	"log"
	"switchcraft/core"
	"switchcraft/types"
	"time"

	"github.com/spf13/cobra"
)

func registerOrgModule(core *core.Core) {
	var orgCmd = &cobra.Command{
		Use:   "organization",
		Short: "SwitchCraft CLI org module",
	}
	orgCreateCmd(core, orgCmd)
	orgGetManyCmd(core, orgCmd)
	orgGetOneCmd(core, orgCmd)
	orgUpdateCmd(core, orgCmd)
	orgDeleteCmd(core, orgCmd)

	rootCmd.AddCommand(orgCmd)
}

func orgCreateCmd(core *core.Core, parentCmd *cobra.Command) {
	args := struct {
		name  string
		slug  string
		owner int64
	}{}
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new organization",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			org, err := core.OrgCreate(opCtx,
				core.NewOrgCreateArgs(
					args.name,
					args.slug,
					args.owner,
				),
			)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(org)
		},
	}
	createCmd.Flags().StringVar(&args.name, "name", "", "org.name")
	createCmd.MarkFlagRequired("name")
	createCmd.Flags().StringVar(&args.slug, "slug", "", "org.slug")
	createCmd.MarkFlagRequired("slug")
	createCmd.Flags().Int64Var(&args.owner, "owner", 0, "org.owner")
	createCmd.MarkFlagRequired("owner")

	parentCmd.AddCommand(createCmd)
}

func orgGetManyCmd(core *core.Core, parentCmd *cobra.Command) {
	getMayCmd := &cobra.Command{
		Use:   "getMany",
		Short: "Get multiple organizations",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			orgs, err := core.OrgGetMany(opCtx)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(orgs)
		},
	}

	parentCmd.AddCommand(getMayCmd)
}

func orgGetOneCmd(core *core.Core, parentCmd *cobra.Command) {
	var orgID int64
	var orgUUID string
	var orgSlug string
	getOneCmd := &cobra.Command{
		Use:   "getOne",
		Short: "Get a single organization",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			var (
				id   *int64
				uuid *string
				slug *string
			)

			if cmd.Flags().Changed("id") {
				id = &orgID
			}
			if cmd.Flags().Changed("uuid") {
				uuid = &orgUUID
			}
			if cmd.Flags().Changed("slug") {
				slug = &orgSlug
			}

			org, err := core.OrgGetOne(opCtx,
				core.NewOrgGetOneArgs(id, uuid, slug),
			)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(org)
		},
	}
	getOneCmd.Flags().Int64Var(&orgID, "id", 0, "org.id")
	getOneCmd.Flags().StringVar(&orgUUID, "uuid", "", "org.uuid")
	getOneCmd.Flags().StringVar(&orgSlug, "slug", "", "org.slug")

	parentCmd.AddCommand(getOneCmd)
}

func orgUpdateCmd(core *core.Core, parentCmd *cobra.Command) {
	args := struct {
		id    int64
		name  string
		slug  string
		owner int64
	}{}
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update an organization",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			org, err := core.OrgUpdate(opCtx,
				core.NewOrgUpdateArgs(
					args.id,
					args.name,
					args.slug,
					args.owner,
				),
			)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(org)
		},
	}
	updateCmd.Flags().Int64Var(&args.id, "id", 0, "org.id")
	updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().StringVar(&args.name, "name", "", "org.name")
	updateCmd.MarkFlagRequired("name")
	updateCmd.Flags().StringVar(&args.slug, "slug", "", "org.slug")
	updateCmd.MarkFlagRequired("slug")
	updateCmd.Flags().Int64Var(&args.owner, "owner", 0, "org.owner")
	updateCmd.MarkFlagRequired("owner")

	parentCmd.AddCommand(updateCmd)
}

func orgDeleteCmd(core *core.Core, parentCmd *cobra.Command) {
	var id int64
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an organization",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			if err := core.OrgDelete(opCtx, id); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Organization '%v' deleted successfully\n", id)
		},
	}
	deleteCmd.Flags().Int64Var(&id, "id", 0, "org.id")
	deleteCmd.MarkFlagRequired("id")

	parentCmd.AddCommand(deleteCmd)
}
