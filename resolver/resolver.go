package resolver

import (
	"context"
	"errors"

	"github.com/muhadyan/blog-graphql-api/graph/model"
)

var (
	errUnauthorized = errors.New("access denied")
)

type userServiceProvider interface {
	Register(ctx context.Context, data model.RegisterUserRequest) (*model.RegisterUserResponse, error)
	Verify(ctx context.Context, data model.VerifyUserRequest) (*model.VerifyUserResponse, error)
	Login(ctx context.Context, data model.LoginRequest) (*model.LoginResponse, error)
}

type articleServiceProvider interface {
	Create(ctx context.Context, data model.CreateArticleRequest, userID int) (*model.CreateArticleResponse, error)
	Update(ctx context.Context, data model.UpdateArticleRequest, userID int) (*model.UpdateArticleResponse, error)
	Delete(ctx context.Context, data model.DeleteArticleRequest, userID int) (*model.DeleteArticleResponse, error)
	View(ctx context.Context, request model.GetArticlesRequest) (*model.GetArticlesResult, error)
	Get(ctx context.Context, request model.GetArticleRequest) (*model.GetArticleResult, error)
}

type commentServiceProvider interface {
	Create(ctx context.Context, data model.CreateCommentRequest, userID int) (*model.CreateCommentResponse, error)
	Approve(ctx context.Context, data model.ApproveCommentRequest, userID int) (*model.ApproveCommentResponse, error)
	View(ctx context.Context, request model.GetCommentsRequest) (*model.GetCommentsResult, error)
	Get(ctx context.Context, request model.GetCommentRequest) (*model.GetCommentResult, error)
}

type childCommentServiceProvider interface {
	Create(ctx context.Context, data model.CreateChildCommentRequest, userID int) (*model.CreateChildCommentResponse, error)
}

type likeServiceProvider interface {
	Create(ctx context.Context, data model.CreateLikeRequest, userID int) (*model.CreateLikeResponse, error)
	View(ctx context.Context, request model.GetLikesRequest) (*model.GetLikesResult, error)
}

// Resolver main struct contain all queries related
type Resolver struct {
	userService         userServiceProvider
	articleService      articleServiceProvider
	commentService      commentServiceProvider
	childCommentService childCommentServiceProvider
	likeService         likeServiceProvider
}

// Options parameter used to create resolver.
type Options struct {
	UserService         userServiceProvider
	ArticleService      articleServiceProvider
	CommentService      commentServiceProvider
	ChildCommentService childCommentServiceProvider
	LikeService         likeServiceProvider
}

// New creating new resolver.
func New(args Options) *Resolver {
	return &Resolver{
		userService:         args.UserService,
		articleService:      args.ArticleService,
		commentService:      args.CommentService,
		childCommentService: args.ChildCommentService,
		likeService:         args.LikeService,
	}
}
