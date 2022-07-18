package transactions

import (
	"net/http"
	"strconv"
	"time"

	"github.com/autocompound/docker_backend/ledger/common"
	"github.com/gin-gonic/gin"
)

// controller file with routes
// register api in this function
func ApisRegister(router *gin.RouterGroup) {
	router.GET("/get", GetTransactionCall)
	router.GET("/profitloss", ProfitLossCall)

	//Authorize Routes
	router.Use(common.AuthMiddleware(true))
	// router.POST("", Add)
	// router.DELETE("/:id", DeleteRecord)
}

// function to get dashboard record
// GetTransactionCall
func GetTransactionCall(c *gin.Context) {
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
	chain_id, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)
	if err != nil {
		chain_id = common.DefaultChainId
	}
	currentTime := time.Now().Unix()
	startTime, err := strconv.ParseInt(c.Query("start_time"), 10, 64)
	if err != nil {
		startTime = currentTime - (24 * 60 * 60) //sub 1 day
	}
	endTime, err := strconv.ParseInt(c.Query("end_time"), 10, 64)
	if err != nil {
		endTime = currentTime
	}
	transactionType := c.Query("type")
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "User wallet address is required"})
		return
	}
	// filter instance
	filters := Filters{
		ChainId:   chain_id,
		Address:   address,
		Type:      transactionType,
		StartTime: startTime,
		EndTime:   endTime,
		Page:      page,
		Limit:     limit,
	}
	records, err := GetTransactions(filters)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": records})
}

// function to get dashboard record
// ProfitLossCall
func ProfitLossCall(c *gin.Context) {
	address := c.Query("address")
	if address == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Strategy address is required"})
		return
	}
	
	records := GetProfitLoss(address)
	c.JSON(http.StatusOK, gin.H{"success": true, "data": records})
}
