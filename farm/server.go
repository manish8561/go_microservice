package main

import (
	// "fmt"

	"net/http"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "github.com/autocompound/docker_backend/farm/articles"
	"github.com/autocompound/docker_backend/farm/common"
	"github.com/autocompound/docker_backend/farm/contracts"
	"github.com/autocompound/docker_backend/farm/farms"
	// "github.com/go-bongo/bongo"
)

// func Migrate(db *bongo.DB) {
// users.AutoMigrate()
// db.AutoMigrate(&articles.ArticleModel{})
// db.AutoMigrate(&articles.TagModel{})
// db.AutoMigrate(&articles.FavoriteModel{})
// db.AutoMigrate(&articles.ArticleUserModel{})
// db.AutoMigrate(&articles.CommentModel{})
// }

func main() {
	// initalize variable from config
	common.InitVariables()

	common.InitDB()

	r := gin.Default()
	r.Use(cors.Default())

	v1 := r.Group("/api/farm_service")
	// farms.UsersRegister(v1.Group("/users"))
	// v1.Use(farms.AuthMiddleware(false))

	v1.Use(farms.AuthMiddleware(false))
	farms.FarmsRegister(v1.Group("/farm"))
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
	r.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.StaticFS("/file", http.Dir("public"))


	r.Run() // listen and serve on 0.0.0.0:8080
}
