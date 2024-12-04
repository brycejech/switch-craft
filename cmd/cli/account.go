package cli

import (
	"fmt"
	"log"
	"switchcraft/core"
	"switchcraft/types"
	"time"

	"github.com/spf13/cobra"
)

func registerAccountModule(core *core.Core) {
	var accountCmd = &cobra.Command{
		Use:   "account",
		Short: "SwitchCraft CLI account module",
	}
	rootCmd.AddCommand(accountCmd)

	/* ------------------------ */
	/* === GET ACCOUNTS CMD === */
	/* ------------------------ */
	var getAccountsOrgSlug string
	var getAccountsCmd = &cobra.Command{
		Use:   "getMany",
		Short: "Get multiple accounts",
		Run: func(cmd *cobra.Command, args []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			var orgSlug *string
			if cmd.Flags().Changed("orgId") {
				orgSlug = &getAccountsOrgSlug
			}

			accounts, err := core.AccountGetMany(opCtx, orgSlug)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(accounts)
		},
	}
	getAccountsCmd.Flags().StringVar(&getAccountsOrgSlug, "orgSlug", "", "Organization slug")
	getAccountsCmd.MarkFlagRequired("orgSlug")
	accountCmd.AddCommand(getAccountsCmd)

	/* -------------------------- */
	/* === CREATE ACCOUNT CMD === */
	/* -------------------------- */
	createAccountCmdArgs := struct {
		OrgSlug   string
		FirstName string
		LastName  string
		Email     string
		Username  string
	}{}
	var createAccountCmd = &cobra.Command{
		Use:   "create",
		Short: "Create new account",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewAccountCreateArgs(
				createAccountCmdArgs.OrgSlug,
				createAccountCmdArgs.FirstName,
				createAccountCmdArgs.LastName,
				createAccountCmdArgs.Email,
				createAccountCmdArgs.Username,
				nil,
			)

			account, err := core.AccountCreate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(account)
		},
	}

	createAccountCmd.Flags().StringVar(&createAccountCmdArgs.OrgSlug, "orgSlug", "", "account.orgSlug")
	createAccountCmd.Flags().StringVar(&createAccountCmdArgs.FirstName, "firstName", "", "account.firstName")
	createAccountCmd.MarkFlagRequired("firstName")
	createAccountCmd.Flags().StringVar(&createAccountCmdArgs.LastName, "lastName", "", "account.lastName")
	createAccountCmd.MarkFlagRequired("lastName")
	createAccountCmd.Flags().StringVar(&createAccountCmdArgs.Email, "email", "", "account.email")
	createAccountCmd.MarkFlagRequired("email")
	createAccountCmd.Flags().StringVar(&createAccountCmdArgs.Username, "username", "", "account.username")
	createAccountCmd.MarkFlagRequired("username")
	accountCmd.AddCommand(createAccountCmd)

	/* ----------------------- */
	/* === GET ACCOUNT CMD === */
	/* ----------------------- */
	var getAccountOrgSlug string
	var getAccountID int64
	var getAccountUUID string
	var getAccountUsername string
	var getAccountCmd = &cobra.Command{
		Use:   "getOne",
		Short: "Get an account by id, uuid, or username",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			var (
				orgSlug  *string
				id       *int64
				uuid     *string
				username *string
			)
			if cmd.Flags().Changed("orgId") {
				orgSlug = &getAccountOrgSlug
			}
			if cmd.Flags().Changed("id") {
				id = &getAccountID
			}
			if cmd.Flags().Changed("uuid") {
				uuid = &getAccountUUID
			}
			if cmd.Flags().Changed("username") {
				username = &getAccountUsername
			}

			args := core.NewAccountGetOneArgs(
				orgSlug,
				id,
				uuid,
				username,
			)

			account, err := core.AccountGetOne(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(account)
		},
	}
	getAccountCmd.Flags().StringVar(&getAccountOrgSlug, "orgSlug", "", "Organization slug")
	getAccountCmd.Flags().Int64Var(&getAccountID, "id", 0, "account.id")
	getAccountCmd.Flags().StringVar(&getAccountUUID, "uuid", "", "account.uuid")
	getAccountCmd.Flags().StringVar(&getAccountUsername, "username", "", "account.username")
	accountCmd.AddCommand(getAccountCmd)

	/* -------------------------- */
	/* === UPDATE ACCOUNT CMD === */
	/* -------------------------- */
	updateAccountCmdArgs := struct {
		orgSlug         string
		id              int64
		isInstanceAdmin bool
		firstName       string
		lastName        string
		email           string
		username        string
	}{}
	var updateAccountCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an existing account",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			var orgSlug *string
			if cmd.Flags().Changed("orgId") {
				orgSlug = &updateAccountCmdArgs.orgSlug
			}

			args := core.NewAccountUpdateArgs(
				orgSlug,
				updateAccountCmdArgs.id,
				updateAccountCmdArgs.isInstanceAdmin,
				updateAccountCmdArgs.firstName,
				updateAccountCmdArgs.lastName,
				updateAccountCmdArgs.email,
				updateAccountCmdArgs.username,
			)

			account, err := core.AccountUpdate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(account)
		},
	}

	updateAccountCmd.Flags().StringVar(&updateAccountCmdArgs.orgSlug, "orgSlug", "", "Organization slug")
	updateAccountCmd.Flags().Int64Var(&updateAccountCmdArgs.id, "id", 0, "account.id")
	updateAccountCmd.MarkFlagRequired("id")
	updateAccountCmd.Flags().BoolVar(&updateAccountCmdArgs.isInstanceAdmin, "isInstanceAdmin", false, "account.isInstanceAdmin")
	updateAccountCmd.MarkFlagRequired("isInstanceAdmin")
	updateAccountCmd.Flags().StringVar(&updateAccountCmdArgs.firstName, "firstName", "", "account.firstName")
	updateAccountCmd.MarkFlagRequired("firstName")
	updateAccountCmd.Flags().StringVar(&updateAccountCmdArgs.lastName, "lastName", "", "account.lastName")
	updateAccountCmd.MarkFlagRequired("lastName")
	updateAccountCmd.Flags().StringVar(&updateAccountCmdArgs.email, "email", "", "account.email")
	updateAccountCmd.MarkFlagRequired("email")
	updateAccountCmd.Flags().StringVar(&updateAccountCmdArgs.username, "username", "", "account.username")
	updateAccountCmd.MarkFlagRequired("username")
	accountCmd.AddCommand(updateAccountCmd)

	/* -------------------------- */
	/* === DELETE ACCOUNT CMD === */
	/* -------------------------- */
	var deleteAccountOrgSlug string
	var deleteAccountID int64
	var deleteAccountCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an account",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			var orgSlug *string
			if cmd.Flags().Changed("orgSlug") {
				orgSlug = &deleteAccountOrgSlug
			}
			if err := core.AccountDelete(opCtx, orgSlug, deleteAccountID); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Account '%v' deleted successfully\n", deleteAccountID)
		},
	}

	deleteAccountCmd.Flags().StringVar(&deleteAccountOrgSlug, "orgSlug", "", "Organization slug")
	deleteAccountCmd.Flags().Int64Var(&deleteAccountID, "id", 0, "account.id")
	deleteAccountCmd.MarkFlagRequired("id")
	accountCmd.AddCommand(deleteAccountCmd)
}
