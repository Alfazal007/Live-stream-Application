package managers

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketManager struct {
	Mutex          sync.RWMutex
	RoomWithPeople map[string]map[string]*websocket.Conn
	Admins         map[string]*websocket.Conn
}

func (wsManager *WebSocketManager) TypeChecker(messageType string, message string) bool {
	switch messageType {
	case JoinAdminMessage:
		var adminMessage AdminJoinType
		err := json.Unmarshal([]byte(message), &adminMessage)
		if err != nil || adminMessage.AdminId == "" || len(adminMessage.RoomId) != 11 || adminMessage.Token == "" {
			return false
		}
		return true
	case JoinUserMessage:
		var userMessage UserJoinType
		err := json.Unmarshal([]byte(message), &userMessage)
		if err != nil || userMessage.UserId == "" || len(userMessage.RoomId) != 11 || userMessage.Token == "" {
			return false
		}
		return true
	case TextMessage:
		var textMessage MessageType
		err := json.Unmarshal([]byte(message), &textMessage)
		if err != nil || textMessage.UserName == "" || len(textMessage.RoomId) != 11 || textMessage.Message == "" {
			return false
		}
		return true
	default:
		return false
	}
}
