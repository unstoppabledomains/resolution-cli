package cmd

import (
	"encoding/json"
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
		var records map[string]string
		var err error
		if strings.HasSuffix(domain, (".crypto")) {
			records, err = CNS.Records(domain, args)
		} else {
			records, err = ZNS.Records(domain, args)
		}
		if err != nil {
			log.Fatal("Error connecting to provider: " + err.Error())
		} else {
			b, _ := json.Marshal(records)
			fmt.Println(string(b))

		}
	},
}
