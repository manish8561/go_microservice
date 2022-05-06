package proposal

import (
	"net/http"
	"strconv"

	"github.com/autocompound/docker_backend/governance/common"
	"github.com/gin-gonic/gin"
)

// controller file with routes
// register api in this function
func ProposalsRegister(router *gin.RouterGroup) {

	router.GET("", ProposalList)
	router.GET("/total", ProposalTotal)
	router.GET("/:id", ProposalRetrieve)
	router.GET("/votecast/", VoteCast)

	router.POST("", ProposalSave)
	router.Use(common.AuthMiddleware(true))
	router.PUT("", ProposalUpdate)

	// router.DELETE("/:username/follow", ProposalUnfollow)
}

/*
function to total proposal counts
*/
func ProposalTotal(c *gin.Context) {
	status := c.Query("status")
	source := c.Query("source")
	token_type := c.Query("token_type")
	name := c.Query("name")
	chain_id, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)
	if err != nil {
		chain_id = 4 //rinkeby
	}
	// filtering
	filters := Filters{
		Source:     source,
		Token_Type: token_type,
		Name:       name,
		Chain_Id:   chain_id,
	}

	num := GetTotal(status, filters)

	c.JSON(http.StatusOK, gin.H{"success": true, "count": num})
}

/*
function to retrive proposal list using get api
*/
func ProposalList(c *gin.Context) {
	//convert string to number
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		page = 1
	}
	if page <= 0 {
		page = 1
	}
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	if err != nil {
		limit = 10
	}
	if limit <= 0 {
		limit = 10
	}
	// filtering
	status := c.Query("status")
	source := c.Query("source")
	token_type := c.Query("token_type")
	name := c.Query("name")
	chain_id, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)
	if err != nil {
		chain_id = 4 //rinkeby
	}
	filters := Filters{
		Source:     source,
		Token_Type: token_type,
		Name:       name,
		Chain_Id:   chain_id,
	}
	//sorting
	sort_by := c.Query("sort_by")
	if sort_by == "" {
		sort_by = "tvl"
	}

	records, err := GetAll(page, limit, status, filters, sort_by)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": records})
}

/*
function to retrive single proposal using get api
*/
func ProposalRetrieve(c *gin.Context) {
	id := c.Param("id")
	record, err := GetRecord(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": record, "success": true})
}

/*
function to save proposal in db
*/
func ProposalSave(c *gin.Context) {
	proposalModelValidator := NewProposalModelValidator()
	if err := proposalModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(), "success": false})
		return
	}
	insertID, err := SaveOne(&(proposalModelValidator.proposalModel))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "proposal inserted", "insertId": insertID, "success": true})
}

/*
function to update proposal in db
*/
func ProposalUpdate(c *gin.Context) {
	proposals := &ProposalModel{}
	if err := common.Bind(c, proposals); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(), "success": false})
		return
	}
	data, err := UpdateOne(proposals)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error(), "success": false})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "proposal updated", "data": data, "success": true})
}

/*
function to total vote cast for proposal
*/
func VoteCast(c *gin.Context) {
	support_query := c.Query("support")
	if support_query == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Support is required", "success": false})
		return
	}
	support := true
	if support_query == "true"  {
		support = true
	} else {
		support = false
	}
	proposalId, err := strconv.ParseInt(c.Query("proposalId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Proposal ID is required", "success": false})
		return
	}
	chain_id, err := strconv.ParseInt(c.Query("chain_id"), 10, 64)
	if err != nil {
		chain_id = 4 //rinkeby
	}
	// filtering
	filters := VoteCast_Filters{
		Support:    support,
		ProposalId: proposalId,
		ChainId:    chain_id,
	}

	num := GetVoteCastTotal(filters)

	c.JSON(http.StatusOK, gin.H{"success": true, "count": num})
}
