package proposal

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	// "github.com/robfig/cron"
	"github.com/autocompound/docker_backend/governance/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// ethcommon "github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/ethclient"
)

const CollectionName = "proposals"

// Models should only be concerned with database schema, more strict checking should be put in validator.
//
// HINT: If you want to split null and "", you should use *string instead of string.
type ProposalModel struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Created          time.Time          `bson:"_created" json:"_created"`
	Modified         time.Time          `bson:"_modified" json:"_modified"`
	Chain_Id         int                `bson:"chain_id" json:"chain_id"`
	Transaction_Hash string             `bson:"transaction_hash" json:"transaction_hash"`
	// Proposal_Type    string             `bson:"proposal_type" json:"proposal_type"`
	Block_Number   int     `bson:"block_number" json:"block_number"`
	Status         string  `bson:"status" json:"status"`
	Proposal_Id    int     `bson:"proposal_id" json:"proposal_id"`
	Proposer       string  `bson:"proposer" json:"proposer"`
	Eta            int     `bson:"eta" json:"eta"`
	Start_Time     int     `bson:"start_time" json:"start_time"`
	End_Time       int     `bson:"end_time" json:"end_time"`
	Description    string  `bson:"description" json:"description"`
	Voting_Period  int     `bson:"voting_period" json:"voting_period"` // in days
	For_Votes      float64 `bson:"for_votes" json:"for_votes"`
	Against_Votes  float64 `bson:"against_votes" json:"against_votes"`
	Canceled       bool    `bson:"canceled" json:"canceled"`
	Executed       bool    `bson:"executed" json:"executed"`
	Title          string  `bson:"title" json:"title"`
	Db_Description string  `bson:"db_description" json:"db_description"`
	Proposal_Type  int     `bson:"proposal_type" json:"proposal_type"` //1 for core 2 for community
	Cron_Status    string  `bson:"cron_status" json:"cron_status"`
}

//struct for filters
type Filters struct {
	// Token_Type string `bson: "token_type", json:"token_type"`
	Status     string `bson: "source", json:"source"`
	Chain_Id   int64  `bson: "chain_id", json:"chain_id"`
}

// init function runs first time
func init() {}

// You could input an ProposalModel which will be saved in database returning with error info
// 	if err := SaveOne(&proposalModel); err != nil { ... }
func SaveOne(data *ProposalModel) (string, error) {
	client := common.GetDB()
	record := &ProposalModel{}
	newID := ""

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// to check for unique email address
	err := collection.FindOne(ctx, bson.M{"transaction_hash": data.Transaction_Hash}).Decode(&record)
	if err != nil {
		res, err := collection.InsertOne(ctx, data)
		fmt.Println(res.InsertedID, "Inserted")
		// newID = res.InsertedID.(string)
		newID = fmt.Sprintf("%s", res.InsertedID)
		newID = strings.Replace(newID, "ObjectID(", "", -1)
		newID = strings.Replace(newID, `"`, "", -1)
		newID = strings.Replace(newID, `)`, "", -1)

		// get transaction data
		go UpdateRecordStatusBackground(newID, data.Transaction_Hash, data.Chain_Id)
		return newID, err
	}
	return newID, errors.New("proposal already exists!")
}

// You could input an ProposalModel which will be saved in database returning with error info
// update the proposal
func UpdateOne(data *ProposalModel) (*mongo.UpdateResult, error) {
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
	if data.Transaction_Hash != "" {
		update["transaction_hash"] = data.Transaction_Hash
	}

	update = bson.M{"$set": update}
	result, err := collection.UpdateOne(ctx, bson.M{"_id": data.ID}, update, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}

// You could input string which will be saved in database returning with error info
// 	if err := FindOne(&proposalModel); err != nil { ... }
func GetRecord(ID string) (ProposalModel, error) {
	client := common.GetDB()
	record := &ProposalModel{}

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

	return *record, err
}

// Farm list api with page and limit
func GetTotal(filters Filters) int64 {
	client := common.GetDB()

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := bson.M{"chain_id": filters.Chain_Id}
	if filters.Status != "" {
		if filters.Status == "active" {
			t := time.Now()
			// fmt.Println(t.Unix(),"time")
			// get unix timestamp
			query["end_time"] = bson.M{"$gte":t.Unix()}
			query["start_time"] = bson.M{"$lt":t.Unix()}
		}
	}

	num, err := collection.CountDocuments(ctx, query)
	if err != nil {
		return 0
	}
	return num
}

// Farm list api with page and limit
func GetAll(page int64, limit int64, filters Filters, sort_by string) ([]*ProposalModel, error) {
	client := common.GetDB()
	var records []*ProposalModel

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	sorting := bson.D{{"_created", -1}}

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	opts := options.Find().SetSort(sorting).SetSkip((page - 1) * limit).SetLimit(limit)
	//SetProjection(bson.M{"_id": 0, "_created": 1, "_modified": 1, "firstname": 1, "lastname": 1, "status": 1, "email": 1, "role": 1, "passwordhash": 0})
	query := bson.M{"chain_id": filters.Chain_Id}
	if filters.Status != "" {
		if filters.Status == "active" {
			t := time.Now()
			// fmt.Println(t.Unix(),"time")
			// get unix timestamp
			query["end_time"] = bson.M{"$gte":t.Unix()}
			query["start_time"] = bson.M{"$lt":t.Unix()}
		}
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

// struct for aggregate response
type Result struct {
	ID     string `bson:"_id" json:"_id"`
	Source string `bson:"source" json:"source"`
	Count  int    `bson:"count" json:"count"`
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

	// options for update
	opts := options.Update().SetUpsert(false)

	modified := time.Now()
	update := bson.M{"_modified": modified}

	txStatus, block_number := common.GetTransaction(transaction_hash, chain_id, 0)
	fmt.Println("Transaction Status", txStatus)

	update["status"] = "reverted"
	update["block_number"] = block_number
	// checking for success

	if txStatus == 1 {
		update["status"] = "Active"
	}
	if txStatus == -1 {
		update["status"] = "Processing"
	}

	update = bson.M{"$set": update}
	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update, opts)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("updated: %v", result)
}
