package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"minecraft-backend/graph/generated"
	"minecraft-backend/graph/model"
)

func (r *mutationResolver) Register(ctx context.Context, userName string, password string, isAdmin bool) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Authenticate(ctx context.Context, password string, userID string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return r.Resolver.healthCheck, nil
}

func (r *queryResolver) Posts(ctx context.Context, jwt string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
