package service

import (
	"context"
	"time"

	"fmt"

	"github.com/dchest/uniuri"
	"github.com/muhadyan/blog-graphql-api/auth"
	"github.com/muhadyan/blog-graphql-api/graph/model"
	"github.com/muhadyan/blog-graphql-api/repository"
	"github.com/muhadyan/blog-graphql-api/service/helper"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	Register(ctx context.Context, data model.RegisterUserRequest) (*model.RegisterUserResponse, error)
	Verify(ctx context.Context, data model.VerifyUserRequest) (*model.VerifyUserResponse, error)
	Login(ctx context.Context, data model.LoginRequest) (*model.LoginResponse, error)
}

type userCtx struct{}

func NewUserService() User {
	return &userCtx{}
}

func (s *userCtx) Register(ctx context.Context, data model.RegisterUserRequest) (*model.RegisterUserResponse, error) {
	err := helper.ValidateRegister(data)
	if err != nil {
		return nil, err
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 12)
	if err != nil {
		err = fmt.Errorf(fmt.Sprintf("error while hashing password. Err: %s", err))
		return nil, err
	}

	token := uniuri.NewLen(12)

	u := model.User{
		Username: data.Username,
		Password: string(hashPassword),
		Email:    data.Email,
		Fullname: data.Fullname,
		Token:    &token,
	}
	user, err := repository.InsertUser(u)
	if err != nil {
		return nil, err
	}

	res := model.RegisterUserResponse{
		Message: "Success",
		User: &model.RegisterDataResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Fullname:  user.Fullname,
			IsActive:  user.IsActive,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
	}

	return &res, nil
}

func (s *userCtx) Verify(ctx context.Context, data model.VerifyUserRequest) (*model.VerifyUserResponse, error) {
	getUser, err := helper.ValidateVerify(data)
	if err != nil {
		return nil, err
	}

	token, err := auth.GenerateJwt(getUser)
	if err != nil {
		return nil, err
	}

	user, err := repository.UpdateUser(model.User{
		ID:        data.UserID,
		Token:     &token,
		IsActive:  true,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	res := model.VerifyUserResponse{
		Message: "Success",
		User: &model.UserDataResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Fullname:  user.Fullname,
			Token:     user.Token,
			IsActive:  user.IsActive,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
	}

	return &res, nil
}

func (s *userCtx) Login(ctx context.Context, data model.LoginRequest) (*model.LoginResponse, error) {
	getUser, err := helper.ValidateLogin(data)
	if err != nil {
		return nil, err
	}

	token, err := auth.GenerateJwt(getUser)
	if err != nil {
		return nil, err
	}

	user, err := repository.UpdateUser(model.User{
		ID:        getUser.ID,
		Token:     &token,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	res := model.LoginResponse{
		Message: "Success",
		User: &model.UserDataResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			Fullname:  user.Fullname,
			Token:     user.Token,
			IsActive:  user.IsActive,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
	}

	return &res, nil
}
