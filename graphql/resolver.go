package graphql

import (
	"context"

	"github.com/dannyrsu/social-media-graphql/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateMessage(ctx context.Context, input models.NewMessage) (*models.Message, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Message(ctx context.Context, userID *int) (*models.Message, error) {
	panic("not implemented")
}
func (r *queryResolver) Messages(ctx context.Context) ([]*models.Message, error) {
	panic("not implemented")
}
