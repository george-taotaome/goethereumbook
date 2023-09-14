package cmd

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

// Client
var chapter1Cmd = &cobra.Command{
	Use:   "chapter1",
	Short: "Demo code for chapter 1: 客户端",

	Run: func(cmd *cobra.Command, args []string) {
		// client, err := ethclient.Dial("https://cloudflare-eth.com")
		client, err := ethclient.Dial("http://localhost:8545")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("we have a connection")
		_ = client // we'll use this in the upcoming sections
	},
}

func init() {
	rootCmd.AddCommand(chapter1Cmd)
}
