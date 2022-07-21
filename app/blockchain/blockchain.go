package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"smartvoting/app/config"
	api "smartvoting/app/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetAccountAuth(client *ethclient.Client, privateKeyAddress string) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(privateKeyAddress)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Invalid key")
	}
	//ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("nounce= ", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = big.NewInt(1000000)
	return auth
}
func Connect() *ethclient.Client {
	conf, _ := config.Init()
	client, err := ethclient.Dial(conf.Blockhain.Host + ":" + conf.Blockhain.Port)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func Init(conf config.Conf) *api.Api {
	client := Connect()
	auth := GetAccountAuth(client, conf.Blockhain.Secret_key)

	address, tx, instance, err := api.DeployApi(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Hex())
	fmt.Println("instance : ", instance)
	fmt.Println("tx : ", tx.Hash().Hex())

	conn, err := api.NewApi(common.HexToAddress(address.Hex()), client)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
