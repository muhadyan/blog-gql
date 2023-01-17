package service

import (
	"context"
	"fmt"

	"github.com/muhadyan/blog-graphql-api/graph/model"
	"github.com/muhadyan/blog-graphql-api/repository"
	"github.com/muhadyan/blog-graphql-api/service/helper"
)

type Like interface {
	Create(ctx context.Context, data model.CreateLikeRequest, userID int) (*model.CreateLikeResponse, error)
	View(ctx context.Context, request model.GetLikesRequest) (*model.GetLikesResult, error)
}

type likeCtx struct{}

func NewLikeService() Like {
	return &likeCtx{}
}

func (s *likeCtx) Create(ctx context.Context, data model.CreateLikeRequest, userID int) (*model.CreateLikeResponse, error) {
	article, err := helper.ValidateCreateLike(data, userID)
	if err != nil {
		return nil, err
	}

	l := model.Like{
		ArticleID: data.ArticleID,
		UserID:    userID,
	}
	like, err := repository.InsertLike(l)
	if err != nil {
		return nil, err
	}

	_, err = repository.UpdateArticle(model.Article{
		ID:    data.ArticleID,
		Likes: article.Likes + 1,
	})
	if err != nil {
		return nil, err
	}

	go func() {
		articleUser, err := repository.GetUser(model.User{
			ID: article.UserID,
		})
		if err != nil {
			return
		}
		likeUser, err := repository.GetUser(model.User{
			ID: userID,
		})
		if err != nil {
			return
		}

		templateEmail := "./template/like.html"
		sendRequest := helper.SendMailModel{
			SendTo:      articleUser.Email,
			ArticleUser: articleUser.Fullname,
			ArticleName: article.Title,
			LikeUser:    likeUser.Username,
		}
		subject := fmt.Sprintf("Like Notification In %s!", article.Title)
		helper.SendMail(templateEmail, sendRequest, subject)
	}()

	res := model.CreateLikeResponse{
		Message: "Success",
		Like: &model.Like{
			ID:        like.ID,
			ArticleID: like.ArticleID,
			UserID:    like.UserID,
			CreatedAt: like.CreatedAt,
		},
	}

	return &res, nil
}

func (s *likeCtx) View(ctx context.Context, request model.GetLikesRequest) (*model.GetLikesResult, error) {
	likes, err := repository.SelectLike(model.Like{
		ArticleID: request.ArticleID,
	})
	if err != nil {
		return nil, err
	}

	res := model.GetLikesResult{
		Message: "Success",
		Likes:   likes,
	}

	return &res, nil
}
