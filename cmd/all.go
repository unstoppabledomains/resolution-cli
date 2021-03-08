package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Resolve all records of a domain",
	Long:  "Resolve all records of a domain",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Resolving all records of %s...\n", domain)
		records, err := CNS.AllRecords(domain)
		if err != nil {
			log.Fatal("Error connecting to provider: " + err.Error())
		} else {
			fmt.Printf("%s records: %s\n", domain, records)
		}
	},
}
