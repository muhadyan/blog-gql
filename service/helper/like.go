package helper

import (
	"fmt"

	"github.com/muhadyan/blog-graphql-api/graph/model"
	"github.com/muhadyan/blog-graphql-api/repository"
)

func ValidateCreateLike(data model.CreateLikeRequest, userID int) (*model.Article, error) {
	user, err := repository.GetUser(model.User{
		ID: userID,
	})
	if err != nil {
		return nil, fmt.Errorf("error while querying db. Err: %s", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user doesn't exist")
	}

	article, err := repository.GetArticle(model.Article{
		ID: data.ArticleID,
	})
	if err != nil {
		return nil, fmt.Errorf("error while querying db. Err: %s", err)
	}
	if article == nil {
		return nil, fmt.Errorf("article doesn't exist")
	}

	like, err := repository.GetLike(model.Like{
		ArticleID: data.ArticleID,
		UserID:    userID,
	})
	if err != nil {
		return nil, fmt.Errorf("error while querying db. Err: %s", err)
	}
	if like != nil {
		return nil, fmt.Errorf("user cant like more than one")
	}

	return article, nil
}
