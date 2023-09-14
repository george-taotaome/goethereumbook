package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	// "yunlabs.com/goethereumbook/contracts/store"
	"yunlabs.com/goethereumbook/contracts/token"
)

func main() {
	// client, err := ethclient.Dial("https://rinkeby.infura.io")
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5")
	// privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
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

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// auth := bind.NewKeyedTransactor(privateKey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units, for Store
	auth.GasPrice = gasPrice

	// input := "1.0"
	// address, tx, instance, err := store.DeployStore(auth, client, input)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	auth.GasLimit = uint64(3000000) // in units，增加gas限额 for erc20, 6.1K需要多点gas

	name := "My Token"
	symbol := "MTK"
	decimals := uint8(18)
	totalSupply := big.NewInt(1000000000000000000)
	address, tx, instance, err := token.DeployToken(auth, client, name, symbol, decimals, totalSupply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0xf3585FCD969502624c6A8ACf73721d1fce214E83
	fmt.Println(tx.Hash().Hex()) // 0x79c08b3f9d93712f6fbe92975191eb4f856ac8a7b42cc2aa21e5da9104d7bce5

	_ = instance
}
