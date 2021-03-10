package cmd

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve record(s) of domain",
	Long:  "Resolve record(s) of a domain. Domain must be specified",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !strings.HasSuffix(domain, (".crypto")) && !strings.HasSuffix(domain, (".zil")) {
			log.Fatal("Domain must end with .crypto or .zil")
		}
	},
}
