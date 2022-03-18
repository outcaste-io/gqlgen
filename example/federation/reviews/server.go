//go:generate go run ../../../testdata/gqlgen.go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/outcaste-io/gqlgen/example/federation/reviews/graph"
	"github.com/outcaste-io/gqlgen/example/federation/reviews/graph/generated"
	"github.com/outcaste-io/gqlgen/graphql/handler"
	"github.com/outcaste-io/gqlgen/graphql/handler/debug"
	"github.com/outcaste-io/gqlgen/graphql/playground"
)

const defaultPort = "4003"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	srv.Use(&debug.Tracer{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
