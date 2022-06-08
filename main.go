package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

var queryFields = graphql.Fields{
	"hello": &graphql.Field{
		Type:    graphql.String,
		Resolve: handleHello,
	},
}

func handleHello(p graphql.ResolveParams) (interface{}, error) {
	return "world", nil
}

func getSchema(fields *graphql.Fields) *graphql.Schema {
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: *fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return &schema
}

func handleQuery(w http.ResponseWriter, r *http.Request) {
	schema := getSchema(&queryFields)
	result := graphql.Do(graphql.Params{
		Schema:        *schema,
		RequestString: r.URL.Query().Get("query"),
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/graphql", handleQuery)
}
