package farms

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	// "github.com/robfig/cron"
	"github.com/autocompound/docker_backend/farm/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

const CollectionName = "farms"

type Token struct {
	Address  string  `bson:"address" json:"address"`
	Name     string  `bson:"name" json:"name"`
	Symbol   string  `bson:"symbol" json:"symbol"`
	Supply   float64 `bson:"supply" json:"supply"`
	Price    float64 `bson:"price" json:"price"`
	Decimals int     `bson:"decimals" json:"decimals"`
	Img      string  `bson:"img" json:"img"`
}

// Models should only be concerned with database schema, more strict checking should be put in validator.
//
// HINT: If you want to split null and "", you should use *string instead of string.
type FarmModel struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Created          time.Time          `bson:"_created" json:"_created"`
	Modified         time.Time          `bson:"_modified" json:"_modified"`
	Chain_Id         int                `bson:"chain_id" json:"chain_id"`
	Transaction_Hash string             `bson:"transaction_hash" json:"transaction_hash"`
	PID              int                `bson:"pid" json:"pid"`
	Address          string             `bson:"address" json:"address"` //address field of strategy
	Name             string             `bson:"name" json:"name"`
	Token_Type       string             `bson:"token_type" json:"token_type"`
	Deposit_Token    string             `bson:"deposit_token" json:"deposit_token"`
	Status           string             `bson:"status" json:"status"`
	Masterchef       string             `bson:"masterchef" json:"masterchef"`
	// masterchef for pancakeswap, stakingRewards address for quickswap
	Router             string  `bson:"router" json:"router"`
	Weth               string  `bson:"weth" json:"weth"`
	Stake              string  `bson:"stake" json:"stake"`       //staking contract address
	AC_Token           string  `bson:"ac_token" json:"ac_token"` //autocompound token
	Reward             string  `bson:"reward" json:"reward"`     //cake address
	Bonus_Multiplier   int     `bson:"bonus_multiplier" json:"bonus_multiplier"`
	Token_Per_Block    float64 `bson:"token_per_block" json:"token_per_block"`
	FarmType           string  `bson:"farmType" json:"farmType"`
	Source             string  `bson:"source" json:"source"`
	Source_Link        string  `bson:"source_link" json:"source_link"`
	Autocompound_Check bool    `bson:"autocompound_check" json:"autocompound_check"`
	Tvl_Staked         float64 `bson:"tvl_staked" json:"tvl_staked"`
	Daily_APR          float64 `bson:"daily_apr" json:"daily_apr"`
	Daily_APY          float64 `bson:"daily_apy" json:"daily_apy"`
	Weekly_APY         float64 `bson:"weekly_apy" json:"weekly_apy"`
	Yearly_APY         float64 `bson:"yearly_apy" json:"yearly_apy"`
	Price_Pool_Token   float64 `bson:"price_pool_token" json:"price_pool_token"`
	Yearly_Swap_Fees   float64 `bson:"yearly_swap_fees" json:"yearly_swap_fees"`
	Token0             Token   `bson:"token0" json:"token0"`
	Token1             Token   `bson:"token1" json:"token1"`
}

//struct for filters
type Filters struct {
	Token_Type string `bson: "token_type", json:"token_type"`
	Source     string `bson: "source", json:"source"`
	Name       string `bson: "name", json:"name"`
	Chain_Id   int64  `bson: "chain_id", json:"chain_id"`
}

// init function runs first time
func init() {
	//creating the index in the collection
	common.AddIndex(os.Getenv("MONGO_DATABASE"), CollectionName, bson.D{{"address", 1}, {"status", 1}, {"chain_id", 1}, {"name", "text"}, {"tvl_staked", 1}, {"token_per_block", 1}, {"token_type", 1}})
}

// You could input an FarmModel which will be saved in database returning with error info
// 	if err := SaveOne(&farmModel); err != nil { ... }
func SaveOne(data *FarmModel) (string, error) {
	client := common.GetDB()
	record := &FarmModel{}
	newID := ""

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// to check for unique email address
	err := collection.FindOne(ctx, bson.M{"deposit_token": data.Deposit_Token}).Decode(&record)
	if err != nil {
		res, err := collection.InsertOne(ctx, data)
		fmt.Println(res.InsertedID, "Inserted")
		// newID = res.InsertedID.(string)
		newID = fmt.Sprintf("%s", res.InsertedID)
		newID = strings.Replace(newID, "ObjectID(", "", -1)
		newID = strings.Replace(newID, `"`, "", -1)
		newID = strings.Replace(newID, `)`, "", -1)
		return newID, err
	}
	return newID, errors.New("farm already exists!")
}

