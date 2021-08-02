package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var emailCmd = &cobra.Command{
	Use:     "email",
	Short:   "Resolve email attached to domain name",
	Args:    cobra.ExactArgs(0),
	Example: "resolution resolve email -d reseller-test-paul019.crypto",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ReturnedValue, err = SelectedNamingService.Email(Domain)
		if err != nil {
			log.Fatal(err)
		}
	},
}
