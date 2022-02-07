package farms

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/autocompound/docker_backend/farm/common"
	"github.com/gin-gonic/gin"
)

func FarmsRegister(router *gin.RouterGroup) {
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.GET("", FarmList)
	router.GET("/:id", FarmRetrieve)
	router.POST("/upload", FileUpload)
	router.POST("", FarmSave)
	router.PUT("", FarmUpdate)
	// router.DELETE("/:username/follow", FarmUnfollow)
}

/*
function to retrive single farm using get api
*/
func FarmList(c *gin.Context) {
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
	if limit <= 10 {
		limit = 10
	}
	farmModel, err := GetAll(page, limit)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("message", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": farmModel})
}

/*
function to retrive single farm using get api
*/
func FarmRetrieve(c *gin.Context) {
	id := c.Param("id")
	farmModel, err := GetFarm(id)
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("message", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": farmModel})
}

/*
function to save farm in db
*/
func FarmSave(c *gin.Context) {
	farmModelValidator := NewFarmModelValidator()
	if err := farmModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("message", err))
		return
	}
	if err := SaveOne(&(farmModelValidator.farmModel)); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "farm inserted"})
}

/*
function to update single farm using put api
*/
func FarmUpdate(c *gin.Context) {
	farmModelValidator := NewFarmModelValidator()
	if err := farmModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("message", err))
		return
	}
	if err := UpdateOne(&(farmModelValidator.farmModel)); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "updated farm successfully"})
}

/*
function to upload file in the farm
*/
func FileUpload(c *gin.Context) {
	// single file
	file, handler, err := c.Request.FormFile("file")
	if file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no file uploaded"})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("file err : %s", err.Error())})
		return
	}
	// convert current timestamp into string
	t := time.Now().Unix()
	n := strconv.FormatInt(t, 10)

	filename := n+"_"+handler.Filename
	
	// file size handler
	if handler.Size > (4 * 1024 * 1024) {// 4MB
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size is greater than 4MB"})
		return
	}

	// file type handler check the type of the image upload
	fileType := handler.Header.Get("Content-Type")
	if !strings.Contains(fileType, "image") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is not an image"})
		return
	}

	// create directory
	out, err := os.Create("public/" + filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("file err : %s", err.Error())})
		return
	}
	defer out.Close()

	// copy file
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("file err : %s", err.Error())})
		return
	}
	filepath := os.Getenv("UPLOAD_URL") + filename

	c.JSON(http.StatusCreated, gin.H{"message": "file uploaded", "filepath": filepath})
}
