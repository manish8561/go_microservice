package common

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

var globalChainId int
var globalEthClient *ethclient.Client

//get eth client connection
func GetEthConnection(chainId int) *ethclient.Client {
	rpc, ok := os.LookupEnv("RPC_ETH_URL")
	if !ok {
		log.Fatalf("end point not found to connect %v", rpc)
	}
	switch chainId {
	case 1:
		rpc, ok = os.LookupEnv("RPC_ETH_URL")
	case 4:
		rpc, ok = os.LookupEnv("RPC_RINKEBY_URL")
	case 56:
		rpc, ok = os.LookupEnv("RPC_BNB_URL")
	case 97:
		rpc, ok = os.LookupEnv("RPC_BSC_TESTNET_URL")
	case 137:
		rpc, ok = os.LookupEnv("RPC_POLYGON_URL")
	case 80001:
		rpc, ok = os.LookupEnv("RPC_POLYGON_TESTNET_URL")
	default:
		rpc, ok = os.LookupEnv("RPC_RINKEBY_URL")
	}

	if !ok {
		log.Printf("end point not found to connect %v", rpc)
	}

	if globalChainId != chainId {
		//create eth client object
		conn, err := ethclient.Dial(rpc)
		globalEthClient = conn
		if err != nil {
			log.Printf("Failed to connect to the Ethereum client: %v", err)
		}
	}

	return globalEthClient
}
