package farms

import (
	"net/http"

	"github.com/autocompound/docker_backend/farm/common"
	"github.com/gin-gonic/gin"
)

func FarmsRegister(router *gin.RouterGroup) {
	router.GET("/", FarmRetrieve)
	// router.POST("/:username/follow", FarmFollow)
	// router.DELETE("/:username/follow", FarmUnfollow)
}

func FarmRetrieve(c *gin.Context) {
	my_user_id, _ := c.Get("my_user_id")
	userModel, err := GetFarm(my_user_id.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"profile": userModel})
}



