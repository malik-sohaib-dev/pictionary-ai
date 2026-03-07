package repo

import (
	"fmt"
	"time"

	"github.com/malik-sohaib-dev/pictionary-ai/backend/types"
)

type RoomRepository interface {
	Create(room types.Room) error
	GetAll() ([]types.Room, error)
	GetById(id string) (types.Room, error)
	UpdateById(id string, room types.Room) (types.Room, error)
}

type RoomRepo struct {
	rooms map[string]types.Room
}

// Room Repo Constructor
func NewRoomRepo() *RoomRepo {
	return &RoomRepo{
		rooms: make(map[string]types.Room),
	}
}

// Create Room
func (r *RoomRepo) Create(room types.Room) error {
	// Latency, just for fun
	time.Sleep(100 * time.Millisecond)

	r.rooms[room.RoomId] = room

	return nil
}

// Get All rooms
func (r *RoomRepo) GetAll() ([]types.Room, error) {
	// Latency, just for fun
	time.Sleep(100 * time.Millisecond)

	rooms := make([]types.Room, 0)

	for _, room := range r.rooms {
		rooms = append(rooms, room)
	}

	return rooms, nil
}

// Get room by id
func (r *RoomRepo) GetById(id string) (types.Room, error) {
	room, exists := r.rooms[id]

	if !exists {
		return types.Room{}, fmt.Errorf("Room not found for id: %s", id)
	}

	return room, nil
}

// Patch room by id
func (r *RoomRepo) UpdateById(id string, room types.Room) (types.Room, error) {
	r.rooms[id] = room

	return r.rooms[id], nil
}