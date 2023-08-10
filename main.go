package main

import (
	"database/sql"
	"go-task/graph"
	"go-task/pkg/generated"
	"go-task/token"
	"go-task/util"
	"log"
	"os"

	repoAuth "go-task/internal/auths/repository/postgres"
	repoLabel "go-task/internal/labels/repository/postgres"
	repoTask "go-task/internal/tasks/repository/postgres"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	useAuth "go-task/internal/auths/usecase"
	useLabel "go-task/internal/labels/usecase"
	useTask "go-task/internal/tasks/usecase"
)

const defaultPort = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Cannot set config: %v", err.Error())
	}
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	// r.Use(middlewares.AuthMiddleware(tokenMaker, config))
	r.POST("/query", graphqlHandler(config, tokenMaker))
	r.GET("/", playgroundHandler())
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	r.Run()
}

// Defining the Graphql handler
func graphqlHandler(config util.Config, tokenMaker token.Maker) gin.HandlerFunc {

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("Cannot connect to db:")
	}
	authRepo := repoAuth.NewAuthentication(conn)
	labelRepo := repoLabel.NewLabel(conn)
	taskLabel := repoTask.NewStore(conn)
	auths := useAuth.NewAuthusecase(authRepo, config, tokenMaker)
	label := useLabel.NewLabelusecase(labelRepo, config, tokenMaker)
	task := useTask.NewTaskusecase(taskLabel, config, tokenMaker)
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Auth:  auths,
		Label: label,
		Task:  task,
	}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// db := store.NewStore()
// 	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	http.Handle("/query", store.WithStore(db, srv))
