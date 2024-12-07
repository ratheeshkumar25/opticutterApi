package chat

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/middleware"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/chat/pb"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/config"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user"
	userpb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user/userpb"
)

type Chat struct {
	cfg        *config.Config
	userClient userpb.UserServiceClient
	//adminClient adminpb.AdminServiceClient
	client pb.ChatServiceClient
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

	// adminClient, err := admin.ClientDial(cfg)
	// if err != nil {
	// 	log.Fatalf("error not connected with grpc admin client: %v", err.Error())
	// }

	chatHandler := &Chat{
		cfg:        &cfg,
		client:     client,
		userClient: userClient,
		//adminClient: adminClient,
	}

	apiVersion := c.Group("/api/v1")

	// Apply authorization middleware
	auth := apiVersion.Group("/auth")
	auth.Use(middleware.Authorization(cfg.SECRETKEY))

	{
		// User routes with authentication
		auth.POST("/submit-review", chatHandler.SubmitReview)             // Submit Review
		auth.GET("/fetch-reviews/:material_id", chatHandler.FetchReviews) // Fetch Reviews
		auth.POST("/upload-video", chatHandler.UploadVideo)               // Upload Video
		auth.GET("/fetch-videos/:material_id", chatHandler.FetchVideos)   // Fetch Videos
	}

	user := apiVersion.Group("/user")
	{
		user.GET("/chat", chatHandler.Chat)
		user.POST("/video-call", chatHandler.VideoCall)
		//user.GET("/video-call", chatHandler.VideoCall)
	}
	c.GET("/chat", chatHandler.ChatPage)
}
