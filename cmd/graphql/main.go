package main

import (
	"fmt"
	"log"
	"net/http"

	graphqlHandler "project/internal/interfaces/graphql"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	schema := graphql.MustParseSchema(graphqlHandler.GetSchema())
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
	http.Handle("/graphql", h)
	fmt.Println("GraphQL server listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
