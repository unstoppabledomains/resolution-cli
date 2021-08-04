package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var resolverCmd = &cobra.Command{
	Use:     "resolver",
	Short:   "Returns domain resolver address",
	Args:    cobra.ExactArgs(0),
	Example: "resolution resolve resolver -d brad.crypto",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ReturnedValue, err = SelectedNamingService.Resolver(Domain)
		if err != nil {
			log.Fatal(err)
		}
	},
}
