package services

import "github.com/malik-sohaib-dev/pictionary-ai/backend/repo"

type RoomService struct {
	repo repo.RoomRepository
}

// Room service constructor
func NewRoomService (repo repo.RoomRepository) RoomService {
	return  RoomService{
		repo: repo,
	}
}