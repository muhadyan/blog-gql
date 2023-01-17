package validation

import (
	"fmt"

	"github.com/muhadyan/blog-graphql-api/graph/model"
)

func ValidateCreateComment(data model.CreateCommentRequest) error {
	if data.ArticleID <= 0 {
		return fmt.Errorf("article id cant be empty")
	}

	if data.Comment == "" {
		return fmt.Errorf("comment cant be empty")
	}

	return nil
}

func ValidateApproveComment(data model.ApproveCommentRequest) error {
	if data.CommentID <= 0 {
		return fmt.Errorf("comment id cant be empty")
	}

	return nil
}

func ValidateGetComments(request model.GetCommentsRequest) error {
	if request.ArticleID <= 0 {
		return fmt.Errorf("article id cant be empty")
	}

	return nil
}

func ValidateGetComment(request model.GetCommentRequest) error {
	if request.CommentID <= 0 {
		return fmt.Errorf("comment id cant be empty")
	}

	return nil
}