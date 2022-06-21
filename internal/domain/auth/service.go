package auth

import (
	"authService/internal/domain/token"
	usr "authService/internal/domain/user"
	"fmt"
)

type Service struct {
	tokenService *token.Service
	userService  *usr.Service
}

func NewAuthService() *Service {
	return &Service{
		tokenService: token.NewTokenService(),
		userService:  usr.NewUserService(),
	}
}

func (s *Service) Authorize(userId string) (*token.JWTTokenDTO, error) {
	user, err := s.userService.FindUser(userId)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	return s.tokenService.GenerateToken(user)
}
