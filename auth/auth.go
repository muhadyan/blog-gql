package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/muhadyan/blog-graphql-api/graph/model"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"userCtxKey"}

const (
	bearerString = "Bearer "
	IDClaimKey   = "id"
)

var (
	ErrUnauthorized = &gqlerror.Error{
		Message: "invalid authorization",

		Extensions: map[string]interface{}{
			"code": http.StatusUnauthorized,
		},
	}
)

type ErrorResponse struct {
	Errors []*gqlerror.Error `json:"errors"`
}

type User struct {
	claims jwt.MapClaims
}

func GenerateJwt(data *model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().AddDate(0, 0, 1).Unix()
	claims["id"] = data.ID
	claims["username"] = data.Username
	claims["email"] = data.Email
	claims["fullname"] = data.Fullname
	claims["created_at"] = data.CreatedAt

	key := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf(fmt.Sprintf("error while generate jwt. Err: %s", err.Error()))
	}

	return tokenString, nil
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			authValue := r.Header.Get("authorization")
			// there is no token, continue the process
			if authValue == "" {
				next.ServeHTTP(w, r)
				return
			}

			if !strings.Contains(authValue, bearerString) {
				next.ServeHTTP(w, r)
				return
			}

			tokenString := strings.Replace(authValue, bearerString, "", 1)

			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_SECRET_KEY")), nil
			})
			if err != nil {
				unauthorized(w, err.Error())
				return
			}

			claim, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				unauthorized(w, "invalid authorization")
				return
			}

			// put it in context
			ctx := RegisterClaimsToContext(r.Context(), claim)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func unauthorized(w http.ResponseWriter, messages ...string) {
	response := ErrorResponse{Errors: []*gqlerror.Error{ErrUnauthorized}}

	for _, message := range messages {
		response.Errors = append(response.Errors, gqlerror.Errorf((message)))
	}

	json, _ := json.Marshal(response)
	http.Error(w, string(json), http.StatusForbidden)
}

func RegisterClaimsToContext(ctx context.Context, claim map[string]interface{}) context.Context {
	user := &User{
		claims: claim,
	}

	return context.WithValue(ctx, userCtxKey, user)
}

func FromContext(ctx context.Context) *User {
	user, exist := ctx.Value(userCtxKey).(*User)
	if !exist {
		return nil
	}

	return user
}

func (user *User) Authorize() bool {
	return user != nil
}

func (user *User) GetUserID() int {
	if user == nil {
		return 0
	}

	userID, ok := user.claims[IDClaimKey].(float64)
	if !ok {
		return 0
	}

	return int(userID)
}
