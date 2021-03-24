package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var ownerCmd = &cobra.Command{
	Use:     "owner",
	Short:   "Returns domain owner address",
	Args:    cobra.ExactArgs(0),
	Example: "resolution resolve owner -d brad.crypto",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ReturnedValue, err = SelectedNamingService.Owner(Domain)
		if err != nil {
			log.Fatal(err)
		}
	},
}
