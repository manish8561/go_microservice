package tokens

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/autocompound/docker_backend/farm/common"
	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const CollectionName = "transfers"
const blockDff = 200

// Models should only be concerned with database schema, more strict checking should be put in validator.
// event Transfer(address indexed from, address indexed to, uint256 value);
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

// struct for graph data
type GraphDataModel struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Value     float64            `bson:"value" json:"value"`
	Timestamp int64              `bson:"timestamp" json:"timestamp"`
}

// struct for graph data
type GraphDataModel2 struct {
	ID    time.Time `json:"_id,omitempty" bson:"_id,omitempty"`
	Value float64   `bson:"value" json:"value"`
	Count int64     `bson:"count" json:"count"`
}

// struct for filters
type Filters struct {
	ChainId int64  `bson:"chainId" json:"chainId"`
	Address string `bson:"address" json:"address"`
}

// init func in go file
func init() {
	// create index
	common.AddIndex(os.Getenv("MONGO_DATABASE"), CollectionName, bson.D{{Key: "address", Value: 1}, {Key: "blockNumber", Value: -1}, {"from", 1}, {"to", 1}, {"createdAt", -1}})

	//start the cron
	StartCall()
}

// cron func call
func StartCall() {
	c := cron.New()
	c.AddFunc("0 */2 * * * *", func() {
		fmt.Println("[Job 1]Every 30 minutes job")
		fmt.Println()
		//calling get autocompounds
		GetAutocompound()
	})
	// Start cron with one scheduled job
	c.Start()
}

// You could input an TransferEventModel which will be saved in database returning with error info
//
//	if err := SaveOne(&farmModel); err != nil { ... }
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
//
//	if err := UpdateOne(&farmModel); err != nil { ... }
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
//
//	if err := FindOne(&farmModel); err != nil { ... }
func GetRecord(chainId int, ac string) (TransferEventModel, error) {
	client := common.GetDB()
	record := &TransferEventModel{}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	opts := options.FindOne().SetSort(bson.D{{Key: "blockNumber", Value: -1}})
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
	opts := options.Find().SetSort(bson.D{{Key: "_created", Value: -1}}).SetSkip((page - 1) * limit).SetLimit(limit)
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

/*
	get last seven days volume

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
	currentDate := getDateWithoutTime(time.Now().Unix())
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
func callingDelete(CollectionName string, ChainId int) {
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
		log.Printf("err in aggregate: %q", err)
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &records)
	if err != nil {
		log.Println("err in aggregate", err)
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
				log.Println("err in aggregate", err)
			}
			fmt.Println("delete response", res)
		}

	}
}

// function to get block timestamp
func GetBlockTimestamp(client *ethclient.Client, blockNum int64) int64 {
	blockNumber := big.NewInt(blockNum)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Println(err)
	}

	// fmt.Printf("%t", block.Time())
	// fmt.Println(block.Time(), block_num, "block timestamp")
	return int64(block.Time())

}

// Get date without time from timestamp
func getDateWithoutTime(timestamp int64) time.Time {
	currentDate := time.Unix(timestamp, 0).UTC()

	//get year month day
	y, m, d := currentDate.Date()
	//convert to date
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

// You could input string which will be saved in database returning with error info
//
//	if err := FindOne(&farmModel); err != nil { ... }
func GetContract(chainId int, ac string, blockNumber int64) error {
	// Create an IPC based RPC connection to a remote node
	conn := common.GetEthConnection(chainId)

	// to get latest blocknumber
	header, err := conn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Println(err)
	}
	lastestBlockNumber := header.Number.Int64()

	contractAddress := ethcommon.HexToAddress(ac)
	token, err := NewTokens(contractAddress, conn)
	if err != nil {
		log.Println("Failed to instantiate a Token contract: ", err)
		return err
	}
	decimals, err := token.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Println("Failed to retrieve token name:", err)
		return err
	}

	record, err := GetRecord(chainId, ac)
	if err != nil {
		log.Println("Failed to retrieve token name: ", err)
		return err
	}

	if (record != TransferEventModel{}) {
		lastBlockNumber := (record.LastBlockNumber + 1)
		newBlockNumber := lastBlockNumber + blockDff
		if newBlockNumber >= lastestBlockNumber {
			newBlockNumber = lastestBlockNumber
		}

		//query the logs
		query := ethereum.FilterQuery{
			FromBlock: big.NewInt(lastBlockNumber),
			ToBlock:   big.NewInt(newBlockNumber),
			Addresses: []ethcommon.Address{
				contractAddress,
			},
		}
		//logs from contract
		logs, err := conn.FilterLogs(context.Background(), query)
		if err != nil {
			log.Printf("Failed to retrieve token name: %q", err)
			return err
		}
		if len(logs) == 0 {
			go UpdateOne(record.ID, newBlockNumber)
			return nil
		}

		logTransferSig := []byte("Transfer(address,address,uint256)")
		// LogApprovalSig := []byte("Approval(address,address,uint256)")
		logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
		// logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

		dd := math.Pow(10, float64(decimals))

		for _, vLog := range logs {
			// fmt.Println(vLog.Topics[0].Hex(), "transfer hex")
			switch vLog.Topics[0].Hex() {
			// Transfer event hex
			case logTransferSigHash.Hex():

				transferEvent, err := token.ParseTransfer(vLog)
				if err != nil {
					log.Println(err)
					return err
				}

				//converting the string to float64
				transferValue, err := strconv.ParseFloat(transferEvent.Value.String(), 64)
				if err != nil {
					log.Println(err)
					return err
				}
				blockTimestamp := GetBlockTimestamp(conn, int64(vLog.BlockNumber))
				d := TransferEventModel{
					ChainId:         chainId,
					Address:         strings.ToLower(ac),
					TransactionHash: vLog.TxHash.Hex(),
					From:            strings.ToLower(transferEvent.From.Hex()),
					To:              strings.ToLower(transferEvent.To.Hex()),
					Value:           (transferValue / dd),
					BlockNumber:     int64(vLog.BlockNumber),
					LastBlockNumber: int64(vLog.BlockNumber),
					Timestamp:       blockTimestamp,
					CreatedAt:       getDateWithoutTime(blockTimestamp),
				}
				err = SaveOne(&d)
				if err != nil {
					log.Printf("Failed to retrieve token name: %v", err)
					return err
				}
			}
		}
		go UpdateOne(record.ID, newBlockNumber)
		//calling delete
		go callingDelete(CollectionName, chainId)
	} else {
		d := TransferEventModel{
			ChainId:         chainId,
			Address:         strings.ToLower(ac),
			TransactionHash: "",
			From:            "",
			To:              "",
			Value:           0,
			BlockNumber:     blockNumber,
			LastBlockNumber: blockNumber,
			Timestamp:       0,
			CreatedAt:       time.Now(),
		}
		err := SaveOne(&d)
		if err != nil {
			log.Println("Failed to retrieve token name: ", err)
			return err
		}
	}
	return err
}

// to get all autocompound address from constant file in a map
func GetAutocompound() {
	for chainId, val := range common.NetworkMap {
		//calling the contract as per chainId
		err := GetContract(chainId, val.AC, val.BlockNumber)
		log.Println(">>>Failed to retrieve token name: ", err)
	}

}
