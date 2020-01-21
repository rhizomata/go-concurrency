package main

import (
	"fmt"
	"github.com/rhizomata/go-concurrency/grpc/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strconv"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test Rhizomata",
	Args: cobra.ExactArgs(2),
	Run:   testRun,
}

func testRun(cmd *cobra.Command, args []string) {
	serverURLs := viper.GetStringSlice("server-urls")
	totalCount, _ := strconv.Atoi(args[0])
	concurrent, _ := strconv.Atoi(args[1])

	testDriver := &client.TestDriver{}
	testDriver.SendHelloTest(serverURLs, totalCount, concurrent)
	fmt.Println("TEST to ", serverURLs)
}
