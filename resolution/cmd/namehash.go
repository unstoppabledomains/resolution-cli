package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var namehashCmd = &cobra.Command{
	Use:     "namehash",
	Example: "resolution namehash -d brad.crypto",
	Short:   "Get namehash of domain",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ReturnedValue, err = SelectedNamingService.Namehash(Domain)
		if err != nil {
			log.Fatal(err)
		}
	}}
