package repository

import (
	"api/graph/model"
	"context"
	"fmt"
)

type PlayerRepository interface {
	CreatePlayer(ctx context.Context, name string) (*model.Player, error)
	DeletePlayer(ctx context.Context, id string) (*model.Player, error)
}

type playerRepositoryInMemory struct {
	players map[string]*model.Player
}

func NewPlayerRepository() PlayerRepository {
	return &playerRepositoryInMemory{
		players: map[string]*model.Player{},
	}
}

func (r *playerRepositoryInMemory) CreatePlayer(ctx context.Context, name string) (*model.Player, error) {
	player := &model.Player{
		ID:   fmt.Sprintf("%d", len(r.players)+1),
		Name: name,
	}
	r.players[player.ID] = player
	return player, nil
}

func (r *playerRepositoryInMemory) DeletePlayer(ctx context.Context, id string) (*model.Player, error) {
	delete(r.players, id)
	return nil, nil
}
