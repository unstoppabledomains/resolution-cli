package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/unstoppabledomains/resolution-go"
	"github.com/unstoppabledomains/resolution-go/namingservice"
	"log"
)

var ReturnedValue interface{}
var SelectedNamingService resolution.NamingService

var resolveCmd = &cobra.Command{
	Use:     "resolve",
	Short:   "Resolve all known records of domain",
	Long:    "Resolve all known records of domain. Domain must be specified",
	Example: "resolution resolve -d brad.crypto",
	Args:    cobra.ExactArgs(0),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if Domain == "" {
			return
		}
		var err error
		namingServiceName, err := resolution.DetectNamingService(Domain)
		if err != nil {
			log.Fatal(err)
		}
		if NamingServices[namingServiceName] == nil {
			log.Fatalf("Naming service %v does not exist in initialized naming services. Supported services are: %v, %v", namingServiceName, namingservice.CNS, namingservice.ZNS)
		}
		SelectedNamingService = NamingServices[namingServiceName]
	},
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		ReturnedValue, err = SelectedNamingService.AllRecords(Domain)
		if err != nil {
			log.Fatal(err)
		}
		ReturnedValue = prepareMultiRecordsOutput(ReturnedValue)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		output, err := formatOutput(ReturnedValue)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(output)
	},
}
