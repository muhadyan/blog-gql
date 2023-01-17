package helper

import (
	"fmt"

	"github.com/muhadyan/blog-graphql-api/graph/model"
	"github.com/muhadyan/blog-graphql-api/repository"
)

func ValidateCreateArticle(userID int) error {
	user, err := repository.GetUser(model.User{
		ID: userID,
	})
	if err != nil {
		return fmt.Errorf("error while querying db. Err: %s", err)
	}
	if user == nil {
		return fmt.Errorf("user doesn't exist")
	}

	return nil
}

func ValidateUpdateArticle(data model.UpdateArticleRequest, userID int) error {
	user, err := repository.GetUser(model.User{
		ID: userID,
	})
	if err != nil {
		return fmt.Errorf("error while querying db. Err: %s", err)
	}
	if user == nil {
		return fmt.Errorf("user doesn't exist")
	}

	article, err := repository.GetArticle(model.Article{
		ID:     data.ArticleID,
		UserID: userID,
	})
	if err != nil {
		return fmt.Errorf("error while querying db. Err: %s", err)
	}
	if article == nil {
		return fmt.Errorf("user cant update this article")
	}

	return nil
}

func ValidateDeleteArticle(data model.DeleteArticleRequest, userID int) error {
	user, err := repository.GetUser(model.User{
		ID: userID,
	})
	if err != nil {
		return fmt.Errorf("error while querying db. Err: %s", err)
	}
	if user == nil {
		return fmt.Errorf("user doesn't exist")
	}

	article, err := repository.GetArticle(model.Article{
		ID:     data.ArticleID,
		UserID: userID,
	})
	if err != nil {
		return fmt.Errorf("error while querying db. Err: %s", err)
	}
	if article == nil {
		return fmt.Errorf("user cant update this article")
	}

	return nil
}
