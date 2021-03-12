package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var addrCmd = &cobra.Command{
	Use:   "addr [address]",
	Short: "Resolve address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var value string
		var err error
		if strings.HasSuffix(domain, (".crypto")) {
			value, err = CNS.Addr(domain, args[0])
		} else {
			value, err = ZNS.Addr(domain, args[0])
		}
		if err != nil {
			log.Fatal("Error connecting to provider: " + err.Error())
		} else {
			b, _ := json.Marshal(map[string]string{args[0]: value})
			fmt.Println(string(b))
		}
	},
}
