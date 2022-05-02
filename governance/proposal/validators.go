package proposal

import (
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
	Chain_Id         int    `form:"chain_id" json:"chain_id" binding:"required"`
	Transaction_Hash string `form:"transaction_hash" 
	json:"transaction_hash" binding:"required"`
	Proposer       string        `form:"proposer" json:"proposer" binding:"required"`
	Voting_Period  int           `form:"voting_period" json:"voting_period binding:"required"` // in days
	Title          string        `form:"title" json:"title" binding:"required"`
	Description    string        `form:"description" json:"description" binding:"required"`
	Db_Description string        `form:"db_description" json:"db_description" binding:"required"`
	Proposal_Type  int        `form:"proposal_type" json:"proposal_type" binding:"required"`
	proposalModel  ProposalModel `json:"-"`
}

// There are some difference when you create or update a model, you need to fill the DataModel before
// update so that you can use your origin data to cheat the validator.
// BTW, you can put your general binding logic here such as setting password.
func (self *ProposalModelValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, self)
	if err != nil {
		return err
	}
	self.proposalModel.Chain_Id = self.Chain_Id
	self.proposalModel.Title = self.Title
	self.proposalModel.Transaction_Hash = self.Transaction_Hash
	self.proposalModel.Description = self.Description //ipfs hash
	self.proposalModel.Db_Description = self.Db_Description
	self.proposalModel.Proposer = self.Proposer
	self.proposalModel.Voting_Period = self.Voting_Period

	self.proposalModel.Proposal_Type = self.Proposal_Type
	self.proposalModel.Status = "pending"
	self.proposalModel.Cron_Status = "pending"
	self.proposalModel.Created = time.Now()
	self.proposalModel.Modified = time.Now()
	self.proposalModel.For_Votes = 0
	self.proposalModel.Against_Votes = 0
	self.proposalModel.Canceled = false
	self.proposalModel.Executed = false

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
