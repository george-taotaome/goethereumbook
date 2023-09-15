package cmd

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/spf13/cobra"
)

var curBlock int64
var runBlock bool
var runTransaction bool
var runTransfer bool
var runTransferToken bool
var runSubscribe bool
var runRawTransaction bool
var runSendRawTransaction bool

// Transaction
var chapter3Cmd = &cobra.Command{
	Use:   "chapter3",
	Short: "Demo code for chapter 3: 交易",

	Run: func(cmd *cobra.Command, args []string) {
		client, err := ethclient.Dial("http://localhost:8545")
		if err != nil {
			log.Fatal(err)
		}

		// 生成block 1
		// ETH转账：以太币数量，gas限额，gas价格，一个随机数(nonce)，接收地址以及可选择性的添加的数据
		if runTransfer {
			// 加载私钥
			privateKey, err := crypto.HexToECDSA("f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5") // ganache-cli
			// privateKey, err := crypto.HexToECDSA("294dae214d4ce7110a0024a565e736ace82ba9620d4a1a62548b7d4e97d38731") // ganache
			if err != nil {
				log.Fatal(err)
			}

			// 帐户的公共地址
			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
			}

			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

			// 读取我们应该用于帐户交易的随机数
			nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Fatal(err)
			}

			// 设置我们将要转移的ETH数量。
			value := big.NewInt(1000000000000000000) // in wei (1 eth)

			// ETH转账的燃气应设上限为“21000”单位。
			gasLimit := uint64(21000) // in units

			// 燃气价格必须以wei为单位设定。 在撰写本文时，将在一个区块中比较快的打包交易的燃气价格为30 gwei。
			// gasPrice := big.NewInt(30000000000) // in wei (30 gwei)

			// 然而，燃气价格总是根据市场需求和用户愿意支付的价格而波动的，因此对燃气价格进行硬编码有时并不理想。 go-ethereum客户端提供SuggestGasPrice函数，用于根据'x'个先前块来获得平均燃气价格。
			gasPrice, err := client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			toAddress := common.HexToAddress("0x68dB32D26d9529B2a142927c6f1af248fc6Ba7e9") // ganache-cli
			// toAddress := common.HexToAddress("0x028EA99Fe457B9Ad405883b9f501cab9a267150F") // ganache

			fmt.Println("tx is:", nonce, fromAddress, value, gasLimit, gasPrice, toAddress)
			// 导入go-ethereumcore/types包并调用NewTransaction来生成我们的未签名以太坊事务，这个函数需要接收nonce，地址，值，燃气上限值，燃气价格和可选发的数据。
			// 发送ETH的数据字段为“nil”。 在与智能合约进行交互时，我们将使用数据字段，仅仅转账以太币是不需要数据字段的。
			tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

			// 下一步是使用发件人的私钥对事务进行签名。
			// 为此，我们调用SignTx方法，该方法接受一个未签名的事务和我们之前构造的私钥。 SignTx方法需要EIP155签名者，这个也需要我们先从客户端拿到链ID。
			networkID, err := client.NetworkID(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			chainID, err := client.ChainID(context.Background())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("chainID: ", networkID, chainID)

			signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
			if err != nil {
				log.Fatal(err)
			}

			// 现在我们终于准备通过在客户端上调用“SendTransaction”来将已签名的事务广播到整个网络。
			err = client.SendTransaction(context.Background(), signedTx)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent:
		}

		// 查询区块
		if runBlock {
			// 调用客户端的HeaderByNumber来返回有关一个区块的头信息, 传入nil，它将返回最新的区块头
			header, err := client.HeaderByNumber(context.Background(), nil)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(header.Number.String())

			// 调用客户端的BlockByNumber方法来获得完整区块。您可以读取该区块的所有内容和元数据，例如，区块号，区块时间戳，区块摘要，区块难度以及交易列表等等。
			// ganache cli客户端启动后，需要先执行go run main.go chapter3 -r，生成第一个区块1，才能查询到
			blockNumber := big.NewInt(curBlock)
			block, err := client.BlockByNumber(context.Background(), blockNumber)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Number().Uint64())     // 1
			fmt.Println(block.Time())                // 1692874584
			fmt.Println(block.Difficulty().Uint64()) // 0
			fmt.Println(block.Hash().Hex())          // 0x277ae95b482a82ae0a5d88eb3b6ddad136b152b27b234d6740f032ed6f895a07
			fmt.Println(len(block.Transactions()))   // 1

			// 调用客户端的 Transaction 方法来获取一个区块中的交易
			count, err := client.TransactionCount(context.Background(), block.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(count) // 144
		}

		// 查询 block =1 交易，应该安排在交易之后
		// ganache cli客户端启动后，需要先执行go run main.go chapter3 -r，生成第一个区块1，才能查询到
		if runTransaction {
			blockNumber := big.NewInt(curBlock)
			block, err := client.BlockByNumber(context.Background(), blockNumber)
			if err != nil {
				log.Fatal(err)
			}

			for _, tx := range block.Transactions() {
				fmt.Println(tx.Hash().Hex())
				fmt.Println(tx.Value().String())
				fmt.Println(tx.Gas())
				fmt.Println(tx.GasPrice().Uint64())
				fmt.Println(tx.Nonce())
				fmt.Println(tx.Data())
				fmt.Println(tx.To().Hex())

				chainID, err := client.ChainID(context.Background())
				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(chainID)      // 1692877060468
				fmt.Println(tx.ChainId()) // 1692877060468 ???

				// 通过交易获取发送者地址 发送方的地址是从交易的签名中恢复出来的
				if fromAddress, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), tx); err == nil {
					fmt.Println(fromAddress.Hex())
				}
				// 以下代码报错, AsMessage方法不存在了
				// if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
				// 	fmt.Println(msg.From().Hex())
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
			fmt.Println("TransactionInBlock", block.Hash().Hex()) // 0xe2e3e100b9da3c9bfa94955517285952311e7a23b0889cde1f571006a5f4e6ac
			blockHash := block.Hash()
			// blockHash := common.HexToHash("0xe2e3e100b9da3c9bfa94955517285952311e7a23b0889cde1f571006a5f4e6ac")
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
			txHash := common.HexToHash("0xa6e572b4298eca0fe306f932c8a614974370a29d05c253e12527fb15930793e5")
			tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("tx: ", tx.Value(), tx.To())
			fmt.Println(isPending) // false
		}

		// ERC20 Token转账
		if runTransferToken {
			privateKey, err := crypto.HexToECDSA("f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5")
			if err != nil {
				log.Fatal(err)
			}

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
			}

			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
			nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Fatal(err)
			}

			value := big.NewInt(0) // in wei (0 eth)
			gasPrice, err := client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			toAddress := common.HexToAddress("0x35bb6eF95c72bf4804334BB9d6A3c77Bef18d81B")
			tokenAddress := common.HexToAddress("0xC28614fEcD3109EFf192DD3cABc7ac9b82C7eD11")

			transferFnSignature := []byte("transfer(address,uint256)")
			// hash := sha3.NewKeccak256()
			// hash.Write(transferFnSignature)
			// methodID := hash.Sum(nil)[:4]
			hash := crypto.Keccak256(transferFnSignature)
			methodID := hash[:4]
			fmt.Println("methodID", hexutil.Encode(methodID))

			paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
			fmt.Println("paddedAddress", hexutil.Encode(paddedAddress))

			amount := new(big.Int)
			amount.SetString("1000000000000000000000", 10) // 1000 tokens
			paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
			fmt.Println("paddedAmount", hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000

			var data []byte
			data = append(data, methodID...)
			data = append(data, paddedAddress...)
			data = append(data, paddedAmount...)

			gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
				To:   &toAddress,
				Data: data,
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("gasLimit", gasLimit) // 23256

			tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)

			// chainID, err := client.NetworkID(context.Background())
			chainID, err := client.ChainID(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
			if err != nil {
				log.Fatal(err)
			}

			err = client.SendTransaction(context.Background(), signedTx)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
		}

		// 订阅，有新block时，打印出来
		if runSubscribe {
			client, err := ethclient.Dial("ws://localhost:8545")
			if err != nil {
				log.Fatal("client ", err)
			}

			headers := make(chan *types.Header)
			sub, err := client.SubscribeNewHead(context.Background(), headers)
			if err != nil {
				log.Fatal("sub ", err)
			}

			for {
				select {
				case err := <-sub.Err():
					log.Fatal("Err ", err)
				case header := <-headers:
					fmt.Println(header.Hash().Hex())

					block, err := client.BlockByHash(context.Background(), header.Hash())
					if err != nil {
						log.Fatal("block ", err)
					}

					fmt.Println(block.Hash().Hex())
					fmt.Println(block.Number().Uint64())
					fmt.Println(block.Time())
					fmt.Println(block.Nonce())
					fmt.Println(len(block.Transactions()))
				}
			}
		}

		// 创建原始交易事务
		if runRawTransaction {
			privateKey, err := crypto.HexToECDSA("f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5")
			if err != nil {
				log.Fatal(err)
			}

			publicKey := privateKey.Public()
			publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
			if !ok {
				log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
			}

			fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
			nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
			if err != nil {
				log.Fatal(err)
			}

			value := big.NewInt(1000000000000000000) // in wei (1 eth)
			gasLimit := uint64(21000)                // in units
			gasPrice, err := client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			toAddress := common.HexToAddress("0x35bb6eF95c72bf4804334BB9d6A3c77Bef18d81B")
			var data []byte
			tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

			chainID, err := client.ChainID(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			// 通过交易获取发送者地址 发送方的地址是从交易的签名中恢复出来的
			signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
			if err != nil {
				log.Fatal(err)
			}

			// ts := types.Transactions{signedTx}
			// rawTxBytes := ts.GetRlp(0) // GetRlp方法已经被弃用
			// rawTxHex := hex.EecodeString(rawTxBytes)

			// 将交易编码为RLP字节
			rawTxBytes, err := rlp.EncodeToBytes(signedTx)
			if err != nil {
				log.Fatal("编码交易为RLP出错:", err)
			}

			fmt.Printf("编码后的交易RLP字节: %x\n", rawTxBytes)
		}

		// 发送原始交易事务
		if runSendRawTransaction {
			rawTx := "0xf86d0484773594008252089435bb6ef95c72bf4804334bb9d6a3c77bef18d81b880de0b6b3a764000080820a95a07cb14afc640715ac92d055cfc9edbc38558ed415844c39402824f1636d3024b9a07c79acbb9821ff982c87e37a7b99d5fa936bffe8d6a170510454e14a1660269b"

			rawTxBytes, err := hexutil.Decode(rawTx)
			if err != nil {
				log.Fatal("rawTxBytes: ", err)
			}

			tx := new(types.Transaction)
			rlp.DecodeBytes(rawTxBytes, &tx)

			err = client.SendTransaction(context.Background(), tx)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("tx sent: %s", tx.Hash().Hex())
		}
	},
}

func init() {
	rootCmd.AddCommand(chapter3Cmd)

	chapter3Cmd.Flags().Int64VarP(&curBlock, "cur", "c", 1, "block number")

	chapter3Cmd.Flags().BoolVarP(&runTransfer, "transfer", "r", false, "run transfer demo, generate block 1")
	chapter3Cmd.Flags().BoolVarP(&runBlock, "block", "b", false, "get block 1 info")
	chapter3Cmd.Flags().BoolVarP(&runTransaction, "transaction", "t", false, "get transaction info from block 1")
	chapter3Cmd.Flags().BoolVarP(&runTransferToken, "transferToken", "o", false, "run transfer token demo")
	chapter3Cmd.Flags().BoolVarP(&runSubscribe, "subscribe", "s", false, "run subscribe demo")
	chapter3Cmd.Flags().BoolVarP(&runRawTransaction, "rawTransaction", "w", false, "run raw transaction demo")
	chapter3Cmd.Flags().BoolVarP(&runSendRawTransaction, "sendRawTransaction", "e", false, "run send raw transaction demo")
}
