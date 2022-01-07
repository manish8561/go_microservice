package common

import (
	"fmt"
	"os"

	"github.com/go-bongo/bongo"
)

var conn *bongo.Connection

// Opening a database and save the reference to `Database` struct.
func Init() *bongo.Connection {
	config := &bongo.Config{
		ConnectionString: os.Getenv("MONGO_HOST"),
		Database:         os.Getenv("MONGO_DATABASE"),
	}
	connection, err := bongo.Connect(config)

	if err != nil {
		fmt.Println("db err: (Init) ", err)
	}
	conn = connection
	return connection
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
func GetDB() *bongo.Connection {
	return conn
}
