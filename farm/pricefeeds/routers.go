package pricefeeds

import (
	"net/http"
	"strconv"

	"github.com/autocompound/docker_backend/farm/common"
	"github.com/gin-gonic/gin"
)

// controller file with routes
// register api in this function
func ApisRegister(router *gin.RouterGroup) {
	router.GET("", GetSymbolPrice)
	router.GET("/update", UpdateSymbolPrice)

	//Authorize Routes
	router.Use(common.AuthMiddleware(true))
	router.GET("/total", Total)
	router.GET("/list", List)
	router.POST("", Add)
	router.DELETE("/:id", DeletePrice)
}

/*
function to insert price feed
*/
func Add(c *gin.Context) {
	record := &PriceFeedModel{}
	if err := common.Bind(c, record); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(), "success": false})
		return
	}
	if err := SaveOne(record); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Insert record successfully", "success": true})
}

/*
function to update price in the collection
*/
func UpdateSymbolPrice(c *gin.Context) {
	symbol := c.Query("symbol")
	val, err := GetTokenPrice(symbol, true)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": val})
}

/*
function to retrive price in the collection
*/
func GetSymbolPrice(c *gin.Context) {
	symbol := c.Query("symbol")
	val, err := GetTokenPrice(symbol, false)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": val})
}

/*
function to total price feed counts
*/
func Total(c *gin.Context) {
	status := c.Query("status")

	num := GetTotal(status)

	c.JSON(http.StatusOK, gin.H{"success": true, "count": num})
}

/*
function to retrive price feed list using get api
*/
func List(c *gin.Context) {
	//convert string to number
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		page = 1
	}
	if page <= 0 {
		page = 1
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 10
	}
	if limit <= 0 {
		limit = 10
	}
	status := c.Query("status")

	records, err := GetAll(page, limit, status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": records})
}

/*
function to delete price feed
*/
func DeletePrice(c *gin.Context) {
	id := c.Param("id")

	ok, err := DeleteRecord(id)
	if ok {
		c.JSON(http.StatusOK, gin.H{"success": true})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "success": false})
}
