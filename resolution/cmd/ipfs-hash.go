package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var ipfsCmd = &cobra.Command{
	Use:     "ipfs-hash",
	Short:   "Resolve ipfs hash",
	Args:    cobra.ExactArgs(0),
	Example: "resolution resolve ipfs-hash -d brad.crypto",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ReturnedValue, err = SelectedNamingService.IpfsHash(Domain)
		if err != nil {
			log.Fatal(err)
		}
	},
}
