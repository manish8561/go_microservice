package userfarms

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// controller file with routes
// register api in this function
func ApisRegister(router *gin.RouterGroup) {
	router.GET("", UserFarmList)
	router.GET("/total", UserFarmTotal)
	router.GET("/:id", UserFarmRetrieve)
	router.POST("", UserFarmSave)
}

/*
function to total record counts
*/
func UserFarmTotal(c *gin.Context) {
	user := strings.ToLower(c.Query("user"))
	if user == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user found", "success": false})
		return
	}
	status := c.Query("status")
	source := c.Query("source")
	tokenType := c.Query("token_type")
	name := c.Query("name")
	chainId, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)

	if err != nil {
		chainId = 4 //rinkeby
	}
	// filtering
	filters := Filters{
		User:      user,
		Source:    source,
		TokenType: tokenType,
		Name:      name,
		ChainId:   chainId,
	}

	num := GetTotal(status, filters)

	c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{"count": num}})
}

/*
function to retrive record list using get api
*/
func UserFarmList(c *gin.Context) {
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
	// filtering query data
	user := strings.ToLower(c.Query("user"))
	if user == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user found", "success": false})
		return
	}

	status := c.Query("status")
	source := c.Query("source")
	tokenType := c.Query("token_type")
	name := c.Query("name")
	chainId, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)
	if err != nil {
		chainId = 4 //rinkeby
	}
	// filter struct instance
	filters := Filters{
		User:      user,
		Source:    source,
		TokenType: tokenType,
		Name:      name,
		ChainId:   chainId,
	}
	//sorting
	sortBy := c.Query("sort_by")

	records, err := GetAll(page, limit, status, filters, sortBy)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": records})
}

/*
function to retrive single record using get api
*/
func UserFarmRetrieve(c *gin.Context) {
	id := c.Param("id")
	record, err := GetRecord(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": record, "success": true})
}

/*
function to save record in db
*/
func UserFarmSave(c *gin.Context) {
	userFarmModelValidator := NewUserFarmsModelValidator()
	if err := userFarmModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(), "success": false})
		return
	}
	insertID, err := SaveOne(&(userFarmModelValidator.userFarmsModel))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Record Inserted", "insertId": insertID, "success": true})
}
