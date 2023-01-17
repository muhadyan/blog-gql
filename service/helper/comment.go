package helper

import (
	"fmt"

	"github.com/muhadyan/blog-graphql-api/graph/model"
	"github.com/muhadyan/blog-graphql-api/repository"
)

func ValidateCreateComment(data model.CreateCommentRequest, userID int) (*model.Article, error) {
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

	return article, nil
}

func ValidateApproveComment(data model.ApproveCommentRequest, userID int) (*model.Article, error) {
	childComment := new(model.ChildComment)
	commentID := data.CommentID

	user, err := repository.GetUser(model.User{
		ID: userID,
	})
	if err != nil {
		return nil, fmt.Errorf("error while querying db. Err: %s", err)
	}
	if user == nil {
		return nil, fmt.Errorf("user doesn't exist")
	}

	if data.IsChild {
		childComment, err = repository.GetChildComment(model.ChildComment{
			ID: data.CommentID,
		})
		if err != nil {
			return nil, fmt.Errorf("error while querying db. Err: %s", err)
		}
		if childComment == nil {
			return nil, fmt.Errorf("comment doesn't exist")
		}

		commentID = childComment.CommentID
	}

	comment, err := repository.GetComment(model.Comment{
		ID: commentID,
	})
	if err != nil {
		return nil, fmt.Errorf("error while querying db. Err: %s", err)
	}
	if comment == nil {
		return nil, fmt.Errorf("comment doesn't exist")
	}

	article, err := repository.GetArticle(model.Article{
		ID:     comment.ArticleID,
		UserID: userID,
	})
	if err != nil {
		return nil, fmt.Errorf("error while querying db. Err: %s", err)
	}
	if article == nil {
		return nil, fmt.Errorf("user cant approve this comment")
	}

	return article, nil
}
