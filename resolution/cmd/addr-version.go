package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var ticker string
var version string

var addrVersionCmd = &cobra.Command{
	Use:     "addr-version -v TICKER_VERSION -c TICKER",
	Example: "resolution resolve addr-version -c USDT -v ERC20 -d udtestdev-usdt.crypto",
	Short:   "Resolve multi-chain address",
	Long: `
Some cryptocurrencies exists in multiple chains. 
This command helps resolve such currencies.
To find details about Unstoppable multi-chain standard follow the link:
https://docs.unstoppabledomains.com/domain-registry-essentials/records-reference#multi-chain-currencies
`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ReturnedValue, err = SelectedNamingService.AddrVersion(Domain, ticker, version)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	addrVersionCmd.Flags().StringVarP(&ticker, "currency", "c", "", "Multi-chain currency ticker (required)")
	addrVersionCmd.Flags().StringVarP(&version, "version", "v", "", "Multi-chain currency version (required)")
}
