package repository

import (
	"api/graph/model"
	"context"
	"fmt"
)

type PlayerRepository interface {
	Create(ctx context.Context, name string) (*model.Player, error)
	Get(ctx context.Context, id string) (*model.Player, error)
	Delete(ctx context.Context, id string) (*model.Player, error)
}

type playerRepositoryInMemory struct {
	players map[string]*model.Player
}

func NewPlayerRepository() PlayerRepository {
	return &playerRepositoryInMemory{
		players: map[string]*model.Player{},
	}
}

func (r *playerRepositoryInMemory) Create(ctx context.Context, name string) (*model.Player, error) {
	player := &model.Player{
		ID:   fmt.Sprintf("%d", len(r.players)+1),
		Name: name,
	}
	r.players[player.ID] = player
	return player, nil
}

func (r *playerRepositoryInMemory) Get(ctx context.Context, id string) (*model.Player, error) {
	player, ok := r.players[id]
	if !ok {
		return nil, fmt.Errorf("player not found")
	}
	return player, nil
}

func (r *playerRepositoryInMemory) Delete(ctx context.Context, id string) (*model.Player, error) {
	delete(r.players, id)
	return nil, nil
}
