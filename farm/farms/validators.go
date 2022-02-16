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
	ID               string `form:"_id" json:"_id"`
	PID              int    `form:"pid" json:"pid" binding:"required"`
	Name             string `form:"name" json:"name" binding:"required,max=255"`
	Token_Type       string `form:"token_type" json:"token_type" binding:"required,max=10"`
	Deposit_Token    string `from:"deposit_token json:"deposit_token"  binding:"required,alphanum,max=255"`
	Masterchef       string `form:"masterchef" json:"masterchef" binding:"required,alphanum,max=255"`
	Router           string `form:"router" json:"router" binding:"required,alphanum,max=255"`
	Reward           string `form:"reward" json:"reward" binding:"required,alphanum,max=255"`
	Stake            string `form:"stake" json:"stake" binding:"required,alphanum,max=255"`
	Token0           Token  `form:"token0" json:"token0" binding:"required"`
	Token1           Token  `form:"token1" json:"token1" `
	Token_Per_Block  int    `form:"token_per_block" json:"token_per_block" binding:"required"`
	Bonus_Multiplier int    `form:"bonus_multiplier" json:"bonus_multiplier" binding:"required"`
	Source           string `form:"source" json:"source" binding:"required,max=255"`
	Source_Link      string `form:"source_link" json:"source_link" binding:"required,max=255"`

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
	self.farmModel.Name = self.Name
	self.farmModel.Token_Type = self.Token_Type
	self.farmModel.Deposit_Token = self.Deposit_Token
	self.farmModel.Masterchef = self.Masterchef
	self.farmModel.Router = self.Router
	self.farmModel.Stake = self.Stake
	self.farmModel.Reward = self.Reward
	self.farmModel.Token_Per_Block = self.Token_Per_Block
	self.farmModel.Bonus_Multiplier = self.Bonus_Multiplier
	self.farmModel.Token0 = self.Token0
	self.farmModel.Token1 = self.Token1
	self.farmModel.Status = "pending"
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
