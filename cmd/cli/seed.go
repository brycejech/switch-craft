package cli

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"switchcraft/core"

	"github.com/spf13/cobra"
)

func registerSeedModule(core *core.Core) {
	seedCmd := &cobra.Command{
		Use:   "seed",
		Short: "SwitchCraft CLI database seeder",
	}
	seedAllCmd(core, seedCmd)

	rootCmd.AddCommand(seedCmd)
}

func seedAllCmd(core *core.Core, parentCmd *cobra.Command) {
	var dataFile string
	allCmd := &cobra.Command{
		Use:   "all",
		Short: "Seed all tables with test data",
		Run: func(cmd *cobra.Command, _ []string) {
			seed := mustParseSeedFile(dataFile)
			fmt.Printf("%+v\n", seed)
		},
	}
	allCmd.Flags().StringVar(&dataFile, "dataFile", "", "Path to json seed file")
	allCmd.MarkFlagRequired("dataFile")

	parentCmd.AddCommand(allCmd)
}

func mustParseSeedFile(filepath string) map[string]any {
	seedFile, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	var seed map[string]any
	if err = json.Unmarshal([]byte(seedFile), &seed); err != nil {
		log.Fatal(err)
	}

	return seed
}
