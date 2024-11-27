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
	var getAccountsOrgID int64
	var getAccountsCmd = &cobra.Command{
		Use:   "getMany",
		Short: "Get multiple accounts",
		Run: func(cmd *cobra.Command, args []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			var orgID *int64
			if cmd.Flags().Changed("orgId") {
				orgID = &getAccountsOrgID
			}

			accounts, err := core.AccountGetMany(opCtx, orgID)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(accounts)
		},
	}
	getAccountsCmd.Flags().Int64Var(&getAccountsOrgID, "orgId", 0, "account.createdBy")
	getAccountsCmd.MarkFlagRequired("orgId")
	accountCmd.AddCommand(getAccountsCmd)

	/* -------------------------- */
	/* === CREATE ACCOUNT CMD === */
	/* -------------------------- */
	createAccountCmdArgs := struct {
		OrgID     int64
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
				createAccountCmdArgs.OrgID,
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

	createAccountCmd.Flags().Int64Var(&createAccountCmdArgs.OrgID, "orgId", 0, "account.orgId")
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
	var getAccountOrgID int64
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
				orgID    *int64
				id       *int64
				uuid     *string
				username *string
			)
			if cmd.Flags().Changed("orgId") {
				orgID = &getAccountOrgID
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
				orgID,
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
	getAccountCmd.Flags().Int64Var(&getAccountOrgID, "orgId", 0, "account.orgId")
	getAccountCmd.Flags().Int64Var(&getAccountID, "id", 0, "account.id")
	getAccountCmd.Flags().StringVar(&getAccountUUID, "uuid", "", "account.uuid")
	getAccountCmd.Flags().StringVar(&getAccountUsername, "username", "", "account.username")
	accountCmd.AddCommand(getAccountCmd)

	/* -------------------------- */
	/* === UPDATE ACCOUNT CMD === */
	/* -------------------------- */
	updateAccountCmdArgs := struct {
		orgID           int64
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

			var orgID *int64
			if cmd.Flags().Changed("orgId") {
				orgID = &updateAccountCmdArgs.orgID
			}
			args := core.NewAccountUpdateArgs(
				orgID,
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

	updateAccountCmd.Flags().Int64Var(&updateAccountCmdArgs.orgID, "orgId", 0, "account.orgId")
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
	var deleteAccountOrgID int64
	var deleteAccountID int64
	var deleteAccountCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an account",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			var orgID *int64
			if cmd.Flags().Changed("orgId") {
				orgID = &deleteAccountOrgID
			}
			if err := core.AccountDelete(opCtx, orgID, deleteAccountID); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Account '%v' deleted successfully\n", deleteAccountID)
		},
	}

	deleteAccountCmd.Flags().Int64Var(&deleteAccountID, "id", 0, "account.id")
	deleteAccountCmd.MarkFlagRequired("id")
	accountCmd.AddCommand(deleteAccountCmd)
}
