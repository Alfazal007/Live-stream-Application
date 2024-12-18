package managers

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
)

type UserWithConnAndType struct {
	UserType string
	Conn     *websocket.Conn
}

type WebSocketManager struct {
	Mutex          sync.RWMutex
	RoomWithPeople map[string]map[string]UserWithConnAndType
}

func (wsManager *WebSocketManager) HandleAdminMessage(message string, conn *websocket.Conn) {
	var adminMessage AdminJoinType
	err := json.Unmarshal([]byte(message), &adminMessage)
	if err != nil || adminMessage.AdminId == "" || len(adminMessage.RoomId) != 11 || adminMessage.Token == "" {
		return
	}
	wsManager.Mutex.Lock()
	defer wsManager.Mutex.Unlock()
	allRooms := wsManager.RoomWithPeople
	_, exists := allRooms[adminMessage.RoomId]
	if exists {
		return
	}
	// TODO:: make an api call to check if this user is this room's admin
	internalMap := make(map[string]UserWithConnAndType)
	internalMap[adminMessage.AdminId] = UserWithConnAndType{
		UserType: "ADMIN",
		Conn:     conn,
	}
	allMaps := wsManager.RoomWithPeople
	allMaps[adminMessage.RoomId] = internalMap
	wsManager.RoomWithPeople = allMaps
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

func (wsManager *WebSocketManager) HandleUserMessage(message string) {}
func (wsManager *WebSocketManager) HandleTextMessage(message string) {}
