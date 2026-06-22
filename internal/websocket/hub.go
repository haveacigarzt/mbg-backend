package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	rooms map[string]map[*websocket.Conn]bool
	mu    sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		rooms: make(map[string]map[*websocket.Conn]bool),
	}
}

func (h *Hub) JoinRoom(roomID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, exists := h.rooms[roomID]; !exists {
		h.rooms[roomID] = make(map[*websocket.Conn]bool)
	}

	h.rooms[roomID][conn] = true
}

func (h *Hub) LeaveRoom(roomID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if room, exists := h.rooms[roomID]; exists {
		delete(room, conn)

		if len(room) == 0 {
			delete(h.rooms, roomID)
		}
	}

	conn.Close()
}

func (h *Hub) BroadcastToRoom(roomID string, data []byte) {
	h.mu.RLock()

	room, exists := h.rooms[roomID]
	if !exists {
		h.mu.RUnlock()
		return
	}

	conns := make([]*websocket.Conn, 0, len(room))
	for conn := range room {
		conns = append(conns, conn)
	}

	h.mu.RUnlock()

	for _, conn := range conns {
		if err := conn.WriteMessage(websocket.TextMessage, data); err != nil {
			h.LeaveRoom(roomID, conn)
		}
	}
}
