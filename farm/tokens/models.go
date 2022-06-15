package tokens

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/autocompound/docker_backend/farm/common"
	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	// "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	// token "./autocompound.go"
)

const CollectionName = "tokens"

// Models should only be concerned with database schema, more strict checking should be put in validator.
// event Transfer(address indexed from, address indexed to, uint256 value);
//
type TokensModel struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Created     time.Time          `bson:"_created" json:"_created,omitempty"`
	Modified    time.Time          `bson:"_modified" json:"_modified,omitempty"`
	Chain_Id    int                `bson:"chain_id" json:"chain_id,omitempty"`
	Address     string             `bson:"address" json:"address,omitempty"`
	Symbol      string             `bson:"symbol" json:"symbol,omitempty"` //address field of strategy
	BlockNumber int                `bson:"blockNumber" json:"blockNumber,omitempty"`
	Status      string             `bson: "status" json: "status"`
}

// init func in go file
func init() {
	// create index
	common.AddIndex(os.Getenv("MONGO_DATABASE"), CollectionName, bson.D{{"address", 1}, {"symbol", 1}})

	//comment
	fmt.Printf("tokens model")

	//start the cron
	// StartCall()
	GetContract(4)
}

// cron func call
func StartCall() {
	c := cron.New()
	c.AddFunc("0 */30 * * * *", func() {
		fmt.Println("[Job 1]Every 30 minutes job\n")
		r := UpdateAll()
		fmt.Println("cron job return value", r)
	})
	// Start cron with one scheduled job
	c.Start()
}

// Get All Symbols
func UpdateAll() bool {
	client := common.GetDB()
	var records []*TokensModel

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	opts := options.Find().SetSort(bson.D{{"_created", -1}})
	//SetProjection(bson.M{"_id": 0, "_created": 1, "_modified": 1, "firstname": 1, "lastname": 1, "status": 1, "email": 1, "role": 1, "passwordhash": 0})
	query := bson.M{"status": "active"}

	cursor, err := collection.Find(ctx, query, opts)
	if err != nil {
		return false
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &records)

	return true
}

// You could input an TokensModel which will be saved in database returning with error info
// 	if err := SaveOne(&farmModel); err != nil { ... }
func SaveOne(data *TokensModel) error {
	client := common.GetDB()
	person := &TokensModel{}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//convert string to uppercase
	data.Symbol = strings.ToUpper(data.Symbol)
	// to check for unique email address
	err := collection.FindOne(ctx, bson.M{"symbol": data.Symbol}).Decode(&person)
	if err != nil {
		data.Created = time.Now()
		data.Modified = time.Now()
		res, err := collection.InsertOne(ctx, data)
		fmt.Println(res, "Inserted")
		return err
	}
	return errors.New("symbol already exists!")
}

// You could input an TokensModel which will be updated in database returning with error info
// 	if err := UpdateOne(&farmModel); err != nil { ... }
func UpdateOne(data *TokensModel) error {
	client := common.GetDB()
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// to check for unique email address
	opts := options.Update().SetUpsert(false)
	// update := bson.D{{"$set", bson.D{{"token", "newemail@example.com"}}}}
	update := bson.D{{"$set", data}}

	res, err := collection.UpdateOne(ctx, bson.M{"_id": data.ID}, update, opts)
	if err != nil {
		return err
	}
	fmt.Println(res, "Updated")
	return err
}

// You could input string which will be saved in database returning with error info
// 	if err := FindOne(&farmModel); err != nil { ... }
func GetFarm(ID string) (TokensModel, error) {
	client := common.GetDB()
	farm := &TokensModel{}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//convert string to objectid
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return *farm, err
	}

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	// opts := options.FindOne().SetProjection(bson.M{"_id": 0, "_created": 1, "_modified": 1, "firstname": 1, "lastname": 1, "status": 1, "email": 1, "role": 1, "passwordhash": 0})
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&farm)

	return *farm, err
}

// Price Feed list api with page and limit
func GetTotal(status string) int64 {
	client := common.GetDB()

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := bson.M{}
	num, err := collection.CountDocuments(ctx, query)
	if err != nil {
		return 0
	}
	return num
}

// Price Feed list api with page and limit
func GetAll(page int64, limit int64, status string) ([]*TokensModel, error) {
	client := common.GetDB()
	var records []*TokensModel

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	opts := options.Find().SetSort(bson.D{{"_created", -1}}).SetSkip((page - 1) * limit).SetLimit(limit)
	//SetProjection(bson.M{"_id": 0, "_created": 1, "_modified": 1, "firstname": 1, "lastname": 1, "status": 1, "email": 1, "role": 1, "passwordhash": 0})

	query := bson.M{}

	cursor, err := collection.Find(ctx, query, opts)
	if err != nil {
		return records, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &records)

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return records, err
}

// delete record from collection
func DeleteRecordModel(ID string) (bool, error) {
	client := common.GetDB()

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)

	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return false, err
	}
	res, err := collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return false, err
	}
	fmt.Println(res, "Delete")
	return true, err
}

// You could input string which will be saved in database returning with error info
// 	if err := FindOne(&farmModel); err != nil { ... }
func GetContract(chain_id int) error {
	// Create an IPC based RPC connection to a remote node
	conn := common.Get_Eth_Connection(chain_id)

	contractAddress := ethcommon.HexToAddress("0x23fc559A2b5c749F87D0Ef30099eb342be3B3150")
	token, err := NewTokens(contractAddress, conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a Token contract: %v", err)
	}
	decimals, err := token.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Printf("Failed to retrieve token name: %v", err)
	}
	fmt.Println("Token decimals:", decimals)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(10826826),
		ToBlock:   big.NewInt(10827728),
		Addresses: []ethcommon.Address{
			contractAddress,
		},
	}
	//logs from contract
	logs, err := conn.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//abi of the contracts
	contractAbi, err := abi.JSON(strings.NewReader(string(TokensABI)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("abi", contractAbi)

	logTransferSig := []byte("Transfer(address,address,uint256)")
	// LogApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	// logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)

	for _, vLog := range logs {
		// fmt.Println("Log Index: ", vLog.Index)
		// fmt.Println("vLog.Topics[0] ", vLog.Topics[0].Hex())

		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Printf("Log Name: Transfer\n")

			transferEvent, err := token.ParseTransfer(vLog)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Log Block Number: ", vLog.BlockNumber)
			fmt.Println("\n Log Transaction: ", vLog.TxHash.Hex())

			fmt.Println("From : ", transferEvent.From)
			fmt.Println("To : ", transferEvent.To)
			fmt.Println("Value: ", transferEvent.Value)

			// case logApprovalSigHash.Hex():
			// 	fmt.Println("Log Name: Approval\n")

			// 	approvalEvent, err := token.ParseApproval(vLog)
			// 	if err != nil {
			// 		log.Fatal(err)
			// 	}
			// 	// approvalEvent.Owner = ethcommon.HexToAddress(vLog.Topics[1].Hex())
			// 	// approvalEvent.Spender = ethcommon.HexToAddress(vLog.Topics[2].Hex())
			// 	fmt.Print("\n value: ", approvalEvent.Value)
			// 	fmt.Print("\n owner: ", approvalEvent.Owner)
			// 	fmt.Print("\n Spender: ", approvalEvent.Spender)

			// 	fmt.Println("\n d: ",approvalEvent)
		}

		fmt.Printf("\n\n")
	}
	return err
}
