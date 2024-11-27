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
	rootCmd.AddCommand(orgCmd)

	/* ----------------------------- */
	/* === CREATE ORG COMMAND === */
	/* ----------------------------- */
	createOrgCmdArgs := struct {
		Name  string
		Slug  string
		Owner int64
	}{}
	var createOrgCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new organization",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewOrgCreateArgs(
				createOrgCmdArgs.Name,
				createOrgCmdArgs.Slug,
				createOrgCmdArgs.Owner,
			)

			org, err := core.OrgCreate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(org)
		},
	}

	createOrgCmd.Flags().StringVar(&createOrgCmdArgs.Name, "name", "", "org.name")
	createOrgCmd.MarkFlagRequired("name")
	createOrgCmd.Flags().StringVar(&createOrgCmdArgs.Slug, "slug", "", "org.slug")
	createOrgCmd.MarkFlagRequired("slug")
	createOrgCmd.Flags().Int64Var(&createOrgCmdArgs.Owner, "owner", 0, "org.owner")
	createOrgCmd.MarkFlagRequired("owner")
	orgCmd.AddCommand(createOrgCmd)

	/* --------------------------- */
	/* === GET ORGS COMMAND === */
	/* --------------------------- */
	var getOrgsCmd = &cobra.Command{
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
	orgCmd.AddCommand(getOrgsCmd)

	/* -------------------------- */
	/* === GET ORG COMMAND === */
	/* -------------------------- */
	var getOrgID int64
	var getOrgUUID string
	var getOrgSlug string
	var getOrgCmd = &cobra.Command{
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
				id = &getOrgID
			}
			if cmd.Flags().Changed("uuid") {
				uuid = &getOrgUUID
			}
			if cmd.Flags().Changed("slug") {
				slug = &getOrgSlug
			}

			args := core.NewOrgGetOneArgs(id, uuid, slug)
			org, err := core.OrgGetOne(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(org)
		},
	}

	getOrgCmd.Flags().Int64Var(&getOrgID, "id", 0, "org.id")
	getOrgCmd.Flags().StringVar(&getOrgUUID, "uuid", "", "org.uuid")
	getOrgCmd.Flags().StringVar(&getOrgSlug, "slug", "", "org.slug")
	orgCmd.AddCommand(getOrgCmd)

	/* ----------------------------- */
	/* === UPDATE ORG COMMAND === */
	/* ----------------------------- */
	updateOrgCmdArgs := struct {
		ID    int64
		Name  string
		Slug  string
		Owner int64
	}{}
	var updateOrgCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an organization",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewOrgUpdateArgs(
				updateOrgCmdArgs.ID,
				updateOrgCmdArgs.Name,
				updateOrgCmdArgs.Slug,
				updateOrgCmdArgs.Owner,
			)

			org, err := core.OrgUpdate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(org)
		},
	}

	updateOrgCmd.Flags().Int64Var(&updateOrgCmdArgs.ID, "id", 0, "org.id")
	updateOrgCmd.MarkFlagRequired("id")
	updateOrgCmd.Flags().StringVar(&updateOrgCmdArgs.Name, "name", "", "org.name")
	updateOrgCmd.MarkFlagRequired("name")
	updateOrgCmd.Flags().StringVar(&updateOrgCmdArgs.Slug, "slug", "", "org.slug")
	updateOrgCmd.MarkFlagRequired("slug")
	updateOrgCmd.Flags().Int64Var(&updateOrgCmdArgs.Owner, "owner", 0, "org.owner")
	updateOrgCmd.MarkFlagRequired("owner")
	orgCmd.AddCommand(updateOrgCmd)

	/* -------------------------- */
	/* === DELETE ACCOUNT CMD === */
	/* -------------------------- */
	var deleteOrgID int64
	var deleteOrgCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an organization",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			if err := core.OrgDelete(opCtx, deleteOrgID); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Organization '%v' deleted successfully\n", deleteOrgID)
		},
	}

	deleteOrgCmd.Flags().Int64Var(&deleteOrgID, "id", 0, "org.id")
	deleteOrgCmd.MarkFlagRequired("id")
	orgCmd.AddCommand(deleteOrgCmd)
}
