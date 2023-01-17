package validation

import (
	"fmt"

	"github.com/muhadyan/blog-graphql-api/graph/model"
)

func ValidateCreateLike(data model.CreateLikeRequest) error {
	if data.ArticleID <= 0 {
		return fmt.Errorf("article id cant be empty")
	}

	return nil
}

func ValidateGetLikes(request model.GetLikesRequest) error {
	if request.ArticleID <= 0 {
		return fmt.Errorf("article id cant be empty")
	}

	return nil
}
