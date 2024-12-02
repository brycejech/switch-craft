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
	rootCmd.AddCommand(appCmd)

	/* -------------------------- */
	/* === CREATE APP COMMAND === */
	/* -------------------------- */
	createAppArgs := struct {
		orgSlug string
		name    string
		slug    string
	}{}
	var createAppCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new application",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewAppCreateArgs(
				createAppArgs.orgSlug,
				createAppArgs.name,
				createAppArgs.slug,
			)

			app, err := core.AppCreate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(app)
		},
	}
	createAppCmd.Flags().StringVar(&createAppArgs.orgSlug, "orgSlug", "", "Organization slug")
	createAppCmd.MarkFlagRequired("orgSlug")
	createAppCmd.Flags().StringVar(&createAppArgs.name, "name", "", "application.name")
	createAppCmd.MarkFlagRequired("name")
	createAppCmd.Flags().StringVar(&createAppArgs.slug, "slug", "", "application.slug")
	createAppCmd.MarkFlagRequired("slug")
	appCmd.AddCommand(createAppCmd)

	/* ------------------------ */
	/* === GET APPS COMMAND === */
	/* ------------------------ */
	var getAppsCmdOrgSlug string
	var getAppsCmd = &cobra.Command{
		Use:   "getMany",
		Short: " Get multiple applications",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			apps, err := core.AppGetMany(opCtx, getAppsCmdOrgSlug)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(apps)
		},
	}
	getAppsCmd.Flags().StringVar(&getAppsCmdOrgSlug, "orgSlug", "", "Organization slug")
	getAppsCmd.MarkFlagRequired("orgSlug")
	appCmd.AddCommand(getAppsCmd)

	/* ----------------------- */
	/* === GET APP COMMAND === */
	/* ----------------------- */
	var getAppCmdArgs = struct {
		orgSlug string
		id      int64
		uuid    string
		slug    string
	}{}
	var getAppCmd = &cobra.Command{
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
				id = &getAppCmdArgs.id
			}
			if cmd.Flags().Changed("uuid") {
				uuid = &getAppCmdArgs.uuid
			}
			if cmd.Flags().Changed("slug") {
				slug = &getAppCmdArgs.slug
			}

			args := core.NewAppGetOneArgs(
				getAppCmdArgs.orgSlug,
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
	getAppCmd.Flags().StringVar(&getAppCmdArgs.orgSlug, "orgSlug", "", "Organization slug")
	getAppCmd.MarkFlagRequired("orgSlug")
	getAppCmd.Flags().Int64Var(&getAppCmdArgs.id, "id", 0, "application.id")
	getAppCmd.Flags().StringVar(&getAppCmdArgs.uuid, "uuid", "", "application.uuid")
	getAppCmd.Flags().StringVar(&getAppCmdArgs.slug, "slug", "", "application.slug")
	appCmd.AddCommand(getAppCmd)

	/* -------------------------- */
	/* === UPDATE APP COMMAND === */
	/* -------------------------- */
	updateAppArgs := struct {
		orgSlug string
		id      int64
		name    string
		slug    string
	}{}
	var updateAppCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an existing application",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewAppUpdateArgs(
				updateAppArgs.orgSlug,
				updateAppArgs.id,
				updateAppArgs.name,
				updateAppArgs.slug,
			)

			app, err := core.AppUpdate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(app)
		},
	}
	updateAppCmd.Flags().StringVar(&updateAppArgs.orgSlug, "orgSlug", "", "application.orgSlug")
	updateAppCmd.MarkFlagRequired("orgSlug")
	updateAppCmd.Flags().Int64Var(&updateAppArgs.id, "id", 0, "application.id")
	updateAppCmd.MarkFlagRequired("id")
	updateAppCmd.Flags().StringVar(&updateAppArgs.name, "name", "", "application.name")
	updateAppCmd.MarkFlagRequired("name")
	updateAppCmd.Flags().StringVar(&updateAppArgs.slug, "slug", "", "application.slug")
	updateAppCmd.MarkFlagRequired("slug")
	appCmd.AddCommand(updateAppCmd)

	/* ---------------------- */
	/* === DELETE APP CMD === */
	/* ---------------------- */
	var deleteAppCmdOrgSlug string
	var deleteAppCmdSlug string
	var deleteAppCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an application",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			if err := core.AppDelete(opCtx, deleteAppCmdOrgSlug, deleteAppCmdSlug); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Application '%v' deleted successfully\n", deleteAppCmdSlug)
		},
	}
	deleteAppCmd.Flags().StringVar(&deleteAppCmdOrgSlug, "orgSlug", "", "application.orgSlug")
	deleteAppCmd.MarkFlagRequired("orgSlug")
	deleteAppCmd.Flags().StringVar(&deleteAppCmdSlug, "slug", "", "application.slug")
	deleteAppCmd.MarkFlagRequired("slug")
	appCmd.AddCommand(deleteAppCmd)
}
