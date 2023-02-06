package models

import (
	"git.garena.com/sea-labs-id/batch-02/andreas-timothy/entry-task-mo-be/dtos"
	"github.com/golang-jwt/jwt/v4"
)

type TokenResponse struct {
	IDToken string `json:"token"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Photo   string `json:"photo"`
	Role    string `json:"role"`
}

type IdTokenClaims struct {
	jwt.RegisteredClaims
	User *dtos.JwtToken `json:"user"`
}
