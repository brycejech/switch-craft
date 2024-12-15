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
	featureFlagCreateCmd(core, featureFlagCmd)
	featureFlagGetManyCmd(core, featureFlagCmd)
	featureFlagGetOneCmd(core, featureFlagCmd)
	featureFlagUpdateCmd(core, featureFlagCmd)
	featureFlagDeleteCmd(core, featureFlagCmd)

	rootCmd.AddCommand(featureFlagCmd)
}

func featureFlagCreateCmd(core *core.Core, parentCmd *cobra.Command) {
	args := struct {
		orgSlug     string
		appSlug     string
		name        string
		label       string
		description string
		isEnabled   bool
	}{}
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new feature flag",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			featureFlag, err := core.FeatFlagCreate(opCtx,
				core.NewFeatFlagCreateArgs(
					args.orgSlug,
					args.appSlug,
					args.name,
					args.label,
					args.description,
					args.isEnabled,
				),
			)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(featureFlag)
		},
	}
	createCmd.Flags().StringVar(&args.orgSlug, "orgSlug", "", "Organization slug")
	createCmd.MarkFlagRequired("orgSlug")
	createCmd.Flags().StringVar(&args.appSlug, "applicationSlug", "", "featureFlag.applicationSlug")
	createCmd.MarkFlagRequired("applicationSlug")
	createCmd.Flags().StringVar(&args.name, "name", "", "featureFlag.name")
	createCmd.MarkFlagRequired("name")
	createCmd.Flags().StringVar(&args.label, "label", "", "featureFlag.label")
	createCmd.MarkFlagRequired("label")
	createCmd.Flags().StringVar(&args.description, "description", "", "featureFlag.description")
	createCmd.Flags().BoolVar(&args.isEnabled, "isEnabled", false, "featureFlag.isEnabled")

	parentCmd.AddCommand(createCmd)
}

func featureFlagGetManyCmd(core *core.Core, parentCmd *cobra.Command) {
	var orgSlug string
	var appSlug string
	getManyCmd := &cobra.Command{
		Use:   "getMany",
		Short: "Get multiple feature flags",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			featureFlags, err := core.FeatFlagGetMany(opCtx, orgSlug, appSlug)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(featureFlags)
		},
	}
	getManyCmd.Flags().StringVar(&orgSlug, "orgSlug", "", "Organization slug")
	getManyCmd.MarkFlagRequired("orgSlug")
	getManyCmd.Flags().StringVar(&appSlug, "applicationSlug", "", "featureFlag.applicationSlug")
	getManyCmd.MarkFlagRequired("applicationSlug")

	parentCmd.AddCommand(getManyCmd)
}

func featureFlagGetOneCmd(core *core.Core, parentCmd *cobra.Command) {
	var args = struct {
		orgSlug string
		appSlug string
		id      int64
		uuid    string
		name    string
	}{}
	getOneCmd := &cobra.Command{
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
				id = &args.id
			}
			if cmd.Flags().Changed("uuid") {
				uuid = &args.uuid
			}
			if cmd.Flags().Changed("name") {
				name = &args.name
			}

			featureFlag, err := core.FeatFlagGetOne(opCtx,
				core.NewFeatFlagGetOneArgs(
					args.orgSlug,
					args.appSlug,
					id,
					uuid,
					name,
				),
			)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(featureFlag)
		},
	}
	getOneCmd.Flags().StringVar(&args.orgSlug, "orgSlug", "", "Organization slug")
	getOneCmd.MarkFlagRequired("orgSlug")
	getOneCmd.Flags().StringVar(&args.appSlug, "applicationSlug", "", "featureFlag.applicationSlug")
	getOneCmd.MarkFlagRequired("applicationSlug")
	getOneCmd.Flags().Int64Var(&args.id, "id", 0, "featureFlag.id")
	getOneCmd.Flags().StringVar(&args.uuid, "uuid", "", "featureFlag.uuid")
	getOneCmd.Flags().StringVar(&args.name, "name", "", "featureFlag.name")

	parentCmd.AddCommand(getOneCmd)
}

func featureFlagUpdateCmd(core *core.Core, parentCmd *cobra.Command) {
	args := struct {
		orgSlug     string
		appSlug     string
		id          int64
		name        string
		label       string
		description string
		isEnabled   bool
	}{}
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing feature flag",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			featureFlag, err := core.FeatFlagUpdate(opCtx,
				core.NewFeatFlagUpdateArgs(
					args.orgSlug,
					args.appSlug,
					args.id,
					args.name,
					args.label,
					args.description,
					args.isEnabled,
				),
			)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(featureFlag)
		},
	}
	updateCmd.Flags().StringVar(&args.orgSlug, "orgSlug", "", "Organization slug")
	updateCmd.MarkFlagRequired("orgSlug")
	updateCmd.Flags().StringVar(&args.appSlug, "applicationSlug", "", "featureFlag.applicationSlug")
	updateCmd.Flags().Int64Var(&args.id, "id", 0, "featureFlag.id")
	updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().StringVar(&args.name, "name", "", "featureFlag.name")
	updateCmd.MarkFlagRequired("name")
	updateCmd.Flags().StringVar(&args.label, "label", "", "featureFlag.label")
	updateCmd.MarkFlagRequired("label")
	updateCmd.Flags().StringVar(&args.description, "description", "", "featureFlag.description")
	updateCmd.Flags().BoolVar(&args.isEnabled, "isEnabled", false, "featureFlag.isEnabled")
	updateCmd.MarkFlagRequired("isEnabled")

	parentCmd.AddCommand(updateCmd)
}

func featureFlagDeleteCmd(core *core.Core, parentCmd *cobra.Command) {
	var orgSlug string
	var appSlug string
	var id int64
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a feature flag",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			err := core.FeatFlagDelete(opCtx, orgSlug, appSlug, id)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Feature flag '%v' deleted successfully\n", id)
		},
	}
	deleteCmd.Flags().StringVar(&orgSlug, "orgSlug", "", "Organization slug")
	deleteCmd.MarkFlagRequired("orgSlug")
	deleteCmd.Flags().StringVar(&appSlug, "applicationSlug", "", "featureFlag.applicationSlug")
	deleteCmd.MarkFlagRequired("applicationSlug")
	deleteCmd.Flags().Int64Var(&id, "id", 0, "featureFlag.id")
	deleteCmd.MarkFlagRequired("id")

	parentCmd.AddCommand(deleteCmd)
}
