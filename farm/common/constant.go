package common

// import "fmt"

// struct for muliple address in a network
type MultipleAddress struct {
	AC          string // capital for export
	BlockNumber int64
}

//network map (object) with chain id
var NetworkMap = map[int]MultipleAddress{
	// 1:     {AC: "0x37DC6cF6A221b6E511EB9fcdeF6cb467c636847b", BlockNumber:10850501},//ethereum
	4: {AC: "0x37DC6cF6A221b6E511EB9fcdeF6cb467c636847b", BlockNumber: 10850501}, //rinkeby
	// 56:    {AC: "0x37DC6cF6A221b6E511EB9fcdeF6cb467c636847b", BlockNumber:10850501},//bsc mainnet
	// 97:    {AC: "0x37DC6cF6A221b6E511EB9fcdeF6cb467c636847b", BlockNumber:10850501},//bsd testnet
	// 137:   {AC: "0x37DC6cF6A221b6E511EB9fcdeF6cb467c636847b", BlockNumber:10850501},//polygon mainnet
	// 80001: {AC: "0x37DC6cF6A221b6E511EB9fcdeF6cb467c636847b", BlockNumber:10850501},//polygon testnet
}

//    networkMap[1] = MultipleAddress{AC:"0x37DC6cF6A221b6E511EB9fcdeF6cb467c636847b"}

// func init() {
// 	fmt.Println("\n\n\n")
// 	fmt.Println("===========================================")
// 	fmt.Println("\n\n\n")
// 	fmt.Println(NetworkMap[4].AC, "before the exec")
// 	fmt.Println("\n\n\n")

// }
