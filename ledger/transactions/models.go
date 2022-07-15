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
	"github.com/ethereum/go-ethereum/ethclient"
	// token "./strategy.go"
)

const CollectionName = "farms_transactions"
const CollectionName2 = "farms_blocks"
const blockDff = 200

// Models should only be concerned with database schema, more strict checking should be put in validator.
// Transaction Model
// event Deposit(address indexed account,  uint256 amount);
// event Withdraw(address indexed account,  uint256 amount);
// Both events
type FarmBlockModel struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Created         time.Time          `bson:"_created" json:"_created"`
	Modified        time.Time          `bson:"_modified" json:"_modified"`
	ChainId         int                `bson:"chainId" json:"chainId"`
	BlockNumber     int64              `bson:"blockNumber" json:"blockNumber"`
	LastBlockNumber int64              `bson:"lastBlockNumber" json:"lastBlockNumber"`
}

type TransactionModel struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ChainId         int                `bson:"chainId" json:"chainId"`
	Strategy        string             `bson:"strategy" json:"strategy"`
	TransactionHash string             `bson:"transactionHash" json:"transactionHash"`
	Type            string             `bson:"type" json:"type"`
	Account         string             `bson:"account" json:"account"`
	Amount          float64            `bson:"amount" json:"amount"`
	BlockNumber     int64              `bson:"blockNumber" json:"blockNumber"`
	Timestamp       int64              `bson:"timestamp" json:"timestamp"`
}

//struct for graph data
type GraphDataModel struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Value     float64            `bson:"value" json:"value"`
	Timestamp int64              `bson:"timestamp" json:"timestamp"`
}

//struct for filters
type Filters struct {
	ChainId int64  `bson: "chainId" json:"chainId"`
	Address string `bson: "address" json:"address"`
	Type    string `bson: "type" json:"type"`
}

// init func in go file
func init() {
	// create index
	common.AddIndex(os.Getenv("MONGO_DATABASE"), CollectionName, bson.D{{"strategy", 1}, {"blockNumber", -1}, {"chainId", 1}, {"account", 1}, {"type", 1}})
	common.AddIndex(os.Getenv("MONGO_DATABASE"), CollectionName2, bson.D{{"blockNumber", -1}, {"chainId", 1}})

	//start the cron
	StartCall()
}

// cron func call
func StartCall() {
	c := cron.New()
	c.AddFunc("*/30 * * * * *", func() {
		fmt.Println("[Job 1]Every 30 minutes job\n")
		//calling get transactions according to farms(strategies)
		GetDetails()
	})
	// Start cron with one scheduled job
	c.Start()
}

//get farms
// func GetFarms(chainId int)
func checkContract(address string) bool {
	strategies := [...]string{
		"0x3349e79dfcc1d80114c37d48a516940f06a2b7d2",
		"0xfaa931e617889a10a2f5d9537a9ff9f4d8cedfb8",
		"0x94764fbaef3804474c583640447e2c2a824d31a6",
		"0x07809cb1c6b275b144fc0bd2b9693f6faa47ea61",
		"0x7e01691b46ecd36b4a0f4f5d1f32dc178c9aa279",
		"0x375746d4c701032a282b4bed951e39b9312f9c6c",
		"0x2d32d65fcd4a2b64e4ffa512ac3d0896b542b0d5",
		"0x3421dfd649b31f5bb48528368a68351014b5029e",
	}
	for _, e := range strategies {
		if strings.ToLower(e) == strings.ToLower(address) {
			return true
		}
	}
	return false
}

