package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var ipfsCmd = &cobra.Command{
	Use:   "ipfs",
	Short: "Resolve ipfs records",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var ipfsHash string
		var err error
		if strings.HasSuffix(domain, (".crypto")) {
			ipfsHash, err = CNS.IpfsHash(domain)
		} else {
			ipfsHash, err = ZNS.IpfsHash(domain)
		}
		if err != nil {
			log.Fatal("Error connecting to provider: " + err.Error())
		} else {
			b, _ := json.Marshal(map[string]string{"ipfs": ipfsHash})

			fmt.Println(string(b))
		}
	},
}
