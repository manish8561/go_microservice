package common

const (
	rpcEthURL      = "RPC_ETH_URL"
	rpcRinkebyURL  = "RPC_RINKEBY_URL"
	rpcBscURL      = "RPC_BNB_URL"
	rpcBscTestnet  = "RPC_BSC_TESTNET_URL"
	rpcPolygonURL  = "RPC_POLYGON_URL"
	rpcPolygonTest = "RPC_POLYGON_TESTNET_URL"
	ErrInAggregate = "err in aggregate: "
)

// struct for muliple address in a network
type MultipleAddress struct {
	AC          string // capital for export
	BlockNumber int64
}

// network map (object) with chain id
var NetworkMap = map[int]MultipleAddress{
	// 1:     {AC: "0x37DC6cF6A221b6E511EB9fcdeF6cb467c636847b", BlockNumber:10850501},//ethereum
	// 4: {AC: "0x37DC6cF6A221b6E511EB9fcdeF6cb467c636847b", BlockNumber: 10850501}, //rinkeby
	// 56:    {AC: "0x37DC6cF6A221b6E511EB9fcdeF6cb467c636847b", BlockNumber:10850501},//bsc mainnet
	// 97:    {AC: "0x37DC6cF6A221b6E511EB9fcdeF6cb467c636847b", BlockNumber:10850501},//bsd testnet
	137: {AC: "0x12e9a9dcDc8f276c71524Ddd102343525ddAbB26", BlockNumber: 30244347}, //polygon mainnet
	// 80001: {AC: "0x6cDb2D638Ed5BCe1791aaaB0e096f047765caDa7", BlockNumber:26942775},//polygon testnet
}

var DefaultChainId int64 = 137 //polygon
