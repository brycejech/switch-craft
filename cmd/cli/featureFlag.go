package cli

import (
	"fmt"
	"log"
	"switchcraft/core"
	"switchcraft/types"
	"time"

	"github.com/spf13/cobra"
)

func registerFeatureFlagModule(core *core.Core) {
	var featureFlagCmd = &cobra.Command{
		Use:   "featureFlag",
		Short: "SwitchCraft CLI feature flag module",
	}
	rootCmd.AddCommand(featureFlagCmd)

	/* ----------------------------------- */
	/* === CREATE FEATURE FLAG COMMAND === */
	/* ----------------------------------- */
	createFeatureFlagCmdArgs := struct {
		orgSlug     string
		appSlug     string
		name        string
		label       string
		description string
		isEnabled   bool
	}{}
	var createFeatureFlagCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new feature flag",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewFeatFlagCreateArgs(
				createFeatureFlagCmdArgs.orgSlug,
				createFeatureFlagCmdArgs.appSlug,
				createFeatureFlagCmdArgs.name,
				createFeatureFlagCmdArgs.label,
				createFeatureFlagCmdArgs.description,
				createFeatureFlagCmdArgs.isEnabled,
			)

			featureFlag, err := core.FeatFlagCreate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(featureFlag)
		},
	}
	createFeatureFlagCmd.Flags().StringVar(&createFeatureFlagCmdArgs.orgSlug, "orgSlug", "", "featureFlag.orgSlug")
	createFeatureFlagCmd.MarkFlagRequired("orgSlug")
	createFeatureFlagCmd.Flags().StringVar(&createFeatureFlagCmdArgs.appSlug, "applicationSlug", "", "featureFlag.applicationSlug")
	createFeatureFlagCmd.MarkFlagRequired("applicationSlug")
	createFeatureFlagCmd.Flags().StringVar(&createFeatureFlagCmdArgs.name, "name", "", "featureFlag.name")
	createFeatureFlagCmd.MarkFlagRequired("name")
	createFeatureFlagCmd.Flags().StringVar(&createFeatureFlagCmdArgs.label, "label", "", "featureFlag.label")
	createFeatureFlagCmd.MarkFlagRequired("label")
	createFeatureFlagCmd.Flags().StringVar(&createFeatureFlagCmdArgs.description, "description", "", "featureFlag.description")
	createFeatureFlagCmd.Flags().BoolVar(&createFeatureFlagCmdArgs.isEnabled, "isEnabled", false, "featureFlag.isEnabled")
	featureFlagCmd.AddCommand(createFeatureFlagCmd)

	/* --------------------------------- */
	/* === GET FEATURE FLAGS COMMAND === */
	/* --------------------------------- */
	var getFeatureFlagsCmdOrgSlug string
	var getFeatureFlagsCmdAppSlug string
	var getFeatureFlagsCmd = &cobra.Command{
		Use:   "getMany",
		Short: "Get multiple feature flags",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			featureFlags, err := core.FeatFlagGetMany(opCtx, getFeatureFlagsCmdOrgSlug, getFeatureFlagsCmdAppSlug)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(featureFlags)
		},
	}
	getFeatureFlagsCmd.Flags().StringVar(&getFeatureFlagsCmdOrgSlug, "orgSlug", "", "featureFlag.orgSlug")
	getFeatureFlagsCmd.MarkFlagRequired("orgSlug")
	getFeatureFlagsCmd.Flags().StringVar(&getFeatureFlagsCmdAppSlug, "applicationSlug", "", "featureFlag.applicationSlug")
	getFeatureFlagsCmd.MarkFlagRequired("applicationSlug")
	featureFlagCmd.AddCommand(getFeatureFlagsCmd)

	/* -------------------------------- */
	/* === GET FEATURE FLAG COMMAND === */
	/* -------------------------------- */
	var getFeatureFlagCmdArgs = struct {
		orgSlug string
		appSlug string
		id      int64
		uuid    string
		name    string
	}{}
	var getFeatureFlagCmd = &cobra.Command{
		Use:   "getOne",
		Short: "Get a single feature flag",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			var (
				id   *int64
				uuid *string
				name *string
			)
			if cmd.Flags().Changed("id") {
				id = &getFeatureFlagCmdArgs.id
			}
			if cmd.Flags().Changed("uuid") {
				uuid = &getFeatureFlagCmdArgs.uuid
			}
			if cmd.Flags().Changed("name") {
				name = &getFeatureFlagCmdArgs.name
			}

			args := core.NewFeatFlagGetOneArgs(
				getFeatureFlagCmdArgs.orgSlug,
				getFeatureFlagCmdArgs.appSlug,
				id,
				uuid,
				name,
			)

			featureFlag, err := core.FeatFlagGetOne(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(featureFlag)
		},
	}
	getFeatureFlagCmd.Flags().StringVar(&getFeatureFlagCmdArgs.orgSlug, "orgSlug", "", "featureFlag.orgSlug")
	getFeatureFlagCmd.MarkFlagRequired("orgSlug")
	getFeatureFlagCmd.Flags().StringVar(&getFeatureFlagCmdArgs.appSlug, "applicationSlug", "", "featureFlag.applicationSlug")
	getFeatureFlagCmd.MarkFlagRequired("applicationSlug")
	getFeatureFlagCmd.Flags().Int64Var(&getFeatureFlagCmdArgs.id, "id", 0, "featureFlag.id")
	getFeatureFlagCmd.Flags().StringVar(&getFeatureFlagCmdArgs.uuid, "uuid", "", "featureFlag.uuid")
	getFeatureFlagCmd.Flags().StringVar(&getFeatureFlagCmdArgs.name, "name", "", "featureFlag.name")
	featureFlagCmd.AddCommand(getFeatureFlagCmd)

	/* ----------------------------------- */
	/* === UPDATE FEATURE FLAG COMMAND === */
	/* ----------------------------------- */
	updateFeatureFlagCmdArgs := struct {
		orgSlug     string
		appSlug     string
		id          int64
		name        string
		label       string
		description string
		isEnabled   bool
	}{}
	var updateFeatureFlagCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an existing feature flag",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewFeatFlagUpdateArgs(
				updateFeatureFlagCmdArgs.orgSlug,
				updateFeatureFlagCmdArgs.appSlug,
				updateFeatureFlagCmdArgs.id,
				updateFeatureFlagCmdArgs.name,
				updateFeatureFlagCmdArgs.label,
				updateFeatureFlagCmdArgs.description,
				updateFeatureFlagCmdArgs.isEnabled,
			)

			featureFlag, err := core.FeatFlagUpdate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(featureFlag)
		},
	}
	updateFeatureFlagCmd.Flags().StringVar(&updateFeatureFlagCmdArgs.orgSlug, "orgSlug", "", "featureFlag.orgSlug")
	updateFeatureFlagCmd.MarkFlagRequired("orgSlug")
	updateFeatureFlagCmd.Flags().StringVar(&updateFeatureFlagCmdArgs.appSlug, "applicationSlug", "", "featureFlag.applicationSlug")
	updateFeatureFlagCmd.Flags().Int64Var(&updateFeatureFlagCmdArgs.id, "id", 0, "featureFlag.id")
	updateFeatureFlagCmd.MarkFlagRequired("id")
	updateFeatureFlagCmd.Flags().StringVar(&updateFeatureFlagCmdArgs.name, "name", "", "featureFlag.name")
	updateFeatureFlagCmd.MarkFlagRequired("name")
	updateFeatureFlagCmd.Flags().StringVar(&updateFeatureFlagCmdArgs.label, "label", "", "featureFlag.label")
	updateFeatureFlagCmd.MarkFlagRequired("label")
	updateFeatureFlagCmd.Flags().StringVar(&updateFeatureFlagCmdArgs.description, "description", "", "featureFlag.description")
	updateFeatureFlagCmd.Flags().BoolVar(&updateFeatureFlagCmdArgs.isEnabled, "isEnabled", false, "featureFlag.isEnabled")
	updateFeatureFlagCmd.MarkFlagRequired("isEnabled")
	featureFlagCmd.AddCommand(updateFeatureFlagCmd)

	/* ------------------------------- */
	/* === DELETE FEATURE FLAG CMD === */
	/* ------------------------------- */
	var deleteFeatureFlagCmdOrgSlug string
	var deleteFeatureFlagCmdApplicationSlug string
	var deleteFeatureFlagCmdID int64
	var deleteFeatureFlagCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a feature flag",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			if err := core.FeatFlagDelete(opCtx,
				deleteFeatureFlagCmdOrgSlug,
				deleteFeatureFlagCmdApplicationSlug,
				deleteFeatureFlagCmdID,
			); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Feature flag '%v' deleted successfully\n", deleteFeatureFlagCmdID)
		},
	}
	deleteFeatureFlagCmd.Flags().StringVar(&deleteFeatureFlagCmdOrgSlug, "orgSlug", "", "featureFlag.orgSlug")
	deleteFeatureFlagCmd.MarkFlagRequired("orgSlug")
	deleteFeatureFlagCmd.Flags().StringVar(&deleteFeatureFlagCmdApplicationSlug, "applicationSlug", "", "featureFlag.applicationSlug")
	deleteFeatureFlagCmd.MarkFlagRequired("applicationSlug")
	deleteFeatureFlagCmd.Flags().Int64Var(&deleteFeatureFlagCmdID, "id", 0, "featureFlag.id")
	deleteFeatureFlagCmd.MarkFlagRequired("id")
	featureFlagCmd.AddCommand(deleteFeatureFlagCmd)
}
