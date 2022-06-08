package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

const PORT uint = 8080

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
		log.Fatalf("Failed to create new schema, error: %v\n", err)
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
		log.Printf("Wrong result, unexpected errors: %v\n", result.Errors)
	}
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/api", handleQuery)
	fmt.Printf("Server running on port %v...\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%v", PORT), nil)
}
