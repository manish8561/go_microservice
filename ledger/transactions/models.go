package transactions

import (
	"context"
	"fmt"
	"log"
	"math"
	"strconv"

	// "math"
	"math/big"
	"os"

	// "strconv"
	"strings"
	"time"

	"github.com/autocompound/docker_backend/ledger/common"
	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	// "github.com/ethereum/go-ethereum"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/crypto"

	ethcommon "github.com/ethereum/go-ethereum/common"

	// token "./strategy.go"
	pb "github.com/autocompound/docker_backend/ledger/helloworld"
)

const CollectionName = "farms_transactions"
const CollectionName2 = "farms_blocks"

// Network Block Number Model
type FarmBlockModel struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Created         time.Time          `bson:"_created" json:"_created"`
	Modified        time.Time          `bson:"_modified" json:"_modified"`
	ChainId         int                `bson:"chainId" json:"chainId"`
	BlockNumber     int64              `bson:"blockNumber" json:"blockNumber"`
	LastBlockNumber int64              `bson:"lastBlockNumber" json:"lastBlockNumber"`
}

// Models should only be concerned with database schema, more strict checking should be put in validator.
// Transaction Model
// event Deposit(address indexed account,  uint256 amount);
// event Withdraw(address indexed account,  uint256 amount);
// Both events
type TransactionModel struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ChainId         int                `bson:"chainId" json:"chainId"`
	Strategy        string             `bson:"strategy" json:"strategy"`
	TransactionHash string             `bson:"transactionHash" json:"transactionHash"`
	Type            string             `bson:"type" json:"type"`
	Account         string             `bson:"account" json:"account"`
	Amount          float64            `bson:"amount" json:"amount"`
	AmountUSD       float64            `bson:"amountUSD" json:"amountUSD"`
	BlockNumber     int64              `bson:"blockNumber" json:"blockNumber"`
	Timestamp       int64              `bson:"timestamp" json:"timestamp"`
}

// struct for api data with total
type EventResult struct {
	Total   int                `bson:"total" json:"total"`
	Records []TransactionModel `bson:"records" json:"records"`
}

// struct for filters
type Filters struct {
	ChainId   int64  `bson:"chainId" json:"chainId"`
	Address   string `bson:"address" json:"address"`
	Type      string `bson:"type" json:"type"`
	StartTime int64  `bson:"startTime" json:"startTime"`
	EndTime   int64  `bson:"endTime" json:"endTime"`
	Page      int64  `bson:"page" json:"page"`
	Limit     int64  `bson:"limit" json:"limit"`
}

// init func in go file
func init() {
	// create index
	common.AddIndex(os.Getenv("MONGO_DATABASE"), CollectionName, bson.D{{Key: "strategy", Value: 1}, {Key: "blockNumber", Value: -1}, {Key: "chainId", Value: 1}, {"account", 1}, {"type", 1}, {"timestamp", -1}, {"amountUSD", 1}})
	common.AddIndex(os.Getenv("MONGO_DATABASE"), CollectionName2, bson.D{{"blockNumber", -1}, {"chainId", 1}})

	//start the cron
	StartCronJob()
}

