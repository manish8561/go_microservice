package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	pb "github.com/autocompound/docker_backend/user/helloworld"
	// "google.golang.org/grpc"

	// "github.com/autocompound/docker_backend/user/articles"
	"github.com/autocompound/docker_backend/user/common"
	"github.com/autocompound/docker_backend/user/users"
)

// cors common function for * n
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Header("Content-Type", "application/json")
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
// server is used to implement helloworld.GreeterServer.(grpc)
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer(grpc)
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

// main func server
func main() {
	// initalize variable from config
	common.InitVariables()
	// call db init function
	common.InitDB()
	// defer conn.Session.Close()

	// endpoint, ok := os.LookupEnv("USER_GRPC_SERVER_PORT")
	// if(!ok){
	// 	endpoint = ":3001"
	// }

	// grpc server as user
	// grpc start
	// lis, err := net.Listen("tcp", endpoint)
	// if err != nil {
	// 	log.Println("ERROR:", err.Error())
	// }
	
	// s := grpc.NewServer()
	// pb.RegisterGreeterServer(s, &server{})
	// log.Printf("server listening at %v", lis.Addr())
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
	// grpc end

	//create server
	r := gin.Default()
	r.Use(CORSMiddleware())

	// Register the middleware
	v1 := r.Group("/api/user_service")
	users.UsersRegister(v1.Group("/users"))
	v1.Use(common.AuthMiddleware(false))

	v1.Use(common.AuthMiddleware(true))
	users.ProfileRegister(v1.Group("/profile"))

	// articles.ArticlesAnonymousRegister(v1.Group("/articles"))
	// articles.TagsAnonymousRegister(v1.Group("/tags"))

	// users.UserRegister(v1.Group("/user"))

	// articles.ArticlesRegister(v1.Group("/articles"))

	testAuth := r.Group("/api/user_service/ping")

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
	r.Run() // listen and serve on 0.0.0.0:8080

	
}
