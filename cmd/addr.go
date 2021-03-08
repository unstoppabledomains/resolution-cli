package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var addrCmd = &cobra.Command{
	Use:   "addr [address]",
	Short: "Resolve address of a domain",
	Long:  "Resolve address of a domain",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Resolving %s address of %s...\n", args[0], domain)
		value, err := CNS.Addr(domain, args[0])
		if err != nil {
			log.Fatal("Error connecting to provider: " + err.Error())
		} else {
			fmt.Printf("%s %s address resolves to %s\n", domain, args[0], value)
		}
	},
}
