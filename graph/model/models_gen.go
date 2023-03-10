// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type ApproveCommentRequest struct {
	CommentID int  `json:"comment_id"`
	IsChild   bool `json:"is_child"`
}

type ApproveCommentResponse struct {
	Message string `json:"message"`
}

type Article struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Likes       int       `json:"likes"`
	Comments    int       `json:"comments"`
	IsModerated bool      `json:"is_moderated"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ChildComment struct {
	ID         int       `json:"id"`
	CommentID  int       `json:"comment_id"`
	UserID     int       `json:"user_id"`
	Comment    string    `json:"comment"`
	IsApproved bool      `json:"is_approved"`
	CreatedAt  time.Time `json:"created_at"`
}

type Comment struct {
	ID         int       `json:"id"`
	ArticleID  int       `json:"article_id"`
	UserID     int       `json:"user_id"`
	Comment    string    `json:"comment"`
	IsApproved bool      `json:"is_approved"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateArticleRequest struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	IsModerated bool   `json:"is_moderated"`
}

type CreateArticleResponse struct {
	Message string   `json:"message"`
	Article *Article `json:"article"`
}

type CreateChildCommentRequest struct {
	CommentID int    `json:"comment_id"`
	Comment   string `json:"comment"`
}

type CreateChildCommentResponse struct {
	Message      string        `json:"message"`
	ChildComment *ChildComment `json:"child_comment"`
}

type CreateCommentRequest struct {
	ArticleID int    `json:"article_id"`
	Comment   string `json:"comment"`
}

type CreateCommentResponse struct {
	Message string   `json:"message"`
	Comment *Comment `json:"comment"`
}

type CreateLikeRequest struct {
	ArticleID int `json:"article_id"`
}

type CreateLikeResponse struct {
	Message string `json:"message"`
	Like    *Like  `json:"like"`
}

type DeleteArticleRequest struct {
	ArticleID int `json:"article_id"`
}

type DeleteArticleResponse struct {
	Message string `json:"message"`
}

type GetArticleRequest struct {
	ArticleID int `json:"article_id"`
}

type GetArticleResult struct {
	Message string   `json:"message"`
	Article *Article `json:"article"`
}

type GetArticlesRequest struct {
	Search  *string        `json:"search"`
	OrderBy ArticleOrderBy `json:"order_by"`
	SortBy  ArticleSortBy  `json:"sort_by"`
}

type GetArticlesResult struct {
	Message  string     `json:"message"`
	Articles []*Article `json:"articles"`
}

type GetCommentRequest struct {
	CommentID int `json:"comment_id"`
}

type GetCommentResult struct {
	Message string          `json:"message"`
	Comment *Comment        `json:"comment"`
	Child   []*ChildComment `json:"child"`
}

type GetCommentsRequest struct {
	ArticleID int            `json:"article_id"`
	OrderBy   CommentOrderBy `json:"order_by"`
	SortBy    CommentSortBy  `json:"sort_by"`
}

type GetCommentsResult struct {
	Message  string     `json:"message"`
	Comments []*Comment `json:"comments"`
}

type GetLikesRequest struct {
	ArticleID int `json:"article_id"`
}

type GetLikesResult struct {
	Message string  `json:"message"`
	Likes   []*Like `json:"likes"`
}

type Like struct {
	ID        int       `json:"id"`
	ArticleID int       `json:"article_id"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string            `json:"message"`
	User    *UserDataResponse `json:"user"`
}

type Pagination struct {
	TotalPage int `json:"totalPage"`
	TotalData int `json:"totalData"`
	PageSize  int `json:"pageSize"`
	Page      int `json:"page"`
}

type RegisterDataResponse struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Fullname  string `json:"fullname"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type RegisterUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

type RegisterUserResponse struct {
	Message string                `json:"message"`
	User    *RegisterDataResponse `json:"user"`
}

type UpdateArticleRequest struct {
	ArticleID   int     `json:"article_id"`
	Title       *string `json:"title"`
	Content     *string `json:"content"`
	IsModerated *bool   `json:"is_moderated"`
}

type UpdateArticleResponse struct {
	Message string   `json:"message"`
	Article *Article `json:"article"`
}

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	IsActive  bool      `json:"is_active"`
	Token     *string   `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserDataResponse struct {
	ID        int     `json:"id"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	Fullname  string  `json:"fullname"`
	IsActive  bool    `json:"is_active"`
	Token     *string `json:"token"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type VerifyUserRequest struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}

type VerifyUserResponse struct {
	Message string            `json:"message"`
	User    *UserDataResponse `json:"user"`
}

type ArticleOrderBy string

const (
	ArticleOrderByID        ArticleOrderBy = "id"
	ArticleOrderByCreatedAt ArticleOrderBy = "created_at"
	ArticleOrderByUserID    ArticleOrderBy = "user_id"
	ArticleOrderByComments  ArticleOrderBy = "comments"
	ArticleOrderByLikes     ArticleOrderBy = "likes"
)

var AllArticleOrderBy = []ArticleOrderBy{
	ArticleOrderByID,
	ArticleOrderByCreatedAt,
	ArticleOrderByUserID,
	ArticleOrderByComments,
	ArticleOrderByLikes,
}

func (e ArticleOrderBy) IsValid() bool {
	switch e {
	case ArticleOrderByID, ArticleOrderByCreatedAt, ArticleOrderByUserID, ArticleOrderByComments, ArticleOrderByLikes:
		return true
	}
	return false
}

func (e ArticleOrderBy) String() string {
	return string(e)
}

func (e *ArticleOrderBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ArticleOrderBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ArticleOrderBy", str)
	}
	return nil
}

func (e ArticleOrderBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ArticleSortBy string

const (
	ArticleSortByAsc  ArticleSortBy = "asc"
	ArticleSortByDesc ArticleSortBy = "desc"
)

var AllArticleSortBy = []ArticleSortBy{
	ArticleSortByAsc,
	ArticleSortByDesc,
}

func (e ArticleSortBy) IsValid() bool {
	switch e {
	case ArticleSortByAsc, ArticleSortByDesc:
		return true
	}
	return false
}

func (e ArticleSortBy) String() string {
	return string(e)
}

func (e *ArticleSortBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ArticleSortBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ArticleSortBy", str)
	}
	return nil
}

func (e ArticleSortBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CommentOrderBy string

const (
	CommentOrderByID        CommentOrderBy = "id"
	CommentOrderByCreatedAt CommentOrderBy = "created_at"
)

var AllCommentOrderBy = []CommentOrderBy{
	CommentOrderByID,
	CommentOrderByCreatedAt,
}

func (e CommentOrderBy) IsValid() bool {
	switch e {
	case CommentOrderByID, CommentOrderByCreatedAt:
		return true
	}
	return false
}

func (e CommentOrderBy) String() string {
	return string(e)
}

func (e *CommentOrderBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CommentOrderBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CommentOrderBy", str)
	}
	return nil
}

func (e CommentOrderBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CommentSortBy string

const (
	CommentSortByAsc  CommentSortBy = "asc"
	CommentSortByDesc CommentSortBy = "desc"
)

var AllCommentSortBy = []CommentSortBy{
	CommentSortByAsc,
	CommentSortByDesc,
}

func (e CommentSortBy) IsValid() bool {
	switch e {
	case CommentSortByAsc, CommentSortByDesc:
		return true
	}
	return false
}

func (e CommentSortBy) String() string {
	return string(e)
}

func (e *CommentSortBy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CommentSortBy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CommentSortBy", str)
	}
	return nil
}

func (e CommentSortBy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
