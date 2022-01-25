package contracts

import (
	"net/http"

	"github.com/autocompound/docker_backend/farm/common"
	"github.com/gin-gonic/gin"
)

func ContractsRegister(router *gin.RouterGroup) {
	// router.GET("/", FarmList)
	router.GET("/", FarmRetrieve)
	// router.POST("/", FarmSave)
	// router.PUT("/", FarmUpdate)
	// router.DELETE("/:username/follow", FarmUnfollow)
}

/*
function to retrive single farm using get api
*/
func FarmRetrieve(c *gin.Context) {
	err := GetContract()
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("message", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "hi"})
}
