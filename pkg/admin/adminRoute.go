package admin

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/middleware"
	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/admin/adminpb"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/config"
)

// Admin represents the admin route handler, containing configuration and gRPC client.
type Admin struct {
	Cfg    *config.Config
	Client pb.AdminServiceClient
}

// NewAdminRoute initializes the admin routes and handlers.
func NewAdminRoute(c *gin.Engine, cfg config.Config) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error not connected with grpc client : %v", err.Error())
	}

	adminHandler := &Admin{
		Cfg:    &cfg,
		Client: client,
	}

	apiVersion := c.Group("/api/v1")

	admin := apiVersion.Group("/admin")
	{
		admin.POST("/login", adminHandler.AdminLogin)

	}

	auth := admin.Group("/auth")
	auth.Use(middleware.AdminAuthorization(cfg.SECRETKEY, "admin"))
	{
		auth.PATCH("/user/:id", adminHandler.BlockUser)
		auth.PATCH("/user/unblock/:id", adminHandler.UnblockUser)

		auth.POST("/material", adminHandler.AddMaterial)
		auth.GET("/material", adminHandler.FindAllMaterial)
		auth.GET("/material/:id", adminHandler.FindMaterialByID)
		auth.PATCH("/material/:id", adminHandler.EditMaterial)
		auth.DELETE("/material/:id", adminHandler.RemoveMaterial)

		auth.GET("/item", adminHandler.FindAllItem)

		auth.GET("/orders", adminHandler.OrderHistory)
		auth.GET("/order/:id", adminHandler.FindOrder)
		auth.GET("/orders/:id", adminHandler.FindOrdersByUser)
	}
}
