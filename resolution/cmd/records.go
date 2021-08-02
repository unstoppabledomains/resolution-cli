package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var recordsCmd = &cobra.Command{
	Use:   "records RECORD_KEY_1 RECORD_KEY_2...",
	Short: "Resolve list of records",
	Long: `
Resolve the list of records. 
Find details about available records and records format in records reference guide: 
https://docs.unstoppabledomains.com/domain-registry-essentials/records-reference
`,
	Example: "resolution resolve records crypto.ETH.address crypto.BTC.address -d brad.crypto",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ReturnedValue, err = SelectedNamingService.Records(Domain, args)
		if err != nil {
			log.Fatal(err)
		}
		ReturnedValue = prepareMultiRecordsOutput(ReturnedValue)
	},
}
