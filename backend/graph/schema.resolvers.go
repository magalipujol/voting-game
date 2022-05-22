package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/graph/generated"
	"api/graph/model"
	"context"
	"fmt"
)

var players = map[string]*model.Player{}
var rooms = map[string]*model.Room{}

func (r *mutationResolver) CreateRoom(ctx context.Context, name string) (*model.Room, error) {
	room := &model.Room{
		ID:      fmt.Sprintf("%d", len(rooms)+1),
		Players: []*model.Player{},
		Votes:   []*model.Vote{},
		Voting:  false,
		Options: []*model.Option{},
	}
	rooms[room.ID] = room
	return room, nil
}

func (r *mutationResolver) CreatePlayer(ctx context.Context, name string) (*model.Player, error) {
	player := &model.Player{
		ID:   fmt.Sprintf("%d", len(players)+1),
		Name: name,
	}
	players[player.ID] = player
	return player, nil
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
	delete(players, id)
	return nil, nil
}

func (r *mutationResolver) DeleteRoom(ctx context.Context, id string) (*model.Room, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Rooms(ctx context.Context) ([]*model.Room, error) {
	roomsArray := make([]*model.Room, 0, len(rooms))
	for _, room := range rooms {
		roomsArray = append(roomsArray, room)
	}
	return roomsArray, nil
}

func (r *queryResolver) Room(ctx context.Context, id string) (*model.Room, error) {
	room, ok := rooms[id]
	if !ok {
		return nil, fmt.Errorf("room not found")
	}
	return room, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
