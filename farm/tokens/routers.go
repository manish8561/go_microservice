package tokens

import (
	"net/http"
	"strconv"

	"github.com/autocompound/docker_backend/farm/common"
	"github.com/gin-gonic/gin"
)

// controller file with routes
// register api in this function
func ApisRegister(router *gin.RouterGroup) {
	router.GET("/getlastseven", GetGraphData)
	router.GET("/getsevendays", GetSevenDays)

	//Authorize Routes
	router.Use(common.AuthMiddleware(true))
	// router.GET("/total", Total)
	// router.GET("/list", List)
	// router.POST("", Add)
	// router.DELETE("/:id", DeleteRecord)
}

// function to get dashboard record
// GetGraphData
func GetGraphData(c *gin.Context) {
	//convert string to number
	
	// filtering
	chain_id, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)
	if err != nil {
		chain_id = 4 //rinkeby
	}
	address := c.Query("address")
	filters := Filters{
		ChainId: chain_id,
		Address: address,
	}
	records, err := GetLastSevenTransaction(filters)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": records})
}

// function to get dashboard data
// GetSevenDays
func GetSevenDays(c *gin.Context) {
	//convert string to number
	
	// filtering
	chain_id, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)
	if err != nil {
		chain_id = 4 //rinkeby
	}
	address := c.Query("address")
	filters := Filters{
		ChainId: chain_id,
		Address: address,
	}
	records, err := GetLastSevenDaysData(filters)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": records})
}