package farms

import (
	"time"

	"github.com/autocompound/docker_backend/farm/common"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// *ModelValidator containing two parts:
// - Validator: write the form/json checking rule according to the doc https://github.com/go-playground/validator
// - DataModel: fill with data from Validator after invoking common.Bind(c, self)
// Then, you can just call model.save() after the data is ready in DataModel.
type FarmModelValidator struct {
	ID         string `form:"_id" json:"_id" `
	PID        int    `form:"pid" json:"pid" binding:"required"`
	Token      string `form:"token" json:"token" binding:"required,alphanum,max=255"`
	TokenType  string `form:"tokenType" json:"TokenType" binding:"required"`
	Vault      string `form:"vault" json:"vault" binding:"required,alphanum,max=255"`
	Masterchef string `form:"masterchef" json:"masterchef" binding:"required,alphanum,max=255"`

	// Image     string    `form:"image" json:"image" binding:"omitempty,url"`
	farmModel FarmModel `json:"-"`
}

// There are some difference when you create or update a model, you need to fill the DataModel before
// update so that you can use your origin data to cheat the validator.
// BTW, you can put your general binding logic here such as setting password.
func (self *FarmModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	self.farmModel.PID = self.PID
	self.farmModel.Token = self.Token
	self.farmModel.TokenType = self.TokenType
	self.farmModel.Vault = self.Vault
	self.farmModel.Masterchef = self.Masterchef
	self.farmModel.Status = "active"
	self.farmModel.Created = time.Now()
	self.farmModel.Modified = time.Now()

	// using _id to update row in db
	if self.ID != "" {
		objID, err := primitive.ObjectIDFromHex(self.ID)
		if err != nil {
			return err
		}
		self.farmModel.ID = objID
	}
	return nil
}

// You can put the default value of a Validator here
func NewFarmModelValidator() FarmModelValidator {
	farmModelValidator := FarmModelValidator{}
	return farmModelValidator
}
