package service

import (
	"context"
	"fmt"
	"time"

	"github.com/muhadyan/blog-graphql-api/graph/model"
	"github.com/muhadyan/blog-graphql-api/repository"
	"github.com/muhadyan/blog-graphql-api/service/helper"
)

type Article interface {
	Create(ctx context.Context, data model.CreateArticleRequest, userID int) (*model.CreateArticleResponse, error)
	Update(ctx context.Context, data model.UpdateArticleRequest, userID int) (*model.UpdateArticleResponse, error)
	Delete(ctx context.Context, data model.DeleteArticleRequest, userID int) (*model.DeleteArticleResponse, error)
	View(ctx context.Context, request model.GetArticlesRequest) (*model.GetArticlesResult, error)
	Get(ctx context.Context, request model.GetArticleRequest) (*model.GetArticleResult, error)
}

type articleCtx struct{}

func NewArticleService() Article {
	return &articleCtx{}
}

func (s *articleCtx) Create(ctx context.Context, data model.CreateArticleRequest, userID int) (*model.CreateArticleResponse, error) {
	err := helper.ValidateCreateArticle(userID)
	if err != nil {
		return nil, err
	}

	a := model.Article{
		UserID:      userID,
		Title:       data.Title,
		Content:     data.Content,
		IsModerated: data.IsModerated,
	}
	article, err := repository.InsertArticle(a)
	if err != nil {
		return nil, err
	}

	res := model.CreateArticleResponse{
		Message: "Success",
		Article: &model.Article{
			ID:          article.ID,
			UserID:      article.UserID,
			Title:       article.Title,
			Content:     article.Content,
			Likes:       article.Likes,
			Comments:    article.Comments,
			IsModerated: article.IsModerated,
			CreatedAt:   article.CreatedAt,
			UpdatedAt:   article.UpdatedAt,
		},
	}

	return &res, nil
}

func (s *articleCtx) Update(ctx context.Context, data model.UpdateArticleRequest, userID int) (*model.UpdateArticleResponse, error) {
	err := helper.ValidateUpdateArticle(data, userID)
	if err != nil {
		return nil, err
	}

	a := model.Article{
		ID:        data.ArticleID,
		UpdatedAt: time.Now(),
	}
	if data.Title != nil {
		a.Title = *data.Title
	}
	if data.Content != nil {
		a.Content = *data.Content
	}
	if data.IsModerated != nil {
		a.IsModerated = *data.IsModerated
	}

	article, err := repository.UpdateArticle(a)
	if err != nil {
		return nil, err
	}

	res := model.UpdateArticleResponse{
		Message: "Success",
		Article: &model.Article{
			ID:          article.ID,
			UserID:      article.UserID,
			Title:       article.Title,
			Content:     article.Content,
			Likes:       article.Likes,
			Comments:    article.Comments,
			IsModerated: article.IsModerated,
			CreatedAt:   article.CreatedAt,
			UpdatedAt:   article.UpdatedAt,
		},
	}

	return &res, nil
}

func (s *articleCtx) Delete(ctx context.Context, data model.DeleteArticleRequest, userID int) (*model.DeleteArticleResponse, error) {
	err := helper.ValidateDeleteArticle(data, userID)
	if err != nil {
		return nil, err
	}

	a := model.Article{
		ID: data.ArticleID,
	}
	err = repository.DeleteArticle(a)
	if err != nil {
		return nil, err
	}

	res := model.DeleteArticleResponse{
		Message: "Success",
	}

	return &res, nil
}

func (s *articleCtx) View(ctx context.Context, request model.GetArticlesRequest) (*model.GetArticlesResult, error) {
	articles, err := repository.SelectArticles(request)
	if err != nil {
		return nil, err
	}

	res := model.GetArticlesResult{
		Message:  "Success",
		Articles: articles,
	}

	return &res, nil
}

func (s *articleCtx) Get(ctx context.Context, request model.GetArticleRequest) (*model.GetArticleResult, error) {
	article, err := repository.GetArticle(model.Article{
		ID: request.ArticleID,
	})
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, fmt.Errorf("article not found")
	}

	res := model.GetArticleResult{
		Message: "Success",
		Article: article,
	}

	return &res, nil
}
