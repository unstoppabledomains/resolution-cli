package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "Displays the Ethereum provider",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if provider == "" {
			fmt.Println("Provider is set to Unstoppable Domains default provider")
		} else {
			fmt.Printf("Provider is set to %s\n", provider)
		}
	},
}
