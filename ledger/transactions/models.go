package transactions

import (
	"context"
	"fmt"
	"log"
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
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/ethclient"
	// token "./autocompound.go"
)

const CollectionName = "transfers"
const blockDff = 200

// Models should only be concerned with database schema, more strict checking should be put in validator.
// event Transfer(address indexed from, address indexed to, uint256 value);
//
type TransferEventModel struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ChainId         int                `bson:"chainId" json:"chainId"`
	Address         string             `bson:"address" json:"address"`
	TransactionHash string             `bson:"transactionHash" json:"transactionHash"`
	From            string             `bson:"from" json:"from"`
	To              string             `bson:"to" json:"to"`
	Value           float64            `bson:"value" json:"value"`
	BlockNumber     int64              `bson:"blockNumber" json:"blockNumber"`
	LastBlockNumber int64              `bson:"lastBlockNumber" json:"lastBlockNumber"`
	Timestamp       int64              `bson:"timestamp" json:"timestamp"`
	CreatedAt       time.Time          `bson:"createdAt" json:"createdAt"`
}

//struct for graph data
type GraphDataModel struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Value     float64            `bson:"value" json:"value"`
	Timestamp int64              `bson:"timestamp" json:"timestamp"`
}

//struct for graph data
type GraphDataModel2 struct {
	ID    time.Time `json:"_id,omitempty" bson:"_id,omitempty"`
	Value float64   `bson:"value" json:"value"`
	Count int64     `bson:"count" json:"count"`
}

//struct for filters
type Filters struct {
	ChainId int64  `bson: "chainId" json:"chainId"`
	Address string `bson: "address" json:"address"`
}

// init func in go file
func init() {
	// create index
	common.AddIndex(os.Getenv("MONGO_DATABASE"), CollectionName, bson.D{{"address", 1}, {"blockNumber", -1}, {"from", 1}, {"to", 1}, {"createdAt", -1}})

	//start the cron
	StartCall()
}

// cron func call
func StartCall() {
	c := cron.New()
	c.AddFunc("0 */1 * * * *", func() {
		fmt.Println("[Job 1]Every 30 minutes job\n")
		//calling get autocompounds
		GetDetails()
	})
	// Start cron with one scheduled job
	c.Start()
}

// You could input string which will be saved in database returning with error info
// 	if err := FindOne(&farmModel); err != nil { ... }
func GetBlockTransactions(chainId int, bN int64) error {
	// Create an IPC based RPC connection to a remote node
	conn := common.Get_Eth_Connection(chainId)

	// to get latest blocknumber
	header, err := conn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	lastestBlockNumber := header.Number.Int64()
	fmt.Println(lastestBlockNumber, "before")

	blockNumber := big.NewInt(bN)
	block, err := conn.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time())                // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions()))   // 144

	fmt.Println("------------------------------------------------")
	strategyContractAddress := "0x375746d4c701032a282b4bed951e39b9312f9c6c"

	for _, tx := range block.Transactions() {
		fmt.Println("Transactions: ", tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		// fmt.Println(tx.Value().String())               // 10000000000000000
		// fmt.Println(tx.Gas())                          // 105000
		// fmt.Println(tx.GasPrice().Uint64())            // 102000000000
		// fmt.Println(tx.Nonce())                        // 110644
		// fmt.Println(tx.Data())                         // []
		// fmt.Println("to address: ", tx.To().Hex()) // 0x375746d4c701032a282b4bed951e39b9312f9c6c
		if strings.ToLower(tx.To().Hex()) == strings.ToLower(strategyContractAddress) {
			contractAddress := ethcommon.HexToAddress(strategyContractAddress)
			strategyContract, err := NewTransactions(contractAddress, conn)
			if err != nil {
				log.Fatalf("Failed to instantiate a Token contract: %v", err)
			}

			logWithdrawSig := []byte("Withdraw(address,uint256)")
			// // LogApprovalSig := []byte("Approval(address,address,uint256)")
			logWithdrawSigHash := crypto.Keccak256Hash(logWithdrawSig)

			receipt, err := conn.TransactionReceipt(context.Background(), tx.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("recept and logs:")

			fmt.Println(receipt.Status) // 1
			fmt.Println(receipt.Logs)   // ...

			for _, vLog := range receipt.Logs {
				fmt.Println(vLog.BlockHash.Hex())
				fmt.Println(vLog.BlockNumber)  // 2394201
				fmt.Println(vLog.TxHash.Hex()) // 

				switch vLog.Topics[0].Hex() {
				// Withdraw event hex
				case logWithdrawSigHash.Hex():
					withdrawEvent, err := strategyContract.ParseWithdraw(*vLog)
					if err != nil {
						log.Fatal(err)
					}

					//converting the string to float64
					transferValue, err := strconv.ParseFloat(withdrawEvent.Amount.String(), 64)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println("----------------------")
					fmt.Println("Account: ",strings.ToLower(withdrawEvent.Account.Hex()))
					fmt.Println("transfer value:", transferValue)
					fmt.Println("----------------------")
					// blockTimestamp := Get_Block_Timestamp(conn, int64(vLog.BlockNumber))
					// d := TransferEventModel{
					// 	ChainId:         chainId,
					// 	Address:         strings.ToLower(ac),
					// 	TransactionHash: vLog.TxHash.Hex(),
					// 	From:            strings.ToLower(transferEvent.From.Hex()),
					// 	To:              strings.ToLower(transferEvent.To.Hex()),
					// 	Value:           (transferValue / dd),
					// 	BlockNumber:     int64(vLog.BlockNumber),
					// 	LastBlockNumber: int64(vLog.BlockNumber),
					// 	Timestamp:       blockTimestamp,
					// 	CreatedAt:       get_date_without_time(blockTimestamp),
					// }
					// err = SaveOne(&d)
					// if err != nil {
					// 	log.Printf("Failed to retrieve token name: %v", err)
					// }
				}
			}
		}
	}

	return err
}

// You could input an TransferEventModel which will be saved in database returning with error info
// 	if err := SaveOne(&farmModel); err != nil { ... }
func SaveOne(data *TransferEventModel) error {
	client := common.GetDB()

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		fmt.Println(res, "Inserted")
		return err
	}
	return nil
}

