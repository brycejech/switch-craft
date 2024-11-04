package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"switchcraft/core"
	"switchcraft/types"

	"github.com/spf13/cobra"
)

/*
	TODO
	----
	For each core method that will require a valid,
	authenticated auth session, call the cli.authn method,
	then types.NewOperationCtx to pass to core methods
*/

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "SwitchCraft CLI",
}

var baseCtx = context.Background()

func Start(switchcraft *core.Core) {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	registerMigrationsModule(switchcraft)
	registerAccountModule(switchcraft)
	registerAuthModule(switchcraft)

	rootCmd.Execute()
}

func mustAuthn(switchcraft *core.Core) (account *types.Account) {
	var (
		username = os.Getenv("SWITCHCRAFT_USER")
		password = os.Getenv("SWITCHCRAFT_PASS")
	)
	if username == "" || password == "" {
		log.Fatal("Must provide SWITCHCRAFT_USER and SWITCHCRAFT_PASS env vars to use CLI")
	}

	account, ok := switchcraft.Authn(baseCtx, username, password)
	if !ok {
		log.Fatal("Unable to authenticate local account")
	}
	return account
}

func printJSON(v interface{}) {
	bytes, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(bytes))
}
