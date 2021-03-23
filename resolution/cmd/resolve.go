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
	Use:   "resolve",
	Short: "Resolve record(s) of Domain",
	Long:  "Resolve record(s) of a Domain. Domain must be specified",
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
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		output, err := formatOutput(ReturnedValue)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(output)
	},
}
