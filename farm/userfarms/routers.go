package userfarms

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/autocompound/docker_backend/farm/common"
	"github.com/gin-gonic/gin"
)

// controller file with routes
// register api in this function
func UserFarmsRegister(router *gin.RouterGroup) {

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.GET("", UserFarmList)
	router.GET("/total", UserFarmTotal)
	router.GET("/:id", UserFarmRetrieve)
	router.POST("", UserFarmSave)

	// enable authentication for below routes
	// router.Use(common.AuthMiddleware(true))
	router.PUT("", UserFarmUpdate)
	router.DELETE("/:id", UserFarmDelete)
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
	token_type := c.Query("token_type")
	name := c.Query("name")
	chain_id, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)

	if err != nil {
		chain_id = 4 //rinkeby
	}
	// filtering
	filters := Filters{
		User:       user,
		Source:     source,
		Token_Type: token_type,
		Name:       name,
		Chain_Id:   chain_id,
	}

	num := GetTotal(status, filters)

	c.JSON(http.StatusOK, gin.H{"success": true, "count": num})
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
	// filtering
	user := strings.ToLower(c.Query("user"))
	if user == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user found", "success": false})
		return
	}

	status := c.Query("status")
	source := c.Query("source")
	token_type := c.Query("token_type")
	name := c.Query("name")
	chain_id, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)
	if err != nil {
		chain_id = 4 //rinkeby
	}
	filters := Filters{
		User:       user,
		Source:     source,
		Token_Type: token_type,
		Name:       name,
		Chain_Id:   chain_id,
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

/*
function to update record in db
*/
func UserFarmUpdate(c *gin.Context) {
	record := &UserFarmsModel{}
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
func UserFarmDelete(c *gin.Context) {
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
