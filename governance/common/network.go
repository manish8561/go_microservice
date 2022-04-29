package common

import (
	"context"
	"log"
	"os"
	"time"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var global_chain_id int
var global_eth_client *ethclient.Client

//get eth client connection
func Get_Eth_Connection(chain_id int) *ethclient.Client {
	rpc, ok := os.LookupEnv("RPC_ETH_URL")
	if !ok {
		log.Fatalf("end point not found to connect %v", rpc)
	}
	switch chain_id {
	case 1:
		rpc, ok = os.LookupEnv("RPC_ETH_URL")
	case 4:
		rpc, ok = os.LookupEnv("RPC_RINKEBY_URL")

	case 56:
		rpc, ok = os.LookupEnv("RPC_BNB_URL")
	}
	if !ok {
		log.Fatalf("end point not found to connect %v", rpc)
	}

	if (global_chain_id !=  chain_id) {
		//create eth client object
		conn, err := ethclient.Dial(rpc)
		global_eth_client = conn
		if err != nil {
			log.Fatalf("Failed to connect to the Ethereum client: %v", err)
		}
	}

	return global_eth_client
}

// Recurrsive function to get Transaction details
func GetTransaction(transaction_hash string, chain_id int, counter int) (int,int64) {
	//get rpc from common file
	conn := Get_Eth_Connection(chain_id)
	

	//convert transaction string to hash
	hash := ethcommon.HexToHash(transaction_hash)

	//get transaction data
	tx, err := conn.TransactionReceipt(context.Background(), hash)
	if err != nil {
		counter = counter + 1
		//after 10 minutes
		if counter >= 100 {
			return -1,0
		}
		time.Sleep(6 * time.Second)
		log.Printf("no transaction found: %v", err)
		return GetTransaction(transaction_hash, chain_id, counter)
	}
	// fmt.Println("tx status:", tx, tx.BlockNumber, tx.Logs)

	return (int)(tx.Status), (tx.BlockNumber).Int64()
}
