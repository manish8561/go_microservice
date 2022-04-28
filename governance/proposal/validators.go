package proposal

import (
	"strings"
	"time"

	"github.com/autocompound/docker_backend/governance/common"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// *ModelValidator containing two parts:
// - Validator: write the form/json checking rule according to the doc https://github.com/go-playground/validator
// - DataModel: fill with data from Validator after invoking common.Bind(c, self)
// Then, you can just call model.save() after the data is ready in DataModel.
type ProposalModelValidator struct {
	ID               string `form:"_id" json:"_id"`
	Address			string `form:"address" json:"address" binding:"required"`
	Chain_Id         int    `form:"chain_id" json:"chain_id" binding:"required"`

	// Image     string    `form:"image" json:"image" binding:"omitempty,url"`
	proposalModel ProposalModel `json:"-"`
}

// There are some difference when you create or update a model, you need to fill the DataModel before
// update so that you can use your origin data to cheat the validator.
// BTW, you can put your general binding logic here such as setting password.
func (self *ProposalModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	self.proposalModel.Address = strings.ToLower(self.Address)
	self.proposalModel.Chain_Id = self.Chain_Id
	
	self.proposalModel.Status = "active"
	self.proposalModel.Created = time.Now()
	self.proposalModel.Modified = time.Now()

	// using _id to update row in db
	if self.ID != "" {
		objID, err := primitive.ObjectIDFromHex(self.ID)
		if err != nil {
			return err
		}
		self.proposalModel.ID = objID
	}
	return nil
}

// You can put the default value of a Validator here
func NewProposalModelValidator() ProposalModelValidator {
	proposalModelValidator := ProposalModelValidator{}
	return proposalModelValidator
}
