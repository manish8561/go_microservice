// Common tools and helper functions
package common

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/joho/godotenv"
	"gopkg.in/go-playground/validator.v8"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"errors"
	"net/http"
	"strings"
	// "github.com/autocompound/docker_backend/user/users"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// init function
func init() {
	//initial variables
	InitVariables()
}

// A helper function to generate random string
func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Keep this two config private, it should not expose to open source
var NBSecretPassword = "A String Very Very Very Strong!!@##$!@#$"
var NBRandomPassword = "A String Very Very Very Niubilty!!@##$!@#4"

// initial values in the project
func InitVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	secret, ok := os.LookupEnv("SECRET")
	if !ok {
		secret = "secret"
	}
	NBSecretPassword = secret
	random_password, ok := os.LookupEnv("RANDOM_PASSWORD")
	if !ok {
		random_password = "random password"
	}
	NBSecretPassword = random_password
}

// A Util function to generate jwt_token which can be used in the request header
func GenToken(id string, role string) string {
	jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Set some claims
	jwt_token.Claims = jwt.MapClaims{
		"id":   id,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"role": role,
	}
	// Sign and get the complete encoded token as a string
	token, _ := jwt_token.SignedString([]byte(NBSecretPassword))
	return token
}

// My own Error type that will help return my customized Error info
//
//	{"database": {"hello":"no such table", error: "not_exists"}}
type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// To handle the error returned by c.Bind in gin framework
// https://github.com/go-playground/validator/blob/v9/_examples/translations/main.go
func NewValidatorError(err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	errs := err.(validator.ValidationErrors)
	for _, v := range errs {
		// can translate each error one at a time.
		//fmt.Println("gg",v.NameNamespace)
		if v.Param != "" {
			res.Errors[v.Field] = fmt.Sprintf("{%v: %v}", v.Tag, v.Param)
		} else {
			res.Errors[v.Field] = fmt.Sprintf("{key: %v}", v.Tag)
		}

	}
	return res
}

// Warp the error info in a object
func NewError(key string, err error) CommonError {
	res := CommonError{}
	res.Errors = make(map[string]interface{})
	res.Errors[key] = err.Error()
	return res
}

// Changed the c.MustBindWith() ->  c.ShouldBindWith().
// I don't want to auto return 400 when error happened.
// origin function is here: https://github.com/gin-gonic/gin/blob/master/context.go
func Bind(c *gin.Context, obj interface{}) error {
	b := binding.Default(c.Request.Method, c.ContentType())
	return c.ShouldBindWith(obj, b)
}

// ------- common middleware code start--------------------
// Strips 'TOKEN ' prefix from token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
		return tok[6:], nil
	}
	return tok, nil
}

// Extract  token from Authorization header
// Uses PostExtractionFilter to strip "TOKEN " prefix from header
var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	request.HeaderExtractor{"Authorization"},
	stripBearerPrefixFromTokenString,
}

// Extractor for OAuth2 access tokens.  Looks in 'Authorization'
// header then 'access_token' argument for a token.
var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

// A helper to write user_id and user_model to the context
func UpdateContextUserModel(c *gin.Context, my_user_id string, user *UserModel) {
	if my_user_id != "" {
		c.Set("my_user_id", my_user_id)
		c.Set("user", user)
	}
	c.Next()
}

// You can custom middlewares yourself as the doc: https://github.com/gin-gonic/gin#custom-middleware
//
//	r.Use(AuthMiddleware(true))
func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if auto401 {
			if c.Request.Header["Authorization"] == nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "No Token Found"})
				c.AbortWithError(http.StatusUnauthorized, errors.New("No Token Found"))
				return
			}
			// UpdateContextUserModel(c, 0)
			token, err := request.ParseFromRequest(c.Request, MyAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
				b := ([]byte(NBSecretPassword))
				return b, nil
			})
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Token Expired"})
				c.AbortWithError(http.StatusUnauthorized, err)
				return
			}
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				//checking for admin role
				if role := claims["role"].(string); role != "admin" && auto401 {
					c.JSON(http.StatusUnauthorized, gin.H{"message": "You dont have the access"})
					c.AbortWithError(http.StatusUnauthorized, errors.New("You dont have the access"))
					return
				}
				my_user_id := claims["id"].(string)

				user, err := GetUserProfile(my_user_id)
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"message": "You dont have the access"})
					c.AbortWithError(http.StatusUnauthorized, errors.New("You dont have the access"))
					return
				}
				fmt.Println(my_user_id, claims["id"])
				fmt.Println("user in common middleware", user)

				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"message": "You dont have the access"})
					c.AbortWithError(http.StatusUnauthorized, errors.New("You dont have the access"))
					return
				}

				UpdateContextUserModel(c, my_user_id, &user)
			}
		} else {
			c.Next()
		}
	}
}

// ------- common middleware code end--------------------
