package cli

import (
	"fmt"
	"log"
	"switchcraft/core"
	"switchcraft/types"
	"time"

	"github.com/spf13/cobra"
)

func registerAppModule(core *core.Core) {
	var appCmd = &cobra.Command{
		Use:   "application",
		Short: "SwitchCraft CLI application module",
	}
	appCreateCmd(core, appCmd)
	appGetManyCmd(core, appCmd)
	appGetOneCmd(core, appCmd)
	appUpdateCmd(core, appCmd)
	appDeleteCmd(core, appCmd)

	rootCmd.AddCommand(appCmd)
}

func appCreateCmd(core *core.Core, parentCmd *cobra.Command) {
	args := struct {
		orgSlug string
		name    string
		slug    string
	}{}
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new application",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			app, err := core.AppCreate(opCtx,
				core.NewAppCreateArgs(
					args.orgSlug,
					args.name,
					args.slug,
				),
			)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(app)
		},
	}
	createCmd.Flags().StringVar(&args.orgSlug, "orgSlug", "", "Organization slug")
	createCmd.MarkFlagRequired("orgSlug")
	createCmd.Flags().StringVar(&args.name, "name", "", "application.name")
	createCmd.MarkFlagRequired("name")
	createCmd.Flags().StringVar(&args.slug, "slug", "", "application.slug")
	createCmd.MarkFlagRequired("slug")

	parentCmd.AddCommand(createCmd)
}

func appGetManyCmd(core *core.Core, parentCmd *cobra.Command) {
	var orgSlug string
	getManyCmd := &cobra.Command{
		Use:   "getMany",
		Short: " Get multiple applications",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			apps, err := core.AppGetMany(opCtx, orgSlug)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(apps)
		},
	}
	getManyCmd.Flags().StringVar(&orgSlug, "orgSlug", "", "Organization slug")
	getManyCmd.MarkFlagRequired("orgSlug")

	parentCmd.AddCommand(getManyCmd)
}

func appGetOneCmd(core *core.Core, parentCmd *cobra.Command) {
	var args = struct {
		orgSlug string
		id      int64
		uuid    string
		slug    string
	}{}
	getOneCmd := &cobra.Command{
		Use:   "getOne",
		Short: " Get a single application",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			var (
				id   *int64
				uuid *string
				slug *string
			)
			if cmd.Flags().Changed("id") {
				id = &args.id
			}
			if cmd.Flags().Changed("uuid") {
				uuid = &args.uuid
			}
			if cmd.Flags().Changed("slug") {
				slug = &args.slug
			}

			args := core.NewAppGetOneArgs(
				args.orgSlug,
				id,
				uuid,
				slug,
			)

			app, err := core.AppGetOne(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(app)
		},
	}
	getOneCmd.Flags().StringVar(&args.orgSlug, "orgSlug", "", "Organization slug")
	getOneCmd.MarkFlagRequired("orgSlug")
	getOneCmd.Flags().Int64Var(&args.id, "id", 0, "application.id")
	getOneCmd.Flags().StringVar(&args.uuid, "uuid", "", "application.uuid")
	getOneCmd.Flags().StringVar(&args.slug, "slug", "", "application.slug")

	parentCmd.AddCommand(getOneCmd)
}

func appUpdateCmd(core *core.Core, parentCmd *cobra.Command) {
	args := struct {
		orgSlug string
		id      int64
		name    string
		slug    string
	}{}
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing application",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			app, err := core.AppUpdate(opCtx,
				core.NewAppUpdateArgs(
					args.orgSlug,
					args.id,
					args.name,
					args.slug,
				),
			)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(app)
		},
	}
	updateCmd.Flags().StringVar(&args.orgSlug, "orgSlug", "", "Organization slug")
	updateCmd.MarkFlagRequired("orgSlug")
	updateCmd.Flags().Int64Var(&args.id, "id", 0, "application.id")
	updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().StringVar(&args.name, "name", "", "application.name")
	updateCmd.MarkFlagRequired("name")
	updateCmd.Flags().StringVar(&args.slug, "slug", "", "application.slug")
	updateCmd.MarkFlagRequired("slug")

	parentCmd.AddCommand(updateCmd)
}

func appDeleteCmd(core *core.Core, parentCmd *cobra.Command) {
	var orgSlug string
	var appSlug string
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an application",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			if err := core.AppDelete(opCtx, orgSlug, appSlug); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Application '%v' deleted successfully\n", appSlug)
		},
	}
	deleteCmd.Flags().StringVar(&orgSlug, "orgSlug", "", "Organization slug")
	deleteCmd.MarkFlagRequired("orgSlug")
	deleteCmd.Flags().StringVar(&appSlug, "slug", "", "application.slug")
	deleteCmd.MarkFlagRequired("slug")

	parentCmd.AddCommand(deleteCmd)
}
