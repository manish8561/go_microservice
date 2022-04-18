package stakes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/autocompound/docker_backend/farm/common"
	"github.com/gin-gonic/gin"
)

// controller file with routes
// register api in this function
func StakesRegister(router *gin.RouterGroup) {

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.GET("", StakeList)
	router.GET("/total", StakeTotal)
	router.GET("/:id", StakeRetrieve)
	router.GET("/chainId/:chain_id", StakeFromChainId)

	// enable authentication for below routes
	router.Use(common.AuthMiddleware(true))
	router.POST("", StakeSave)
	router.PUT("", StakeUpdate)
	router.DELETE("/:id", StakeDelete)
}

/*
function to total record counts
*/
func StakeTotal(c *gin.Context) {
	status := c.Query("status")
	chain_id, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)
	if err != nil {
		chain_id = 4 //rinkeby
	}
	// filtering
	filters := Filters{
		Chain_Id: chain_id,
	}

	num := GetTotal(status, filters)

	c.JSON(http.StatusOK, gin.H{"success": true, "count": num})
}

/*
function to retrive record list using get api
*/
func StakeList(c *gin.Context) {
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
	// filtering
	status := c.Query("status")
	chain_id, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)
	if err != nil {
		chain_id = 4 //rinkeby
	}
	filters := Filters{
		Chain_Id: chain_id,
	}
	//sorting
	sort_by := c.Query("sort_by")

	records, err := GetAll(page, limit, status, filters, sort_by)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": records})
}

/*
function to retrive single record using get api
*/
func StakeRetrieve(c *gin.Context) {
	id := c.Param("id")
	record, err := GetRecord(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": record, "success": true})
}
/*
function to retrive single record using get api
*/
func StakeFromChainId(c *gin.Context) {
	chain_id, err := strconv.ParseInt(c.Param("chain_id"), 10, 64)
	fmt.Printf("chainid", chain_id)
	if err != nil {
		chain_id = 4 //rinkeby
	}
	record, err := GetStakeFromChainId(chain_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": record, "success": true})
}
/*
function to save record in db
*/
func StakeSave(c *gin.Context) {
	stakeModelValidator := NewStakeModelValidator()
	if err := stakeModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(), "success": false})
		return
	}
	insertID, err := SaveOne(&(stakeModelValidator.stakeModel))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Record Inserted", "insertId": insertID, "success": true})
}

/*
function to update record in db
*/
func StakeUpdate(c *gin.Context) {
	record := &StakeModel{}
	if err := common.Bind(c, record); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(), "success": false})
		return
	}
	data, err := UpdateOne(record)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Record Updated", "data": data, "success": true})
}

/*
function to delete single record using delete api
*/
func StakeDelete(c *gin.Context) {
	id := c.Param("id")
	record, err := DeleteRecord(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "success": false})
		return
	}
	if record {
		c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully", "success": true})
		return
	}
}
