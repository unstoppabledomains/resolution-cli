package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var recordsCmd = &cobra.Command{
	Use:   "records [records]",
	Short: "Resolve records of a domain",
	Long:  "Resolve records of a domain",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Resolving records of %s\n", domain)
		values, err := CNS.Records(domain, args)
		if err != nil {
			log.Fatal("Error connecting to provider: " + err.Error())
		} else {
			fmt.Printf("%s records: %s\n", domain, values)
		}
	},
}
