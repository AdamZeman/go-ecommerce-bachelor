package handlers

import (
	"context"
	"ecomm-go/db/goqueries"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func (h DataHandler) HandleWebSocket(c echo.Context) error {
	fmt.Println("WebSocket connection attempt...")

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		fmt.Println("WebSocket Upgrade HandleError:", err)
		return err
	}

	user, err := GetUserFromSession(h, c)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	registerClient(conn, user.Email)
	fmt.Println("WebSocket Connected!")
	defer disconnectClient(conn)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read HandleError:", err)
			break
		}
		fmt.Println("Received:", string(msg))

		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Printf("HandleError unmarshalling JSON: %v", err)
			return err
		}
		message.SenderEmail = user.Email
		jsonMessage, err := json.Marshal(message)
		if err != nil {
			return err
		}
		broadcast <- string(jsonMessage)
	}
	return nil
}

func BroadcastMessages(queries *goqueries.Queries) error {
	for {
		select {
		case msg := <-broadcast:
			fmt.Println("Broadcasting client message:", msg)

			var message Message
			if err := json.Unmarshal([]byte(msg), &message); err != nil {
				log.Printf("HandleError unmarshalling JSON: %v", err)
				return err
			}
			sendToServerAndClients(message, queries)
		}
	}
}

func sendToServerAndClients(message Message, queries *goqueries.Queries) {
	mutex.Lock()
	defer mutex.Unlock()

	numIdParam, _ := ParseToInt32(message.ChatRoomId)

	insertMessage := goqueries.InsertMessageParams{
		ConversationID: numIdParam,
		Email:          message.SenderEmail,
		Content:        message.ChatMessage,
	}

	messageFromDatabase, err := queries.InsertMessage(context.Background(), insertMessage)
	if err != nil {
		fmt.Println("issue writing to database")
		return
	}

	currentUserMsg := renderMessage("current", messageFromDatabase)
	otherUserMsg := renderMessage("other", messageFromDatabase)

	for client, email := range clients {
		fmt.Println("clients")

		msg := currentUserMsg
		if email != message.SenderEmail {
			msg = otherUserMsg
		}
		if err := client.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
			log.Printf("Write HandleError: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}
