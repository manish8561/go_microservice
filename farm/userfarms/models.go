package userfarms

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	// "github.com/robfig/cron"
	"github.com/autocompound/docker_backend/farm/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "go.mongodb.org/mongo-driver/mongo/readpref"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	//implementation of other module struct (directly)
	FarmsModule "github.com/autocompound/docker_backend/farm/farms"
)

const CollectionName = "user_farms"

// Models should only be concerned with database schema, more strict checking should be put in validator.
// HINT: If you want to split null and "", you should use *string instead of string.
type UserFarmsModel struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Created          time.Time          `bson:"_created" json:"_created"`
	Modified         time.Time          `bson:"_modified" json:"_modified"`
	Chain_Id         int                `bson:"chain_id" json:"chain_id"`
	User             string             `bson:"user" json:"user"`         //address field of user wallet
	Strategy         string             `bson:"strategy" json:"strategy"` //address field of strategy
	Transaction_Hash string             `bson:"transaction_hash" json:"transaction_hash"`
	Status           string             `bson:"status" json:"status"`
}

//struct for filters
type Filters struct {
	User       string `bson: "user", json:"user"`
	Chain_Id   int64  `bson: "chain_id", json:"chain_id"`
	Token_Type string `bson: "token_type", json:"token_type"`
	Source     string `bson: "source", json:"source"`
	Name       string `bson: "name", json:"name"`
}

// init function runs first time
func init() {}

// You could input an UserFarmsModel which will be saved in database returning with error info
// 	if err := SaveOne(&stakeModel); err != nil { ... }
func SaveOne(data *UserFarmsModel) (string, error) {
	client := common.GetDB()
	record := &UserFarmsModel{}
	newID := ""

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// to check for unique email address
	err := collection.FindOne(ctx, bson.M{"strategy": data.Strategy}).Decode(&record)
	if err != nil {
		res, err := collection.InsertOne(ctx, data)
		fmt.Println(res.InsertedID, "Inserted")
		// newID = res.InsertedID.(string)
		newID = fmt.Sprintf("%s", res.InsertedID)
		newID = strings.Replace(newID, "ObjectID(", "", -1)
		newID = strings.Replace(newID, `"`, "", -1)
		newID = strings.Replace(newID, `)`, "", -1)

		// update transaction status in the record
		go UpdateRecordStatusBackground(newID, data.Transaction_Hash, data.Chain_Id)

		return newID, err
	}
	//sending old record ID
	newID = fmt.Sprintf("%s", record.ID)
	newID = strings.Replace(newID, "ObjectID(", "", -1)
	newID = strings.Replace(newID, `"`, "", -1)
	newID = strings.Replace(newID, `)`, "", -1)
	return newID, nil
}

// You could input string which will be saved in database returning with error info
// 	if err := FindOne(&stakeModel); err != nil { ... }
func GetRecord(ID string) (UserFarmsModel, error) {
	client := common.GetDB()
	record := &UserFarmsModel{}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//convert string to objectid
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return *record, err
	}

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	// opts := options.FindOne().SetProjection(bson.M{"_id": 0, "_created": 1, "_modified": 1, "firstname": 1, "lastname": 1, "status": 1, "email": 1, "role": 1, "passwordhash": 0})
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&record)

	go UpdateRecordStatusBackground(ID, record.Transaction_Hash, record.Chain_Id)
	// fmt.Println("Transaction Status", txStatus)
	return *record, err
}

// Record list api with page and limit
func GetTotal(status string, filters Filters) int64 {
	client := common.GetDB()
	var records []bson.M

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//filters on farms
	query := bson.M{}
	if filters.Token_Type != "" {
		query["farms.token_type"] = filters.Token_Type
		// checking for the stable in token type
		if strings.Contains(filters.Token_Type, "stable") {
			query["farms.token_type"] = primitive.Regex{Pattern: "^" + filters.Token_Type + "*", Options: "i"}
		}
	}
	if filters.Source != "" {
		query["farms.source"] = filters.Source
	}
	if filters.Name != "" {
		query["farms.name"] = primitive.Regex{Pattern: "^" + filters.Name + "*", Options: "i"}
	}

	// Specify a pipeline that will return the number of times each name appears
	// in the collection.
	pipeline := []bson.M{
		{"$match": bson.M{"status": status, "chain_id": filters.Chain_Id, "user": filters.User}},
		{"$lookup": bson.M{"from": "farms", "localField": "strategy", "foreignField": "address", "as": "farmsData"}},
		{"$project": bson.M{
			"chain_id":  1,
			"user":      1,
			"strategy":  1,
			"_created":  1,
			"_modified": 1,
			"farms":     bson.M{"$arrayElemAt": bson.A{"$farmsData", 0}}}},
		{"$match": query},
		{"$group": bson.M{"_id": "$user", "count": bson.M{"$sum": 1}}},
	}

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	opts := options.Aggregate()

	cursor, err := collection.Aggregate(ctx, pipeline, opts)

	if err != nil {
		return 0
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &records)

	if records == nil {
		return 0
	}

	//convert int32
	num := int64((records[0]["count"]).(int32))
	// n := int64(num)
	return num
}

