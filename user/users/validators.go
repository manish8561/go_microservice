package users

import (
	"time"

	"github.com/autocompound/docker_backend/user/common"
	"github.com/gin-gonic/gin"
)

// *ModelValidator containing two parts:
// - Validator: write the form/json checking rule according to the doc https://github.com/go-playground/validator
// - DataModel: fill with data from Validator after invoking common.Bind(c, self)
// Then, you can just call model.save() after the data is ready in DataModel.
type UserModelValidator struct {
	Firstname string `form:"firstname" json:"firstname" binding:"required,alphanum,min=4,max=255"`
	Lastname  string `form:"lastname" json:"lastname" binding:"required,alphanum,min=4,max=255"`
	Email     string `form:"email" json:"email" binding:"required,email"`
	Password  string `form:"password" json:"password" binding:"required,min=8,max=255"`

	// Image     string    `form:"image" json:"image" binding:"omitempty,url"`
	userModel UserModel `json:"-"`
}

// There are some difference when you create or update a model, you need to fill the DataModel before
// update so that you can use your origin data to cheat the validator.
// BTW, you can put your general binding logic here such as setting password.
func (self *UserModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	self.userModel.Firstname = self.Firstname
	self.userModel.Lastname = self.Lastname
	self.userModel.Email = self.Email
	self.userModel.Role = "user"
	self.userModel.Status = "active"
	self.userModel.Created = time.Now()
	self.userModel.Modified = time.Now()

	if self.Password != common.NBRandomPassword {
		self.userModel.setPassword(self.Password)
	}
	// if self.Image != "" {
	// 	self.userModel.Image = &self.Image
	// }
	return nil
}

// You can put the default value of a Validator here
func NewUserModelValidator() UserModelValidator {
	userModelValidator := UserModelValidator{}
	//userModelValidator.User.Email ="w@g.cn"
	return userModelValidator
}

func NewUserModelValidatorFillWith(userModel UserModel) UserModelValidator {
	userModelValidator := NewUserModelValidator()
	userModelValidator.Firstname = userModel.Firstname
	userModelValidator.Lastname = userModel.Lastname
	userModelValidator.Email = userModel.Email
	userModelValidator.Password = common.NBRandomPassword

	// if userModel.Image != nil {
	// 	userModelValidator.Image = *userModel.Image
	// }
	return userModelValidator
}

type LoginValidator struct {
	Email     string    `form:"email" json:"email" binding:"required,email"`
	Password  string    `form:"password" json:"password" binding:"required,min=8,max=255"`
	userModel UserModel `json:"-"`
}

func (self *LoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}

	self.userModel.Email = self.Email
	return nil
}

// You can put the default value of a Validator here
func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}

type ChangePasswordValidator struct {
	OldPassword string    `form:"oldPassword" json:"oldPassword" binding:"required,min=8,max=255"`
	Password    string    `form:"password" json:"password" binding:"required,min=8,max=255"`
	userModel   UserModel `json:"-"`
}

func (self *ChangePasswordValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}

	// self.userModel.PasswordHash = self.Password
	return nil
}

// You can put the default value of a Validator here
func NewChangePasswordValidator() ChangePasswordValidator {
	changePasswordValidator := ChangePasswordValidator{}
	return changePasswordValidator
}
