package cmd

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"yunlabs.com/goethereumbook/contracts/store"
	"yunlabs.com/goethereumbook/contracts/token"
)

var curAddress string
var runLoad bool
var runSetItem bool
var runCodeAt bool
var runERC20 bool

// Client
var chapter4Cmd = &cobra.Command{
	Use:   "chapter4",
	Short: "Demo code for chapter 4: 智能合约",

	Run: func(cmd *cobra.Command, args []string) {
		client, err := ethclient.Dial("http://localhost:8545")
		if err != nil {
			log.Fatal(err)
		}

		if runLoad {
			address := common.HexToAddress("0x2e144aF3Bde9B518C7C65FBE170c07c888f1fF1a")
			// 从地址加载合约
			instance, err := store.NewStore(address, client)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("contract is loaded")

			// 查询合约版本
			version, err := instance.Version(nil)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("version: ", version) // "1.0"
		}

		// 写入智能合约
		if runSetItem {
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

			gasPrice, err := client.SuggestGasPrice(context.Background())
			if err != nil {
				log.Fatal(err)
			}

			auth := bind.NewKeyedTransactor(privateKey)
			auth.Nonce = big.NewInt(int64(nonce))
			auth.Value = big.NewInt(0)     // in wei
			auth.GasLimit = uint64(300000) // in units
			auth.GasPrice = gasPrice

			address := common.HexToAddress("0x2e144aF3Bde9B518C7C65FBE170c07c888f1fF1a")
			instance, err := store.NewStore(address, client)
			if err != nil {
				log.Fatal(err)
			}

			key := [32]byte{}
			value := [32]byte{}
			copy(key[:], []byte("foo"))
			copy(value[:], []byte("bar"))

			tx, err := instance.SetItem(auth, key, value)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("tx sent: %s \n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

			result, err := instance.Items(nil, key)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(result[:])) // "bar"
		}

		// 读取智能合约的字节码
		if runCodeAt {
			contractAddress := common.HexToAddress("0x2e144aF3Bde9B518C7C65FBE170c07c888f1fF1a")
			bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(hex.EncodeToString(bytecode))
		}

		// 读取ERC20代币
		if runERC20 {
			// My Token/MTK address on ganache
			tokenAddress := common.HexToAddress("0xC28614fEcD3109EFf192DD3cABc7ac9b82C7eD11")
			instance, err := token.NewToken(tokenAddress, client)
			if err != nil {
				log.Fatal("token", err)
			}

			address := common.HexToAddress(curAddress)
			bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
			if err != nil {
				log.Fatal("BalanceOf", err)
			}

			name, err := instance.Name(&bind.CallOpts{})
			if err != nil {
				log.Fatal(err)
			}

			symbol, err := instance.Symbol(&bind.CallOpts{})
			if err != nil {
				log.Fatal(err)
			}

			decimals, err := instance.Decimals(&bind.CallOpts{})
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("name: %s\n", name)         // "name: Golem Network"
			fmt.Printf("symbol: %s\n", symbol)     // "symbol: GNT"
			fmt.Printf("decimals: %v\n", decimals) // "decimals: 18"

			fmt.Printf("wei: %s\n", bal) // "wei: 74605500647408739782407023"

			fbal := new(big.Float)
			fbal.SetString(bal.String())
			value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))

			fmt.Printf("balance: %f", value)
		}

	},
}

func init() {
	rootCmd.AddCommand(chapter4Cmd)

	chapter4Cmd.Flags().StringVarP(&curAddress, "address", "a", "0xE280029a7867BA5C9154434886c241775ea87e53", "account address")

	chapter4Cmd.Flags().BoolVarP(&runLoad, "load", "l", false, "load contract and query version")
	chapter4Cmd.Flags().BoolVarP(&runSetItem, "set", "s", false, "set item")
	chapter4Cmd.Flags().BoolVarP(&runCodeAt, "code", "c", false, "get code")
	chapter4Cmd.Flags().BoolVarP(&runERC20, "erc20", "e", false, "erc20 token")
}
