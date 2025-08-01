package common

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	//initialize the db
	InitDB()
}

// Opening a database and save the reference to `Database` struct.
func InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//function to monitor mongodb logs mongodb queries and data
	monitor := &event.CommandMonitor{
		Started: func(_ context.Context, e *event.CommandStartedEvent) {
			fmt.Println(e.Command)
		},
		// Succeeded: func(_ context.Context, e *event.CommandSucceededEvent) {
		// 	fmt.Println(e.Reply)
		// },
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
	fmt.Println("Mongodb is connected.")
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

// AddIndex adds an index to a collection in the specified database
// dbName: name of the database
// collection: name of the collection
// indexKeys: keys to be indexed, e.g., bson.D{{"field", 1}}
// Returns an error if the index creation fails	
func AddIndex(dbName string, collection string, indexKeys interface{}) error {
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db := GetDB() // get clients of mongodb connection
	serviceCollection := db.Database(dbName).Collection(collection)
	indexName, err := serviceCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: indexKeys,
	})
	if err != nil {
		fmt.Println("error", err)
		return err
	}
	fmt.Println(indexName)
	return nil
}
