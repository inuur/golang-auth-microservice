package user

import (
	"context"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	storage Storage
}

func NewUserService() *Service {
	userStorage := NewStorage()
	return &Service{
		storage: userStorage,
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (s *Service) CreateUser(dto CreateUserDTO) (*User, error) {
	hashedPassword, _ := HashPassword(dto.Password)

	user := User{
		ID:           "",
		Username:     dto.Username,
		PasswordHash: hashedPassword,
		Email:        dto.Email,
	}

	return s.storage.Create(context.Background(), user)
}

func (s *Service) FindUser(id string) (*User, error) {
	return s.storage.FindOne(context.Background(), id)
}
