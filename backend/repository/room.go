package repository

import (
	"api/graph/model"
	"context"
	"fmt"
	"time"
)

type RoomRepository interface {
	List(ctx context.Context) ([]*model.Room, error)
	Get(ctx context.Context, id string) (*model.Room, error)
	Create(ctx context.Context, name string) (*model.Room, error)
	Delete(ctx context.Context, id string) (*model.Room, error)
	Join(ctx context.Context, player *model.Player, roomID string) (*model.Room, error)
	StartGame(ctx context.Context, roomID string) (*model.Room, error)
	Vote(ctx context.Context, playerID string, option string) (*model.Room, error)
	Leave(ctx context.Context, id string) (*model.Room, error)
}

type roomRepositoryInMemory struct {
	counter int
	rooms   map[string]*model.Room
}

func NewRoomRepository() RoomRepository {
	return &roomRepositoryInMemory{
		rooms: map[string]*model.Room{},
	}
}

func (r *roomRepositoryInMemory) List(ctx context.Context) ([]*model.Room, error) {
	roomsArray := make([]*model.Room, 0, len(r.rooms))
	for _, room := range r.rooms {
		roomsArray = append(roomsArray, room)
	}
	return roomsArray, nil
}

func (r *roomRepositoryInMemory) Get(ctx context.Context, id string) (*model.Room, error) {
	room, ok := r.rooms[id]
	if !ok {
		return nil, fmt.Errorf("room not found")
	}
	return room, nil
}

func (r *roomRepositoryInMemory) Create(ctx context.Context, name string) (*model.Room, error) {
	room := &model.Room{
		ID:      fmt.Sprintf("%d", r.counter),
		Players: []*model.Player{},
		Votes:   []*model.Vote{},
		Voting:  false,
		Options: []*model.Option{},
	}
	r.counter++
	r.rooms[room.ID] = room
	return room, nil
}

func (r *roomRepositoryInMemory) Delete(ctx context.Context, id string) (*model.Room, error) {
	delete(r.rooms, id)
	return nil, nil
}

func (r *roomRepositoryInMemory) Join(ctx context.Context, player *model.Player, roomID string) (*model.Room, error) {
	playersInRoom := r.rooms[roomID].Players
	for _, playerInRoom := range playersInRoom {
		if playerInRoom.ID == player.ID {
			return nil, fmt.Errorf("player already in room")
		}
	}
	r.rooms[roomID].Players = append(r.rooms[roomID].Players, player)
	return r.rooms[roomID], nil
}

func (r *roomRepositoryInMemory) StartGame(ctx context.Context, roomID string) (*model.Room, error) {
	room := r.rooms[roomID]
	if room.Voting {
		return nil, fmt.Errorf("game already started")
	}
	room.Voting = true
	go func() {
		intervalCheckTime := time.Second
		votingFinishTime := time.Now().Add(10 * time.Second)
		for votingFinishTime.Before(time.Now()) {
			if len(room.Votes) == len(room.Players) {
				break
			}
			time.Sleep(intervalCheckTime)
		}
		room.Voting = false
	}()
	return room, nil
}

func (r *roomRepositoryInMemory) Vote(ctx context.Context, playerID string, option string) (*model.Room, error) {
	var player *model.Player
	var room *model.Room

	for _, searchingRoom := range r.rooms {
		for _, playerInRoom := range searchingRoom.Players {
			if playerInRoom.ID == playerID {
				player = playerInRoom
				room = searchingRoom
				break
			}
		}
	}
	if player == nil {
		return nil, fmt.Errorf("player not found")
	}

	for _, vote := range room.Votes {
		if vote.Player.ID == player.ID {
			return nil, fmt.Errorf("player already voted")
		}
	}

	room.Votes = append(room.Votes, &model.Vote{
		Player: player,
		Option: &model.Option{
			Value: option,
		},
		Room:      room,
		TimeStamp: time.Now().UTC().Format(time.RFC3339),
	})

	return room, nil
}

func (r *roomRepositoryInMemory) Leave(ctx context.Context, id string) (*model.Room, error) {
	panic(fmt.Errorf("not implemented"))
}
