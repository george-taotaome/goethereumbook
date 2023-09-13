package cmd

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

// Client
var chapter3TokenCmd = &cobra.Command{
	Use:   "chapter3_token",
	Short: "Demo code for chapter 3: ERC20 Token转账",

	Run: func(cmd *cobra.Command, args []string) {
		// client, err := ethclient.Dial("https://cloudflare-eth.com")
		client, err := ethclient.Dial("http://localhost:8545")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("we have a connection")
		_ = client // we'll use this in the upcoming sections

		// 代币传输不需要传输ETH，因此将交易“值”设置为“0”。
		// value := big.NewInt(0)

		// 发送代币的地址
		// toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	},
}

func init() {
	rootCmd.AddCommand(chapter3TokenCmd)
}
