package cmd

import (
	"encoding/hex"
	"log"
	"math/big"
	"strings"

	"github.com/spf13/cobra"
	"github.com/unstoppabledomains/resolution-go/v2/namingservice"
)

var unhashCmd = &cobra.Command{
	Use:     "unhash TOKEN_ID",
	Example: "resolution unhash 0x756e4e998dbffd803c21d23b06cd855cdc7a4b57706c95964a37e24b47c10fc9",
	Short:   "Unhash a token_id or namehash to its domain name",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var domainHash = args[0]
		if !strings.HasPrefix(domainHash, "0x") {
			var erc721TokenID big.Int
			erc721TokenID.SetString(domainHash, 10)
			domainHash = hex.EncodeToString(erc721TokenID.Bytes())
		}
		var err error
		uns := NamingServices[namingservice.UNS]
		ReturnedValue, err = uns.Unhash(domainHash)
		if err != nil {
			log.Fatal(err)
		}
	}}
