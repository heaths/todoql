package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/heaths/todoql/graphql"
	"github.com/heaths/todoql/graphql/generated"
)

const (
	defaultPort   = "8080"
	graphUrl      = "/graphql"
	playgroundUrl = "/graphql/play"
)

func main() {
	port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if port == "" {
		port = defaultPort
	}

	resolver, err := graphql.NewResolver(context.Background(), "TODO")
	if err != nil {
		log.Fatalf("failed to initialize resolver: %s", err)
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle(graphUrl, srv)
	http.Handle(playgroundUrl, playground.Handler("GraphQL playground", graphUrl))

	log.Printf("Connect to http://localhost:%s%s for GraphQL playground", port, playgroundUrl)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
