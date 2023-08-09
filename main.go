package main

import (
	"database/sql"
	"go-task/graph"
	"go-task/token"
	"go-task/util"
	"log"
	"net/http"
	"os"

	repoAuth "go-task/internal/auths/repository/postgres"
	repoLabel "go-task/internal/labels/repository/postgres"
	repoTask "go-task/internal/tasks/repository/postgres"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	useAuth "go-task/internal/auths/usecase"
	useLabel "go-task/internal/labels/usecase"
	useTask "go-task/internal/tasks/usecase"
)

const defaultPort = "8080"

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("Cannot set config: %v", err.Error())
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal(err)
	}
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
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Auth:  auths,
		Label: label,
		Task:  task,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
