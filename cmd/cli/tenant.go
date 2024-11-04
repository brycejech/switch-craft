package cli

import (
	"fmt"
	"log"
	"switchcraft/core"
	"switchcraft/types"
	"time"

	"github.com/spf13/cobra"
)

func registerTenantModule(switchcraft *core.Core) {
	var tenantCmd = &cobra.Command{
		Use:   "tenant",
		Short: "SwitchCraft CLI tenant module",
	}
	rootCmd.AddCommand(tenantCmd)

	/* === CREATE TENANT COMMAND === */
	createTenantCmdArgs := struct {
		Name  string
		Slug  string
		Owner int64
	}{}
	var createTenantCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new tenant",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(switchcraft)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewTenantCreateArgs(
				createTenantCmdArgs.Name,
				createTenantCmdArgs.Slug,
				createTenantCmdArgs.Owner,
			)

			tenant, err := switchcraft.TenantCreate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(tenant)
		},
	}

	createTenantCmd.Flags().StringVar(&createTenantCmdArgs.Name, "name", "", "tenant.name")
	createTenantCmd.MarkFlagRequired("name")
	createTenantCmd.Flags().StringVar(&createTenantCmdArgs.Slug, "slug", "", "tenant.slug")
	createTenantCmd.MarkFlagRequired("slug")
	createTenantCmd.Flags().Int64Var(&createTenantCmdArgs.Owner, "owner", 0, "tenant.owner")
	createTenantCmd.MarkFlagRequired("owner")
	tenantCmd.AddCommand(createTenantCmd)

	/* === GET TENANTS COMMAND */
	var getTenantsCmd = &cobra.Command{
		Use:   "getMany",
		Short: "Get multiple tenants",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(switchcraft)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			tenants, err := switchcraft.TenantGetMany(opCtx)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(tenants)
		},
	}
	tenantCmd.AddCommand(getTenantsCmd)

	/* === GET TENANT COMMAND === */
	var getTenantID int64
	var getTenantCmd = &cobra.Command{
		Use:   "getOne",
		Short: "Get a single tenant",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(switchcraft)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			tenant, err := switchcraft.TenantGetOne(opCtx, getTenantID)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(tenant)
		},
	}

	getTenantCmd.Flags().Int64Var(&getTenantID, "id", 0, "tenant.id")
	getTenantCmd.MarkFlagRequired("id")
	tenantCmd.AddCommand(getTenantCmd)

	/* === UPDATE TENANT COMMAND === */
	updateTenantCmdArgs := struct {
		ID    int64
		Name  string
		Slug  string
		Owner int64
	}{}
	var updateTenantCmd = &cobra.Command{
		Use:   "update",
		Short: "Update a tenant",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(switchcraft)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewTenantUpdateArgs(
				updateTenantCmdArgs.ID,
				updateTenantCmdArgs.Name,
				updateTenantCmdArgs.Slug,
				updateTenantCmdArgs.Owner,
			)

			tenant, err := switchcraft.TenantUpdate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(tenant)
		},
	}

	updateTenantCmd.Flags().Int64Var(&updateTenantCmdArgs.ID, "id", 0, "tenant.id")
	updateTenantCmd.MarkFlagRequired("id")
	updateTenantCmd.Flags().StringVar(&updateTenantCmdArgs.Name, "name", "", "tenant.name")
	updateTenantCmd.MarkFlagRequired("name")
	updateTenantCmd.Flags().StringVar(&updateTenantCmdArgs.Slug, "slug", "", "tenant.slug")
	updateTenantCmd.MarkFlagRequired("slug")
	updateTenantCmd.Flags().Int64Var(&updateTenantCmdArgs.Owner, "owner", 0, "tenant.owner")
	updateTenantCmd.MarkFlagRequired("owner")
	tenantCmd.AddCommand(updateTenantCmd)

	/* === DELETE ACCOUNT CMD === */
	/* -------------------------- */
	var deleteTenantID int64
	var deleteTenantCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a tenant",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(switchcraft)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			if err := switchcraft.TenantDelete(opCtx, deleteTenantID); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Tenant '%v' deleted successfully\n", deleteTenantID)
		},
	}

	deleteTenantCmd.Flags().Int64Var(&deleteTenantID, "id", 0, "tenant.id")
	deleteTenantCmd.MarkFlagRequired("id")
	tenantCmd.AddCommand(deleteTenantCmd)
}
