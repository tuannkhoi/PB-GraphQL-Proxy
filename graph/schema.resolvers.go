package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/tuannkhoi/PB-GraphQL-Proxy/graph/generated"
	"github.com/tuannkhoi/PB-GraphQL-Proxy/graph/model"
)

// CreateNewUser is the resolver for the createNewUser field.
func (r *mutationResolver) CreateNewUser(_ context.Context, input *model.NewUser) (*model.User, error) {
	newUser := &model.User{
		Username:    input.Username,
		Password:    input.Password,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
	}

	r.users = append(r.users, newUser)
	return newUser, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(_ context.Context) ([]*model.User, error) {
	usersToReturn := r.users
	// hide every user's password
	for i := range usersToReturn {
		usersToReturn[i].Password = "TOP SECRET"
	}

	return usersToReturn, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