// Record list api with page and limit
func GetAll(page int64, limit int64, status string, filters Filters, sort_by string) ([]*FarmsModule.FarmModel, error) {
	client := common.GetDB()
	var records []*FarmsModule.FarmModel

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//filters on farms
	query := bson.M{}
	if filters.Token_Type != "" {
		query["token_type"] = filters.Token_Type
		// checking for the stable in token type
		if strings.Contains(filters.Token_Type, "stable") {
			query["token_type"] = primitive.Regex{Pattern: "^" + filters.Token_Type + "*", Options: "i"}
		}
	}
	if filters.Source != "" {
		query["source"] = filters.Source
	}
	if filters.Name != "" {
		query["name"] = primitive.Regex{Pattern: "^" + filters.Name + "*", Options: "i"}
	}

	//sorting from userfarms
	sorting := bson.M{"_created": -1}
	if sort_by == "recent" {
		sorting = bson.M{"_created": -1}
	}
	if sort_by == "apy" {
		sorting = bson.M{"daily_apy": -1}
	}
	if sort_by == "tvl" {
		sorting = bson.M{"tvl_staked": -1}
	}
	if sort_by == "yourTvl" {
		sorting = bson.M{"tvl_staked": -1}
	}
	// Specify a pipeline that will return the number of times each name appears
	// in the collection.
	pipeline := []bson.M{
		{"$match": bson.M{"status": status, "chain_id": filters.Chain_Id, "user": filters.User}},
		{"$lookup": bson.M{"from": "farms", "localField": "strategy", "foreignField": "address", "as": "farmsData"}},
		{"$project": bson.M{
			"_id":   0,
			"user":  1,
			"farms": bson.M{"$arrayElemAt": bson.A{"$farmsData", 0}}},
		},
		{"$project": bson.M{
			"user":               1,
			"_created":           "$farms._created",
			"_modified":          "$farms._modified",
			"chain_id":           "$farms.chain_id",
			"pid":                "$farms.pid",
			"address":            "$farms.address",
			"name":               "$farms.name",
			"token_type":         "$farms.token_type",
			"deposit_token":      "$farms.deposit_token",
			"status":             "$farms.status",
			"masterchef":         "$farms.masterchef",
			"router":             "$farms.router",
			"weth":               "$farms.weth",
			"stake":              "$farms.stake",
			"ac_token":           "$farms.ac_token",
			"reward":             "$farms.reward",
			"bonus_multiplier":   "$farms.bonus_multiplier",
			"token_per_block":    "$farms.token_per_block",
			"source":             "$farms.source",
			"source_link":        "$farms.source_link",
			"autocompound_check": "$farms.autocompound_check",
			"tvl_staked":         "$farms.tvl_staked",
			"daily_apr":          "$farms.daily_apr",
			"daily_apy":          "$farms.daily_apy",
			"weekly_apy":         "$farms.weekly_apy",
			"yearly_apy":         "$farms.yearly_apy",
			"price_pool_token":   "$farms.price_pool_token",
			"yearly_swap_fees":   "$farms.yearly_swap_fees",
			"token0":             "$farms.token0",
			"token1":             "$farms.token1",
		}},
		{"$match": query},
		{"$sort": sorting},
		{"$skip": ((page - 1) * limit)},
		{"$limit": limit},
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

// go background function to update the transaction status
func UpdateRecordStatusBackground(ID string, transaction_hash string, chain_id int) {
	client := common.GetDB()

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctx := context.Background()

	//convert string to objectid
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		fmt.Printf("error %v", err)
		return
	}

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	// opts := options.FindOne().SetProjection(bson.M{"_id": 0, "_created": 1, "_modified": 1, "firstname": 1, "lastname": 1, "status": 1, "email": 1, "role": 1, "passwordhash": 0})

	// options for update
	opts := options.Update().SetUpsert(false)

	modified := time.Now()
	update := bson.M{"_modified": modified}

	txStatus := GetTransaction(transaction_hash, chain_id, 0)
	fmt.Println("Transaction Status", txStatus)

	update["status"] = "reverted"
	// checking for success

	if txStatus == 1 {
		update["status"] = "success"
	}
	if txStatus == -1 {
		update["status"] = "processing"
	}

	update = bson.M{"$set": update}
	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update, opts)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("updated: %v", result)
}

// Recurrsive function to get Transaction details
func GetTransaction(transaction_hash string, chain_id int, counter int) int {
	//get rpc from common file
	rpc := common.Get_RPC_ChainId(chain_id)
	//create eth client object
	conn, err := ethclient.Dial(rpc)
	if err != nil {
		fmt.Println("Failed to connect to the Ethereum client: %v", err)
	}

	//convert transaction string to hash
	hash := ethcommon.HexToHash(transaction_hash)

	//get transaction data
	tx, err := conn.TransactionReceipt(context.Background(), hash)
	if err != nil {
		counter = counter + 1
		//after 10 minutes
		if counter >= 100 {
			return -1
		}
		time.Sleep(6 * time.Second)
		fmt.Printf("no transaction found: %v", err)
		return GetTransaction(transaction_hash, chain_id, counter)
	}
	// fmt.Println("Token balance:", tx.Status, "-------------------------")
	// fmt.Println("tx status:", tx.Status, tx.BlockNumber)

	return (int)(tx.Status)
}
