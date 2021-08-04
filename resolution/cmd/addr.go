package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var addrCmd = &cobra.Command{
	Use:     "addr CURRENCY_TICKER",
	Example: "resolution resolve addr ETH -d brad.crypto",
	Short:   "Resolve address",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ReturnedValue, err = SelectedNamingService.Addr(Domain, args[0])
		if err != nil {
			log.Fatal(err)
		}
	}}
