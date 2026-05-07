package main

import (
	"graphqlGO/config"
	"graphqlGO/graph"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
)


func main() {
	config.ConnectDB()
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}),
	)
http.Handle("/graphql", srv)
log.Fatal(http.ListenAndServe(":8086", nil))	
}