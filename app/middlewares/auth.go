package middleware

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/hashfyre/sample-go-app/app/context"
	"github.com/hashfyre/sample-go-app/app/controllers"
	"github.com/hashfyre/sample-go-app/app/types"
	"github.com/hashfyre/sample-go-app/pkg/header"
)

// BasicAuth - middleware that validates access-token and refresh-token
func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// parse the accessToken from Auth header
		auth := c.Request.Header.Get("Authorization")
		userCreds, err := header.ParseAuth("basic", auth)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, types.ResponseError{
				Code:    "ERR_UNAUTHORIZED_REQUEST",
				Message: err.Error(),
			})
			return
		}

		userEmail, userPass, err := parseCreds(userCreds)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, types.ResponseError{
				Code:    "ERR_UNAUTHORIZED_REQUEST",
				Message: err.Error(),
			})
			return
		}

		// validate accessToken
		user, err := controllers.UserLogin(userEmail, userPass)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, types.ResponseError{
				Code:    "ERR_UNAUTHORIZED_REQUEST",
				Message: errUserUnauthorized.Error(),
			})
			return
		}

		// Set valid userInfo from accessToken into the context; chain
		c.Set(context.UserIDKey, user.BaseModel.ID)
		c.Next()
	}
}

func parseCreds(creds string) (string, string, error) {
	decodedCreds, err := base64.StdEncoding.DecodeString(creds)
	if err != nil {
		return "", "", errMalformedAuthHeader
	}

	credSlice := strings.Split(string(decodedCreds), ":")
	if len(credSlice) != 2 {
		return "", "", errMalformedAuthHeader
	}
	email := credSlice[0]
	pass := credSlice[1]

	return email, pass, nil
}
