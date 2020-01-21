package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "rhizome-srv",
		Short: "rhizomata server",
		Long:  `Rhizomata server CLI commands.`,
	}
)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .rhizome.toml)")
	rootCmd.PersistentFlags().StringArrayP("server", "", []string{"127.0.0.1:12345"}, "server url - ip:port")
	viper.BindPFlag("server-urls", rootCmd.PersistentFlags().Lookup("server"))
	viper.SetDefault("server-urls", []string{"127.0.0.1:12345"})

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(testCmd)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".rhizome" (without extension).
		viper.AddConfigPath("./")
		viper.SetConfigName("rhizome")
		viper.SetConfigType("toml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
