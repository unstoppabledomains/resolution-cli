package cmd

import (
	"log"
	"math/big"
	"strings"

	"github.com/spf13/cobra"
)

var namehashCmd = &cobra.Command{
	Use:     "namehash",
	Example: "resolution namehash -d brad.crypto",
	Short:   "Get namehash of domain",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		namehash, err := SelectedNamingService.Namehash(Domain)
		if err != nil {
			log.Fatal(err)
		}
		ReturnedValue = namehash
		if Decimal {
			tokenId := new(big.Int)
			tokenId.SetString(strings.Replace(namehash, "0x", "", -1), 16)
			ReturnedValue = tokenId
		}
	}}
