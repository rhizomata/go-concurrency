package main

import (
	"fmt"
	"github.com/rhizomata/go-concurrency/grpc/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Rhizome Server",
	Run:   startServer,
}

func startServer(cmd *cobra.Command, args []string) {
	port := viper.GetUint("listen-port")
	fmt.Println("Rhizome server START on ", port)

	server.Start(port)
}
