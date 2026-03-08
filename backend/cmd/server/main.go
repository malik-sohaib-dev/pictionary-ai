package main

import (
	"fmt"
	"net/http"

	"github.com/malik-sohaib-dev/pictionary-ai/backend/handlers"
	"github.com/malik-sohaib-dev/pictionary-ai/backend/repo"
	"github.com/malik-sohaib-dev/pictionary-ai/backend/services"
)

func main() {
	roomRepo := repo.NewRoomRepo()

	roomService := services.NewRoomService(roomRepo)

	roomHandler := handlers.NewRoomHandler(roomService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/rooms", roomHandler.GetAll)
	mux.HandleFunc("POST /api/rooms", roomHandler.Create)
	mux.HandleFunc("PATCH /api/rooms/{id}", roomHandler.Patch)
	mux.HandleFunc("GET /api/rooms/{id}", roomHandler.GetById)

	fmt.Println("Listening on port 8080")

	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		fmt.Println("Server Crashed")
	}
}