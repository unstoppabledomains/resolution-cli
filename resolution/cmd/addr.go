package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var addrCmd = &cobra.Command{
	Use:   "addr [address]",
	Short: "Resolve address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ReturnedValue, err = SelectedNamingService.Addr(Domain, args[1])
		if err != nil {
			log.Fatal(err)
		}
	},
}
