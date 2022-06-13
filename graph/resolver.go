package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	healthCheck string
}

func NewResolver() *Resolver {
	return &Resolver{healthCheck: "GraphQL Server is healthy!"}
}
