package chat

import (
	"log"

	"github.com/gin-gonic/gin"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/chat/pb"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/config"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user"
	userpb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user/userpb"
)

type Chat struct {
	cfg        *config.Config
	userClient userpb.UserServiceClient
	client     pb.ChatServiceClient
}

func NewChatRoutes(c *gin.Engine, cfg config.Config) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error not connected with grpc client : %v", err.Error())
	}

	userClient, err := user.ClientDial(cfg)
	if err != nil {
		log.Fatalf("error not connected with grpc user client : %v", err.Error())
	}
	chatHandler := &Chat{
		cfg:        &cfg,
		client:     client,
		userClient: userClient,
	}

	apiVersion := c.Group("/api/v1")

	user := apiVersion.Group("/user")
	{
		user.GET("/chat", chatHandler.Chat)
		user.POST("/video-call", chatHandler.VideoCall)
		//user.GET("/video-call", chatHandler.VideoCall)
	}
	c.GET("/chat", chatHandler.ChatPage)
}
