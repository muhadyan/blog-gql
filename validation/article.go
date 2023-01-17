package validation

import (
	"fmt"

	"github.com/muhadyan/blog-graphql-api/graph/model"
)

func ValidateCreateArticle(data model.CreateArticleRequest) error {
	if data.Title == "" {
		return fmt.Errorf("title cant be empty")
	}

	if data.Content == "" {
		return fmt.Errorf("content cant be empty")
	}

	return nil
}

func ValidateUpdateArticle(data model.UpdateArticleRequest) error {
	if data.ArticleID <= 0 {
		return fmt.Errorf("article id cant be empty")
	}

	if data.Title != nil && *data.Title == "" {
		return fmt.Errorf("title cant be empty")
	}

	if data.Content != nil && *data.Content == "" {
		return fmt.Errorf("content cant be empty")
	}

	return nil
}

func ValidateDeleteArticle(data model.DeleteArticleRequest) error {
	if data.ArticleID <= 0 {
		return fmt.Errorf("article id cant be empty")
	}

	return nil
}

func ValiationGetArticles(request model.GetArticlesRequest) error {
	if request.Search != nil && *request.Search == "" {
		return fmt.Errorf("search must contain at least 1 character")
	}

	return nil
}

func ValidationGetArticle(request model.GetArticleRequest) error {
	if request.ArticleID <= 0 {
		return fmt.Errorf("article id cant be empty")
	}

	return nil
}