// You could input string which will be saved in database returning with error info
// 	if err := FindOne(&record); err != nil { ... }
func GetBlockTransactions(chainId int, bN int64) (int64, error) {
	// Create an IPC based RPC connection to a remote node
	conn := common.Get_Eth_Connection(chainId)

	// to get latest blocknumber
	header, err := conn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	lastestBlockNumber := header.Number.Int64()
	// if blockNumber is greater than current blocknumber
	if bN > lastestBlockNumber {
		bN = lastestBlockNumber
	}

	blockNumber := big.NewInt(bN)
	block, err := conn.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	blockTimestamp := int64(block.Time())
	// fmt.Println(block.Number().Uint64()) // 5671744
	// fmt.Println(block.Time())            // 1527211625

	fmt.Println("------------------------------------------------")
	fmt.Println("Total Transactions: ", len(block.Transactions()))
	dd := math.Pow(10, float64(18))

	for _, tx := range block.Transactions() {
		// fmt.Println("Transaction: ", tx.Hash().Hex())
		if tx.To() != nil && checkContract(tx.To().Hex()) {
			strategyAddress := strings.ToLower(tx.To().Hex())
			contractAddress := ethcommon.HexToAddress(strategyAddress)
			strategyContract, err := NewTransactions(contractAddress, conn)
			if err != nil {
				fmt.Printf("Failed to instantiate a Token contract: %v", err)
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
							log.Println(err)
						}

						//converting the string to float64
						transferValue, err := strconv.ParseFloat(withdrawEvent.Amount.String(), 64)
						if err != nil {
							fmt.Println("Float conversion", err)
						}

						d := TransactionModel{
							ChainId:         chainId,
							Strategy:        strategyAddress,
							TransactionHash: vLog.TxHash.Hex(),
							Type:            "deposit",
							Account:         strings.ToLower(withdrawEvent.Account.Hex()),
							Amount:          (transferValue / dd),
							BlockNumber:     bN,
							Timestamp:       blockTimestamp,
						}
						go SaveDataBackground(&d, CollectionName)
						
					//Withdraw Event
					case logWithdrawSigHash.Hex():
						withdrawEvent, err := strategyContract.ParseWithdraw(*vLog)
						if err != nil {
							log.Println(err)
						}

						//converting the string to float64
						transferValue, err := strconv.ParseFloat(withdrawEvent.Amount.String(), 64)
						if err != nil {
							fmt.Println("Float conversion", err)
						}
						fmt.Println("----------------------")
						fmt.Println("Account: ", strings.ToLower(withdrawEvent.Account.Hex()))
						fmt.Println("transfer value:", transferValue)
						fmt.Println("----------------------")
						// blockTimestamp := Get_Block_Timestamp(conn, int64(vLog.BlockNumber))
						fmt.Println("block times: ", blockTimestamp)
						d := TransactionModel{
							ChainId:         chainId,
							Strategy:        strategyAddress,
							TransactionHash: vLog.TxHash.Hex(),
							Type:            "withdraw",
							Account:         strings.ToLower(withdrawEvent.Account.Hex()),
							Amount:          (transferValue / dd),
							BlockNumber:     bN,
							Timestamp:       blockTimestamp,
						}
						go SaveDataBackground(&d, CollectionName)
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

// You could input an TransactionModel which will be updated in database returning with error info
// 	if err := UpdateOne(&record); err != nil { ... }
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
// 	if err := FindOne(&record); err != nil { ... }
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

// Price Feed list api with page and limit
func GetTotal(status string) int64 {
	client := common.GetDB()

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := bson.M{}
	num, err := collection.CountDocuments(ctx, query)
	if err != nil {
		return 0
	}
	return num
}

// Price Feed list api with page and limit
func GetAll(page int64, limit int64, status string) ([]*TransactionModel, error) {
	client := common.GetDB()
	var records []*TransactionModel

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	opts := options.Find().SetSort(bson.D{{"_created", -1}}).SetSkip((page - 1) * limit).SetLimit(limit)
	//SetProjection(bson.M{"_id": 0, "_created": 1, "_modified": 1, "firstname": 1, "lastname": 1, "status": 1, "email": 1, "role": 1, "passwordhash": 0})

	query := bson.M{}

	cursor, err := collection.Find(ctx, query, opts)
	if err != nil {
		return records, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &records)

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return records, err
}

// delete record from collection
func DeleteRecordModel(ID string) (bool, error) {
	client := common.GetDB()

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)

	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return false, err
	}
	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return false, err
	}
	fmt.Println(res, "Delete")
	return true, err
}

