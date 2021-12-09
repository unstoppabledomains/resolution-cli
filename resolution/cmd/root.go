package cmd

import (
	"fmt"
	"log"

	"github.com/Zilliqa/gozilliqa-sdk/provider"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/unstoppabledomains/resolution-go/v2"
	"github.com/unstoppabledomains/resolution-go/v2/namingservice"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const ethereumUrlKey = "ETHEREUM_PROVIDER_URL"
const ethereumL2UrlKey = "ETHEREUM_L2_PROVIDER_URL"
const ethereumNetworkIdKey = "ETHEREUM_NETWORK_ID"
const ethereumL2NetworkIdKey = "ETHEREUM_L2_NETWORK_ID"
const zilliqaUrlKey = "ZILLIQA_PROVIDER_URL"

var Domain string
var Decimal bool
var NamingServices map[string]resolution.NamingService
var ReturnedValue interface{}
var SelectedNamingService resolution.NamingService

var ethereumProviderUrlFlag string
var ethereumL2ProviderUrlFlag string
var ethereumNetworkIdFlag string
var ethereumL2NetworkIdFlag string
var zilliqaProviderUrlFlag string

var (
	rootCmd = &cobra.Command{
		Version: "v1.1.0",
		Use:     "resolution",
		Short:   "Resolution is a simple blockchain Domain resolution tool",
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
				log.Fatalf("Naming service %v does not exist in initialized naming services. Supported services are: %v, %v", namingServiceName, namingservice.UNS, namingservice.ZNS)
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
	rootCmd.PersistentFlags().StringVar(&ethereumL2ProviderUrlFlag, "ethereum-l2-provider-url", "", "Ethereum L2 JSON RPC endpoint url (could be set via RESOLUTION_ETHEREUM_L2_PROVIDER_URL environment variable)")
	rootCmd.PersistentFlags().StringVar(&ethereumNetworkIdFlag, "ethereum-network-id", "", "Ethereum network id (could be set via RESOLUTION_ETHEREUM_NETWORK_ID environment variable)")
	rootCmd.PersistentFlags().StringVar(&ethereumL2NetworkIdFlag, "ethereum-l2-network-id", "", "Ethereum L2 network id (could be set via RESOLUTION_ETHEREUM_L2_NETWORK_ID environment variable)")
	rootCmd.PersistentFlags().StringVar(&zilliqaProviderUrlFlag, "zilliqa-provider-url", "", "Zilliqa JSON RPC endpoint url (could be set via RESOLUTION_ZILLIQA_PROVIDER_URL environment variable)")
	resolveCmd.PersistentFlags().StringVarP(&Domain, "domain", "d", "", ".crypto or .zil domain to resolve (required)")
	namehashCmd.PersistentFlags().StringVarP(&Domain, "domain", "d", "", ".crypto or .zil domain to resolve (required)")
	namehashCmd.PersistentFlags().BoolVar(&Decimal, "decimal", false, "return namehash as decimal")
	err := resolveCmd.MarkPersistentFlagRequired("domain")
	if err != nil {
		log.Fatal(err)
	}
	err = namehashCmd.MarkPersistentFlagRequired("domain")
	if err != nil {
		log.Fatal(err)
	}
	resolveCmd.AddCommand(addrCmd, ipfsCmd, recordsCmd, httpUrlCmd, ownerCmd, resolverCmd, emailCmd, dnsCmd, addrVersionCmd, locationCmd)
	rootCmd.AddCommand(resolveCmd)
	rootCmd.AddCommand(unhashCmd)
	rootCmd.AddCommand(namehashCmd)
}

func initConfig() {
	viper.SetEnvPrefix("RESOLUTION")
	viper.AutomaticEnv()
	if ethereumProviderUrlFlag != "" {
		viper.Set(ethereumUrlKey, ethereumProviderUrlFlag)
	}
	if ethereumL2ProviderUrlFlag != "" {
		viper.Set(ethereumL2UrlKey, ethereumL2ProviderUrlFlag)
	}
	if ethereumNetworkIdFlag != "" {
		viper.Set(ethereumNetworkIdKey, ethereumNetworkIdFlag)
	}
	if ethereumL2NetworkIdFlag != "" {
		viper.Set(ethereumL2NetworkIdKey, ethereumL2NetworkIdFlag)
	}
	if zilliqaProviderUrlFlag != "" {
		viper.Set(zilliqaUrlKey, zilliqaProviderUrlFlag)
	}
	initNamingServices()
}

func initNamingServices() {
	var err error
	ethereumUrl := viper.GetString(ethereumUrlKey)
	ethereumL2Url := viper.GetString(ethereumL2UrlKey)
	ethereumNetworkId := viper.GetString(ethereumNetworkIdKey)
	ethereumL2NetworkId := viper.GetString(ethereumL2NetworkIdKey)
	zilliqaUrl := viper.GetString(zilliqaUrlKey)
	unsBuilder := resolution.NewUnsBuilder()
	znsBuilder := resolution.NewZnsBuilder()
	if ethereumUrl != "" || ethereumL2Url != "" {
		if ethereumUrl == "" || ethereumL2Url == "" {
			log.Fatalf("Specify both L1 and L2 ethereum url when defining your own networks")
		}
		if ethereumNetworkId != "mainnet" && ethereumNetworkId != "rinkeby" {
			log.Fatalf("Specify ethereum network id ('mainnet' or 'rinkeby')")
		}
		if ethereumL2NetworkId != "polygon" && ethereumL2NetworkId != "mumbai" {
			log.Fatalf("Specify ethereum L2 network id ('polygon' or 'mumbai')")
		}
		backend, err := ethclient.Dial(ethereumUrl)
		if err != nil {
			log.Fatalf("Error connecting to Ethereum provider. Provider: %v. Error: %v", ethereumUrl, err.Error())
		}
		unsBuilder.SetContractBackend(backend).SetEthereumNetwork(ethereumNetworkId)
		backendL2, err := ethclient.Dial(ethereumL2Url)
		if err != nil {
			log.Fatalf("Error connecting to Ethereum L2 provider. Provider: %v. Error: %v", ethereumL2Url, err.Error())
		}
		unsBuilder.SetL2ContractBackend(backendL2).SetL2EthereumNetwork(ethereumL2NetworkId)
	} else {
		unsBuilder.SetEthereumNetwork("mainnet").SetL2EthereumNetwork("polygon")
	}
	unsService, err := unsBuilder.Build()
	if err != nil {
		log.Fatalf("Error with initiation UNS naming service. Provider: %v. Error: %v", ethereumUrl, err.Error())
	}
	if zilliqaUrl != "" {
		zilliqaProvider := provider.NewProvider(zilliqaUrl)
		znsBuilder.SetProvider(zilliqaProvider)
	}
	znsService, err := znsBuilder.Build()
	if err != nil {
		log.Fatalf("Error with initiation ZNS naming service. Provider: %v. Error: %v", zilliqaUrl, err.Error())
	}
	NamingServices = map[string]resolution.NamingService{namingservice.UNS: unsService, namingservice.ZNS: znsService}
}