// get active farms
func GetFarmFromService(chainId int) *pb.FarmReply {
	c := int64(chainId)
	grpc_server_conn := common.Get_GRPC_Conn()
	cc := pb.NewGreeterClient(grpc_server_conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := cc.GetFarms(ctx, &pb.FarmRequest{ChainId: c, Status: "active"})
	if err != nil {
		fmt.Println("grpc ledger error", err)
		return nil
	}

	return result
}

// cron func call
func StartCronJob() {
	c := cron.New()
	c.AddFunc("*/3 * * * * *", func() {
		fmt.Println("[Job 1]Every 30 minutes job\n")
		//calling get transactions according to farms(strategies)
		GetDetails()
	})
	// Start cron with one scheduled job
	c.Start()
}

// get farms
// func GetFarms(chainId int)
func checkContract(address string, farmReply *pb.FarmReply) (bool, float64) {
	// strategies := [...]string{
	// 	"0x3349e79dfcc1d80114c37d48a516940f06a2b7d2",
	// 	"0xfaa931e617889a10a2f5d9537a9ff9f4d8cedfb8",
	// 	"0x94764fbaef3804474c583640447e2c2a824d31a6",
	// 	"0x07809cb1c6b275b144fc0bd2b9693f6faa47ea61",
	// 	"0x7e01691b46ecd36b4a0f4f5d1f32dc178c9aa279",
	// 	"0x375746d4c701032a282b4bed951e39b9312f9c6c",
	// 	"0x2d32d65fcd4a2b64e4ffa512ac3d0896b542b0d5",
	// 	"0x3421dfd649b31f5bb48528368a68351014b5029e",
	// 	"0x12e9a9dcDc8f276c71524Ddd102343525ddAbB26",
	// }
	var val float64 = 0.0

	for _, e := range farmReply.Items {
		if strings.ToLower(e.Address) == strings.ToLower(address) {
			// check = e
			val = e.TokenPrice
			return true, val
		}
	}
	return false, val
}

// You could input string which will be saved in database returning with error info
func GetBlockTransactions(chainId int, bN int64) (int64, error) {
	//GRPC call
	r := GetFarmFromService(chainId)
	// error from grpc
	if r == nil {
		return bN, nil
	}
	fmt.Println("------------------------------------------------")
	fmt.Println("gprc result: ", len(r.Items), "chainId: ", chainId)

	// Create an IPC based RPC connection to a remote node
	conn := common.Get_Eth_Connection(chainId)

	// to get latest blocknumber
	header, err := conn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Println(err)
		return bN, err
	}
	lastestBlockNumber := header.Number.Int64()
	// if blockNumber is greater than current blocknumber
	if bN > lastestBlockNumber {
		bN = lastestBlockNumber
	}

	blockNumber := big.NewInt(bN)
	block, err := conn.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Println(err)
		return bN, err
	}
	blockTimestamp := int64(block.Time())
	// fmt.Println(block.Number().Uint64()) // 5671744
	// fmt.Println(block.Time())            // 1527211625

	fmt.Println("------------------------------------------------")
	fmt.Println("chainId: ", chainId, "blockNumber: ", bN)
	fmt.Println("Total Transactions: ", len(block.Transactions()))
	dd := math.Pow(10, float64(18))

	for _, tx := range block.Transactions() {
		// fmt.Println("Transaction: ", tx.Hash().Hex())
		if tx.To() != nil {
			check, tokenPrice := checkContract(tx.To().Hex(), r)
			if check {
				strategyAddress := strings.ToLower(tx.To().Hex())
				contractAddress := ethcommon.HexToAddress(strategyAddress)
				strategyContract, err := NewTransactions(contractAddress, conn)
				if err != nil {
					fmt.Printf("Failed to instantiate a Token contract: %v", err)
					return bN, err
				}
				//Withdraw
				logWithdrawSig := []byte("Withdraw(address,uint256)")
				logWithdrawSigHash := crypto.Keccak256Hash(logWithdrawSig)

				//Deposit
				logDepositSig := []byte("Deposit(address,uint256)")
				logDepositSigHash := crypto.Keccak256Hash(logDepositSig)

				receipt, err := conn.TransactionReceipt(context.Background(), tx.Hash())
				if err != nil {
					fmt.Println(err)
					return bN, err
				}

				// 1 success status
				if receipt.Status == 1 {
					// fmt.Println("Logs: ", len(receipt.Logs)) // ...

					for _, vLog := range receipt.Logs {
						// fmt.Println(vLog.BlockHash.Hex())
						// fmt.Println(vLog.BlockNumber)  // 2394201
						switch vLog.Topics[0].Hex() {
						// Deposit event hex
						case logDepositSigHash.Hex():
							withdrawEvent, err := strategyContract.ParseDeposit(*vLog)
							if err != nil {
								log.Println("Parse deposit error", err)
								return bN, err
							}

							//converting the string to float64
							transferValue, err := strconv.ParseFloat(withdrawEvent.Amount.String(), 64)
							if err != nil {
								fmt.Println("Float conversion", err)
								return bN, err
							}
							// convert to dollar price
							USDValue := (transferValue / dd) * (tokenPrice)

							d := TransactionModel{
								ChainId:         chainId,
								Strategy:        strategyAddress,
								TransactionHash: vLog.TxHash.Hex(),
								Type:            "deposit",
								Account:         strings.ToLower(withdrawEvent.Account.Hex()),
								Amount:          (transferValue / dd),
								AmountUSD:       USDValue,
								BlockNumber:     bN,
								Timestamp:       blockTimestamp,
							}
							go SaveDataBackground(&d, CollectionName)

						//Withdraw Event
						case logWithdrawSigHash.Hex():
							withdrawEvent, err := strategyContract.ParseWithdraw(*vLog)
							if err != nil {
								log.Println("Parse withdraw event", err)
								return bN, err
							}

							//converting the string to float64
							transferValue, err := strconv.ParseFloat(withdrawEvent.Amount.String(), 64)
							if err != nil {
								fmt.Println("Float conversion", err)
								return bN, err
							}
							USDValue := (transferValue / dd) * (tokenPrice)
							// USDValue := (transferValue / dd) * tokenPrice

							d := TransactionModel{
								ChainId:         chainId,
								Strategy:        strategyAddress,
								TransactionHash: vLog.TxHash.Hex(),
								Type:            "withdraw",
								Account:         strings.ToLower(withdrawEvent.Account.Hex()),
								Amount:          (transferValue / dd),
								AmountUSD:       USDValue,

								BlockNumber: bN,
								Timestamp:   blockTimestamp,
							}
							go SaveDataBackground(&d, CollectionName)
						}
					}
				}
			}
		}
	}
	return (bN + 1), err
}