// You could input an FarmModel which will be saved in database returning with error info
// update the farm
func UpdateOne(data *FarmModel) (*mongo.UpdateResult, error) {
	client := common.GetDB()

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := primitive.ObjectIDFromHex("")
	// if err != nil {
	// 	return nil, err
	// }
	if data.ID == res {
		return nil, errors.New("Object ID is required field")
	}
	if data.Token_Type == "" {
		return nil, errors.New("Token Type is required field")
	}

	// options for update
	opts := options.Update().SetUpsert(false)

	modified := time.Now()
	update := bson.M{"_modified": modified, "token_type": data.Token_Type}

	if data.Address != "" {
		update["address"] = strings.ToLower(data.Address)
	}
	if data.Status != "" {
		update["status"] = data.Status
	}
	if data.Chain_Id > 0 {
		update["chain_id"] = data.Chain_Id
	}
	if data.Transaction_Hash != "" {
		update["transaction_hash"] = data.Transaction_Hash
	}
	if data.PID > 0 {
		update["pid"] = data.PID
	}
	if data.Name != "" {
		update["name"] = data.Name
	}
	if data.Deposit_Token != "" {
		update["deposit_token"] = strings.ToLower(data.Deposit_Token)
	}
	if data.Masterchef != "" {
		update["masterchef"] = strings.ToLower(data.Masterchef)
	}
	if data.Router != "" {
		update["router"] = strings.ToLower(data.Router)
	}
	if data.Weth != "" {
		update["weth"] = strings.ToLower(data.Weth)
	}
	if data.Stake != "" {
		update["stake"] = strings.ToLower(data.Stake)
	}
	if data.AC_Token != "" {
		update["ac_token"] = strings.ToLower(data.AC_Token)
	}
	if data.Reward != "" {
		update["reward"] = strings.ToLower(data.Reward)
	}
	if data.FarmType != "" {
		update["farmType"] = data.FarmType
	}
	if data.Source != "" {
		update["source"] = data.Source
	}
	if data.Source_Link != "" {
		update["source_link"] = data.Source_Link
	}
	if data.Bonus_Multiplier > 1 {
		update["bonus_multiplier"] = data.Bonus_Multiplier
	}
	if data.Token_Per_Block > 0 {
		update["token_per_block"] = data.Token_Per_Block
	}

	if data.Tvl_Staked > 0 {
		update["tvl_staked"] = data.Tvl_Staked
	}
	if data.Daily_APR > 0 {
		update["daily_apr"] = data.Daily_APR
	}
	if data.Daily_APR > 0 {
		update["daily_apr"] = data.Daily_APR
	}
	if data.Daily_APY > 0 {
		update["daily_apy"] = data.Daily_APY
	}
	if data.Weekly_APY > 0 {
		update["weekly_apy"] = data.Weekly_APY
	}
	if data.Yearly_APY > 0 {
		update["yearly_apy"] = data.Yearly_APY
	}
	if data.Price_Pool_Token > 0 {
		update["price_pool_token"] = data.Price_Pool_Token
	}
	if data.Yearly_Swap_Fees > 0 {
		update["yearly_swap_fees"] = data.Yearly_Swap_Fees
	}
	// check struct for empty value
	if (data.Token0 != Token{}) {
		update["token0"] = data.Token0
	}
	if (data.Token1 != Token{}) {
		update["token1"] = data.Token1
	}
	if data.Token_Type == "token" || data.Token_Type == "stable" {
		update["token1"] = nil
	}
	update = bson.M{"$set": update}
	result, err := collection.UpdateOne(ctx, bson.M{"_id": data.ID}, update, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}

// You could input an FarmModel which will be saved in database returning with error info
// update stake in the farm
func UpdateStake(data *FarmModel) (*mongo.UpdateResult, error) {
	client := common.GetDB()

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := primitive.ObjectIDFromHex("")

	if data.ID == res {
		return nil, errors.New("Object ID is required field")
	}
	// options for update
	opts := options.Update().SetUpsert(false)

	modified := time.Now()
	update := bson.M{"_modified": modified, "token_type": data.Token_Type}

	if data.Stake != "" {
		update["stake"] = data.Stake
	}
	update = bson.M{"$set": update}
	result, err := collection.UpdateOne(ctx, bson.M{"_id": data.ID}, update, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}

// You could input an FarmModel which will be updated in database returning with error info
// 	if err := UpdateOne(&farmModel); err != nil { ... }
func TransactionUpdate(data *FarmModel) error {
	client := common.GetDB()
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// to check for unique email address
	opts := options.Update().SetUpsert(false)
	// update := bson.D{{"$set", bson.D{{"token", "newemail@example.com"}}}}
	// update := bson.D{{"$set", data}}
	address := ""
	status := "processing"
	modified := time.Now()
	if data.Address != "" {
		address = strings.ToLower(data.Address)
		status = "active"
	}
	update := bson.M{"$set": bson.M{"transaction_hash": data.Transaction_Hash, "status": status, "_modified": modified, "address": address}}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": data.ID}, update, opts)
	if err != nil {
		return err
	}
	fmt.Println(res, "Updated")
	return err
}

// SetOperator for autocompound check
func SetOperator(data *FarmModel) error {
	client := common.GetDB()
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// to check for unique email address
	opts := options.Update().SetUpsert(false)
	// update := bson.D{{"$set", bson.D{{"token", "newemail@example.com"}}}}
	// update := bson.D{{"$set", data}}

	modified := time.Now()
	update := bson.M{"$set": bson.M{"autocompound_check": data.Autocompound_Check, "_modified": modified}}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": data.ID}, update, opts)
	if err != nil {
		return err
	}
	fmt.Println(res, "Updated")
	return err
}

// You could input string which will be saved in database returning with error info
// 	if err := FindOne(&farmModel); err != nil { ... }
func GetFarm(ID string) (FarmModel, error) {
	client := common.GetDB()
	farm := &FarmModel{}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//convert string to objectid
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return *farm, err
	}

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	// opts := options.FindOne().SetProjection(bson.M{"_id": 0, "_created": 1, "_modified": 1, "firstname": 1, "lastname": 1, "status": 1, "email": 1, "role": 1, "passwordhash": 0})
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&farm)

	return *farm, err
}

