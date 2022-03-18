package main

import (
	"log"
	"net/http"

	"github.com/outcaste-io/gqlgen/example/scalars"
	"github.com/outcaste-io/gqlgen/graphql/handler"
	"github.com/outcaste-io/gqlgen/graphql/playground"
)

func main() {
	http.Handle("/", playground.Handler("Starwars", "/query"))
	http.Handle("/query", handler.NewDefaultServer(scalars.NewExecutableSchema(scalars.Config{Resolvers: &scalars.Resolver{}})))

	log.Fatal(http.ListenAndServe(":8084", nil))
}
