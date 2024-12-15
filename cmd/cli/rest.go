package cli

import (
	"switchcraft/cmd/rest"
	"switchcraft/core"
	"switchcraft/types"

	"github.com/spf13/cobra"
)

func registerRestModule(logger *types.Logger, core *core.Core) {
	var restPort string
	var restCmd = &cobra.Command{
		Use:   "serve",
		Short: "SwitchCraft REST API server",
		Run: func(_ *cobra.Command, _ []string) {
			rest.Start(logger, core, restPort)
		},
	}
	restCmd.Flags().StringVar(&restPort, "port", "8080", "REST API server port")

	rootCmd.AddCommand(restCmd)
}
