package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	//"fmt"

	"github.com/aleksandarmilanovic/gqlgen-todos/graph/generated"
	"github.com/aleksandarmilanovic/gqlgen-todos/graph/model"
	"github.com/aleksandarmilanovic/gqlgen-todos/database"
)

var db = database.Connect()

func (r *mutationResolver) CreateLion(ctx context.Context, input model.NewLion) (*model.Lion, error) {
	return db.Save(&input), nil
}

func (r *queryResolver) Lion(ctx context.Context, id string) (*model.Lion, error) {
	return db.GetByID(id),nil
}

func (r *queryResolver) Lions(ctx context.Context) ([]*model.Lion, error) {
	return db.GetAllLions(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

