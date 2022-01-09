package users

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/autocompound/docker_backend/user/common"
	"github.com/gin-gonic/gin"
)

func UsersRegister(router *gin.RouterGroup) {
	router.POST("/", UsersRegistration)
	router.POST("/login", UsersLogin)
}

// func UserRegister(router *gin.RouterGroup) {
// 	router.GET("/", UserRetrieve)
// 	router.PUT("/", UserUpdate)
// }

func ProfileRegister(router *gin.RouterGroup) {
	router.GET("/", ProfileRetrieve)
	// router.POST("/:username/follow", ProfileFollow)
	// router.DELETE("/:username/follow", ProfileUnfollow)
}

func ProfileRetrieve(c *gin.Context) {
	my_user_id, _ := c.Get("my_user_id")
	userModel, err := GetProfile(my_user_id.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("profile", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{"profile": userModel})
}

// func ProfileFollow(c *gin.Context) {
// 	username := c.Param("username")
// 	userModel, err := FindOneUser(&UserModel{Username: username})
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("Invalid username")))
// 		return
// 	}
// 	myUserModel := c.MustGet("my_user_model").(UserModel)
// 	err = myUserModel.following(userModel)
// 	if err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
// 		return
// 	}
// 	serializer := ProfileSerializer{c, userModel}
// 	c.JSON(http.StatusOK, gin.H{"profile": serializer.Response()})
// }

// func ProfileUnfollow(c *gin.Context) {
// 	username := c.Param("username")
// 	userModel, err := FindOneUser(&UserModel{Username: username})
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, common.NewError("profile", errors.New("Invalid username")))
// 		return
// 	}
// 	myUserModel := c.MustGet("my_user_model").(UserModel)

// 	err = myUserModel.unFollowing(userModel)
// 	if err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
// 		return
// 	}
// 	serializer := ProfileSerializer{c, userModel}
// 	c.JSON(http.StatusOK, gin.H{"profile": serializer.Response()})
// }

func UsersRegistration(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	if err := SaveOne(&(userModelValidator.userModel)); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	// c.Set("my_user_model", userModelValidator.userModel)
	// serializer := UserSerializer{c}
	c.JSON(http.StatusCreated, gin.H{"user": "success"})
}

// user login function for jwt token
func UsersLogin(c *gin.Context) {
	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	userModel, err := FindOneUser(loginValidator.Email)

	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}

	if userModel.checkPassword(loginValidator.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered email or invalid password")))
		return
	}
	fmt.Println(userModel)
	c.JSON(http.StatusOK, gin.H{"token":common.GenToken(userModel.ID.Hex())})
}

// func UserRetrieve(c *gin.Context) {
// 	serializer := UserSerializer{c}
// 	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
// }

// func UserUpdate(c *gin.Context) {
// 	myUserModel := c.MustGet("my_user_model").(UserModel)
// 	userModelValidator := NewUserModelValidatorFillWith(myUserModel)
// 	if err := userModelValidator.Bind(c); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
// 		return
// 	}

// 	userModelValidator.userModel.ID = myUserModel.ID
// 	if err := myUserModel.Update(userModelValidator.userModel); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
// 		return
// 	}
// 	UpdateContextUserModel(c, myUserModel.ID)
// 	serializer := UserSerializer{c}
// 	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
// }
