package services

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/malik-sohaib-dev/pictionary-ai/backend/repo"
	"github.com/malik-sohaib-dev/pictionary-ai/backend/types"
)

type RoomService struct {
	repo repo.RoomRepository
}

// Room service constructor
func NewRoomService(repo repo.RoomRepository) *RoomService {
	return &RoomService{
		repo: repo,
	}
}

// Default values
var defaultRoom = types.Room{
	Visibility:      types.AllowedVisibilityPublic,
	MaxPlayers:      8,
	Hints:           3,
	WordChoiceCount: 3,
	GameDuration:    80,
}

// Random 6 char Room ID geenrator
func generateRandomRoomId() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	const length = 6

	result := make([]byte, length)

	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))

		if err != nil {
			return "", err
		}

		result[i] = charset[num.Int64()]
	}

	return string(result), nil
}

// Create User
func (s *RoomService) Create(createRoom types.CreateRoom) (string, error) {

	if createRoom.Owner == "" {
		return "", fmt.Errorf("Invalid room owner")
	}

	var roomId string
	// Generating random room ID with uniqueness check
	// 6 characters long, can contain letters and numbers, case insensitive
	for {
		randId, err := generateRandomRoomId()
		if err != nil {
			return "", err
		}

		_, found := s.repo.GetById(randId)
		if found != nil {
			roomId = randId
			break
		}
	}

	if !createRoom.Visibility.Valid() {
		createRoom.Visibility = defaultRoom.Visibility
	}

	roomData := types.Room{
		Visibility:      createRoom.Visibility,
		Owner:           createRoom.Owner,
		RoomId:          roomId,
		MaxPlayers:      defaultRoom.MaxPlayers,
		Hints:           defaultRoom.Hints,
		GameDuration:    defaultRoom.GameDuration,
		WordChoiceCount: defaultRoom.WordChoiceCount,
	}

	err := s.repo.Create(roomData)

	if err != nil {
		return "", err
	}

	return roomId, nil
}

func (s *RoomService) Patch(roomId string, patchRoom types.PatchRoom) (types.Room, error) {
	room, err := s.repo.GetById(roomId)

	if err != nil {
		return types.Room{}, err
	}

	// 1. Serialize the patch data (ignores nil fields if json tags use omitempty,
	// or simply passes them as null which encoding/json skips during unmarshal into concrete values)
	patchBytes, err := json.Marshal(patchRoom)
	if err != nil {
		return types.Room{}, err
	}

	// 2. Unmarshal directly into the existing room object
	// This will only overwrite the fields present in patchBytes
	if err := json.Unmarshal(patchBytes, &room); err != nil {
		return types.Room{}, err
	}

	return s.repo.UpdateById(roomId, room)
}

func (s *RoomService) GetById(roomId string) (types.Room, error) {
	return s.repo.GetById(roomId)
}

func (s *RoomService) GetAll() ([]types.Room, error) {
	return s.repo.GetAll()
}
