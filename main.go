package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/muhadyan/blog-graphql-api/auth"
	"github.com/muhadyan/blog-graphql-api/config"
	"github.com/muhadyan/blog-graphql-api/graph/generated"
	logUtil "github.com/muhadyan/blog-graphql-api/log"
	"github.com/muhadyan/blog-graphql-api/resolver"
	"github.com/muhadyan/blog-graphql-api/service"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	_, err = config.InitDB()
	if err != nil {
		log.Fatalf("Error db connection. Err: %s", err)
	}

	svc := initService()
	resolver := initResolver(svc)

	schema := generated.Config{
		Resolvers: resolver,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(schema))

	router := chi.NewRouter()
	startTime := time.Now()
	router.Use(auth.Middleware())
	router.Use(logUtil.MiddlewareLog(startTime))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

type serviceProvider struct {
	user         service.User
	article      service.Article
	comment      service.Comment
	childComment service.ChildComment
	like         service.Like
}

func initService() serviceProvider {
	return serviceProvider{
		user:         service.NewUserService(),
		article:      service.NewArticleService(),
		comment:      service.NewCommentService(),
		childComment: service.NewChildCommentService(),
		like:         service.NewLikeService(),
	}
}

func initResolver(svc serviceProvider) *resolver.Resolver {
	return resolver.New(resolver.Options{
		UserService:         svc.user,
		ArticleService:      svc.article,
		CommentService:      svc.comment,
		ChildCommentService: svc.childComment,
		LikeService:         svc.like,
	})
}
