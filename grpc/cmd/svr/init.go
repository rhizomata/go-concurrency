package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Rhizomata",
	RunE:  initFiles,
}

func initFiles(cmd *cobra.Command, args []string) error {
	err := viper.SafeWriteConfig()
	if err != nil {
		fmt.Println("Write Config ", err)
	}
	return err
}
