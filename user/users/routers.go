// Package users provides HTTP route handlers and registration functions for user-related operations
// such as registration, login, token refresh, profile retrieval, and password change.
//
// Functions:
//
//   - UsersRegister: Registers user-related routes (registration, login, refresh token, change password) to the provided Gin router group.
//   - ProfileRegister: Registers the user profile retrieval route to the provided Gin router group.
//   - ProfileRetrieve: Retrieves the authenticated user's profile from the request context and returns it as JSON.
//   - UsersRegistration: Handles user registration by validating input, saving the user, and returning a success response.
//   - UsersLogin: Handles user login, validates credentials, and returns JWT and refresh tokens upon success.
//   - UsersRefreshToken: Handles refresh token requests, validates the refresh token, and issues a new access token.
//   - ChangePassword: Allows authenticated users to change their password after validating the old password and ensuring the new password is different.
//
// Note: Some functions related to user profile following/unfollowing and user update/retrieve are commented out and not currently active.
package users

import (
	"errors"
	"net/http"
	"time"

	"github.com/autocompound/docker_backend/user/common"
	"github.com/gin-gonic/gin"
)

// controller file with routes
// register api(s) in this function
func UsersRegister(router *gin.RouterGroup) {
	router.POST("", UsersRegistration)
	router.POST("/login", UsersLogin)
	router.POST("/refresh", UsersRefreshToken)

	router.Use(common.AuthMiddleware(true))
	router.POST("/changePassword", ChangePassword)
}

//	func UserRegister(router *gin.RouterGroup) {
//		router.GET("/", UserRetrieve)
//		router.PUT("/", UserUpdate)
//	}
//
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
		c.JSON(http.StatusForbidden, common.NewError("message", errors.New("invalid  email or password")))
		return
	}

	if userModel.checkPassword(loginValidator.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("message", errors.New("invalid email or password")))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":         common.GenToken(userModel.ID.Hex(), userModel.Role),
		"refresh_token": UpdateRefreshToken(userModel.ID.Hex()),
	})
}
//
func UsersRefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	// Find user by refresh token, check expiry
	user, err := FindUserByRefreshToken(req.RefreshToken)
	if err != nil || user.RefreshTokenExpiry.Before(time.Now()) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}
	// Generate new access token
	accessToken := common.GenToken(user.ID.Hex(), user.Role)
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
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

// user change password
func ChangePassword(c *gin.Context) {
	changePasswordValidator := NewChangePasswordValidator()
	if err := changePasswordValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "error": err.Error()})
		return
	}
	// get user object from request
	user, _ := c.Get("user")
	// convert to common userModel
	u := (user).(*common.UserModel)
	//get user from db
	userModel, err := FindOneUser(u.Email)

	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("message", errors.New("invalid  email or password")))
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
