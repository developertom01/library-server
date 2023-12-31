package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/developertom01/library-server/app/graphql/dataloader"
	"github.com/developertom01/library-server/app/graphql/resolvers"
	"github.com/developertom01/library-server/app/middleware"
	"github.com/developertom01/library-server/app/socket"
	"github.com/developertom01/library-server/generated"
	"github.com/developertom01/library-server/internals/db"
	"github.com/developertom01/library-server/internals/object"

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
	dataLoader := dataloader.NewDataLoader(database)
	resolver := resolvers.NewResolver(database, socket, objectStorage)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	r.Use(middleware.DataLoaderMiddleware(dataLoader))
	r.Use(middleware.AuthenticationMiddleware(database))
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)
	r.Handle("/socket.io/", socket.Server)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
