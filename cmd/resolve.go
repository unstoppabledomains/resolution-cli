package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve a domain",
	Long:  "Resolve records of a domain. Domain must be specified",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Resolving " + domain)
	},
}
