package main

import (
	"log"
	"net/http"
	"os"

	"github.com/outcaste-io/gqlgen/graphql/playground"

	extension "github.com/outcaste-io/gqlgen/example/type-system-extension"
	"github.com/outcaste-io/gqlgen/graphql/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", handler.NewDefaultServer(
		extension.NewExecutableSchema(
			extension.Config{
				Resolvers: extension.NewRootResolver(),
				Directives: extension.DirectiveRoot{
					EnumLogging:   extension.EnumLogging,
					FieldLogging:  extension.FieldLogging,
					InputLogging:  extension.InputLogging,
					ObjectLogging: extension.ObjectLogging,
					ScalarLogging: extension.ScalarLogging,
					UnionLogging:  extension.UnionLogging,
				},
			},
		),
	))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
