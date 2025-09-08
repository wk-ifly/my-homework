package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	// Make sure the following import points to the generated Go bindings for your contract,
	// and that this package does not import your main package to avoid import cycles.
	counter "github.com/wk-ifly/my-homework/counter"
)

func main() {
	//连接sepolia测试网络
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/06c5cf58ca5e49bb88f7ca9be2abc9ab")
	if err != nil {
		log.Fatal(err)
	}

	//设置合约地址
	address := common.HexToAddress("0xfD43820363f414F1B968e9e7fE95BEcf3b3e0985")

	//实例化合约
	counterContract, err := counter.NewCounter(address, client)
	if err != nil {
		log.Fatal(err)
	}
	// 根据私钥hex生成esdsa
	privateKey, err := crypto.HexToECDSA("your private key")
	if err != nil {
		log.Fatal(err)
	}

	// 设置交易签名选项
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}

	//调用合约的getCount方法
	count, err := counterContract.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current count:", count)
	//调用合约的increment方法
	tx, err := counterContract.Increment(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction sent:", tx.Hash().Hex())
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("交易确认失败: %v", err)
	}
	tx, err = counterContract.Increment(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction sent:", tx.Hash().Hex())
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("交易确认失败: %v", err)
	}
	//再次调用合约的getCount方法
	count, err = counterContract.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("New count:", count)
	//调用合约的decrement方法
	tx, err = counterContract.Decrement(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Transaction sent:", tx.Hash().Hex())
	_, err = bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatalf("交易确认失败: %v", err)
	}
	//再次调用合约的getCount方法
	count1, err := counterContract.GetCount(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Final count:", count1)
}
