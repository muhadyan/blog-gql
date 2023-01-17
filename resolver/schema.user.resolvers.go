package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"github.com/muhadyan/blog-graphql-api/graph/model"
	"github.com/muhadyan/blog-graphql-api/validation"
)

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, data model.RegisterUserRequest) (*model.RegisterUserResponse, error) {
	err := validation.ValidateRegister(data)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return nil, err
	}

	return r.userService.Register(ctx, data)
}

// VerifyUser is the resolver for the verifyUser field.
func (r *mutationResolver) VerifyUser(ctx context.Context, data model.VerifyUserRequest) (*model.VerifyUserResponse, error) {
	err := validation.ValidateVerify(data)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return nil, err
	}

	return r.userService.Verify(ctx, data)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, data model.LoginRequest) (*model.LoginResponse, error) {
	err := validation.ValidateLogin(data)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return nil, err
	}

	return r.userService.Login(ctx, data)
}