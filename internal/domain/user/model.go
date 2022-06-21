package user

import "fmt"

type User struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
	Email        string `json:"email" bson:"email"`
}

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (userDto *CreateUserDTO) OK() error {
	if len(userDto.Username) == 0 {
		return fmt.Errorf("username is required")
	}
	if len(userDto.Email) == 0 {
		return fmt.Errorf("email is required")
	}
	if len(userDto.Password) == 0 {
		return fmt.Errorf("password is required")
	}
	return nil
}
