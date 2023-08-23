package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

var runBlock bool
var runTransaction bool

// Transaction
var chapter3Cmd = &cobra.Command{
	Use:   "chapter3",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		client, err := ethclient.Dial("https://cloudflare-eth.com")
		// client, err := ethclient.Dial("http://localhost:8545")
		if err != nil {
			log.Fatal(err)
		}

		if runBlock {
			// 调用客户端的HeaderByNumber来返回有关一个区块的头信息, 传入nil，它将返回最新的区块头
			header, err := client.HeaderByNumber(context.Background(), nil)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(header.Number.String()) // 17975292

			// 调用客户端的BlockByNumber方法来获得完整区块。您可以读取该区块的所有内容和元数据，例如，区块号，区块时间戳，区块摘要，区块难度以及交易列表等等。
			blockNumber := big.NewInt(5671744)
			block, err := client.BlockByNumber(context.Background(), blockNumber)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Number().Uint64())     // 5671744
			fmt.Println(block.Time())                // 1527211625
			fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
			fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
			fmt.Println(len(block.Transactions()))   // 144

			// 调用客户端的 Transaction 方法来获取一个区块中的交易
			count, err := client.TransactionCount(context.Background(), block.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(count) // 144
		}

		if runTransaction {
			blockNumber := big.NewInt(5671744)
			block, err := client.BlockByNumber(context.Background(), blockNumber)
			if err != nil {
				log.Fatal(err)
			}

			for _, tx := range block.Transactions() {
				fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
				fmt.Println(tx.Value().String())    // 10000000000000000
				fmt.Println(tx.Gas())               // 105000
				fmt.Println(tx.GasPrice().Uint64()) // 102000000000
				fmt.Println(tx.Nonce())             // 110644
				fmt.Println(tx.Data())              // []
				fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

				chainID, err := client.NetworkID(context.Background())
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(chainID)      // 1
				fmt.Println(tx.ChainId()) // 0 ???

				// 通过交易获取发送者地址 发送方的地址是从交易的签名中恢复出来的
				if fromAddress, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx); err == nil {
					fmt.Println(fromAddress.Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
				}
				// 以下代码报错, AsMessage方法不存在了
				// if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
				// 	fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
				// }

				receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(receipt.Status) // 1
				fmt.Println(receipt.Logs)   // []

				if tx.Hash().Hex() == "0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2" {
					break //Test 仅打印第一个交易
				}
			}

			// 在不获取块的情况下遍历事务的另一种方法是调用客户端的TransactionInBlock方法。 此方法仅接受块哈希和块内事务的索引值。 您可以调用TransactionCount来了解块中有多少个事务。
			// fmt.Println("TransactionInBlock", block.Hash().Hex()) // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
			// blockHash := block.Hash()
			blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
			count, err := client.TransactionCount(context.Background(), blockHash)
			if err != nil {
				log.Fatal(err)
			}

			for idx := uint(0); idx < count; idx++ {
				tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
				if tx.Hash().Hex() == "0x36368eb4665367100bcb46427e8ac39b7873abfca2015116c478f84642a8812d" {
					break //Test 仅打印前三个交易
				}
			}

			// 还可以使用TransactionByHash在给定具体事务哈希值的情况下直接查询单个事务
			txHash := common.HexToHash("0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2")
			tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
			fmt.Println(isPending)       // false

		}

	},
}

func init() {
	rootCmd.AddCommand(chapter3Cmd)

	chapter3Cmd.Flags().BoolVarP(&runBlock, "block", "b", false, "run block demo")
	chapter3Cmd.Flags().BoolVarP(&runTransaction, "transaction", "t", false, "run transaction demo")
}