// function for background process
func SaveDataBackground(data interface{}, CollectionName string) {
	err := common.SaveOne(data, CollectionName)
	if err != nil {
		log.Printf("Failed to retrieve token name: %v", err)
	}
}

// to get all autocompound address from constant file in a map
func GetDetails() {
	for chainId, val := range common.NetworkMap {
		//calling the contract as per chainId

		// get record from db
		result, err := GetRecord(chainId, CollectionName2)
		if err != nil {
			// if not found add block number for specific chainId
			d := FarmBlockModel{
				ChainId:         chainId,
				Created:         time.Now(),
				Modified:        time.Now(),
				BlockNumber:     val.Strategy.BlockNumber,
				LastBlockNumber: val.Strategy.BlockNumber,
			}
			err := common.SaveOne(&d, CollectionName2)
			if err != nil {
				log.Fatalf("Failed to retrieve token name: %v", err)
			}
			return
		}
		r := result.(FarmBlockModel)

		//get block transactions
		bn, err := GetBlockTransactions(r.ChainId, r.LastBlockNumber)
		if err != nil {
			fmt.Println(err, "get transaction err")
			return
		}
		fmt.Println("before update blockNumber:", bn)
		fmt.Println("------------------------------------")
		//if block number is > 0
		if bn > 0 {
			go UpdateOne(r.ID, bn, CollectionName2)
		}
	}
}

// You could input an TransactionModel which will be updated in database returning with error info
//
//	if err := UpdateOne(&record); err != nil { ... }
func UpdateOne(ID primitive.ObjectID, lastBlockNumber int64, CollectionName string) error {
	client := common.GetDB()
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// to check for unique email address
	opts := options.Update().SetUpsert(false)
	update := bson.M{"lastBlockNumber": lastBlockNumber}
	update = bson.M{"$set": update}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": ID}, update, opts)
	if err != nil {
		return err
	}
	fmt.Println(res, "Updated")
	return nil
}

// You could input string which will be saved in database returning with error info
//
//	if err := FindOne(&record); err != nil { ... }
func GetRecord(chainId int, CN string) (interface{}, error) {
	client := common.GetDB()

	record := &TransactionModel{}
	record2 := &FarmBlockModel{}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CN)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	opts := options.FindOne().SetSort(bson.D{{"blockNumber", -1}})
	if CN == CollectionName {
		err := collection.FindOne(ctx, bson.M{"chainId": chainId}, opts).Decode(&record)
		return *record, err
	} else {
		err := collection.FindOne(ctx, bson.M{"chainId": chainId}, opts).Decode(&record2)
		return *record2, err
	}
}

