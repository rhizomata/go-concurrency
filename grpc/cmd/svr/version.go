package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Rhizomata",
	Long:  `Rhizomata's version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Rhizomata version 0.0.1")
	},
}
