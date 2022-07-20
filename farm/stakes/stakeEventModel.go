package stakes

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

	// "github.com/robfig/cron"
	"github.com/autocompound/docker_backend/farm/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robfig/cron"
)

const StakeEventCollection = "stakeEvents"
const UnStakeEventCollection = "unstakeEvents"
const blockDiff = 100

// Models should only be concerned with database schema, more strict checking should be put in validator.
//
// HINT: If you want to split null and "", you should use *string instead of string.
//struct for stake event
type StakeEventModel struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ChainId         int                `bson:"chainId" json:"chainId"`
	Staking         string             `bson:"staking" json:"staking"`
	TransactionHash string             `bson:"transactionHash" json:"transactionHash"`
	Account         string             `bson:"account" json:"account"` //address field of strategy
	Amount          float64            `bson:"amount" json:"amount"`
	BlockNumber     int64              `bson:"blockNumber" json:"blockNumber"`
	Timestamp       int64              `bson:"timestamp" json:"timestamp"`
}

//struct for unstake event
type UnstakeEventModel struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ChainId         int                `bson:"chainId" json:"chainId"`
	Staking         string             `bson:"staking" json:"staking"`
	TransactionHash string             `bson:"transactionHash" json:"transactionHash"`
	Account         string             `bson:"account" json:"account"` //address field of strategy
	Amount          float64            `bson:"amount" json:"amount"`
	BlockNumber     int64              `bson:"blockNumber" json:"blockNumber"`
	Timestamp       int64              `bson:"timestamp" json:"timestamp"`
}

//struct for filters
type EventFilters struct {
	ChainId   int64  `bson: "chainId" json:"chainId"`
	Account   string `bson: "account" json:"account"`
	Staking   string `bson: "staking" json:"staking"`
	EventType string `bson:"eventType" json:"eventType"`
}

//struct for votes with total
type EventResult struct {
	Total   int               `bson:"total" json:"total"`
	Records []StakeEventModel `bson:"records" json:"records"`
	// UnstakeEvent []UnstakeEventModel `bson:"unstakeEvent" json:"unstakeEvent"`
}

// init function runs first time
func init() {
	common.AddIndex(os.Getenv("MONGO_DATABASE"), StakeEventCollection, bson.D{{"blockNumber", "-1"}, {"account", "1"}, {"chainId", "1"}})
	common.AddIndex(os.Getenv("MONGO_DATABASE"), UnStakeEventCollection, bson.D{{"blockNumber", "-1"}, {"account", "1"}, {"chainId", "1"}})

	StartCall()
}

// cron func call
func StartCall() {
	c := cron.New()
	c.AddFunc("0 */2 * * * *", func() {
		fmt.Println("[Job 1]Every 2 minutes job\n")
		getStakingContracts()
		fmt.Println("cron job return value")
	})
	// Start cron with one scheduled job
	c.Start()
}

// You could input an StakeEventModel which will be saved in database returning with error info
// 	if err := SaveStakeEventOne(&StakeEventModel); err != nil { ... }
func SaveStakeEventOne(data *StakeEventModel) (string, error) {
	client := common.GetDB()
	newID := ""

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(StakeEventCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, data)
	fmt.Println(res.InsertedID, "Inserted")
	// newID = res.InsertedID.(string)
	newID = fmt.Sprintf("%s", res.InsertedID)
	newID = strings.Replace(newID, "ObjectID(", "", -1)
	newID = strings.Replace(newID, `"`, "", -1)
	newID = strings.Replace(newID, `)`, "", -1)
	return newID, err
}

