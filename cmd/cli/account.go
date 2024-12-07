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
	rootCmd.AddCommand(orgAccountCmd)

	/* ---------------------------- */
	/* === GET ORG ACCOUNTS CMD === */
	/* ---------------------------- */
	var getOrgAccountsOrgSlug string
	var getOrgAccountsCmd = &cobra.Command{
		Use:   "getMany",
		Short: "Get multiple organization accounts",
		Run: func(cmd *cobra.Command, args []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			accounts, err := core.OrgAccountGetMany(opCtx, getOrgAccountsOrgSlug)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(accounts)
		},
	}
	getOrgAccountsCmd.Flags().StringVar(&getOrgAccountsOrgSlug, "orgSlug", "", "Organization slug")
	getOrgAccountsCmd.MarkFlagRequired("orgSlug")
	orgAccountCmd.AddCommand(getOrgAccountsCmd)

	/* ------------------------------ */
	/* === CREATE ORG ACCOUNT CMD === */
	/* ------------------------------ */
	createOrgAccountCmdArgs := struct {
		OrgSlug   string
		FirstName string
		LastName  string
		Email     string
		Username  string
	}{}
	var createOrgAccountCmd = &cobra.Command{
		Use:   "create",
		Short: "Create new organization account",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewOrgAccountCreateArgs(
				createOrgAccountCmdArgs.OrgSlug,
				createOrgAccountCmdArgs.FirstName,
				createOrgAccountCmdArgs.LastName,
				createOrgAccountCmdArgs.Email,
				createOrgAccountCmdArgs.Username,
				nil,
			)

			account, err := core.OrgAccountCreate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(account)
		},
	}

	createOrgAccountCmd.Flags().StringVar(&createOrgAccountCmdArgs.OrgSlug, "orgSlug", "", "account.orgSlug")
	createOrgAccountCmd.Flags().StringVar(&createOrgAccountCmdArgs.FirstName, "firstName", "", "account.firstName")
	createOrgAccountCmd.MarkFlagRequired("firstName")
	createOrgAccountCmd.Flags().StringVar(&createOrgAccountCmdArgs.LastName, "lastName", "", "account.lastName")
	createOrgAccountCmd.MarkFlagRequired("lastName")
	createOrgAccountCmd.Flags().StringVar(&createOrgAccountCmdArgs.Email, "email", "", "account.email")
	createOrgAccountCmd.MarkFlagRequired("email")
	createOrgAccountCmd.Flags().StringVar(&createOrgAccountCmdArgs.Username, "username", "", "account.username")
	createOrgAccountCmd.MarkFlagRequired("username")
	orgAccountCmd.AddCommand(createOrgAccountCmd)

	/* --------------------------- */
	/* === GET ORG ACCOUNT CMD === */
	/* --------------------------- */
	var getOrgAccountOrgSlug string
	var getOrgAccountID int64
	var getOrgAccountUUID string
	var getOrgAccountUsername string
	var getOrgAccountCmd = &cobra.Command{
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
				id = &getOrgAccountID
			}
			if cmd.Flags().Changed("uuid") {
				uuid = &getOrgAccountUUID
			}
			if cmd.Flags().Changed("username") {
				username = &getOrgAccountUsername
			}

			args := core.NewOrgAccountGetOneArgs(
				getOrgAccountOrgSlug,
				id,
				uuid,
				username,
			)

			account, err := core.OrgAccountGetOne(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(account)
		},
	}
	getOrgAccountCmd.Flags().StringVar(&getOrgAccountOrgSlug, "orgSlug", "", "Organization slug")
	getOrgAccountCmd.MarkFlagRequired("orgSlug")
	getOrgAccountCmd.Flags().Int64Var(&getOrgAccountID, "id", 0, "account.id")
	getOrgAccountCmd.Flags().StringVar(&getOrgAccountUUID, "uuid", "", "account.uuid")
	getOrgAccountCmd.Flags().StringVar(&getOrgAccountUsername, "username", "", "account.username")
	orgAccountCmd.AddCommand(getOrgAccountCmd)

	/* ------------------------------ */
	/* === UPDATE ORG ACCOUNT CMD === */
	/* ------------------------------ */
	updateOrgAccountCmdArgs := struct {
		orgSlug   string
		id        int64
		firstName string
		lastName  string
		email     string
		username  string
	}{}
	var updateOrgAccountCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an existing organization account",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewOrgAccountUpdateArgs(
				updateOrgAccountCmdArgs.orgSlug,
				updateOrgAccountCmdArgs.id,
				updateOrgAccountCmdArgs.firstName,
				updateOrgAccountCmdArgs.lastName,
				updateOrgAccountCmdArgs.email,
				updateOrgAccountCmdArgs.username,
			)

			account, err := core.OrgAccountUpdate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(account)
		},
	}

	updateOrgAccountCmd.Flags().StringVar(&updateOrgAccountCmdArgs.orgSlug, "orgSlug", "", "Organization slug")
	updateOrgAccountCmd.MarkFlagRequired("orgSlug")
	updateOrgAccountCmd.Flags().Int64Var(&updateOrgAccountCmdArgs.id, "id", 0, "account.id")
	updateOrgAccountCmd.MarkFlagRequired("id")
	updateOrgAccountCmd.Flags().StringVar(&updateOrgAccountCmdArgs.firstName, "firstName", "", "account.firstName")
	updateOrgAccountCmd.MarkFlagRequired("firstName")
	updateOrgAccountCmd.Flags().StringVar(&updateOrgAccountCmdArgs.lastName, "lastName", "", "account.lastName")
	updateOrgAccountCmd.MarkFlagRequired("lastName")
	updateOrgAccountCmd.Flags().StringVar(&updateOrgAccountCmdArgs.email, "email", "", "account.email")
	updateOrgAccountCmd.MarkFlagRequired("email")
	updateOrgAccountCmd.Flags().StringVar(&updateOrgAccountCmdArgs.username, "username", "", "account.username")
	updateOrgAccountCmd.MarkFlagRequired("username")
	orgAccountCmd.AddCommand(updateOrgAccountCmd)

	/* ------------------------------ */
	/* === DELETE ORG ACCOUNT CMD === */
	/* ------------------------------ */
	var deleteOrgAccountOrgSlug string
	var deleteOrgAccountAccountID int64
	var deleteOrgAccountCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an organization account",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			if err := core.OrgAccountDelete(opCtx, deleteOrgAccountOrgSlug, deleteOrgAccountAccountID); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Organization account '%v' deleted successfully\n", deleteOrgAccountAccountID)
		},
	}

	deleteOrgAccountCmd.Flags().StringVar(&deleteOrgAccountOrgSlug, "orgSlug", "", "Organization slug")
	deleteOrgAccountCmd.MarkFlagRequired("orgSlug")
	deleteOrgAccountCmd.Flags().Int64Var(&deleteOrgAccountAccountID, "id", 0, "account.id")
	deleteOrgAccountCmd.MarkFlagRequired("id")
	orgAccountCmd.AddCommand(deleteOrgAccountCmd)
}
