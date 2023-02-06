package middlewares

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/helpers"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func validateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(os.Getenv("SECRETKEY")), nil
	})
}

func AuthorizeJWT(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if os.Getenv("STAGE") == "testing" {
		user := models.User{}
		c.Set("user", user)
		c.Next()
		return
	}

	if authHeader == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			helpers.FailedResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)),
		)
		return
	}

	tokenString := strings.Split(authHeader, " ")[1]

	token, err := validateToken(tokenString)

	if err != nil || !token.Valid {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			helpers.FailedResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)),
		)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			helpers.FailedResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)),
		)
		return
	}

	userJson, _ := json.Marshal(claims["user"])
	user := models.User{}
	err = json.Unmarshal(userJson, &user)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			helpers.FailedResponse(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized)),
		)
		return
	}
	c.Set("user", user)
}
