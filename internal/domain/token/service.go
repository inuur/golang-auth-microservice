package token

import (
	"authService/internal/config"
	"authService/internal/domain/user"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Service struct {
	storage     Storage
	userService *user.Service
}

func NewTokenService() *Service {
	return &Service{
		storage:     NewStorage(),
		userService: user.NewUserService(),
	}
}

func HashToken(token string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(token), 5)
	return string(bytes), err
}

func (s *Service) GenerateToken(user *user.User) (*JWTTokenDTO, error) {
	secret := config.GetConfig().JWT.SecretKey
	atExpTime := config.GetConfig().JWT.AccessTokenExpTime
	rtExpTime := config.GetConfig().JWT.RefreshTokenExpTime

	atExpirationTime := time.Now().Add(time.Duration(atExpTime) * time.Minute)
	atUUID := uuid.New().String()
	atClaims := &Claims{
		UserId:   user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: atExpirationTime.Unix(),
			Id:        atUUID,
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, atClaims)
	accessTokenString, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	rtExpirationTime := time.Now().Add(time.Duration(rtExpTime) * time.Minute)
	rtClaims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: rtExpirationTime.Unix(),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS512, rtClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	tokenLast50Symbols := refreshTokenString[len(refreshTokenString)-50:]

	hashedRefreshToken, _ := HashToken(tokenLast50Symbols)
	_ = s.storage.SaveRefreshToken(context.Background(), hashedRefreshToken, atUUID)

	return &JWTTokenDTO{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func (s *Service) RefreshToken(tokenDto *JWTTokenDTO) (*JWTTokenDTO, error) {
	secret := config.GetConfig().JWT.SecretKey
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenDto.AccessToken, claims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

	if err != nil {
		return nil, fmt.Errorf("invalid access_token")
	}

	parsedRefreshToken, err := jwt.Parse(tokenDto.RefreshToken,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
	if err != nil || !parsedRefreshToken.Valid {
		return nil, fmt.Errorf("invalid refresh_token")
	}

	atUUID := claims.Id
	token, err := s.storage.FindRefreshToken(context.Background(), atUUID)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh_token")
	}

	tokenLast50Symbols := tokenDto.RefreshToken[len(tokenDto.RefreshToken)-50:]

	err = bcrypt.CompareHashAndPassword([]byte(token.Token), []byte(tokenLast50Symbols))
	if err != nil {
		return nil, fmt.Errorf("invalid refresh_token")
	}

	err = s.storage.DeleteRefreshToken(context.Background(), atUUID)
	if err != nil {
		return nil, err
	}

	foundUser, err := s.userService.FindUser(claims.UserId)
	if err != nil {
		return nil, err
	}

	return s.GenerateToken(foundUser)
}
