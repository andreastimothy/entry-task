package helpers

import (
	"os"
	"strconv"
	"time"

	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/dtos"
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/models"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateJwtToken(data *dtos.JwtToken) (string, error) {
	idExp, err := strconv.ParseInt(os.Getenv("TIME"), 10, 64)

	if err != nil {
		return "", err
	}

	unixTime := time.Now().Unix()
	tokenExp := unixTime + idExp

	claims := &models.IdTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Entry Task MO FE",
			ExpiresAt: &jwt.NumericDate{Time: time.Unix(tokenExp, 0)},
			IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		},
		User: data,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRETKEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}