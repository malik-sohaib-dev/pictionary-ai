package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/coder/websocket"
)

type SocketHandler struct {}

func NewSocketHandler() SocketHandler {
	return SocketHandler{}
}

func (h *SocketHandler) EchoHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Accept the WebSocket handshake
	// In production, configure Options to restrict Origins for security
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true, // Allows cross-origin requests for testing
	})
	if err != nil {
		log.Printf("Failed to accept websocket connection: %v", err)
		return
	}
	// Always ensure the connection gets closed when the handler exits
	defer conn.Close(websocket.StatusInternalError, "connection closed unexpectedly")

	log.Println("Client connected!")

	// 2. Loop to continuously read and write messages
	for {
		// Set a timeout context or use r.Context() to handle client disconnection
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Minute)
		
		// Read message from client
		messageType, data, err := conn.Read(ctx)
		if err != nil {
			log.Printf("Read error (client likely disconnected): %v", err)
			cancel()
			break
		}

		log.Printf("Received: %s", string(data))

		// Echo the same message back to the client
		err = conn.Write(ctx, messageType, data)
		if err != nil {
			log.Printf("Write error: %v", err)
			cancel()
			break
		}
		
		cancel() // Clean up context resources after each loop iteration
	}
}