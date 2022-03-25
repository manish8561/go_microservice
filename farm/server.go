package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/autocompound/docker_backend/farm/common"
	"github.com/autocompound/docker_backend/farm/contracts"
	"github.com/autocompound/docker_backend/farm/farms"
	"github.com/autocompound/docker_backend/farm/pricefeeds"
	
	// pb "github.com/autocompound/docker_backend/farm/helloworld"
)

// cors common function for * n
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		// c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization, accept, origin, Cache-Control, X-Requested-With")

		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			// c.AbortWithStatus(204)
			c.Status(http.StatusOK)
			return
		}
		c.Next()
	}
}

// init function in every file
func init() {
	// initalize variable from config
	common.InitVariables()

	//init db function
	common.InitDB()
}

// main function
func main() {
	// Set up a connection to the grpc client for user .
	// grpc start
	// endpoint, ok := os.LookupEnv("USER_GRPC_SERVER_PORT")
	// if(!ok){
	// 	log.Fatalf("end point not found to connect")
	// }
	// conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }

	// defer conn.Close()
	// c := pb.NewGreeterClient(conn)
	// fmt.Printf("grpc", c)
	// // Contact the server and print out its response.
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// defer cancel()
	// rr, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	// if err != nil {
	// 	log.Fatalf("could not greet: %v", err)
	// }
	// log.Printf("Greeting: %s", rr.GetMessage())
	// grpc end


	//create server
	r := gin.Default()
	r.Use(CORSMiddleware())

	v1 := r.Group("/api/farm_service")
	// farms.UsersRegister(v1.Group("/users"))
	// v1.Use(farms.AuthMiddleware(false))

	v1.Use(common.AuthMiddleware(false))
	farms.FarmsRegister(v1.Group("/farm"))
	pricefeeds.PriceFeedsRegister(v1.Group("/pricefeeds"))
	contracts.ContractsRegister(v1.Group("/contract"))

	testAuth := r.Group("/api/farm_service/ping")

	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// test 1 to 1
	// tx1 := db.Begin()
	// userA := users.UserModel{
	// 	Username: "AAAAAAAAAAAAAAAA",
	// 	Email:    "aaaa@g.cn",
	// 	Bio:      "hehddeda",
	// 	Image:    nil,
	// }
	// tx1.Save(&userA)
	// tx1.Commit()
	// fmt.Println(userA)

	//db.Save(&ArticleUserModel{
	//    UserModelID:userA.ID,
	//})
	//var userAA ArticleUserModel
	//db.Where(&ArticleUserModel{
	//    UserModelID:userA.ID,
	//}).First(&userAA)
	//fmt.Println(userAA)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.StaticFS("/api/farm_service/file", http.Dir("public"))
	// r.StaticFS("/api/farm_service/file", http.Dir("/data/dvolumes/autocompound/public"))

	r.Run() // listen and serve on 0.0.0.0:8080
}
