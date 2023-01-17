package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"fmt"

	"github.com/muhadyan/blog-graphql-api/auth"
	"github.com/muhadyan/blog-graphql-api/graph/model"
	"github.com/muhadyan/blog-graphql-api/validation"
)

// CreateArticle is the resolver for the createArticle field.
func (r *mutationResolver) CreateArticle(ctx context.Context, data model.CreateArticleRequest) (*model.CreateArticleResponse, error) {
	user := auth.FromContext(ctx)
	if !user.Authorize() {
		return nil, errUnauthorized
	}

	err := validation.ValidateCreateArticle(data)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return nil, err
	}

	return r.articleService.Create(ctx, data, user.GetUserID())
}

// UpdateArticle is the resolver for the updateArticle field.
func (r *mutationResolver) UpdateArticle(ctx context.Context, data model.UpdateArticleRequest) (*model.UpdateArticleResponse, error) {
	user := auth.FromContext(ctx)
	if !user.Authorize() {
		return nil, errUnauthorized
	}

	err := validation.ValidateUpdateArticle(data)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return nil, err
	}

	return r.articleService.Update(ctx, data, user.GetUserID())
}

// DeleteArticle is the resolver for the deleteArticle field.
func (r *mutationResolver) DeleteArticle(ctx context.Context, data model.DeleteArticleRequest) (*model.DeleteArticleResponse, error) {
	user := auth.FromContext(ctx)
	if !user.Authorize() {
		return nil, errUnauthorized
	}

	err := validation.ValidateDeleteArticle(data)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return nil, err
	}

	return r.articleService.Delete(ctx, data, user.GetUserID())
}

// GetArticles is the resolver for the getArticles field.
func (r *queryResolver) GetArticles(ctx context.Context, request model.GetArticlesRequest) (*model.GetArticlesResult, error) {
	user := auth.FromContext(ctx)
	if !user.Authorize() {
		return nil, errUnauthorized
	}

	err := validation.ValiationGetArticles(request)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return nil, err
	}

	return r.articleService.View(ctx, request)
}

// GetArticle is the resolver for the getArticle field.
func (r *queryResolver) GetArticle(ctx context.Context, request model.GetArticleRequest) (*model.GetArticleResult, error) {
	user := auth.FromContext(ctx)
	if !user.Authorize() {
		return nil, errUnauthorized
	}

	err := validation.ValidationGetArticle(request)
	if err != nil {
		err = fmt.Errorf(err.Error())
		return nil, err
	}

	return r.articleService.Get(ctx, request)
}