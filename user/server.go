package main

import (
	// "fmt"
	// "time"
	"fmt"
	"net/http"

	// "github.com/gin-contrib/cors"
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		// c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		// c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			// c.AbortWithStatus(204)
			c.Status(http.StatusOK)
			return
		}
		c.Next()
	}
}

// Middlewares
func CORSMiddleware_old() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("manish")
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusOK)
			return
		}
		c.Next()
	}
}

func main() {
	// initalize variable from config
	common.InitVariables()

	common.InitDB()
	// defer conn.Session.Close()

	r := gin.Default()
	r.Use(CORSMiddleware())

	// r.Use(cors.Default())
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"*"}
	// // config.AllowAllOrigins = true
	// config.AllowCredentials = true
	// config.AddAllowHeaders("authorization")
	// r.Use(cors.New(config))

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	// AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
	// 	AllowHeaders:     []string{"*"},
	// 	// ExposeHeaders:    []string{"Content-Length"},
	// 	// AllowCredentials: true,
	// 	// AllowOriginFunc: func(origin string) bool {
	// 	// 	return origin == "https://github.com"
	// 	// },
	// 	// MaxAge: 12 * time.Hour,
	// }))

	// Register the middleware
	v1 := r.Group("/api/user_service")
	users.UsersRegister(v1.Group("/users"))
	v1.Use(users.AuthMiddleware(false))
	
	v1.Use(users.AuthMiddleware(true))
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
