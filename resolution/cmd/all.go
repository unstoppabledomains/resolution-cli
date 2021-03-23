package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Resolve all records",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ReturnedValue, err = SelectedNamingService.AllRecords(Domain) // todo output should comply with resolution api standard
		if err != nil {
			log.Fatal(err)
		}
	},
}
