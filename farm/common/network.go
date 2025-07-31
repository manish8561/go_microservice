package common

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

var globalChainId int
var globalEthClient *ethclient.Client

// SetGlobalChainId sets the global chain ID
func GetEthConnection(chainId int) *ethclient.Client {
	rpc, ok := os.LookupEnv("RPC_ETH_URL")
	if !ok {
		log.Fatalf("end point not found to connect %v", rpc)
	}
	switch chainId {
	case 1:
		rpc, ok = os.LookupEnv(rpcEthURL)
	case 4:
		rpc, ok = os.LookupEnv(rpcRinkebyURL)
	case 56:
		rpc, ok = os.LookupEnv(rpcBscURL)
	case 97:
		rpc, ok = os.LookupEnv(rpcBscTestnet)
	case 137:
		rpc, ok = os.LookupEnv(rpcPolygonURL)
	case 80001:
		rpc, ok = os.LookupEnv(rpcPolygonTest)
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
