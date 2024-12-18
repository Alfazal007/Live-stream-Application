package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/Alfazal007/websocket/managers"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsHandler(w http.ResponseWriter, r *http.Request, wsManager *managers.WebSocketManager) {
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade the websocket connection")
		return
	}
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("Connected and read ", string(message))
	}
}

func main() {
	wsManager := managers.WebSocketManager{
		Mutex:          sync.RWMutex{},
		RoomWithPeople: make(map[string][]*websocket.Conn),
		Admins:         make(map[string]*websocket.Conn),
	}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsHandler(w, r, &wsManager)
	})

	err := http.ListenAndServe("0.0.0.0:8001", nil)
	fmt.Println(err)
}
