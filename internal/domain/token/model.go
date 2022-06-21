package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type JWTTokenDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (tokenDTO *JWTTokenDTO) OK() error {
	if len(tokenDTO.AccessToken) == 0 {
		return fmt.Errorf("access_token is required")
	}
	if len(tokenDTO.RefreshToken) == 0 {
		return fmt.Errorf("refresh_token is required")
	}
	return nil
}

type Claims struct {
	UserId   string
	Username string
	jwt.StandardClaims
}

type RefreshToken struct {
	Id            string `bson:"_id,omitempty"`
	AccessTokenID string `bson:"access_token_id"`
	Token         string `bson:"token"`
}
