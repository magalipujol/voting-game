package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/graph/generated"
	"api/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateRoom(ctx context.Context, name string) (*model.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePlayer(ctx context.Context, name string) (*model.Player, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) JoinRoom(ctx context.Context, id string) (*model.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) StartGame(ctx context.Context, roomID string) (*model.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Vote(ctx context.Context, playerID string, option string) (*model.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LeaveRoom(ctx context.Context, id string) (*model.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePlayer(ctx context.Context, id string) (*model.Player, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteRoom(ctx context.Context, id string) (*model.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Rooms(ctx context.Context) ([]*model.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Room(ctx context.Context, id string) (*model.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
