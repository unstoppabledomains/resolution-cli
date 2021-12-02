package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/unstoppabledomains/resolution-go/v2/dnsrecords"
)

var dnsCmd = &cobra.Command{
	Use:     "dns",
	Short:   "Resolve dns records attached to domain name",
	Args:    cobra.MinimumNArgs(1),
	Example: "resolution resolve dns A AAAA CNAME -d udtestdev-dns-cname.crypto",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var dnsRecordsKeys []dnsrecords.Type
		for _, argument := range args {
			dnsRecordsKeys = append(dnsRecordsKeys, dnsrecords.Type(argument))
		}
		ReturnedValue, err = SelectedNamingService.DNS(Domain, dnsRecordsKeys)
		if err != nil {
			log.Fatal(err)
		}
	},
}
