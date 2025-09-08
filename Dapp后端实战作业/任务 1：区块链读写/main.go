package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/06c5cf58ca5e49bb88f7ca9be2abc9ab")
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block hash:", block.Hash().Hex())
	fmt.Println("timestamp:", block.Time())
	fmt.Println("number of txs:", len(block.Transactions()))
	privatekey, err := crypto.HexToECDSA("your private key")
	if err != nil {
		log.Fatal(err)
	}
	publickeyecdsa := privatekey.Public()
	publickey := publickeyecdsa.(*ecdsa.PublicKey)
	fromaddress := crypto.PubkeyToAddress(*publickey)
	fmt.Println("fromAddress hex:", fromaddress.Hex())
	toaddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	value := big.NewInt(1000000000) // in wei (1 gwei = 10^-9 eth)
	gaslimit := uint64(21000)       // in units
	gasprice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var data []byte
	nonce, err := client.PendingNonceAt(context.Background(), fromaddress)
	if err != nil {
		log.Fatal(err)
	}
	tx := types.NewTransaction(nonce, toaddress, value, gaslimit, gasprice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)

	}
	signedtx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privatekey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedtx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", signedtx.Hash().Hex())
}
