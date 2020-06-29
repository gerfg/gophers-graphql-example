package main

import (
	"log"
	"net/http"

	"github.com/gerfg/gophers-graphql/resolver.go"
	"github.com/gerfg/gophers-graphql/schema"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

func main() {
	schemaString := schema.ReadSchemaFile("schema.graphql")
	schema := graphql.MustParseSchema(schemaString, &resolver.Resolver{})
	http.Handle("/", &relay.Handler{Schema: schema})
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
