package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/library-server/app/middleware"
	"github.com/library-server/internals/db"
	"github.com/library-server/internals/object"

	"github.com/go-chi/chi/v5"
)

const defaultPort = "8080"

func main() {
	r := chi.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	database := db.NewDatabase()
	socket := socket.NewSocket(database)
	objectStorage, err := object.NewObjectStorage()
	if err != nil {
		panic("Failed to initialize object storage")
	}
	defer socket.Close()
	resolver := resolvers.NewResolver(database, socket, objectStorage)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	r.Use(middleware.AuthenticationMiddleware(database))
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)
	r.Handle("/socket.io/", socket.Server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
