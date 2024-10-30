package cli

import (
	"switchcraft/core"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "SwitchCraft CLI",
}

func Start(core *core.Core) {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	initMigrationsModule(core)

	rootCmd.Execute()
}
