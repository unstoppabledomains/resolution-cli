package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Resolve all records",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
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
			b, _ := json.Marshal(records)
			fmt.Println(string(b))
		}
	},
}
