package stakes

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
type StakeModelValidator struct {
	ID          string `form:"_id" json:"_id"`
	Address     string `form:"address" json:"address" binding:"required"`
	Chain_Id    int    `form:"chain_id" json:"chain_id" binding:"required"`
	BlockNumber int64   `form:"blockNumber" json:"blockNumber" binding:"required"`

	// Image     string    `form:"image" json:"image" binding:"omitempty,url"`
	stakeModel StakeModel `json:"-"`
}

// There are some difference when you create or update a model, you need to fill the DataModel before
// update so that you can use your origin data to cheat the validator.
// BTW, you can put your general binding logic here such as setting password.
func (self *StakeModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	self.stakeModel.Address = strings.ToLower(self.Address)
	self.stakeModel.Chain_Id = self.Chain_Id
	self.stakeModel.BlockNumber = self.BlockNumber
	self.stakeModel.LastBlockNumber = self.BlockNumber

	self.stakeModel.Status = "active"
	self.stakeModel.Created = time.Now()
	self.stakeModel.Modified = time.Now()

	// using _id to update row in db
	if self.ID != "" {
		objID, err := primitive.ObjectIDFromHex(self.ID)
		if err != nil {
			return err
		}
		self.stakeModel.ID = objID
	}
	return nil
}

// You can put the default value of a Validator here
func NewStakeModelValidator() StakeModelValidator {
	stakeModelValidator := StakeModelValidator{}
	return stakeModelValidator
}
