package common

import (
	"log"
	"os"

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
