package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var locationCmd = &cobra.Command{
	Use:     "location",
	Example: "resolution resolve location -d brad.crypto",
	Short:   "Resolve location of domain",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		location, err := SelectedNamingService.Locations([]string{Domain})
		if err != nil {
			log.Fatal(err)
		}
		ReturnedValue = location[Domain]
	}}
