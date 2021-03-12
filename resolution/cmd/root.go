package cmd

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/unstoppabledomains/resolution-go"
)

var (
	// Used for flags.
	userLicense string
	provider    string
	domain      string
	// CNS Resolution instance
	CNS *resolution.Cns
	// ZNS Resolution instance
	ZNS *resolution.Zns

	rootCmd = &cobra.Command{
		Use:   "resolution",
		Short: "Resolution is a simple blockchain domain resolution tool",
		Long: `A simple blockchain domain resolution cli tool built by the Unstoppable Domains team. 
Complete documentation is available at http://docs.unstoppabledomains.com`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func initCNS() {
	if provider == "" {
		cns, err := resolution.NewCnsWithDefaultBackend()
		if err != nil {
			log.Fatal("Error connecting to provider: " + err.Error())
		}
		CNS = cns
	} else {
		backend, err := ethclient.Dial(provider)
		if err != nil {
			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		}
		CNS, err = resolution.NewCns(backend)
		if err != nil {
			log.Fatal("Error connecting to provider: " + err.Error())
			return
		}
	}
}

func initZNS() {
	ZNS = resolution.NewZnsWithDefaultProvider()
}

func init() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&provider, "provider", "p", "", "Ethereum JSON RPC endpoint to retrieve records from. Overrides set-provider")
	viper.BindPFlag("provider", rootCmd.PersistentFlags().Lookup("provider"))
	setProviderCmd.MarkPersistentFlagRequired("provider")

	initCNS()
	initZNS()

	resolveCmd.PersistentFlags().StringVarP(&domain, "domain", "d", "", ".crypto or .zil domain to resolve")
	resolveCmd.MarkPersistentFlagRequired("domain")
	resolveCmd.AddCommand(addrCmd)
	resolveCmd.AddCommand(ipfsCmd)
	resolveCmd.AddCommand(recordsCmd)
	resolveCmd.AddCommand(allCmd)

	rootCmd.AddCommand(providerCmd)
	rootCmd.AddCommand(setProviderCmd)
	rootCmd.AddCommand(resolveCmd)

}

func initConfig() {
	viper.AutomaticEnv()
	provider = viper.GetString("RESOLUTION_PROVIDER")
}
