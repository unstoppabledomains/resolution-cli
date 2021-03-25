package cmd

import (
	"fmt"
	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/unstoppabledomains/resolution-go"
	"github.com/unstoppabledomains/resolution-go/namingservice"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const ethereumUrlKey = "ETHEREUM_PROVIDER_URL"
const zilliqaUrlKey = "ZILLIQA_PROVIDER_URL"

var Domain string
var NamingServices map[string]resolution.NamingService
var ReturnedValue interface{}
var SelectedNamingService resolution.NamingService

var ethereumProviderUrlFlag string
var zilliqaProviderUrlFlag string

var (
	rootCmd = &cobra.Command{
		Use:   "resolution",
		Short: "Resolution is a simple blockchain Domain resolution tool",
		Long: `A simple blockchain Domain resolution cli tool built by the Unstoppable Domains team. 
Complete documentation is available at http://docs.unstoppabledomains.com`,
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
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	rootCmd.PersistentFlags().StringVar(&ethereumProviderUrlFlag, "ethereum-provider-url", "", "Ethereum JSON RPC endpoint url (could be set via RESOLUTION_ETHEREUM_PROVIDER_URL environment variable)")
	rootCmd.PersistentFlags().StringVar(&zilliqaProviderUrlFlag, "zilliqa-provider-url", "", "Zilliqa JSON RPC endpoint url (could be set via RESOLUTION_ZILLIQA_PROVIDER_URL environment variable)")
	resolveCmd.PersistentFlags().StringVarP(&Domain, "domain", "d", "", ".crypto or .zil domain to resolve (required)")
	err := resolveCmd.MarkPersistentFlagRequired("domain")
	if err != nil {
		log.Fatal(err)
	}
	resolveCmd.AddCommand(addrCmd, ipfsCmd, recordsCmd, httpUrlCmd, ownerCmd, resolverCmd, emailCmd, dnsCmd, addrVersionCmd)
	rootCmd.AddCommand(resolveCmd)
}

func initConfig() {
	viper.SetEnvPrefix("RESOLUTION")
	viper.AutomaticEnv()
	if ethereumProviderUrlFlag != "" {
		viper.Set(ethereumUrlKey, ethereumProviderUrlFlag)
	}
	if zilliqaProviderUrlFlag != "" {
		viper.Set(zilliqaUrlKey, zilliqaProviderUrlFlag)
	}
	initNamingServices()
}

func initNamingServices() {
	var err error
	ethereumUrl := viper.GetString(ethereumUrlKey)
	zilliqaUrl := viper.GetString(zilliqaUrlKey)
	cnsBuilder := resolution.NewCnsBuilder()
	znsBuilder := resolution.NewZnsBuilder()
	if ethereumUrl != "" {
		backend, err := ethclient.Dial(ethereumUrl)
		if err != nil {
			log.Fatalf("Error connecting to Ethereum provider. Provider: %v. Error: %v", ethereumUrl, err.Error())
		}
		cnsBuilder.SetContractBackend(backend)
	}
	cnsService, err := cnsBuilder.Build()
	if err != nil {
		log.Fatalf("Error with initiation CNS naming service. Provider: %v. Error: %v", ethereumUrl, err.Error())
	}
	if zilliqaUrl != "" {
		zilliqaProvider := provider.NewProvider(zilliqaUrl)
		znsBuilder.SetProvider(zilliqaProvider)
	}
	znsService, err := znsBuilder.Build()
	if err != nil {
		log.Fatalf("Error with initiation ZNS naming service. Provider: %v. Error: %v", zilliqaUrl, err.Error())
	}
	NamingServices = map[string]resolution.NamingService{namingservice.CNS: cnsService, namingservice.ZNS: znsService}
}
