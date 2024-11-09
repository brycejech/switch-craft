package cli

import (
	"fmt"
	"log"
	"switchcraft/core"
	"switchcraft/types"
	"time"

	"github.com/spf13/cobra"
)

func registerAppModule(switchcraft *core.Core) {
	var appCmd = &cobra.Command{
		Use:   "application",
		Short: "SwitchCraft CLI application module",
	}
	rootCmd.AddCommand(appCmd)

	/* -------------------------- */
	/* === CREATE APP COMMAND === */
	/* -------------------------- */
	createAppArgs := struct {
		tenantID int64
		name     string
		slug     string
	}{}
	var createAppCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new application",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(switchcraft)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewAppCreateArgs(
				createAppArgs.tenantID,
				createAppArgs.name,
				createAppArgs.slug,
			)

			app, err := switchcraft.AppCreate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(app)
		},
	}
	createAppCmd.Flags().Int64Var(&createAppArgs.tenantID, "tenantId", 0, "application.tenantId")
	createAppCmd.MarkFlagRequired("tenantId")
	createAppCmd.Flags().StringVar(&createAppArgs.name, "name", "", "application.name")
	createAppCmd.MarkFlagRequired("name")
	createAppCmd.Flags().StringVar(&createAppArgs.slug, "slug", "", "application.slug")
	createAppCmd.MarkFlagRequired("slug")
	appCmd.AddCommand(createAppCmd)

	/* ------------------------ */
	/* === GET APPS COMMAND === */
	/* ------------------------ */
	var getAppsCmdTenantID int64
	var getAppsCmd = &cobra.Command{
		Use:   "getMany",
		Short: " Get multiple applications",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(switchcraft)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			apps, err := switchcraft.AppGetMany(opCtx, getAppsCmdTenantID)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(apps)
		},
	}
	getAppsCmd.Flags().Int64Var(&getAppsCmdTenantID, "tenantId", 0, "application.tenantId")
	getAppsCmd.MarkFlagRequired("tenantId")
	appCmd.AddCommand(getAppsCmd)

	/* ----------------------- */
	/* === GET APP COMMAND === */
	/* ----------------------- */
	var getAppCmdArgs = struct {
		tenantID int64
		id       int64
		uuid     string
		slug     string
	}{}
	var getAppCmd = &cobra.Command{
		Use:   "getOne",
		Short: " Get a single application",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(switchcraft)
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
				getAppCmdArgs.tenantID,
				id,
				uuid,
				slug,
			)

			app, err := switchcraft.AppGetOne(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(app)
		},
	}
	getAppCmd.Flags().Int64Var(&getAppCmdArgs.tenantID, "tenantId", 0, "application.tenantId")
	getAppCmd.MarkFlagRequired("tenantId")
	getAppCmd.Flags().Int64Var(&getAppCmdArgs.id, "id", 0, "application.id")
	getAppCmd.Flags().StringVar(&getAppCmdArgs.uuid, "uuid", "", "application.uuid")
	getAppCmd.Flags().StringVar(&getAppCmdArgs.slug, "slug", "", "application.slug")
	appCmd.AddCommand(getAppCmd)

	/* -------------------------- */
	/* === UPDATE APP COMMAND === */
	/* -------------------------- */
	updateAppArgs := struct {
		tenantID int64
		id       int64
		name     string
		slug     string
	}{}
	var updateAppCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an existing application",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(switchcraft)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewAppUpdateArgs(
				updateAppArgs.tenantID,
				updateAppArgs.id,
				updateAppArgs.name,
				updateAppArgs.slug,
			)

			app, err := switchcraft.AppUpdate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(app)
		},
	}
	updateAppCmd.Flags().Int64Var(&updateAppArgs.tenantID, "tenantId", 0, "application.tenantId")
	updateAppCmd.MarkFlagRequired("tenantId")
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
	var deleteAppCmdTenantID int64
	var deleteAppCmdID int64
	var deleteAppCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an application",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(switchcraft)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			if err := switchcraft.AppDelete(opCtx, deleteAppCmdTenantID, deleteAppCmdID); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Application '%v' deleted successfully\n", deleteAppCmdID)
		},
	}
	deleteAppCmd.Flags().Int64Var(&deleteAppCmdTenantID, "tenantId", 0, "application.tenantId")
	deleteAppCmd.MarkFlagRequired("tenantId")
	deleteAppCmd.Flags().Int64Var(&deleteAppCmdID, "id", 0, "application.id")
	deleteAppCmd.MarkFlagRequired("id")
	appCmd.AddCommand(deleteAppCmd)
}
