package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	dto "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/DTO"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/chat/pb"
	userpb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user/userpb"
)

// Upgrader variable specifies the parmeters of upgrading HTTP request
var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// HandleWebSocketConnection handles the weboscket connection and bidirectional streaming
func HandleWebSocketConnection(c *gin.Context, client pb.ChatServiceClient, userClient userpb.UserServiceClient) {
	ctx := c.Request.Context()

	conn, err := Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	log.Println("WebSocket connection established")

	for {
		select {
		case <-ctx.Done():
			// Context canceled, stop processing messages
			log.Println("WebSocket connection closed")
			return
		default:
			_, msg, err := conn.ReadMessage()
			if err != nil {
				log.Println("Error reading message:", err)
				return
			}

			var message dto.Message
			err = json.Unmarshal(msg, &message)
			if err != nil {
				log.Println("Error decoding JSON:", err)
				continue
			}

			// Checking the user and receiver IDs
			_, err = userClient.ViewProfile(ctx, &userpb.ID{ID: uint32(message.UserID)})
			if err != nil {
				log.Println("Error fetching user or invalid userID:", err)
				continue
			}

			_, err = userClient.ViewProfile(ctx, &userpb.ID{ID: uint32(message.ReceiverID)})
			if err != nil {
				log.Println("Error fetching receiver or invalid receiverID:", err)
				continue
			}

			stream, err := client.Connect(ctx)
			if err != nil {
				log.Println("Error calling chat service:", err)
				continue
			}
			ch := &clientHandle{
				stream:     stream,
				userID:     uint32(message.UserID),
				receiverID: uint32(message.ReceiverID),
			}

			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Error writing message:", err)
				return
			}

			go ch.sentMessage(message.Message)
			go ch.receiveMessage(conn, uint32(message.UserID), uint32(message.ReceiverID))
		}
	}
}

// ChatPage loads the chat page.
func ChatPage(c *gin.Context, client pb.ChatServiceClient) {
	timeout := time.Second * 1000
	ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
	defer cancel()

	id := c.Query("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error converting id to int:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest, "Message": "Error converting id to int", "Error": err.Error()})
		return
	}

	receiverIDStr := c.Query("receiverId")
	receiverID, err := strconv.Atoi(receiverIDStr)
	if err != nil {
		log.Println("Error converting receiverId to int:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Status": http.StatusBadRequest, "Message": "Error converting receiverId to int", "Error": err.Error()})
		return
	}

	response, err := client.FetchHistory(ctx, &pb.ChatID{User_ID: uint32(userID), Receiver_ID: uint32(receiverID)})
	if err != nil {
		log.Println("Error calling chat client:", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Status": http.StatusInternalServerError, "Message": "Error calling chat client", "Error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "chat.html", gin.H{"response": response.Chats, "id": userID})
}

type clientHandle struct {
	userID     uint32
	receiverID uint32
	stream     pb.ChatService_ConnectClient
}

func (c *clientHandle) sentMessage(msg string) {
	message := &pb.Message{
		User_ID:     c.userID,
		Receiver_ID: c.receiverID,
		Content:     msg,
	}

	err := c.stream.Send(message)
	if err != nil {
		log.Printf("Error while sending message to server: %v", err)
	}
}

func (c *clientHandle) receiveMessage(conn *websocket.Conn, userID, receiverID uint32) {
	for {
		mssg, err := c.stream.Recv()
		if err != nil {
			log.Printf("Error receiving message from server: %v", err)
			return
		}

		if userID == mssg.Receiver_ID && receiverID == mssg.User_ID {
			dom := &dto.Message{
				UserID:     uint(mssg.User_ID),
				ReceiverID: uint(mssg.Receiver_ID),
				Message:    mssg.Content,
			}
			msg, err := json.Marshal(dom)
			if err != nil {
				log.Println("Error encoding JSON:", err)
				return
			}

			err = conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Error writing message:", err)
				return
			}
		}
	}
}

func StartVideoCall(c *gin.Context, client pb.ChatServiceClient) {
	var req struct {
		UserID     uint32 `json:"user_id" binding:"required"`
		ReceiverID uint32 `json:"receiver_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Invalid request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	response, err := client.StartVideoCall(ctx, &pb.VideoCallRequest{
		User_ID:     req.UserID,
		Receiver_ID: req.ReceiverID,
	})
	if err != nil {
		log.Printf("Failed to start video call: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start video call", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"room_url": response.RoomUrl})
}
