package graph

import (
	"api/graph/model"
	"api/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PlayerRepository repository.PlayerRepository
	RoomRepository   repository.RoomRepository
	RoomObservers    map[string]chan *model.Room
}
