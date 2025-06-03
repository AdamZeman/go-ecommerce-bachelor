package handlers

import (
	"bytes"
	"context"
	"ecomm-go/app/views/components"
	"ecomm-go/db/goqueries"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var (
	upgrader  = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	clients   = make(map[*websocket.Conn]string)
	broadcast = make(chan string)
	mutex     sync.Mutex
)

type Message struct {
	SenderEmail string `json:"sender_email"`
	ChatMessage string `json:"chat_message"`
	ChatRoomId  string `json:"chat_room_id"`
}

func registerClient(conn *websocket.Conn, email string) {
	mutex.Lock()
	clients[conn] = email
	mutex.Unlock()
}

func disconnectClient(conn *websocket.Conn) {
	mutex.Lock()
	delete(clients, conn)
	mutex.Unlock()
	conn.Close()
	fmt.Println("WebSocket Disconnected")
}

func renderMessage(senderType string, messageFromDatabase goqueries.ConversationMessage) string {
	buffer := &bytes.Buffer{}
	msgComponent := components.Message(senderType, messageFromDatabase.Content, messageFromDatabase.CreatedAt.Format("2006-01-02 15:04:05"), true)
	msgComponent.Render(context.Background(), buffer)
	return buffer.String()
}
