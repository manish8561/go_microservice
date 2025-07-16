package userfarms

import (
	"strings"
	"time"

	"github.com/autocompound/docker_backend/farm/common"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// *ModelValidator containing two parts:
// - Validator: write the form/json checking rule according to the doc https://github.com/go-playground/validator
// - DataModel: fill with data from Validator after invoking common.Bind(c, self)
// Then, you can just call model.save() after the data is ready in DataModel.
type UserFarmsModelValidator struct {
	ID               string `form:"_id" json:"_id"`
	Chain_Id         int    `form:"chain_id" json:"chain_id" binding:"required"`
	Strategy         string `form:"strategy" json:"strategy" binding:"required,alphanum,max=255"`
	User             string `form:"user" json:"user" binding:"required,alphanum,max=255"`
	Transaction_Hash string `form:"transaction_hash" json:"transaction_hash" binding:"required,alphanum,max=255"`

	userFarmsModel UserFarmsModel `json:"-"`
}

// There are some difference when you create or update a model, you need to fill the DataModel before
// update so that you can use your origin data to cheat the validator.
// BTW, you can put your general binding logic here such as setting password.
func (v *UserFarmsModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}
	v.userFarmsModel.Strategy = strings.ToLower(v.Strategy)
	v.userFarmsModel.User = strings.ToLower(v.User)
	v.userFarmsModel.Transaction_Hash = v.Transaction_Hash
	v.userFarmsModel.Chain_Id = v.Chain_Id

	v.userFarmsModel.Status = "pending"
	v.userFarmsModel.Created = time.Now()
	v.userFarmsModel.Modified = time.Now()

	// using _id to update row in db
	if v.ID != "" {
		objID, err := primitive.ObjectIDFromHex(v.ID)
		if err != nil {
			return err
		}
		v.userFarmsModel.ID = objID
	}
	return nil
}

// You can put the default value of a Validator here
func NewUserFarmsModelValidator() UserFarmsModelValidator {
	validator := UserFarmsModelValidator{}
	return validator
}
