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
		orgID     int64
		appID     int64
		name      string
		isEnabled bool
	}{}
	var createFeatureFlagCmd = &cobra.Command{
		Use:   "create",
		Short: "Create a new feature flag",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewFeatureFlagCreateArgs(
				createFeatureFlagCmdArgs.orgID,
				createFeatureFlagCmdArgs.appID,
				createFeatureFlagCmdArgs.name,
				createFeatureFlagCmdArgs.isEnabled,
			)

			featureFlag, err := core.FeatureFlagCreate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(featureFlag)
		},
	}
	createFeatureFlagCmd.Flags().Int64Var(&createFeatureFlagCmdArgs.orgID, "orgId", 0, "featureFlag.orgId")
	createFeatureFlagCmd.MarkFlagRequired("orgId")
	createFeatureFlagCmd.Flags().Int64Var(&createFeatureFlagCmdArgs.appID, "applicationId", 0, "featureFlag.applicationId")
	createFeatureFlagCmd.MarkFlagRequired("applicationId")
	createFeatureFlagCmd.Flags().StringVar(&createFeatureFlagCmdArgs.name, "name", "", "featureFlag.name")
	createFeatureFlagCmd.MarkFlagRequired("name")
	createFeatureFlagCmd.Flags().BoolVar(&createFeatureFlagCmdArgs.isEnabled, "isEnabled", false, "featureFlag.isEnabled")
	createFeatureFlagCmd.MarkFlagRequired("isEnabled")
	featureFlagCmd.AddCommand(createFeatureFlagCmd)

	/* --------------------------------- */
	/* === GET FEATURE FLAGS COMMAND === */
	/* --------------------------------- */
	var getFeatureFlagsCmdOrgID int64
	var getFeatureFlagsCmdAppID int64
	var getFeatureFlagsCmd = &cobra.Command{
		Use:   "getMany",
		Short: "Get multiple feature flags",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			featureFlags, err := core.FeatureFlagGetMany(opCtx, getFeatureFlagsCmdOrgID, getFeatureFlagsCmdAppID)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(featureFlags)
		},
	}
	getFeatureFlagsCmd.Flags().Int64Var(&getFeatureFlagsCmdOrgID, "orgId", 0, "featureFlag.orgId")
	getFeatureFlagsCmd.MarkFlagRequired("orgId")
	getFeatureFlagsCmd.Flags().Int64Var(&getFeatureFlagsCmdAppID, "applicationId", 0, "featureFlag.applicationId")
	getFeatureFlagsCmd.MarkFlagRequired("applicationId")
	featureFlagCmd.AddCommand(getFeatureFlagsCmd)

	/* -------------------------------- */
	/* === GET FEATURE FLAG COMMAND === */
	/* -------------------------------- */
	var getFeatureFlagCmdArgs = struct {
		orgID int64
		appID int64
		id    int64
		uuid  string
		name  string
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

			args := core.NewFeatureFlagGetOneArgs(
				getFeatureFlagCmdArgs.orgID,
				getFeatureFlagCmdArgs.appID,
				id,
				uuid,
				name,
			)

			featureFlag, err := core.FeatureFlagGetOne(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(featureFlag)
		},
	}
	getFeatureFlagCmd.Flags().Int64Var(&getFeatureFlagCmdArgs.orgID, "orgId", 0, "featureFlag.orgId")
	getFeatureFlagCmd.MarkFlagRequired("orgId")
	getFeatureFlagCmd.Flags().Int64Var(&getFeatureFlagCmdArgs.appID, "applicationId", 0, "featureFlag.applicationId")
	getFeatureFlagCmd.MarkFlagRequired("applicationId")
	getFeatureFlagCmd.Flags().Int64Var(&getFeatureFlagCmdArgs.id, "id", 0, "featureFlag.id")
	getFeatureFlagCmd.Flags().StringVar(&getFeatureFlagCmdArgs.uuid, "uuid", "", "featureFlag.uuid")
	getFeatureFlagCmd.Flags().StringVar(&getFeatureFlagCmdArgs.name, "name", "", "featureFlag.name")
	featureFlagCmd.AddCommand(getFeatureFlagCmd)

	/* ----------------------------------- */
	/* === UPDATE FEATURE FLAG COMMAND === */
	/* ----------------------------------- */
	updateFeatureFlagCmdArgs := struct {
		orgID     int64
		appID     int64
		id        int64
		name      string
		isEnabled bool
	}{}
	var updateFeatureFlagCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an existing feature flag",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewFeatureFlagUpdateArgs(
				updateFeatureFlagCmdArgs.orgID,
				updateFeatureFlagCmdArgs.appID,
				updateFeatureFlagCmdArgs.id,
				updateFeatureFlagCmdArgs.name,
				updateFeatureFlagCmdArgs.isEnabled,
			)

			featureFlag, err := core.FeatureFlagUpdate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(featureFlag)
		},
	}
	updateFeatureFlagCmd.Flags().Int64Var(&updateFeatureFlagCmdArgs.orgID, "orgId", 0, "featureFlag.orgId")
	updateFeatureFlagCmd.MarkFlagRequired("orgId")
	updateFeatureFlagCmd.Flags().Int64Var(&updateFeatureFlagCmdArgs.appID, "applicationId", 0, "featureFlag.applicationId")
	updateFeatureFlagCmd.Flags().Int64Var(&updateFeatureFlagCmdArgs.id, "id", 0, "featureFlag.id")
	updateFeatureFlagCmd.MarkFlagRequired("id")
	updateFeatureFlagCmd.Flags().StringVar(&updateFeatureFlagCmdArgs.name, "name", "", "featureFlag.name")
	updateFeatureFlagCmd.MarkFlagRequired("name")
	updateFeatureFlagCmd.Flags().BoolVar(&updateFeatureFlagCmdArgs.isEnabled, "isEnabled", false, "featureFlag.isEnabled")
	updateFeatureFlagCmd.MarkFlagRequired("isEnabled")
	featureFlagCmd.AddCommand(updateFeatureFlagCmd)

	/* ------------------------------- */
	/* === DELETE FEATURE FLAG CMD === */
	/* ------------------------------- */
	var deleteFeatureFlagCmdOrgID int64
	var deleteFeatureFlagCmdApplicationID int64
	var deleteFeatureFlagCmdID int64
	var deleteFeatureFlagCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete a feature flag",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			if err := core.FeatureFlagDelete(opCtx,
				deleteFeatureFlagCmdOrgID,
				deleteFeatureFlagCmdApplicationID,
				deleteFeatureFlagCmdID,
			); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Feature flag '%v' deleted successfully\n", deleteFeatureFlagCmdID)
		},
	}
	deleteFeatureFlagCmd.Flags().Int64Var(&deleteFeatureFlagCmdOrgID, "orgId", 0, "featureFlag.orgId")
	deleteFeatureFlagCmd.MarkFlagRequired("orgId")
	deleteFeatureFlagCmd.Flags().Int64Var(&deleteFeatureFlagCmdApplicationID, "applicationId", 0, "featureFlag.applicationId")
	deleteFeatureFlagCmd.MarkFlagRequired("applicationId")
	deleteFeatureFlagCmd.Flags().Int64Var(&deleteFeatureFlagCmdID, "id", 0, "featureFlag.id")
	deleteFeatureFlagCmd.MarkFlagRequired("id")
	featureFlagCmd.AddCommand(deleteFeatureFlagCmd)
}
