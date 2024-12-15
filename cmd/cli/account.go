package cli

import (
	"fmt"
	"log"
	"switchcraft/core"
	"switchcraft/types"
	"time"

	"github.com/spf13/cobra"
)

func registerOrgAccountModule(core *core.Core) {
	var orgAccountCmd = &cobra.Command{
		Use:   "orgAccount",
		Short: "SwitchCraft CLI organization account module",
	}
	orgAccountCreateCmd(core, orgAccountCmd)
	orgAccountGetManyCmd(core, orgAccountCmd)
	orgAccountGetOneCmd(core, orgAccountCmd)
	orgAccountUpdateCmd(core, orgAccountCmd)
	orgAccountDeleteCmd(core, orgAccountCmd)

	rootCmd.AddCommand(orgAccountCmd)

}

func orgAccountCreateCmd(core *core.Core, parentCmd *cobra.Command) {
	args := struct {
		orgSlug   string
		firstName string
		lastName  string
		email     string
		username  string
	}{}
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create new organization account",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			account, err := core.OrgAccountCreate(opCtx,
				core.NewOrgAccountCreateArgs(
					args.orgSlug,
					args.firstName,
					args.lastName,
					args.email,
					args.username,
					nil,
				),
			)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(account)
		},
	}
	createCmd.Flags().StringVar(&args.orgSlug, "orgSlug", "", "Organization slug")
	createCmd.Flags().StringVar(&args.firstName, "firstName", "", "account.firstName")
	createCmd.MarkFlagRequired("firstName")
	createCmd.Flags().StringVar(&args.lastName, "lastName", "", "account.lastName")
	createCmd.MarkFlagRequired("lastName")
	createCmd.Flags().StringVar(&args.email, "email", "", "account.email")
	createCmd.MarkFlagRequired("email")
	createCmd.Flags().StringVar(&args.username, "username", "", "account.username")
	createCmd.MarkFlagRequired("username")

	parentCmd.AddCommand(createCmd)
}

func orgAccountGetManyCmd(core *core.Core, parentCmd *cobra.Command) {
	var slug string
	getManyCmd := &cobra.Command{
		Use:   "getMany",
		Short: "Get multiple organization accounts",
		Run: func(cmd *cobra.Command, args []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			accounts, err := core.OrgAccountGetMany(opCtx, slug)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(accounts)
		},
	}
	getManyCmd.Flags().StringVar(&slug, "orgSlug", "", "Organization slug")
	getManyCmd.MarkFlagRequired("orgSlug")

	parentCmd.AddCommand(getManyCmd)
}

func orgAccountGetOneCmd(core *core.Core, parentCmd *cobra.Command) {
	var orgSlug string
	var accountID int64
	var accountUUID string
	var accountUsername string
	getOneCmd := &cobra.Command{
		Use:   "getOne",
		Short: "Get an organization account by id, uuid, or username",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			var (
				id       *int64
				uuid     *string
				username *string
			)
			if cmd.Flags().Changed("id") {
				id = &accountID
			}
			if cmd.Flags().Changed("uuid") {
				uuid = &accountUUID
			}
			if cmd.Flags().Changed("username") {
				username = &accountUsername
			}

			account, err := core.OrgAccountGetOne(opCtx,
				core.NewOrgAccountGetOneArgs(
					orgSlug,
					id,
					uuid,
					username,
				),
			)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(account)
		},
	}
	getOneCmd.Flags().StringVar(&orgSlug, "orgSlug", "", "Organization slug")
	getOneCmd.MarkFlagRequired("orgSlug")
	getOneCmd.Flags().Int64Var(&accountID, "id", 0, "account.id")
	getOneCmd.Flags().StringVar(&accountUUID, "uuid", "", "account.uuid")
	getOneCmd.Flags().StringVar(&accountUsername, "username", "", "account.username")

	parentCmd.AddCommand(getOneCmd)
}

func orgAccountUpdateCmd(core *core.Core, parentCmd *cobra.Command) {
	args := struct {
		orgSlug   string
		id        int64
		firstName string
		lastName  string
		email     string
		username  string
	}{}
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Update an existing organization account",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			account, err := core.OrgAccountUpdate(opCtx,
				core.NewOrgAccountUpdateArgs(
					args.orgSlug,
					args.id,
					args.firstName,
					args.lastName,
					args.email,
					args.username,
				),
			)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(account)
		},
	}

	updateCmd.Flags().StringVar(&args.orgSlug, "orgSlug", "", "Organization slug")
	updateCmd.MarkFlagRequired("orgSlug")
	updateCmd.Flags().Int64Var(&args.id, "id", 0, "account.id")
	updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().StringVar(&args.firstName, "firstName", "", "account.firstName")
	updateCmd.MarkFlagRequired("firstName")
	updateCmd.Flags().StringVar(&args.lastName, "lastName", "", "account.lastName")
	updateCmd.MarkFlagRequired("lastName")
	updateCmd.Flags().StringVar(&args.email, "email", "", "account.email")
	updateCmd.MarkFlagRequired("email")
	updateCmd.Flags().StringVar(&args.username, "username", "", "account.username")
	updateCmd.MarkFlagRequired("username")

	parentCmd.AddCommand(updateCmd)
}

func orgAccountDeleteCmd(core *core.Core, parentCmd *cobra.Command) {
	var orgSlug string
	var accountID int64
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an organization account",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			if err := core.OrgAccountDelete(opCtx, orgSlug, accountID); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Organization account '%v' deleted successfully\n", accountID)
		},
	}
	deleteCmd.Flags().StringVar(&orgSlug, "orgSlug", "", "Organization slug")
	deleteCmd.MarkFlagRequired("orgSlug")
	deleteCmd.Flags().Int64Var(&accountID, "id", 0, "account.id")
	deleteCmd.MarkFlagRequired("id")

	parentCmd.AddCommand(deleteCmd)
}
