// Common tools and helper functions
package common

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/go-playground/validator.v8"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"errors"
	"net/http"
	"strings"

	// "github.com/autocompound/docker_backend/user/users"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
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
	randomPassword, ok := os.LookupEnv("RANDOM_PASSWORD")
	if !ok {
		randomPassword = "random password"
	}
	NBRandomPassword = randomPassword
}

// A Util function to generate jwt_token which can be used in the request header

// CustomClaims extends the default claims
type CustomClaims struct {
	ID   string `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateToken generates a JWT token
func GenToken(id string, role string) string {
	claims := CustomClaims{
		ID:   id,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "my_app",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	str, err := token.SignedString(NBSecretPassword)
	if err != nil {
		fmt.Println("generate token error: ", err)
	}
	return str
}

func GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// My own Error type that will help return my customized Error info
//	{"database": {"hello":"no such table", error: "not_exists"}}
type CommonError struct {
	Errors map[string]interface{} `json:"errors"`
}

// ValidateToken parses and validates a JWT token
func ValidateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return NBSecretPassword, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
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
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromTokenString,
}

// Extractor for OAuth2 access tokens.  Looks in 'Authorization'
// header then 'access_token' argument for a token.
var MyAuth2Extractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

// A helper to write user_id and user_model to the context
func UpdateContextUserModel(c *gin.Context, MyUserID string, user *UserModel) {
	if MyUserID != "" {
		c.Set("my_user_id", MyUserID)
		c.Set("user", user)
	}
	c.Next()
}

// ExtractTokenFromHeader extracts the token from Authorization header
func ExtractTokenFromHeader(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	// Expecting header format: "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
		return parts[1]
	}

	return ""
}

// You can custom middlewares yourself as the doc: https://github.com/gin-gonic/gin#custom-middleware
//
//	r.Use(AuthMiddleware(true))
func respondUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, gin.H{"message": message})
	c.AbortWithError(http.StatusUnauthorized, errors.New(message))
}

func processTokenAndUser(c *gin.Context, token string, auto401 bool) bool {
	claims, err := ValidateToken(token)
	if err != nil {
		respondUnauthorized(c, ErrAccessDeniedMsg)
		return false
	}
	//checking for admin role
	if role := claims.Role; role != "admin" && auto401 {
		respondUnauthorized(c, ErrAccessDeniedMsg)
		return false
	}
	MyUserID := claims.ID
	user, err := GetUserProfile(MyUserID)
	if err != nil {
		respondUnauthorized(c, ErrAccessDeniedMsg)
		return false
	}
	fmt.Println(MyUserID, claims.ID)
	fmt.Println("user in common middleware", user)
	// Update the context with user_id and user_model
	UpdateContextUserModel(c, MyUserID, &user)
	return true
}

func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !auto401 {
			c.Next()
			return
		}
		token := ExtractTokenFromHeader(c)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}
		if !processTokenAndUser(c, token, auto401) {
			return
		}
	}
}

// ------- common middleware code end--------------------
