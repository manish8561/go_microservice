package farms

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
type FarmModelValidator struct {
	ID            string `form:"_id" json:"_id"`
	Address       string `form:"address" json:"address" binding:"required"`
	Chain_Id      int    `form:"chain_id" json:"chain_id" binding:"required"`
	PID           int    `form:"pid" json:"pid" binding:"required"`
	Name          string `form:"name" json:"name" binding:"required,max=255"`
	Token_Type    string `form:"token_type" json:"token_type" binding:"required,max=20"`
	Deposit_Token string `from:"deposit_token json:"deposit_token"  binding:"required,alphanum,max=255"`
	Masterchef    string `form:"masterchef" json:"masterchef" binding:"required,alphanum,max=255"`
	Router        string `form:"router" json:"router" binding:"required,alphanum,max=255"`
	Weth          string `form:"weth" json:"weth" binding:"required,alphanum,max=255"`
	Reward        string `form:"reward" json:"reward" binding:"required,alphanum,max=255"`
	RewardImage   string `form:"rewardImage" json:"rewardImage" binding:"required"`
	Stake         string `form:"stake" json:"stake" binding:"required,alphanum,max=255"`
	AC_Token      string `form:"ac_token" json:"ac_token" binding:"required,alphanum,max=255"`
	Token0        Token  `form:"token0" json:"token0" binding:"required"`
	Token1        Token  `form:"token1" json:"token1"`
	FarmType      string `form:"farmType" json:"farmType" binding:"required,max=255"`
	Source        string `form:"source" json:"source" binding:"required,max=255"`
	Source_Link   string `form:"source_link" json:"source_link" binding:"required,max=255"`

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
	self.farmModel.Address = strings.ToLower(self.Address)
	self.farmModel.Chain_Id = self.Chain_Id
	self.farmModel.PID = self.PID
	self.farmModel.Name = self.Name
	self.farmModel.Token_Type = self.Token_Type
	self.farmModel.Deposit_Token = strings.ToLower(self.Deposit_Token)
	self.farmModel.Masterchef = strings.ToLower(self.Masterchef)
	self.farmModel.Router = strings.ToLower(self.Router)
	self.farmModel.Weth = strings.ToLower(self.Weth)
	self.farmModel.Stake = strings.ToLower(self.Stake)
	self.farmModel.AC_Token = strings.ToLower(self.AC_Token)
	self.farmModel.Reward = strings.ToLower(self.Reward)
	self.farmModel.RewardImage = self.RewardImage
	self.farmModel.Token0 = self.Token0
	self.farmModel.Token1 = self.Token1
	self.farmModel.FarmType = self.FarmType
	self.farmModel.Source = self.Source
	self.farmModel.Source_Link = self.Source_Link

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
