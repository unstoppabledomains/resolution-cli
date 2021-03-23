package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var recordsCmd = &cobra.Command{
	Use:   "records [records]",
	Short: "Resolve list of records",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ReturnedValue, err = SelectedNamingService.Records(Domain, args) // todo output should comply with resolution api standard
		if err != nil {
			log.Fatal(err)
		}
	},
}
