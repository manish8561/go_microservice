package userfarms

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	// "github.com/robfig/cron"
	"github.com/autocompound/docker_backend/farm/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "go.mongodb.org/mongo-driver/mongo/readpref"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const CollectionName = "user_farms"

// Models should only be concerned with database schema, more strict checking should be put in validator.
//
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
	Chain_Id int64 `bson: "chain_id", json:"chain_id"`
}

// init function runs first time
func init() {
	// common.AddIndex(os.Getenv("MONGO_DATABASE"), CollectionName, bson.D{{"deposit_token", "text"}, {"name", "text"}})
}

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
		return newID, err
	}
	//sending old record ID
	newID = fmt.Sprintf("%s", record.ID)
	newID = strings.Replace(newID, "ObjectID(", "", -1)
	newID = strings.Replace(newID, `"`, "", -1)
	newID = strings.Replace(newID, `)`, "", -1)
	return newID, nil
}

// You could input an UserFarmsModel which will be saved in database returning with error info
// update the farm
func UpdateOne(data *UserFarmsModel) (*mongo.UpdateResult, error) {
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
	// options for update
	opts := options.Update().SetUpsert(false)

	modified := time.Now()
	update := bson.M{"_modified": modified}

	if data.Status != "" {
		update["status"] = data.Status
	}
	if data.Chain_Id > 0 {
		update["chain_id"] = data.Chain_Id
	}
	update = bson.M{"$set": update}
	result, err := collection.UpdateOne(ctx, bson.M{"_id": data.ID}, update, opts)
	if err != nil {
		return result, err
	}
	return result, nil
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

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := bson.M{"chain_id": filters.Chain_Id}
	if status != "" {
		query["status"] = status
	}

	num, err := collection.CountDocuments(ctx, query)
	if err != nil {
		return 0
	}
	return num
}

// Record list api with page and limit
func GetAll(page int64, limit int64, status string, filters Filters, sort_by string) ([]*UserFarmsModel, error) {
	client := common.GetDB()
	var records []*UserFarmsModel

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sorting := bson.D{{"_created", -1}}

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	opts := options.Find().SetSort(sorting).SetSkip((page - 1) * limit).SetLimit(limit)
	query := bson.M{"chain_id": filters.Chain_Id}
	if status != "" {
		query["status"] = status
	}

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

// Record delete function
func DeleteRecord(ID string) (bool, error) {
	client := common.GetDB()
	record := false

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//convert string to objectid
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return false, err
	}

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	// opts := options.FindOne().SetProjection(bson.M{"_id": 0, "_created": 1, "_modified": 1, "firstname": 1, "lastname": 1, "status": 1, "email": 1, "role": 1, "passwordhash": 0})
	res, err := collection.DeleteOne(ctx, bson.M{"_id": objID})

	if res != nil {
		record = true
	}

	return record, err
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
		fmt.Println("----------------------------------counter:", counter)
		time.Sleep(6 * time.Second)
		fmt.Printf("no transaction found: %v", err)
		return GetTransaction(transaction_hash, chain_id, counter)
	}
	fmt.Println("Token balance:", tx.Status, "-------------------------")
	fmt.Println("tx status:", tx.Status, tx.BlockNumber)

	return (int)(tx.Status)
}
