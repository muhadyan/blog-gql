package validation

import (
	"fmt"
	"net/mail"

	"github.com/muhadyan/blog-graphql-api/graph/model"
)

func ValidateRegister(data model.RegisterUserRequest) error {
	if data.Username == "" {
		return fmt.Errorf("username cant be empty")
	}

	if data.Email == "" {
		return fmt.Errorf("email cant be empty")
	}

	_, err := mail.ParseAddress(data.Email)
	if err != nil {
		return fmt.Errorf("invalid email format")
	}

	if data.Fullname == "" {
		return fmt.Errorf("fullname cant be empty")
	}

	if data.Password == "" {
		return fmt.Errorf("password cant be empty")
	}

	return nil
}

func ValidateVerify(data model.VerifyUserRequest) error {
	if data.UserID <= 0 {
		return fmt.Errorf("user id cant be empty")
	}

	if data.Token == "" {
		return fmt.Errorf("token verification cant be empty")
	}

	return nil
}

func ValidateLogin(data model.LoginRequest) error {
	if data.Username == "" {
		return fmt.Errorf("username cant be empty")
	}

	if data.Password == "" {
		return fmt.Errorf("password cant be empty")
	}

	return nil
}