// You could input an TransferEventModel which will be updated in database returning with error info
// 	if err := UpdateOne(&farmModel); err != nil { ... }
func UpdateOne(ID primitive.ObjectID, lastBlockNumber int64) error {
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
// 	if err := FindOne(&farmModel); err != nil { ... }
func GetRecord(chainId int, ac string) (TransferEventModel, error) {
	client := common.GetDB()
	record := &TransferEventModel{}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	opts := options.FindOne().SetSort(bson.D{{"blockNumber", -1}})
	err := collection.FindOne(ctx, bson.M{"chainId": chainId, "address": strings.ToLower(ac)}, opts).Decode(&record)

	return *record, err
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
func GetAll(page int64, limit int64, status string) ([]*TransferEventModel, error) {
	client := common.GetDB()
	var records []*TransferEventModel

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

/* get last seven days volume
db.getCollection('transfers').aggregate([
{$sort: {"timestamp":-1}},
{$match:{chainId:4}},
{$group:{
    _id:"$createdAt",
    count:{$sum:1},
    value:{$sum:"$value"}
    }},
{$sort:{_id:1}}
])
*/
func GetLastSevenDaysData(filters Filters) ([]*GraphDataModel2, error) {
	client := common.GetDB()
	var records []*GraphDataModel2

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//last seven day timestamp
	currentDate := get_date_without_time(time.Now().Unix())
	lastWeekDate := currentDate.Unix() - (7 * 24 * 60 * 60)

	// Specify a pipeline that will return the number of times each name appears
	// in the collection.
	pipeline := []bson.M{
		{"$sort": bson.M{"timestamp": -1}},
		{"$match": bson.M{"chainId": filters.ChainId, "address": strings.ToLower(filters.Address), "timestamp": bson.M{"$gt": lastWeekDate, "$lte": currentDate.Unix()}}},
		{"$group": bson.M{
			"_id":   "$createdAt",
			"count": bson.M{"$sum": 1},
			"value": bson.M{"$sum": "$value"},
		}},
		{"$sort": bson.M{"_id": 1}},
		{"$project": bson.M{
			"_id":   1,
			"value": 1,
			"count": 1,
		}},
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
		fmt.Println(chainId, val,val.AC.Address, val.AC.BlockNumber,"----------")
		// GetBlockTransactions(137, 30591094)

	}
}
