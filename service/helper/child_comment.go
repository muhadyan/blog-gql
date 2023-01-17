package helper

import (
	"fmt"

	"github.com/muhadyan/blog-graphql-api/graph/model"
	"github.com/muhadyan/blog-graphql-api/repository"
)

func ValidateCreateChildComment(data model.CreateChildCommentRequest, userID int) (*model.Article, *model.Comment, error) {
	user, err := repository.GetUser(model.User{
		ID: userID,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("error while querying db. Err: %s", err)
	}
	if user == nil {
		return nil, nil, fmt.Errorf("user doesn't exist")
	}

	comment, err := repository.GetComment(model.Comment{
		ID: data.CommentID,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("error while querying db. Err: %s", err)
	}
	if comment == nil {
		return nil, nil, fmt.Errorf("comment doesn't exist")
	}

	article, err := repository.GetArticle(model.Article{
		ID: comment.ArticleID,
	})
	if err != nil {
		return nil, nil, fmt.Errorf("error while querying db. Err: %s", err)
	}
	if article == nil {
		return nil, nil, fmt.Errorf("article doesn't exist")
	}

	return article, comment, nil
}
