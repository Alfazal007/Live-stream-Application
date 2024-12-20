package managers

import (
	"encoding/json"
	"fmt"
	"sync"

	apicalls "github.com/Alfazal007/websocket/apiCalls"
	"github.com/Alfazal007/websocket/utils"
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
	isValidUser := apicalls.AuthenticateUser(apicalls.AuthenticateStruct{
		Token:  adminMessage.Token,
		UserId: adminMessage.AdminId,
	})
	if !isValidUser {
		return
	}
	isValidStreamAdmin := apicalls.AuthenticateAdminFunction(apicalls.AuthenticateAdmin{
		AdminId:  adminMessage.AdminId,
		StreamId: adminMessage.RoomId,
		Secret:   utils.LoadEnvFiles().Secret,
	})
	if !isValidStreamAdmin {
		return
	}
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

func (wsManager *WebSocketManager) HandleUserMessage(message string, conn *websocket.Conn) {
	var userMessage UserJoinType
	err := json.Unmarshal([]byte(message), &userMessage)
	if err != nil || userMessage.UserId == "" || len(userMessage.RoomId) != 11 || userMessage.Token == "" {
		return
	}
	wsManager.Mutex.Lock()
	defer wsManager.Mutex.Unlock()
	allRooms := wsManager.RoomWithPeople
	requiredRoom, exists := allRooms[userMessage.RoomId]
	if !exists {
		conn.Close()
		return
	}
	isValidUser := apicalls.AuthenticateUser(apicalls.AuthenticateStruct{
		Token:  userMessage.Token,
		UserId: userMessage.UserId,
	})
	if !isValidUser {
		return
	}
	requiredRoom[userMessage.UserId] = UserWithConnAndType{
		Conn:     conn,
		UserType: "USER",
	}

	allMaps := wsManager.RoomWithPeople
	allMaps[userMessage.RoomId] = requiredRoom
	wsManager.RoomWithPeople = allMaps
}

func (wsManager *WebSocketManager) HandleTextMessage(message string, conn *websocket.Conn, messageType int) {
	var messageSentIn MessageType
	err := json.Unmarshal([]byte(message), &messageSentIn)
	if err != nil || messageSentIn.UserName == "" || len(messageSentIn.RoomId) != 11 || messageSentIn.Message == "" || messageSentIn.UserId == "" {
		return
	}
	wsManager.Mutex.RLock()
	defer wsManager.Mutex.RUnlock()
	allRooms := wsManager.RoomWithPeople
	requiredRoom, exists := allRooms[messageSentIn.RoomId]
	if !exists {
		conn.Close()
		return
	}
	userConnectionStuff, userExists := requiredRoom[messageSentIn.UserId]
	if !userExists || userConnectionStuff.Conn != conn {
		conn.Close()
		return
	}

	messageToBeSent := BroadCast{
		TypeOfMessage: "NORMALMESSAGE",
		Message:       messageSentIn.Message,
		Sender:        messageSentIn.UserName,
	}

	bytesSendingMessage, _ := json.Marshal(messageToBeSent)
	for _, userWithConn := range requiredRoom {
		if userWithConn.Conn != conn {
			fmt.Println("sending message")
			userWithConn.Conn.WriteMessage(messageType, bytesSendingMessage)
		}
	}
}

func (wsManager *WebSocketManager) CleanUp(conn *websocket.Conn) {
	wsManager.Mutex.Lock()
	defer wsManager.Mutex.Unlock()
outer:
	for roomId, roomInner := range wsManager.RoomWithPeople {
		for userId, userConnectionData := range roomInner {
			if conn == userConnectionData.Conn {
				if userConnectionData.UserType == "ADMIN" {
					// disconnect the whole room
					allRooms := wsManager.RoomWithPeople
					roomRemoved := allRooms[roomId]
					delete(allRooms, roomId)
					wsManager.RoomWithPeople = allRooms
					for _, roomId := range roomRemoved {
						roomId.Conn.Close()
					}
				} else {
					// remove from the map and close the connection
					newMap := roomInner
					delete(newMap, userId)
					wsManager.RoomWithPeople[roomId] = roomInner
					conn.Close()
					break outer
				}
			}
		}
	}
}
