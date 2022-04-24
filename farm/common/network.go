package common

import (
	"log"
	"os"
)

//get rpc from chain id
func Get_RPC_ChainId(chain_id int) string {
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
	return rpc
}
