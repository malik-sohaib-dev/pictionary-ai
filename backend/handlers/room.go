package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/malik-sohaib-dev/pictionary-ai/backend/services"
	"github.com/malik-sohaib-dev/pictionary-ai/backend/types"
)

type RoomHandler struct {
	service *services.RoomService
}

// Handler constructor
func NewRoomHandler(s *services.RoomService) *RoomHandler {
	return &RoomHandler{
		service: s,
	}
}

// Get All Rooms
func (h *RoomHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	rooms, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(rooms)
}

// Create Room
func (h *RoomHandler) Create(w http.ResponseWriter, r *http.Request) {
	var createRoomBody types.CreateRoom

	err := json.NewDecoder(r.Body).Decode(&createRoomBody)
	if err != nil {
		http.Error(w, "Bad request JSON", http.StatusBadRequest)
		return
	}

	roomId, err := h.service.Create(createRoomBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := types.CreateRoomResponse{
		RoomId: roomId,
	}

	json.NewEncoder(w).Encode(response)
}

// Patch
func (h *RoomHandler) Patch(w http.ResponseWriter, r *http.Request) {
	var patchRoomBody types.PatchRoom

	roomId := r.PathValue("id")
	if roomId == "" {
		http.Error(w, "room id is required", http.StatusBadRequest)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&patchRoomBody)
	if err != nil {
		http.Error(w, "Bad request JSON", http.StatusBadRequest)
		return
	}

	room, err := h.service.Patch(roomId, patchRoomBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(room)
}

// Get By ID
func (h *RoomHandler) GetById(w http.ResponseWriter, r *http.Request) {
	roomId := r.PathValue("id")
	if roomId == "" {
		http.Error(w, "room id is required", http.StatusBadRequest)
		return
	}

	room, err := h.service.GetById(roomId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(room)
}

