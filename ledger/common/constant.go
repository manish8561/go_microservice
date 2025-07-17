package common

// struct for muliple address in a network
type ContractDetails struct {
	Address     string // capital for export
	BlockNumber int64
}

type MultipleAddress struct {
	AC       ContractDetails
	Strategy ContractDetails
}

// var aa = ContractDetails{Address: "0x12e9a9dcDc8f276c71524Ddd102343525ddAbB26", BlockNumber: 30244347}

//network map (object) with chain id
var NetworkMap = map[int]MultipleAddress{
	// 1:     {
	// 	AC: ContractDetails{Address: "0x12e9a9dcDc8f276c71524Ddd102343525ddAbB26", BlockNumber: 30244347},
	// 	Strategy: ContractDetails{Address: "", BlockNumber: 30244347},
	// },//ethereum
	// 4: {AC: "0x37DC6cF6A221b6E511EB9fcdeF6cb467c636847b", BlockNumber: 10850501}, //rinkeby
	56:    {
		AC: ContractDetails{Address: "", BlockNumber: 19540751},
		Strategy: ContractDetails{Address: "", BlockNumber: 19540751},
	},//bsc mainnet
	// 97:    {
	// 	AC: ContractDetails{Address: "0x12e9a9dcDc8f276c71524Ddd102343525ddAbB26", BlockNumber: 30244347},
	// 	Strategy: ContractDetails{Address: "", BlockNumber: 30244347},
	// },//bsd testnet
	137: {
		AC:       ContractDetails{Address: "0x12e9a9dcDc8f276c71524Ddd102343525ddAbB26", BlockNumber: 30244347},
		Strategy: ContractDetails{Address: "", BlockNumber: 30244347},
	}, //polygon mainnet
	// 80001:
	// 	AC: ContractDetails{Address: "0x12e9a9dcDc8f276c71524Ddd102343525ddAbB26", BlockNumber: 30244347},
	// 	Strategy: ContractDetails{Address: "", BlockNumber: 30244347},
	// },//polygon testnet
}
//default network
var DefaultChainId int64 = 137 //polygon