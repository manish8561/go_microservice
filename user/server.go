package main

import (
	// "fmt"

	"github.com/gin-gonic/gin"

	// "github.com/autocompound/docker_backend/user/articles"
	"github.com/autocompound/docker_backend/user/common"
	"github.com/autocompound/docker_backend/user/users"
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
	// defer conn.Session.Close()

	r := gin.Default()

	v1 := r.Group("/api")
	users.UsersRegister(v1.Group("/users"))
	v1.Use(users.AuthMiddleware(false))

	v1.Use(users.AuthMiddleware(true))
	users.ProfileRegister(v1.Group("/profile"))

	// articles.ArticlesAnonymousRegister(v1.Group("/articles"))
	// articles.TagsAnonymousRegister(v1.Group("/tags"))

	// users.UserRegister(v1.Group("/user"))
	

	// articles.ArticlesRegister(v1.Group("/articles"))

	testAuth := r.Group("/api/ping")

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
