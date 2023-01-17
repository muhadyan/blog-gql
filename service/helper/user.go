package helper

import (
	"fmt"

	"github.com/muhadyan/blog-graphql-api/graph/model"
	"github.com/muhadyan/blog-graphql-api/repository"
	"golang.org/x/crypto/bcrypt"
)

func ValidateRegister(data model.RegisterUserRequest) error {
	userByUsername, err := repository.GetUser(model.User{
		Username: data.Username,
	})
	if err != nil {
		return fmt.Errorf("error while querying db. Err: %s", err)
	}
	if userByUsername != nil {
		return fmt.Errorf("username already exist")
	}

	userByEmail, err := repository.GetUser(model.User{
		Email: data.Email,
	})
	if err != nil {
		return fmt.Errorf("error while querying db. Err: %s", err)
	}
	if userByEmail != nil {
		return fmt.Errorf("email already exist")
	}

	return nil
}

func ValidateVerify(data model.VerifyUserRequest) (*model.User, error) {
	user, err := repository.GetUser(model.User{
		ID:    data.UserID,
		Token: &data.Token,
	})
	if err != nil {
		return nil, fmt.Errorf("error while querying db. Err: %s", err)
	}
	if user == nil {
		return nil, fmt.Errorf("id and token not match")
	}
	if user.IsActive {
		return nil, fmt.Errorf("user already active")
	}

	return user, nil
}

func ValidateLogin(data model.LoginRequest) (*model.User, error) {
	user, err := repository.GetUser(model.User{
		Username: data.Username,
		IsActive: true,
	})
	if err != nil {
		return nil, fmt.Errorf("error while querying db. Err: %s", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user doesn't exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return nil, fmt.Errorf("wrong password")
	}

	return user, nil
}
