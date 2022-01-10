package common

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Opening a database and save the reference to `Database` struct.
func InitDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+os.Getenv("MONGO_HOST")+":27017"))
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
