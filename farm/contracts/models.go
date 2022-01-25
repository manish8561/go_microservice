package contracts

import (
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	// "github.com/autocompound/docker_backend/farm/common"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

// You could input string which will be saved in database returning with error info
// 	if err := FindOne(&farmModel); err != nil { ... }
func GetContract() error {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/6839cbc8fb81452da72f56af6deebb7a")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	
	
	token, err := NewToken(common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	name, err := token.Name(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	fmt.Println("Token name:", name)

	bal, err := token.BalanceOf(nil, common.HexToAddress("0xbebc44782c7db0a1a60cb6fe97d0b483032ff1c7"))
	if err != nil {
		log.Fatalf("Failed to retrieve token name: %v", err)
	}
	fmt.Println("Token balance:", bal)
	return err
}
