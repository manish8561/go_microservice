package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/autocompound/docker_backend/governance/common"
	"github.com/autocompound/docker_backend/governance/proposal"

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
// func init() {}

// main function
func main() {
	//calling grpc common server
	// common.Call_GRPC_Server()

	//create server
	r := gin.Default()
	r.Use(CORSMiddleware())

	v1 := r.Group("/api/governance_service")

	// v1.Use(proposals.AuthMiddleware(false))

	v1.Use(common.AuthMiddleware(false))
	proposal.ProposalsRegister(v1.Group("/proposal"))

	testAuth := r.Group("/api/governance_service/ping")

	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})



	r.Run() // listen and serve on 0.0.0.0:8080
}
