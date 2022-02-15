package pricefeeds

import (
	"net/http"

	// "github.com/autocompound/docker_backend/farm/common"
	"github.com/gin-gonic/gin"
)

// controller file with routes
// register api in this function
func FarmsRegister(router *gin.RouterGroup) {

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.GET("/", GetSymbolPrice)
	router.GET("/update", UpdateSymbolPrice)
	// router.GET("/total", FarmTotal)
	// router.Use(common.AuthMiddleware(true))

	// router.GET("/:id", FarmRetrieve)
	// router.POST("", FarmSave)
	// router.PUT("", FarmUpdate)
	// router.DELETE("/:username/follow", FarmUnfollow)
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
