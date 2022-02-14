package farms

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

const CollectionName = "farms"

type Token struct {
	Address  string `bson:"address" json:"address"`
	Name     string `bson:"name" json:"name"`
	Symbol   string `bson:"symbol" json:"symbol"`
	Supply   int    `bson:"supply" json:"supply"`
	Price    int    `bson:"price" json:"price"`
	Decimals int    `bson:"decimals" json:"decimals"`
	Img      string `bson:"img" json:"img"`
}

// Models should only be concerned with database schema, more strict checking should be put in validator.
//
// HINT: If you want to split null and "", you should use *string instead of string.
type FarmModel struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Created          time.Time          `bson:"_created" json:"_created"`
	Modified         time.Time          `bson:"_modified" json:"_modified"`
	Strategy_ABI     string             `bson:"strategy_abi" json:"strategy_abi"`
	PID              int                `bson:"pid" json:"pid"`
	Address          string             `bson:"address" json:"address"` //address field of strategy
	Name             string             `bson:"name" json:"name"`
	Token_Type       string             `bson:"token_type" json:"token_type"`
	Deposit_Token    string             `bson:"deposit_token" json:"deposit_token"`
	Status           string             `bson:"status" json:"status"`
	Masterchef       string             `bson:"masterchef" json:"masterchef"`
	Router           string             `bson:"router" json:"router"`
	Stake            string             `bson:"stake" json:"stake"`
	Reward           string             `bson:"reward" json:"reward"`
	Bonus_Multiplier int                `bson:"bonus_multiplier" json:"bonus_multiplier"`
	Token_Per_Block  int                `bson:"token_per_block" json:"token_per_block"`
	Source           string             `bson:"source" json:"source"`
	Source_Link      string             `bson:"source_link" json:"source_link"`

	Tvl_Staked       int    `bson:"tvl_staked" json:"tvl_staked"`
	Daily_APR        int    `bson:"daily_apr" json:"daily_apr"`
	Daily_APY        int    `bson:"daily_apy" json:"daily_apy"`
	Weekly_APY       int    `bson:"weekly_apy" json:"weekly_apy"`
	Yearly_APY       int    `bson:"yearly_apy" json:"yearly_apy"`
	Price_Pool_Token int    `bson:"price_pool_token" json:"price_pool_token"`
	Yearly_Swap_Fees int    `bson:"yearly_swap_fees" json:"yearly_swap_fees"`
	Token0           Token  `bson:"token0" json:"token0"`
	Token1           Token  `bson:"token1" json:"token1"`
	Gauge_Info       string `bson:"gauge_info" json:"gauge_info"`

	// "gaugeInfo": {
	//     "address": "0x5d1E2Ad05A946Ac05f188dAa0BF8c9b010dE356d",
	//     "tvlStaked": 32410.18832923819,
	//     "snobDailyAPR": 0.01243831494737943,
	//     "snobWeeklyAPR": 0.08706820463165602,
	//     "snobYearlyAPR": 4.539984955793492,
	//     "fullDailyAPY": 0.12256110767415393,
	//     "fullWeeklyAPY": 0.8604791091782902,
	//     "fullYearlyAPY": 53.81597928044311,
	//     "snobAllocation": 0.0025946792694176986,
	//     "__typename": "GaugeInfo"
	// },
	// PasswordHash string `json:"-"` // to hide filed in json
}

// You could input an FarmModel which will be saved in database returning with error info
// 	if err := SaveOne(&farmModel); err != nil { ... }
func SaveOne(data *FarmModel) error {
	client := common.GetDB()
	person := &FarmModel{}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// to check for unique email address
	err := collection.FindOne(ctx, bson.M{"deposit_token": data.Deposit_Token}).Decode(&person)
	if err != nil {
		res, err := collection.InsertOne(ctx, data)
		fmt.Println(res, "Inserted")
		return err
	}
	return errors.New("farm already exists!")
}

// You could input an FarmModel which will be updated in database returning with error info
// 	if err := UpdateOne(&farmModel); err != nil { ... }
func UpdateOne(data *FarmModel) error {
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
func GetTotal(status string) int64 {
	client := common.GetDB()

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := bson.M{"status": status}
	if status == "" {
		query = bson.M{}
	}

	num, err := collection.CountDocuments(ctx, query)
	if err != nil {
		return 0
	}
	return num
}

// Farm list api with page and limit
func GetAll(page int64, limit int64, status string) ([]*FarmModel, error) {
	client := common.GetDB()
	var farms []*FarmModel

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	opts := options.Find().SetSort(bson.D{{"_created", -1}}).SetSkip((page - 1) * limit).SetLimit(limit)
	//SetProjection(bson.M{"_id": 0, "_created": 1, "_modified": 1, "firstname": 1, "lastname": 1, "status": 1, "email": 1, "role": 1, "passwordhash": 0})
	query := bson.M{"status": status}
	if status == "" {
		query = bson.M{}
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
