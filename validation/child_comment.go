package validation

import (
	"fmt"

	"github.com/muhadyan/blog-graphql-api/graph/model"
)

func ValidateCreateChildComment(data model.CreateChildCommentRequest) error {
	if data.CommentID <= 0 {
		return fmt.Errorf("comment id cant be empty")
	}

	if data.Comment == "" {
		return fmt.Errorf("comment cant be empty")
	}

	return nil
}
