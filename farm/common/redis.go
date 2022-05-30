package common

import (
	"fmt"
	"os"
	"strconv"

	// "time"

	"github.com/go-redis/redis"
)

var clientRedis *redis.Client

func init() {
	//initialize the db
	InitRedisDB()
}

// Opening a database and save the reference to `Database` struct.
func InitRedisDB() {
	db, err := strconv.ParseInt(os.Getenv("REDIS_DB"), 10, 32)
	if err != nil {
		fmt.Println("error in redis", )
	}
	clientRedis = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       int(db),
	})

	pong, err := clientRedis.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Redis connected successfully. Ping:", pong)
}

// Using this function to get a connection, you can create your connection pool here.
func GetRedisDB() *redis.Client {
	return clientRedis
}
