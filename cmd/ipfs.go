package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var ipfsCmd = &cobra.Command{
	Use:   "ipfs",
	Short: "Resolve ipfs records of a domain",
	Long:  "Resolve ipfs records of a domain",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Resolving ipfs hash of %s\n", domain)
		ipfsHash, err := CNS.IpfsHash(domain)
		if err != nil {
			log.Fatal("Error connecting to provider: " + err.Error())
		} else {
			fmt.Printf("%s ipfs hash: %s\n", domain, ipfsHash)

		}
	},
}