//aggregate function for graph in dashboard

/*
db.getCollection('transfers').aggregate([{$match:{chainId:4, address:"0x37dc6cf6a221b6e511eb9fcdef6cb467c636847b"}},{$limit:7},{$sort:{timestamp:-1}},{$project:{value:1, timestamp:1}}])
*/
func GetLastSevenTransaction(filters Filters) ([]*GraphDataModel, error) {
	client := common.GetDB()
	var records []*GraphDataModel

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Specify a pipeline that will return the number of times each name appears
	// in the collection.
	pipeline := []bson.M{
		{"$match": bson.M{"chainId": filters.ChainId, "address": strings.ToLower(filters.Address)}},
		{"$sort": bson.M{"timestamp": -1}},
		{"$project": bson.M{
			// "_id":   0,
			"value":     1,
			"timestamp": 1,
		}},
		{"$limit": 7},
		{"$sort": bson.M{"timestamp": 1}},
	}
	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	opts := options.Aggregate()

	cursor, err := collection.Aggregate(ctx, pipeline, opts)
	if err != nil {
		return records, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &records)

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return records, err
}

/**
* delete the duplicates
* @param  {string} d
 */
func callingDelete(CollectionName string, ChainId int) error {
	type IDResult struct {
		TransactionHash string
	}

	type DeleteResult struct {
		ID    IDResult `json:"_id" bson:"_id"`
		Dups  []primitive.ObjectID
		Count float64
	}
	var records []*DeleteResult

	client := common.GetDB()
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := bson.M{"chainId": ChainId}

	// Specify a pipeline that will return the number of times each name appears
	// in the collection.
	pipeline := []bson.M{
		{"$match": query},
		{"$group": bson.M{
			"_id":   bson.M{"transactionHash": "$transactionHash"},
			"dups":  bson.M{"$addToSet": "$_id"},
			"count": bson.M{"$sum": 1},
		}},
		{"$match": bson.M{"count": bson.M{"$gt": 1}}},
	}
	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	opts := options.Aggregate()

	cursor, err := collection.Aggregate(ctx, pipeline, opts)
	if err != nil {
		log.Fatalf("err in aggregate", err)
		return err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &records)
	if err != nil {
		return err
	}

	// fmt.Println("records", len(records))
	// fmt.Println("records", records[0].ID.TransactionHash)
	// delete the duplicate ids at once
	if len(records) > 0 {
		for _, element := range records {
			//slice array
			slicedArr := element.Dups[1:]

			res, err := collection.DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$in": slicedArr}})
			if err != nil {
				return err
			}
			fmt.Println("delete response", res)
		}

	}
	return nil
}

// function to get block timestamp
func Get_Block_Timestamp(client *ethclient.Client, block_num int64) int64 {
	blockNumber := big.NewInt(block_num)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%t", block.Time())
	// fmt.Println(block.Time(), block_num, "block timestamp")
	return int64(block.Time())
}

// Get date without time from timestamp
func get_date_without_time(timestamp int64) time.Time {
	currentDate := time.Unix(timestamp, 0).UTC()

	//get year month day
	y, m, d := currentDate.Date()
	//convert to date
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

//to get all autocompound address from constant file in a map
func GetDetails() {
	for chainId, val := range common.NetworkMap {
		//calling the contract as per chainId
		// GetContract(chainId, val.AC, val.BlockNumber)
		// fmt.Println(chainId, val, val.AC.Address, val.AC.BlockNumber, "----------")

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
		}
		go UpdateOne(r.ID, bn, CollectionName2)
		// GetBlockTransactions(chainId, 30591094)

	}
}
