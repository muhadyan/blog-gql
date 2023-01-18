package log

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/muhadyan/blog-graphql-api/auth"
	"github.com/muhadyan/blog-graphql-api/config"
	"github.com/muhadyan/blog-graphql-api/graph/model"
	"github.com/muhadyan/blog-graphql-api/repository"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Log struct {
	Id        primitive.ObjectID `json:"id"`
	Path      string             `json:"path"`
	User      *model.User        `json:"user"`
	Duration  string             `json:"duration"`
	Request   string             `json:"request"`
	Response  string             `json:"response"`
	Timestamp time.Time          `json:"timestamp"`
}

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte
	StartTime  time.Time
	Request    string
}

var logCollection *mongo.Collection = config.GetCollection(config.MongoDB, "logs")

func MiddlewareLog(startTime time.Time) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqRec, err := httputil.DumpRequest(r, true)
			if err != nil {
				http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
				return
			}

			rec := &ResponseRecorder{
				ResponseWriter: w,
				StatusCode:     http.StatusOK,
				StartTime:      startTime,
				Request:        string(reqRec),
			}
			next.ServeHTTP(rec, r)

			defer CreateLog(r.Context(), r, rec)
		})
	}
}

func CreateLog(ctx context.Context, r *http.Request, rec *ResponseRecorder) {
	user := auth.FromContext(ctx)

	userData, _ := repository.GetUser(model.User{
		ID: user.GetUserID(),
	})

	newLog := Log{
		Id:        primitive.NewObjectID(),
		Path:      r.RequestURI,
		User:      userData,
		Duration:  time.Since(rec.StartTime).String(),
		Request:   rec.Request,
		Response:  string(rec.Body),
		Timestamp: time.Now(),
	}

	_, err := logCollection.InsertOne(ctx, newLog)
	if err != nil {
		log.Print(err)
	}
}