// Farm list api with page and limit
func GetTotal(status string, filters Filters) int64 {
	client := common.GetDB()

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := bson.M{"chain_id": filters.Chain_Id}
	if status != "" {
		query["status"] = status
	}
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

	num, err := collection.CountDocuments(ctx, query)
	if err != nil {
		return 0
	}
	return num
}

// Farm list api with page and limit
func GetAll(page int64, limit int64, status string, filters Filters, sort_by string) ([]*FarmModel, error) {
	client := common.GetDB()
	var farms []*FarmModel

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sorting := bson.D{{"_created", -1}}
	if sort_by == "recent" {
		sorting = bson.D{{"_created", -1}}
	}
	if sort_by == "apy" {
		sorting = bson.D{{"daily_apy", -1}}
	}
	if sort_by == "tvl" {
		sorting = bson.D{{"tvl_staked", -1}}
	}
	if sort_by == "yourTvl" {
		sorting = bson.D{{"tvl_staked", -1}}
	}

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	opts := options.Find().SetSort(sorting).SetSkip((page - 1) * limit).SetLimit(limit)
	//SetProjection(bson.M{"_id": 0, "_created": 1, "_modified": 1, "firstname": 1, "lastname": 1, "status": 1, "email": 1, "role": 1, "passwordhash": 0})
	query := bson.M{"chain_id": filters.Chain_Id}
	if status != "" {
		query["status"] = status
	}
	if filters.Token_Type != "" {
		// checking for the stable in token type
		query["token_type"] = filters.Token_Type
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

	cursor, err := collection.Find(ctx, query, opts)
	if err != nil {
		return farms, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &farms)

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return farms, err
}

// struct for aggregate response
type Result struct {
	ID     string `bson:"_id" json:"_id"`
	Source string `bson:"source" json:"source"`
	Count  int    `bson:"count" json:"count"`
}

// get multiple tags of source from farms
func GetSource() ([]*Result, error) {
	client := common.GetDB()

	var records []*Result

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := []bson.M{
		{"$group": bson.M{"_id": "$source", "source": bson.M{"$first": "$source"}, "count": bson.M{"$sum": 1}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
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

// get total tvl from farms
func GetTvl() float64 {
	client := common.GetDB()
	clientRedis := common.GetRedisDB()

	//get data from redis
	val, err := clientRedis.Get("totalTvl").Result()
	if s, err := strconv.ParseFloat(val, 64); err == nil {
		return s
	}
	if err != nil {
		fmt.Println(err, "get data from redis")
	}
	type TVL struct {
		Total float64 `bson:"total" json:"total"`
	}
	var results []TVL

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := []bson.M{
		{"$match": bson.M{"status": "active"}},
		{"$group": bson.M{"_id": nil, "total": bson.M{"$sum": "$tvl_staked"}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &results)

	if err := cursor.Err(); err != nil {
		fmt.Println(cursor)
		return 0
	}
	if len(results) == 0 {
		return 0
	}
	//adding in redis
	err = clientRedis.Set("totalTvl", results[0].Total, 10*time.Minute).Err()
	if err != nil {
		fmt.Println(err, "saving in redis")
	}
	return results[0].Total
}

//get avg autocompound per block
func GetACPerBlock(chainId int64) float64 {
	client := common.GetDB()
	clientRedis := common.GetRedisDB()

	//get data from redis
	val, err := clientRedis.Get("acPerToken").Result()
	if s, err := strconv.ParseFloat(val, 64); err == nil {
		return s
	}
	if err != nil {
		fmt.Println(err, "get data from redis")
	}

	type Avg struct {
		// ID      nil     `bson:"_id" bson:"_id"`
		Average float64 `bson:"average" json:"average"`
	}
	var results []Avg

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := []bson.M{
		{"$match": bson.M{"status": "active", "chain_id": chainId}},
		{"$group": bson.M{"_id": nil, "average": bson.M{"$avg": "$token_per_block"}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &results)
	if err := cursor.Err(); err != nil {
		return 0
	}

	// checking length of array
	if len(results) == 0 {
		return 0
	}

	//adding in redis
	err = clientRedis.Set("acPerToken", results[0].Average, 10*time.Minute).Err()
	if err != nil {
		fmt.Println(err, "saving in redis")
	}

	return results[0].Average
}
