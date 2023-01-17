package service

import (
	"context"
	"fmt"

	"github.com/muhadyan/blog-graphql-api/graph/model"
	"github.com/muhadyan/blog-graphql-api/repository"
	"github.com/muhadyan/blog-graphql-api/service/helper"
)

type ChildComment interface {
	Create(ctx context.Context, data model.CreateChildCommentRequest, userID int) (*model.CreateChildCommentResponse, error)
}

type childCommentCtx struct{}

func NewChildCommentService() ChildComment {
	return &childCommentCtx{}
}

func (s *childCommentCtx) Create(ctx context.Context, data model.CreateChildCommentRequest, userID int) (*model.CreateChildCommentResponse, error) {
	article, comment, err := helper.ValidateCreateChildComment(data, userID)
	if err != nil {
		return nil, err
	}

	cc := model.ChildComment{
		CommentID:  data.CommentID,
		UserID:     userID,
		Comment:    data.Comment,
		IsApproved: false,
	}
	if !article.IsModerated {
		cc.IsApproved = true

		_, err = repository.UpdateArticle(model.Article{
			ID:       comment.ArticleID,
			Comments: article.Comments + 1,
		})
		if err != nil {
			return nil, err
		}
	}

	childComment, err := repository.InsertChildComment(cc)
	if err != nil {
		return nil, err
	}

	go func() {
		commentUser, err := repository.GetUser(model.User{
			ID: comment.UserID,
		})
		if err != nil {
			return
		}
		childCommentUser, err := repository.GetUser(model.User{
			ID: userID,
		})
		if err != nil {
			return
		}

		templateEmail := "./template/child_comment.html"
		sendRequest := helper.SendMailModel{
			SendTo:          commentUser.Email,
			CommentUser:     commentUser.Fullname,
			ArticleName:     article.Title,
			ChilCommentUser: childCommentUser.Username,
		}
		subject := fmt.Sprintf("Comment Of Your Comment In %s!", article.Title)
		helper.SendMail(templateEmail, sendRequest, subject)
	}()

	res := model.CreateChildCommentResponse{
		Message: "Success",
		ChildComment: &model.ChildComment{
			ID:         childComment.ID,
			CommentID:  childComment.CommentID,
			UserID:     childComment.UserID,
			Comment:    childComment.Comment,
			IsApproved: childComment.IsApproved,
			CreatedAt:  childComment.CreatedAt,
		},
	}

	return &res, nil
}
