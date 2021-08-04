package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:     "resolve",
	Short:   "Resolve all known records of domain",
	Long:    "Resolve all known records of domain. Domain must be specified",
	Example: "resolution resolve -d brad.crypto",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ReturnedValue, err = SelectedNamingService.AllRecords(Domain)
		if err != nil {
			log.Fatal(err)
		}
		ReturnedValue = prepareMultiRecordsOutput(ReturnedValue)
	},
}
