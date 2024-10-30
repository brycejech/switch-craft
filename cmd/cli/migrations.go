package cli

import (
	"fmt"
	"switchcraft/core"

	"github.com/spf13/cobra"
)

func initMigrationsModule(core *core.Core) {
	var migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "SwitchCraft CLI migrations module",
	}
	rootCmd.AddCommand(migrateCmd)

	var upCmd = &cobra.Command{
		Use:   "up",
		Short: "Migrate database all the way up",
		Run: func(cmd *cobra.Command, args []string) {
			if err := core.MigrateUp(); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Successfully ran up migration(s)")
		},
	}
	migrateCmd.AddCommand(upCmd)

	var downCmd = &cobra.Command{
		Use:   "down",
		Short: "Migrate database down by a single migration",
		Run: func(cmd *cobra.Command, args []string) {
			if err := core.MigrateDown(); err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Successfully ran down migration")
		},
	}
	migrateCmd.AddCommand(downCmd)
}
