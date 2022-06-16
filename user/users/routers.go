package users

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/autocompound/docker_backend/user/common"
	"github.com/gin-gonic/gin"
)

// controller file with routes
// register api(s) in this function
func UsersRegister(router *gin.RouterGroup) {
	router.POST("", UsersRegistration)
	router.POST("/login", UsersLogin)

	router.Use(common.AuthMiddleware(true))
	router.POST("/changePassword", ChangePassword)

}

// func UserRegister(router *gin.RouterGroup) {
// 	router.GET("/", UserRetrieve)
// 	router.PUT("/", UserUpdate)
// }
// register the user profile route
func ProfileRegister(router *gin.RouterGroup) {
	router.GET("", ProfileRetrieve)
	// router.POST("/:username/follow", ProfileFollow)
	// router.DELETE("/:username/follow", ProfileUnfollow)
}

// get user profile from middleware
func ProfileRetrieve(c *gin.Context) {
	userModel, _ := c.Get("user")
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"success": false, "error": err.Error()})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{"profile": userModel, "success": true})
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
// 		c.JSON(http.StatusUnprocessableEntity, err.Error())
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
// 		c.JSON(http.StatusUnprocessableEntity, err.Error())
// 		return
// 	}
// 	serializer := ProfileSerializer{c, userModel}
// 	c.JSON(http.StatusOK, gin.H{"profile": serializer.Response()})
// }

func UsersRegistration(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "error": common.NewValidatorError(err)})
		return
	}
	if err := SaveOne(&(userModelValidator.userModel)); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "error": err.Error()})
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
		c.JSON(http.StatusForbidden, common.NewError("message", errors.New("Invalid  email or password")))
		return
	}

	if userModel.checkPassword(loginValidator.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("message", errors.New("Invalid email or password")))
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": common.GenToken(userModel.ID.Hex(), userModel.Role)})
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
// 		c.JSON(http.StatusUnprocessableEntity, err.Error())
// 		return
// 	}
// 	UpdateContextUserModel(c, myUserModel.ID)
// 	serializer := UserSerializer{c}
// 	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
// }

//user change password
func ChangePassword(c *gin.Context) {

	changePasswordValidator := NewChangePasswordValidator()
	if err := changePasswordValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "error": common.NewValidatorError(err)})
		return
	}
	// get user object from request
	user, _ := c.Get("user")
	// convert to common userModel
	u := (user).(*common.UserModel)

	fmt.Println("compare", u.Email)
	userModel, err := FindOneUser(u.Email)

	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("message", errors.New("Invalid  email or password")))
		return
	}
	//checking old password
	if userModel.checkPassword(changePasswordValidator.OldPassword) != nil {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "error": "Incorrect Old Password"})
		return
	}
	//checking old password is not equal to new one
	if userModel.checkPassword(changePasswordValidator.Password) == nil {
		c.JSON(http.StatusForbidden, gin.H{"success": false, "error": "Enter different new password."})
		return
	}
	userModel.setPassword(changePasswordValidator.Password)

	res, err := ChangePasswordOne(&userModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": true, "data": res})
}
