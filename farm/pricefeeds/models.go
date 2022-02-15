package pricefeeds

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/autocompound/docker_backend/farm/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

const CollectionName = "pricefeeds"

// Models should only be concerned with database schema, more strict checking should be put in validator.
//
// HINT: If you want to split null and "", you should use *string instead of string.
type PriceFeedModel struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Created      time.Time          `bson:"_created" json:"_created"`
	Modified     time.Time          `bson:"_modified" json:"_modified"`
	Coingeeko_Id string             `bson:"coingeeko_id" json:"coingeeko_id"`
	Symbol       string             `bson:"symbol" json:"symbol"` //address field of strategy
	Price        float64            `bson:"price" json:"price"`
}

// You could input an PriceFeedModel which will be saved in database returning with error info
// 	if err := SaveOne(&farmModel); err != nil { ... }
func SaveOne(data *PriceFeedModel) error {
	client := common.GetDB()
	person := &PriceFeedModel{}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// to check for unique email address
	err := collection.FindOne(ctx, bson.M{}).Decode(&person)
	if err != nil {
		res, err := collection.InsertOne(ctx, data)
		fmt.Println(res, "Inserted")
		return err
	}
	return errors.New("farm already exists!")
}

// You could input an PriceFeedModel which will be updated in database returning with error info
// 	if err := UpdateOne(&farmModel); err != nil { ... }
func UpdateOne(data *PriceFeedModel) error {
	client := common.GetDB()
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// to check for unique email address
	opts := options.Update().SetUpsert(false)
	// update := bson.D{{"$set", bson.D{{"token", "newemail@example.com"}}}}
	update := bson.D{{"$set", data}}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": data.ID}, update, opts)
	if err != nil {
		return err
	}
	fmt.Println(res, "Updated")
	return err
}

// You could input string which will be saved in database returning with error info
// 	if err := FindOne(&farmModel); err != nil { ... }
func GetFarm(ID string) (PriceFeedModel, error) {
	client := common.GetDB()
	farm := &PriceFeedModel{}

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

// price feed update and get return value
func GetTokenPrice(symbol string, check bool) (PriceFeedModel, error) {
	client := common.GetDB()
	record := &PriceFeedModel{}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//find record in db
	err := collection.FindOne(ctx, bson.M{"symbol": symbol}).Decode(&record)
	if(!check){
		return	*record, err
	}
	//get price from coingeeko
	str := record.Coingeeko_Id
	//calling from utils file
	f := common.GetPrice(str)
	fmt.Println(str,f)

	opts := options.Update().SetUpsert(false)
	record.Created = time.Now()
	record.Modified = time.Now()
	record.Price = f
	//update price in collection
	update := bson.D{{"$set", record}}
	//update query
	res, err := collection.UpdateOne(ctx, bson.M{"_id": record.ID}, update, opts)
	if err != nil {
		return *record, err
	}
	fmt.Println(res, "Updated", f)

	return *record, err
}