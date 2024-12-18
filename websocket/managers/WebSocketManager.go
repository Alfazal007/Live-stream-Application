package managers

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketManager struct {
	Mutex          sync.RWMutex
	RoomWithPeople map[string][]*websocket.Conn
	Admins         map[string]*websocket.Conn
}