// You could input an UnstakeEventModel which will be saved in database returning with error info
// 	if err := SaveUnstakeEventOne(&UnstakeEventModel); err != nil { ... }
func SaveUnstakeEventOne(data *UnstakeEventModel) (string, error) {
	client := common.GetDB()
	newID := ""

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(UnStakeEventCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, data)
	fmt.Println(res.InsertedID, "Inserted")
	// newID = res.InsertedID.(string)
	newID = fmt.Sprintf("%s", res.InsertedID)
	newID = strings.Replace(newID, "ObjectID(", "", -1)
	newID = strings.Replace(newID, `"`, "", -1)
	newID = strings.Replace(newID, `)`, "", -1)
	return newID, err
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
		log.Println("err in aggregate", err)

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

// GetAllStakeEvents list api with page and limit
func GetAllStakeEvents(page int64, limit int64, filters EventFilters) (*EventResult, error) {
	var records []*EventResult

	client := common.GetDB()
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(StakeEventCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := bson.M{"chainId": filters.ChainId, "staking": strings.ToLower(filters.Staking), "account": strings.ToLower(filters.Account)}

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
				{"$skip": (page - 1) * limit},
				{"$limit": limit},
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

// GetAllUnstakeEvents list api with page and limit
func GetAllUnstakeEvents(page int64, limit int64, filters EventFilters) (*EventResult, error) {
	var records []*EventResult

	client := common.GetDB()
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(UnStakeEventCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := bson.M{"chainId": filters.ChainId, "staking": strings.ToLower(filters.Staking), "account": strings.ToLower(filters.Account)}

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
				{"$skip": (page - 1) * limit},
				{"$limit": limit},
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

//get staking contract address from db
func getStakingContracts() {
	records, err := GetAllActive()
	if err != nil {
		fmt.Println("Error to get active strategy:", err)
		return
	}
	for _, element := range records {
		//call events from network
		err := GetContractEvent(element.Chain_Id, element.Address, int64(element.LastBlockNumber), element.ID)
		if err == nil {
			go callingDelete(StakeEventCollection, element.Chain_Id)
			go callingDelete(UnStakeEventCollection, element.Chain_Id)
		}

	}
}

// function to get block timestamp
func Get_Block_Timestamp(client *ethclient.Client, block_num int64) int64 {
	blockNumber := big.NewInt(block_num)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Println("get timestamp err", err)
		return 0
	}

	// fmt.Printf("%t", block.Time())
	// fmt.Println(block.Time(), block_num, "block timestamp")
	return int64(block.Time())

}

// You could input string which will be saved in database returning with error info
// 	if err := FindOne(&farmModel); err != nil { ... }
func GetContractEvent(chainId int, staking string, lastBlockNumber int64, ID primitive.ObjectID) error {
	// Create an IPC based RPC connection to a remote node
	conn := common.Get_Eth_Connection(chainId)

	// to get latest blocknumber
	header, err := conn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	lastestBlockNumber := header.Number.Int64()

	contractAddress := ethcommon.HexToAddress(staking)
	stakeObj, err := NewStakes(contractAddress, conn)
	if err != nil {
		log.Println("Failed to instantiate a Token contract: %v", err)
		return err
	}
	decimals, err := stakeObj.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Printf("Failed to retrieve token name: %v", err)
		return err
	}
	//traversal
	newBlockNumber := (lastBlockNumber + 1) + blockDiff

	if newBlockNumber >= lastestBlockNumber {
		newBlockNumber = lastestBlockNumber
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(lastBlockNumber + 1),
		ToBlock:   big.NewInt(newBlockNumber),
		Addresses: []ethcommon.Address{
			contractAddress,
		},
	}
	//logs from contract
	logs, err := conn.FilterLogs(context.Background(), query)
	if err != nil {
		log.Println(err)
		return err
	}
	fmt.Println("length event:", len(logs))

	if len(logs) == 0 {
		go UpdateLastBlockNumberOne(ID, newBlockNumber)
		return nil
	}

	logStakeSig := []byte("Stake(address,uint256)")
	logUnstakeSig := []byte("Unstake(address,uint256)")
	logStakeSigHash := crypto.Keccak256Hash(logStakeSig)
	logUnstakeSigHash := crypto.Keccak256Hash(logUnstakeSig)

	for _, vLog := range logs {
		// fmt.Println("Log Index: ", vLog.Index)
		// fmt.Println("vLog.Topics[0] ", vLog.Topics[0].Hex())

		switch vLog.Topics[0].Hex() {
		case logStakeSigHash.Hex():
			stakeEvent, err := stakeObj.ParseStake(vLog)
			if err != nil {
				log.Println(err)
				return err
			}

			blockTimestamp := Get_Block_Timestamp(conn, int64(vLog.BlockNumber))

			d := math.Pow(10, float64(decimals))
			//converting the string to float64
			s, err := strconv.ParseFloat(stakeEvent.Amount.String(), 64)
			if err != nil {
				log.Println(err)
				return err
			}
			//saving the event
			data := StakeEventModel{
				ChainId:         chainId,
				Staking:         strings.ToLower(staking),
				TransactionHash: vLog.TxHash.Hex(),
				Account:         strings.ToLower(stakeEvent.Account.Hex()),
				Amount:          (s / d),
				BlockNumber:     int64(vLog.BlockNumber),
				Timestamp:       blockTimestamp,
			}
			//saving data in the db
			res, err := SaveStakeEventOne(&data)
			fmt.Println("data after saving the event", res)

		case logUnstakeSigHash.Hex():

			unstakeEvent, err := stakeObj.ParseUnstake(vLog)
			if err != nil {
				log.Println(err)
				return err
			}

			blockTimestamp := Get_Block_Timestamp(conn, int64(vLog.BlockNumber))

			d := math.Pow(10, float64(decimals))
			//converting the string to float64
			s, err := strconv.ParseFloat(unstakeEvent.Amount.String(), 64)
			if err != nil {
				log.Println(err)
				return err
			}
			//saving the event
			data := UnstakeEventModel{
				ChainId:         chainId,
				Staking:         strings.ToLower(staking),
				TransactionHash: vLog.TxHash.Hex(),
				Account:         strings.ToLower(unstakeEvent.Account.Hex()),
				Amount:          (s / d),
				BlockNumber:     int64(vLog.BlockNumber),
				Timestamp:       blockTimestamp,
			}
			//saving data in the db
			res, err := SaveUnstakeEventOne(&data)
			fmt.Println("data after saving the event", res)
		}

		fmt.Printf("\n\n")
	}
	//calling delete duplicate for each event

	go UpdateLastBlockNumberOne(ID, newBlockNumber)

	go callingDelete(StakeEventCollection, chainId)
	go callingDelete(UnStakeEventCollection, chainId)
	//update event
	return err
}
