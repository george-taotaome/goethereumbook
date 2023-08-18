package cmd

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var chapter2Cmd = &cobra.Command{
	Use:   "chapter2",
	Short: "Demo code for chapter 2",

	Run: func(cmd *cobra.Command, args []string) {
		client, err := ethclient.Dial("http://localhost:8545")
		if err != nil {
			log.Fatal(err)
		}

		// 以太坊上的账户要么是钱包地址要么是智能合约地址。
		// 要使用go-ethereum的账户地址，您必须先将它们转化为go-ethereum中的common.Address类型。
		account := common.HexToAddress("0xE280029a7867BA5C9154434886c241775ea87e53")
		// fmt.Println(account)              // 0xE280029a7867BA5C9154434886c241775ea87e53
		// fmt.Println(account.Hex())        // 0xE280029a7867BA5C9154434886c241775ea87e53
		// fmt.Println(account.Hash().Hex()) // 0x000000000000000000000000e280029a7867ba5c9154434886c241775ea87e53
		// fmt.Println(account.Bytes())      // [226 128 2 154 120 103 186 92 145 84 67 72 134 194 65 119 94 168 126 83]

		// 读取一个账户的余额相当简单。调用客户端的BalanceAt方法，给它传递账户地址和可选的区块号。将区块号设置为nil将返回最新的余额。
		// 传区块号能让您读取该区块时的账户余额。区块号必须是big.Int类型。
		balance, err := client.BalanceAt(context.Background(), account, nil)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(balance) // 100000000000000000000 wei (100 ether)

		// 传递区块号能让您读取该区块时的账户余额。区块号必须是big.Int类型。
		// blockNumber := big.NewInt(5532993)
		// balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// fmt.Println(balanceAt)

		// 以太坊中的所有值都是以wei为单位的。wei是以太坊中的最小单位。1 ether = 10^18 wei。
		fbalance := new(big.Float)
		fbalance.SetString(balance.String())
		ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18))) // wei/10^18
		fmt.Println(ethValue)                                                  // 100 ether

		// 待处理的账户余额是指账户的余额加上所有待处理的交易的总和。
		pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(pendingBalance) // 100000000000000000000
	},
}

func init() {
	rootCmd.AddCommand(chapter2Cmd)
}
