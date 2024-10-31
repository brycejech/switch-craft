package cli

import (
	"context"
	"fmt"
	"log"
	"switchcraft/core"

	"github.com/spf13/cobra"
)

func registerAccountModule(switchcraft *core.Core) {
	ctx := context.Background()

	var accountCmd = &cobra.Command{
		Use:   "account",
		Short: "SwitchCraft CLI account module",
	}
	rootCmd.AddCommand(accountCmd)

	/* === GET ACCOUNTS CMD === */
	/* ------------------------ */
	var getAccountsCmd = &cobra.Command{
		Use:   "getMany",
		Short: "Get multiple accounts",
		Run: func(cmd *cobra.Command, args []string) {
			accounts, err := switchcraft.AccountGetMany(ctx)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(accounts)
		},
	}
	accountCmd.AddCommand(getAccountsCmd)

	/* === CREATE ACCOUNT CMD === */
	/* -------------------------- */
	createAccountCmdArgs := struct {
		FirstName string
		LastName  string
		Email     string
		Username  string
		CreatedBy int64
	}{}
	var createAccountCmd = &cobra.Command{
		Use:   "create",
		Short: "Create new account",
		Run: func(cmd *cobra.Command, _ []string) {
			args := core.NewAccountCreateArgs(
				createAccountCmdArgs.FirstName,
				createAccountCmdArgs.LastName,
				createAccountCmdArgs.Email,
				createAccountCmdArgs.Username,
				createAccountCmdArgs.CreatedBy,
			)

			account, err := switchcraft.AccountCreate(ctx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(account)
		},
	}

	createAccountCmd.Flags().StringVar(&createAccountCmdArgs.FirstName, "firstName", "", "Account.FirstName")
	createAccountCmd.MarkFlagRequired("firstName")
	createAccountCmd.Flags().StringVar(&createAccountCmdArgs.LastName, "lastName", "", "Account.LastName")
	createAccountCmd.MarkFlagRequired("lastName")
	createAccountCmd.Flags().StringVar(&createAccountCmdArgs.Email, "email", "", "Account.Email")
	createAccountCmd.MarkFlagRequired("email")
	createAccountCmd.Flags().StringVar(&createAccountCmdArgs.Username, "username", "", "Account.Username")
	createAccountCmd.MarkFlagRequired("username")
	createAccountCmd.Flags().Int64Var(&createAccountCmdArgs.CreatedBy, "createdBy", 0, "Account.CreatedBy")
	createAccountCmd.MarkFlagRequired("createdBy")
	accountCmd.AddCommand(createAccountCmd)

	/* === GET ACCOUNT CMD === */
	/* ----------------------- */
	var getAccountID int64
	var getAccountUUID string
	var getAccountUsername string
	var getAccountCmd = &cobra.Command{
		Use:   "getOne",
		Short: "Get an account by id, uuid, or username",
		Run: func(cmd *cobra.Command, _ []string) {
			var (
				id       *int64
				uuid     *string
				username *string
			)
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
				id,
				uuid,
				username,
			)

			account, err := switchcraft.AccountGetOne(ctx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(account)
		},
	}
	getAccountCmd.Flags().Int64Var(&getAccountID, "id", 0, "Account.ID")
	getAccountCmd.Flags().StringVar(&getAccountUUID, "uuid", "", "Account.UUID")
	getAccountCmd.Flags().StringVar(&getAccountUsername, "username", "", "Account.Username")
	accountCmd.AddCommand(getAccountCmd)

	/* === GET ACCOUNT CMD === */
	/* ----------------------- */
	updateAccountCmdArgs := struct {
		ID         int64
		FirstName  string
		LastName   string
		Email      string
		Username   string
		ModifiedBy int64
	}{}
	var updateAccountCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an existing account",
		Run: func(cmd *cobra.Command, _ []string) {
			args := core.NewAccountUpdateArgs(
				updateAccountCmdArgs.ID,
				updateAccountCmdArgs.FirstName,
				updateAccountCmdArgs.LastName,
				updateAccountCmdArgs.Email,
				updateAccountCmdArgs.Username,
				updateAccountCmdArgs.ModifiedBy,
			)

			account, err := switchcraft.AccountUpdate(ctx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(account)
		},
	}

	updateAccountCmd.Flags().Int64Var(&updateAccountCmdArgs.ID, "id", 0, "Account.ID")
	updateAccountCmd.MarkFlagRequired("id")
	updateAccountCmd.Flags().StringVar(&updateAccountCmdArgs.FirstName, "firstName", "", "Account.FirstName")
	updateAccountCmd.MarkFlagRequired("firstName")
	updateAccountCmd.Flags().StringVar(&updateAccountCmdArgs.LastName, "lastName", "", "Account.LastName")
	updateAccountCmd.MarkFlagRequired("lastName")
	updateAccountCmd.Flags().StringVar(&updateAccountCmdArgs.Email, "email", "", "Account.Email")
	updateAccountCmd.MarkFlagRequired("email")
	updateAccountCmd.Flags().StringVar(&updateAccountCmdArgs.Username, "username", "", "Account.Username")
	updateAccountCmd.MarkFlagRequired("username")
	updateAccountCmd.Flags().Int64Var(&updateAccountCmdArgs.ModifiedBy, "modifiedBy", 0, "Account.ModifiedBy")
	updateAccountCmd.MarkFlagRequired("modifiedBy")
	accountCmd.AddCommand(updateAccountCmd)

	/* === DELETE ACCOUNT CMD === */
	/* -------------------------- */
	var deleteAccountID int64
	var deleteAccountCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an account",
		Run: func(cmd *cobra.Command, _ []string) {
			if err := switchcraft.AccountDelete(ctx, deleteAccountID); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Account '%v' deleted successfully\n", deleteAccountID)
		},
	}
	deleteAccountCmd.Flags().Int64Var(&deleteAccountID, "id", 0, "Account.ID")
	deleteAccountCmd.MarkFlagRequired("id")
	accountCmd.AddCommand(deleteAccountCmd)
}
