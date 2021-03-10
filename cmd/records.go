package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var recordsCmd = &cobra.Command{
	Use:   "records [records]",
	Short: "Resolve list of records",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Resolving records of %s\n", domain)

		var values map[string]string
		var err error
		if strings.HasSuffix(domain, (".crypto")) {
			values, err = CNS.Records(domain, args)
		} else {
			values, err = ZNS.Records(domain, args)
		}
		if err != nil {
			log.Fatal("Error connecting to provider: " + err.Error())
		} else {
			fmt.Printf("%s records: %s\n", domain, values)
		}
	},
}
