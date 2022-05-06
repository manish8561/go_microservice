package proposal

import (
	"context"
	"os"
	"time"

	"github.com/autocompound/docker_backend/governance/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CollectionName2 = "votecasts"

// Models should only be concerned with database schema, more strict checking should be put in validator.
//
// HINT: If you want to split null and "", you should use *string instead of string.
type VoteCastModel struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ChainId         int                `bson:"chainId" json:"chainId"`
	TransactionHash string             `bson:"transactionHash" json:"transactionHash"`
	BlockNumber     int                `bson:"blockNumber" json:"blockNumber"`
	LastBlockNumber int                `bson:"lastBlockNumber" json:"lastBlockNumber"`
	Contract        string             `bson:"contract" json:"contract`
	ContractName    string             `bson:"contractName" json:"contractName`
	ProposalId      int                `bson:"proposalId" json:"proposalId"`
	Support         bool               `bson:"support" json:"support"`
	Voter           string             `bson:"voter" json:"voter"`
	Votes           float64            `bson:"votes" json:"votes"`
}

//struct for filters
type VoteCast_Filters struct {
	Support    bool  `bson: "support" json:"support"`
	ProposalId int64 `bson: "proposalId" json:"proposalId"`
	ChainId    int64 `bson: "chainId" json:"chainId"`
}

// init function runs first time
func init() {}

// Farm list api with page and limit
func GetVoteCastTotal(filters VoteCast_Filters) float64 {
	var records []struct {
		ID    int     `bson: "_id"`
		Count float64 `bson: "count"`
	}
	client := common.GetDB()

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName2)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := bson.M{"chainId": filters.ChainId, "support": filters.Support}

	if filters.ProposalId > 0 {
		query["proposalId"] = filters.ProposalId
	}
	// Specify a pipeline that will return the number of times each name appears
	// in the collection.
	pipeline := []bson.M{
		{"$match": query},
		{"$group": bson.M{"_id": "proposalId", "count": bson.M{"$sum": "$votes"}}},
	}
	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	opts := options.Aggregate()

	cursor, err := collection.Aggregate(ctx, pipeline, opts)
	if err != nil {
		return 0
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &records)
	if err != nil {
		return 0
	}
	//record found
	if records != nil {
		return records[0].Count
	}
	return 0
}
