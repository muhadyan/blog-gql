package service

import (
	"context"
	"fmt"

	"github.com/muhadyan/blog-graphql-api/graph/model"
	"github.com/muhadyan/blog-graphql-api/repository"
	"github.com/muhadyan/blog-graphql-api/service/helper"
)

type Comment interface {
	Create(ctx context.Context, data model.CreateCommentRequest, userID int) (*model.CreateCommentResponse, error)
	Approve(ctx context.Context, data model.ApproveCommentRequest, userID int) (*model.ApproveCommentResponse, error)
	View(ctx context.Context, request model.GetCommentsRequest) (*model.GetCommentsResult, error)
	Get(ctx context.Context, request model.GetCommentRequest) (*model.GetCommentResult, error)
}

type commentCtx struct{}

func NewCommentService() Comment {
	return &commentCtx{}
}

func (s *commentCtx) Create(ctx context.Context, data model.CreateCommentRequest, userID int) (*model.CreateCommentResponse, error) {
	article, err := helper.ValidateCreateComment(data, userID)
	if err != nil {
		return nil, err
	}

	c := model.Comment{
		ArticleID:  data.ArticleID,
		UserID:     userID,
		Comment:    data.Comment,
		IsApproved: false,
	}
	if !article.IsModerated {
		c.IsApproved = true

		_, err = repository.UpdateArticle(model.Article{
			ID:       data.ArticleID,
			Comments: article.Comments + 1,
		})
		if err != nil {
			return nil, err
		}
	}

	comment, err := repository.InsertComment(c)
	if err != nil {
		return nil, err
	}

	res := model.CreateCommentResponse{
		Message: "Success",
		Comment: &model.Comment{
			ID:         comment.ID,
			ArticleID:  comment.ArticleID,
			UserID:     comment.UserID,
			Comment:    comment.Comment,
			IsApproved: comment.IsApproved,
			CreatedAt:  comment.CreatedAt,
		},
	}

	return &res, nil
}

func (s *commentCtx) Approve(ctx context.Context, data model.ApproveCommentRequest, userID int) (*model.ApproveCommentResponse, error) {
	article, err := helper.ValidateApproveComment(data, userID)
	if err != nil {
		return nil, err
	}

	if data.IsChild {
		err := repository.UpdateChildComment(model.ChildComment{
			ID:         data.CommentID,
			IsApproved: true,
		})
		if err != nil {
			return nil, err
		}
	} else {
		err := repository.UpdateComment(model.Comment{
			ID:         data.CommentID,
			IsApproved: true,
		})
		if err != nil {
			return nil, err
		}
	}

	_, err = repository.UpdateArticle(model.Article{
		ID:       article.ID,
		Comments: article.Comments + 1,
	})
	if err != nil {
		return nil, err
	}

	res := model.ApproveCommentResponse{
		Message: "Success",
	}

	return &res, nil
}

func (s *commentCtx) View(ctx context.Context, request model.GetCommentsRequest) (*model.GetCommentsResult, error) {
	comments, err := repository.SelectComment(request)
	if err != nil {
		return nil, err
	}

	res := model.GetCommentsResult{
		Message:  "Success",
		Comments: comments,
	}

	return &res, nil
}

func (s *commentCtx) Get(ctx context.Context, request model.GetCommentRequest) (*model.GetCommentResult, error) {
	comment, err := repository.GetComment(model.Comment{
		ID:         request.CommentID,
		IsApproved: true,
	})
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, fmt.Errorf("comment not found")
	}

	childComments, err := repository.GetChildComments(model.ChildComment{
		CommentID:  request.CommentID,
		IsApproved: true,
	})
	if err != nil {
		return nil, err
	}

	res := model.GetCommentResult{
		Message: "Success",
		Comment: comment,
		Child:   childComments,
	}

	return &res, nil
}
