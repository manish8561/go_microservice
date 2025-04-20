package contracts

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/autocompound/docker_backend/farm/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5/request"
)

// Strips 'TOKEN ' prefix from token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 5 && strings.ToUpper(tok[0:6]) == "TOKEN " {
		return tok[6:], nil
	}
	return tok, nil
}

// Extract token from Authorization header
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
func UpdateContextUserModel(c *gin.Context, MyUserID string) {
	// var myUserModel UserModel
	if MyUserID != "" {
		// db := common.GetDB()
		// db.First(&myUserModel, MyUserID)
	}
	c.Set("my_user_id", MyUserID)
	// c.Set("my_user_model", myUserModel)
	c.Next()
}

// You can custom middlewares yourself as the doc: https://github.com/gin-gonic/gin#custom-middleware
//
//	r.Use(AuthMiddleware(true))
func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// UpdateContextUserModel(c, 0)
		token := common.ExtractTokenFromHeader(c)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}

		if claims, err := common.ValidateToken(token); err == nil {

			// checking for admin role
			if role := claims.Role; role != "admin" {
				if auto401 {
					c.AbortWithError(http.StatusUnauthorized, err)
				}
				return
			}
			MyUserID := claims.ID
			fmt.Println(MyUserID, claims.ID)
			UpdateContextUserModel(c, MyUserID)
		}
	}
}
