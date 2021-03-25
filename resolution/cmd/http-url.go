package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var httpUrlCmd = &cobra.Command{
	Use:     "http-url",
	Short:   "Resolve http redirect url",
	Args:    cobra.ExactArgs(0),
	Example: "resolution resolve http-url -d brad.crypto",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ReturnedValue, err = SelectedNamingService.HTTPUrl(Domain)
		if err != nil {
			log.Fatal(err)
		}
	},
}
