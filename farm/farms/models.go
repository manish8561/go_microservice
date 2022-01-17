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

// Models should only be concerned with database schema, more strict checking should be put in validator.
//
// HINT: If you want to split null and "", you should use *string instead of string.
type FarmModel struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Created    time.Time          `bson:"_created" json:"_created"`
	Modified   time.Time          `bson:"_modified" json:"_modified"`
	PID        int
	Token      string
	TokenType  string
	Status     string
	Masterchef string
	Vault      string
	// PasswordHash string `json:"-"` // to hide filed in json
}


// You could input the conditions and it will return an FarmModel in database with error info.
// 	farmModel, err := FindOneUser(&FarmModel{Username: "username0"})
func FindOneUser(token string) (FarmModel, error) {
	person := &FarmModel{}

	client := common.GetDB()

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"token": token}).Decode(&person)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Found user:", person.ID.Hex())
	}
	return *person, err
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
	err := collection.FindOne(ctx, bson.M{"token": data.Token}).Decode(&person)
	if err != nil {
		res, err := collection.InsertOne(ctx, data)
		fmt.Println(res, "Inserted")
		return err
	}
	return errors.New("farm already exists!")
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
func GetAll(page int64, limit int64) ([]*FarmModel, error) {
	client := common.GetDB()
	var farms  []*FarmModel

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	opts := options.Find().SetSkip((page - 1) * limit).SetLimit(limit)
	//SetProjection(bson.M{"_id": 0, "_created": 1, "_modified": 1, "firstname": 1, "lastname": 1, "status": 1, "email": 1, "role": 1, "passwordhash": 0})

	cursor, err := collection.Find(ctx, bson.M{},opts)
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

// You could update properties of an FarmModel to database returning with error info.
//  err := db.Model(farmModel).Update(FarmModel{Username: "wangzitian0"}).Error
// func (model *FarmModel) Update(data interface{}) error {
// 	db := common.GetDB()
// 	err := db.Model(model).Update(data).Error
// 	return err
// }
