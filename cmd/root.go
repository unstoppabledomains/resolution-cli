package cmd

import (
	"fmt"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	userLicense string
	provider    string
	domain      string

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

func init() {
	cobra.OnInitialize(initConfig)

	// ryan <ryan@unstoppabledomains.com>, kiryl <kiryl@unstoppabledomains.com>, johnny <johnny@unstoppabledomains>")
	rootCmd.PersistentFlags().StringVarP(&provider, "provider", "p", "", "Ethereum JSON RPC endpoint to retrieve records from")
	viper.BindPFlag("provider", rootCmd.PersistentFlags().Lookup("provider"))
	setProviderCmd.MarkPersistentFlagRequired("provider")

	rootCmd.AddCommand(providerCmd)
	rootCmd.AddCommand(setProviderCmd)
	rootCmd.AddCommand(resolveCmd)

	resolveCmd.Flags().StringVarP(&domain, "domain", "d", "", "Domain to resolve (required)")
	resolveCmd.MarkFlagRequired("domain")
}

func initConfig() {
	// Find home directory.
	home, err := homedir.Dir()
	cobra.CheckErr(err)

	filename := "cobra"
	viper.AddConfigPath(home)
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		viper.WriteConfigAs(home + "/" + filename)
	} else {
		provider = viper.GetString("provider")
	}
}
