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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .rhizome.yaml)")
	rootCmd.PersistentFlags().Uint("port", 12345, "server listen port")
	viper.BindPFlag("listen-port", rootCmd.PersistentFlags().Lookup("port"))
	viper.SetDefault("listen-port", 12345)

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(startCmd)

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
