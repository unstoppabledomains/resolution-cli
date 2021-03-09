package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Resolve all records of a domain",
	Long:  "Resolve all records of a domain",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Resolving all records of %s...\n", domain)
		var records map[string]string
		var err error
		if strings.HasSuffix(domain, (".crypto")) {
			records, err = CNS.AllRecords(domain)
		} else {
			records, err = ZNS.AllRecords(domain)
		}
		if err != nil {
			log.Fatal("Error connecting to provider: " + err.Error())
		} else {
			fmt.Printf("%s records: %s\n", domain, records)
		}
	},
}
