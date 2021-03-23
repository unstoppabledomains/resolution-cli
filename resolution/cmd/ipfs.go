package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var ipfsCmd = &cobra.Command{
	Use:   "ipfs",
	Short: "Resolve ipfs records",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ReturnedValue, err = SelectedNamingService.IpfsHash(Domain)
		if err != nil {
			log.Fatal(err)
		}
	},
}
