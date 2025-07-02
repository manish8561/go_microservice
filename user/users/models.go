package users

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"log"

	"github.com/autocompound/docker_backend/user/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "go.mongodb.org/mongo-driver/mongo/readpref"
	// pb "github.com/autocompound/docker_backend/user/helloworld"
	"golang.org/x/crypto/bcrypt"
)

const CollectionName = "users"

// var UserModel= common.UserModel
// Models should only be concerned with database schema, more strict checking should be put in validator.
//
// HINT: If you want to split null and "", you should use *string instead of string.
type UserModel struct {
	ID                 primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Created            time.Time          `bson:"_created" json:"_created"`
	Modified           time.Time          `bson:"_modified" json:"_modified"`
	Firstname          string             `bson:"firstname" json:"firstname"`
	Lastname           string             `bson:"lastname" json:"lastname"`
	Status             string             `bson:"status" json:"status"`
	Email              string             `bson:"email" json:"email"`
	Role               string             `bson:"role" json:"role"`
	PasswordHash       string             `json:"-"` // to hide filed in json
	RefreshToken       string             `bson:"refresh_token"`
	RefreshTokenExpiry time.Time          `bson:"refresh_token_expiry"`
	// Image              *string
}

// initialize function
func init() {
	//create index
	common.AddIndex(os.Getenv("MONGO_DATABASE"), CollectionName, bson.M{"email": "text"})
}

// A hack way to save ManyToMany relationship,
// DB schema looks like: id, created_at, updated_at, deleted_at, following_id, followed_by_id.
//
// Retrieve them by:
// 	db.Where(FollowModel{ FollowingID:  v.ID, FollowedByID: u.ID, }).First(&follow)
// 	db.Where(FollowModel{ FollowedByID: u.ID, }).Find(&follows)
//
// type FollowModel struct {
// 	gorm.Model
// 	Following    UserModel
// 	FollowingID  uint
// 	FollowedBy   UserModel
// 	FollowedByID uint
// }

// What's bcrypt? https://en.wikipedia.org/wiki/Bcrypt
// Golang bcrypt doc: https://godoc.org/golang.org/x/crypto/bcrypt
// You can change the value in bcrypt.DefaultCost to adjust the security index.
//
//	err := userModel.setPassword("password0")
func (u *UserModel) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("password should not be empty")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

// Database will only save the hashed string, you should check it by util function.
//
//	if err := serModel.checkPassword("password0"); err != nil { password error }
func (u *UserModel) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// You could input the conditions and it will return an UserModel in database with error info.
//
//	userModel, err := FindOneUser(&UserModel{Username: "username0"})
func FindOneUser(email string) (UserModel, error) {
	person := &UserModel{}

	client := common.GetDB()

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&person)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Found user:", person.ID.Hex())
	}
	return *person, err
}

// You could input an UserModel which will be saved in database returning with error info
//
//	if err := SaveOne(&userModel); err != nil { ... }
func SaveOne(data *UserModel) error {
	client := common.GetDB()
	person := &UserModel{}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// to check for unique email address
	err := collection.FindOne(ctx, bson.M{"email": data.Email}).Decode(&person)
	if err != nil {
		refreshToken, _ := common.GenerateRefreshToken()
		data.RefreshToken = refreshToken
		data.RefreshTokenExpiry = time.Now().Add(7 * 24 * time.Hour) // 7 days
		// Save user to DB
		res, err := collection.InsertOne(ctx, data)
		fmt.Println(res, "Inserted")
		return err
	}
	return errors.New("user already exists")
}

// You could input string which will be saved in database returning with error info
//
//	if err := FindOne(&userModel); err != nil { ... }
func GetProfile(ID string) (UserModel, error) {
	client := common.GetDB()
	person := &UserModel{}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//convert string to objectid
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		// panic(err)
		return *person, err
	}

	// Find the document for which the _id field matches id.
	// Specify the Sort option to sort the documents by age.
	// The first document in the sorted order will be returned.
	// opts := options.FindOne().SetProjection(bson.M{"_id": 0, "_created": 1, "_modified": 1, "firstname": 1, "lastname": 1, "status": 1, "email": 1, "role": 1, "passwordhash": 0})
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&person)

	return *person, err
}

// You could update properties of an UserModel to database returning with error info.
//  err := db.Model(userModel).Update(UserModel{Username: "wangzitian0"}).Error
// func (model *UserModel) Update(data interface{}) error {
// 	db := common.GetDB()
// 	err := db.Model(model).Update(data).Error
// 	return err
// }

// You could add a following relationship as userModel1 following userModel2
// 	err = userModel1.following(userModel2)
// func (u UserModel) following(v UserModel) error {
// 	db := common.GetDB()
// 	var follow FollowModel
// 	err := db.FirstOrCreate(&follow, &FollowModel{
// 		FollowingID:  v.ID,
// 		FollowedByID: u.ID,
// 	}).Error
// 	return err
// }

// You could check whether  userModel1 following userModel2
// 	followingBool = myUserModel.isFollowing(self.UserModel)
// func (u UserModel) isFollowing(v UserModel) bool {
// 	db := common.GetDB()
// 	var follow FollowModel
// 	db.Where(FollowModel{
// 		FollowingID:  v.ID,
// 		FollowedByID: u.ID,
// 	}).First(&follow)
// 	return follow.ID != 0
// }

// You could delete a following relationship as userModel1 following userModel2
//
//	err = userModel1.unFollowing(userModel2)
//
//	func (u UserModel) unFollowing(v UserModel) error {
//		db := common.GetDB()
//		err := db.Where(FollowModel{
//			FollowingID:  v.ID,
//			FollowedByID: u.ID,
//		}).Delete(FollowModel{}).Error
//		return err
//	}
func ChangePasswordOne(data *UserModel) (*mongo.UpdateResult, error) {
	client := common.GetDB()

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// options for update
	opts := options.Update().SetUpsert(false)

	modified := time.Now()
	update := bson.M{"_modified": modified, "passwordhash": data.PasswordHash}
	update = bson.M{"$set": update}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": data.ID}, update, opts)
	if err != nil {
		return result, err
	}
	return result, nil
}

// FindUserByRefreshToken finds a user by their refresh token and checks expiry
func FindUserByRefreshToken(refreshToken string) (UserModel, error) {
	client := common.GetDB()
	person := &UserModel{}

	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"refresh_token": refreshToken}).Decode(&person)
	if err != nil {
		return *person, err
	}
	if person.RefreshTokenExpiry.Before(time.Now()) {
		return *person, errors.New("refresh token expired")
	}
	return *person, nil
}

func UpdateRefreshToken(userID string) string {
	client := common.GetDB()
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(CollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	refreshToken, err := common.GenerateRefreshToken()
	if err != nil {
		log.Println("Error generating refresh token", err)
		return ""
	}

	update := bson.M{
		"$set": bson.M{
			"refresh_token":        refreshToken,
			"refresh_token_expiry": time.Now().Add(7 * 24 * time.Hour), // 7 days
		},
	}

	objID, _ := primitive.ObjectIDFromHex(userID)
	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		log.Println("Error updating refresh token", "error", err)
		return ""
	}
	if result.MatchedCount == 0 {
		log.Println("No user found with the given user ID", userID)
		return ""
	}
	return refreshToken
}
