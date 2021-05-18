package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type M map[string]interface{}

const (
	MessageNewUser = "New User"
	MessageChat    = "Chat"
	MessageLeave   = "Leave"
)

var connections = make([]*WebSocketConnection, 0)

type (
	WebSocketConnection struct {
		*websocket.Conn
		Username string
	}
	SocketPayload struct {
		Message string `json:"message"`
	}
	SocketResponse struct {
		From    string `json:"from"`
		Type    string `json:"type"`
		Message string `json:"message"`
	}
)

func broadcastMessage(currentConn *WebSocketConnection, kind, message string) {
	for _, eachConn := range connections {
		if eachConn == currentConn {
			continue
		}
		err := eachConn.WriteJSON(&SocketResponse{
			From:    currentConn.Username,
			Type:    kind,
			Message: message,
		})
		if err != nil {
			fmt.Printf("Error write json : %v", err)
		}
	}
}

func ejectConnection(currentConn *WebSocketConnection) {
	for k, v := range connections {
		if v == currentConn {
			connections[k] = v
		}
	}
}

func handleIO(currentConn *WebSocketConnection, connections []*WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", fmt.Sprintf("%v", r))
		}
	}()

	broadcastMessage(currentConn, MessageNewUser, "")
	for {
		payload := SocketPayload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				broadcastMessage(currentConn, MessageLeave, "")
				ejectConnection(currentConn)
				return
			}
			log.Println("ERROR", err.Error())
			continue
		}
		broadcastMessage(currentConn, MessageChat, payload.Message)
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("index.html")
		if err != nil {
			http.Error(w, "Could not open requested file", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s", content)
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		currentConnection, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		}
		username := r.URL.Query().Get("username")
		wsConn := WebSocketConnection{Conn: currentConnection, Username: username}
		connections = append(connections, &wsConn)

		go handleIO(&wsConn, connections)
	})

	// Start server
	fmt.Println("Server starting at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
