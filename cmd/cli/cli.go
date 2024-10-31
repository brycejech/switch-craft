package cli

import (
	"encoding/json"
	"fmt"
	"switchcraft/core"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "SwitchCraft CLI",
}

func Start(switchcraft *core.Core) {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	registerMigrationsModule(switchcraft)
	registerAccountModule(switchcraft)

	rootCmd.Execute()
}

func printJSON(v interface{}) {
	bytes, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(string(bytes))
}
