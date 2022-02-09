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
	ID              string `form:"_id" json:"_id"`
	StrategyABI     string `form: "strategyABI" json:"strategyABI" binding:"required"`
	PID             int    `form:"pid" json:"pid" binding:"required"`
	Token           string `form:"token" json:"token" binding:"required,alphanum,max=255"`
	TokenType       string `form:"tokenType" json:"TokenType" binding:"required,max=50"`
	Strategy        string `form:"strategy" json:"strategy" binding:"max=255"`
	Masterchef      string `form:"masterchef" json:"masterchef" binding:"required,alphanum,max=255"`
	Router          string `form:"router" json:"router" binding:"required,alphanum,max=255"`
	Reward          string `form:"reward" json:"reward" binding:"required,alphanum,max=255"`
	Stake           string `form:"stake" json:"stake" binding:"required,alphanum,max=255"`
	Token0Img       string `form:"token0Img" json:"token0Img" binding:"max=255"`
	Token1Img       string `form:"token1Img" json:"token1Img" binding:"max=255"`
	TokenPerBlock   int    `form:"tokenPerBlock" json:"tokenPerBlock" binding:"required"`
	StakePercentage int    `form:"stakePercentage" json:"stakePercentage" binding:"required"`
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
	self.farmModel.StrategyABI = self.StrategyABI
	self.farmModel.PID = self.PID
	self.farmModel.Token = self.Token
	self.farmModel.TokenType = self.TokenType
	self.farmModel.Strategy = self.Strategy
	self.farmModel.Masterchef = self.Masterchef
	self.farmModel.Router = self.Router
	self.farmModel.Stake = self.Stake
	self.farmModel.Reward = self.Reward
	self.farmModel.TokenPerBlock = self.TokenPerBlock
	self.farmModel.StakePercentage = self.StakePercentage
	self.farmModel.Token0Img = self.Token0Img
	self.farmModel.Token1Img = self.Token1Img
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
