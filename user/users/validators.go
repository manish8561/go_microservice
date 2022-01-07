package users

import (
	"github.com/autocompound/docker_backend/user/common"
	"github.com/gin-gonic/gin"
)

// *ModelValidator containing two parts:
// - Validator: write the form/json checking rule according to the doc https://github.com/go-playground/validator
// - DataModel: fill with data from Validator after invoking common.Bind(c, self)
// Then, you can just call model.save() after the data is ready in DataModel.
type UserModelValidator struct {
	Username  string    `form:"username" json:"username" binding:"exists,alphanum,min=4,max=255"`
	Email     string    `form:"email" json:"email" binding:"exists,email"`
	Password  string    `form:"password" json:"password" binding:"exists,min=8,max=255"`
	Bio       string    `form:"bio" json:"bio" binding:"max=1024"`
	Image     string    `form:"image" json:"image" binding:"omitempty,url"`
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
	self.userModel.Username = self.Username
	self.userModel.Email = self.Email
	self.userModel.Bio = self.Bio

	if self.Password != common.NBRandomPassword {
		self.userModel.setPassword(self.Password)
	}
	if self.Image != "" {
		self.userModel.Image = &self.Image
	}
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
	userModelValidator.Username = userModel.Username
	userModelValidator.Email = userModel.Email
	userModelValidator.Bio = userModel.Bio
	userModelValidator.Password = common.NBRandomPassword

	if userModel.Image != nil {
		userModelValidator.Image = *userModel.Image
	}
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
