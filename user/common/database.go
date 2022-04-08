package common

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

)

var client *mongo.Client

// Opening a database and save the reference to `Database` struct.
func InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//function to monitor mongodb logs mongodb queries and data
	monitor := &event.CommandMonitor{
		Started: func(_ context.Context, e *event.CommandStartedEvent) {
			fmt.Println(e.Command)
		},
		Succeeded: func(_ context.Context, e *event.CommandSucceededEvent) {
			fmt.Println(e.Reply)
		},
		Failed: func(_ context.Context, e *event.CommandFailedEvent) {
			fmt.Println(e)
		},
	}
	// client options
	clientOpts := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	if os.Getenv("GIN_MODE") != "release" {
		clientOpts.SetMonitor(monitor)
	}
	// connect
	c, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	client = c
}

// This function will create a temporarily database for running testing cases
// func TestDBInit() *gorm.DB {
// 	test_db, err := gorm.Open("sqlite3", "./../gorm_test.db")
// 	if err != nil {
// 		fmt.Println("db err: (TestDBInit) ", err)
// 	}
// 	test_db.DB().SetMaxIdleConns(3)
// 	test_db.LogMode(true)
// 	DB = test_db
// 	return DB
// }

// // Delete the database after running testing cases.
// func TestDBFree(test_db *gorm.DB) error {
// 	test_db.Close()
// 	err := os.Remove("./../gorm_test.db")
// 	return err
// }

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *mongo.Client {
	return client
}

type UserModel struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Created   time.Time          `bson:"_created" json:"_created"`
	Modified  time.Time          `bson:"_modified" json:"_modified"`
	Firstname string
	Lastname  string
	Status    string
	Username  string
	Email     string
	Role      string
	// Image              *string
	PasswordHash string `json:"-"` // to hide filed in json
}

// You could input string which will be saved in database returning with error info
// 	if err := FindOne(&userModel); err != nil { ... }
func GetUserProfile(ID string) (UserModel, error) {
	client := GetDB()
	person := &UserModel{}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//convert string to objectid
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		// panic(err)
		return *person, err
	}

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	// opts := options.FindOne().SetProjection(bson.M{"_id": 0, "_created": 1, "_modified": 1, "firstname": 1, "lastname": 1, "status": 1, "email": 1, "role": 1, "passwordhash": 0})
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&person)

	return *person, err
}
