package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"switchcraft/core"
	"switchcraft/types"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "SwitchCraft CLI",
}

var baseCtx = context.Background()

func Start(logger *types.Logger, core *core.Core) {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	registerMigrationsModule(core)
	registerAccountModule(core)
	registerTenantModule(core)
	registerAuthModule(core)
	registerAppModule(core)
	registerFeatureFlagModule(core)
	registerRestModule(logger, core)

	rootCmd.Execute()
}

// TODO: This should just return opCtx with account embedded
func mustAuthn(core *core.Core) (account *types.Account) {
	var (
		username = os.Getenv("SWITCHCRAFT_USER")
		password = os.Getenv("SWITCHCRAFT_PASS")
	)
	if username == "" || password == "" {
		log.Fatal("Must provide SWITCHCRAFT_USER and SWITCHCRAFT_PASS env vars to use CLI")
	}

	opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), types.Account{})
	account, ok := core.Authn(opCtx, username, password)
	if !ok {
		log.Fatal("Unable to authenticate local account")
	}
	return account
}

func printJSON(v interface{}) {
	bytes, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(bytes))
}
