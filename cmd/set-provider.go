package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setProviderCmd = &cobra.Command{
	Use:   "set-provider [ethereum-url]",
	Short: "Sets the Ethereum provider",
	Long:  "Sets Ethereum provider. Default is Unstoppable Domains Infura provider",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set("provider", args[0])
		viper.WriteConfig()
		fmt.Println("Successfully set provider to " + args[0])
	},
}