// get all transaction list api with page and limit
func GetTransactions(filters Filters) (*EventResult, error) {
	var records []*EventResult

	client := common.GetDB()
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := bson.M{"chainId": filters.ChainId, "type": strings.ToLower(filters.Type), "account": strings.ToLower(filters.Address), "timestamp": bson.M{"$gte": filters.StartTime, "$lte": filters.EndTime}}

	// Specify a pipeline that will return the number of times each name appears
	// in the collection.
	pipeline := []bson.M{
		{"$facet": bson.M{
			"total": []bson.M{
				{"$match": query},
				{"$count": "total"},
			},
			"records": []bson.M{
				{"$match": query},
				{"$skip": (filters.Page - 1) * filters.Limit},
				{"$limit": filters.Limit},
				{"$sort": bson.M{"blockNumber": -1}},
			},
		}},
		{"$project": bson.M{
			"total": bson.M{"$cond": bson.M{
				"if": bson.M{"$gt": bson.A{bson.M{"$size": "$total"}, 0}}, "then": bson.M{"$first": "$total.total"}, "else": 0}}, "records": 1,
		}},
	}
	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	opts := options.Aggregate()

	cursor, err := collection.Aggregate(ctx, pipeline, opts)
	if err != nil {
		return records[0], err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &records)
	if err != nil {
		return records[0], err
	}
	return records[0], nil
}

// Get date without time from timestamp
func get_date_without_time(timestamp int64) time.Time {
	currentDate := time.Unix(timestamp, 0).UTC()

	//get year month day
	y, m, d := currentDate.Date()
	//convert to date
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

//to get the profit and loss
/*
db.getCollection('farms_transactions').aggregate([
{$facet:{
    deposit:[{$match:{type:"deposit"}},{$group:{_id:null, total:{$sum:"$amountUSD"}}}],
    withdraw:[{$match:{type:"withdraw"}},{$group:{_id:null, total:{$sum:"$amountUSD"}}}]}
}, {$project:{
        desposit:{$first:"$deposit.total"},
        withdraw:{$first:"$withdraw.total"},
    } },
])
*/
func GetProfitLoss(chainId int64) (float64, float64) {
	client := common.GetDB()
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	type Result struct {
		Deposit  float64 `bson:"deposit" json:"deposit"`
		Withdraw float64 `bson:"withdraw" json:"withdraw"`
	}

	var records []*Result
	// Specify a pipeline that will return the number of times each name appears
	// in the collection.
	pipeline := []bson.M{
		{"$facet": bson.M{
			"deposit": []bson.M{
				{"$match": bson.M{"type": "deposit", "chainId": chainId}},
				{"$group": bson.M{
					"_id":   nil,
					"total": bson.M{"$sum": "$amountUSD"},
				}},
			},
			"withdraw": []bson.M{
				{"$match": bson.M{"type": "withdraw", "chainId": chainId}},
				{"$group": bson.M{
					"_id":   nil,
					"total": bson.M{"$sum": "$amountUSD"},
				}},
			},
		}},
		{"$project": bson.M{
			"deposit":  bson.M{"$first": "$deposit.total"},
			"withdraw": bson.M{"$first": "$withdraw.total"},
		}},
	}
	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	opts := options.Aggregate()

	cursor, err := collection.Aggregate(ctx, pipeline, opts)
	if err != nil {
		return 0, 0
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &records)
	if err != nil {
		return 0, 0
	}

	if len(records) == 1 {
		v := records[0].Deposit - records[0].Withdraw
		if records[0].Withdraw > 0 {
			t := (v / records[0].Withdraw) * 100
			return t, v
		}
		if records[0].Deposit > 0 {
			return 100, v
		}
		return 0, 0

	}

	return 0, 0
}
