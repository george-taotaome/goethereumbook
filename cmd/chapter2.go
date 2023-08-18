package cmd

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/sha3"
)

var chapter2Cmd = &cobra.Command{
	Use:   "chapter2",
	Short: "Demo code for chapter 2: 以太坊账户",

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

		// 生成新钱包，需要导入go-ethereumcrypto包，该包提供用于生成随机私钥的GenerateKey方法。
		privateKey, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(`privateKey is:`, privateKey)
		//然后可以通过导入golangcrypto/ecdsa包并使用FromECDSA方法将其转换为字节
		privateKeyBytes := crypto.FromECDSA(privateKey)
		fmt.Println(`privateKeyBytes is:`, hexutil.Encode(privateKeyBytes)[2:])
		//这就是用于签署交易的私钥，将被视为密码，永远不应该被共享给别人，因为谁拥有它可以访问你的所有资产。
		//由于公钥是从私钥派生的，因此go-ethereum的加密私钥具有一个返回公钥的Public方法
		publicKey := privateKey.Public()
		fmt.Println(`publicKey is:`, publicKey)
		//将其转换为十六进制的过程与我们使用转化私钥的过程类似。 我们剥离了0x和前2个字符04，它始终是EC前缀，不是必需的
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}
		publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
		fmt.Println(`publicKeyBytes is:`, hexutil.Encode(publicKeyBytes)[4:])

		//拥有公钥，就可以轻松生成你经常看到的公共地址。 为了做到这一点，go-ethereum加密包有一个PubkeyToAddress方法，它接受一个ECDSA公钥，并返回公共地址。
		address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
		fmt.Println(`address is:`, address)
		// 公共地址其实就是公钥的Keccak-256哈希，然后我们取最后40个字符（20个字节）并用“0x”作为前缀。 以下是使用 golang.org/x/crypto/sha3 的 Keccak256函数手动完成的方法。
		hash := sha3.NewLegacyKeccak256()
		hash.Write(publicKeyBytes[1:])
		fmt.Println(`keccak-address is:`, hexutil.Encode(hash.Sum(nil)[12:]))
	},
}

func init() {
	rootCmd.AddCommand(chapter2Cmd)
}
