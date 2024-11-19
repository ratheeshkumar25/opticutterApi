package user

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/middleware"
	"github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/config"

	pb "github.com/ratheeshkumar25/opti_cut_api_gateway/pkg/user/userpb"
)

// User represents the User route handler, containing configuration and gRPC client.
type User struct {
	cfg    *config.Config
	client pb.UserServiceClient
}

// NewUserRoute initializes the user routes and handlers.
func NewUserRoute(c *gin.Engine, cfg config.Config) {
	client, err := ClientDial(cfg)
	if err != nil {
		log.Fatalf("error not connected with grpc client : %v", err.Error())
	}

	userHandler := &User{
		cfg:    &cfg,
		client: client,
	}

	apiVersion := c.Group("/api/v1")

	user := apiVersion.Group("/user")

	{
		user.POST("/signup", userHandler.UserSignup)
		user.POST("/verify", userHandler.UserVerify)
		user.POST("/login", userHandler.UserLogin)
	}
	auth := user.Group("/auth")
	auth.Use(middleware.Authorization(cfg.SECRETKEY))
	{
		auth.POST("/address", userHandler.AddAddress)
		auth.GET("/address", userHandler.ViewAllAddress)
		auth.PATCH("/address/:id", userHandler.EditAddress)
		auth.DELETE("/address/:id", userHandler.RemoveAddress)

		auth.GET("/profile", userHandler.ViewProfile)
		auth.PATCH("/profile", userHandler.EditProfile)
		auth.PATCH("/password", userHandler.ChangePassword)

		auth.POST("/item", userHandler.AddItem)
		auth.PATCH("item/:id", userHandler.EditItem)
		//auth.GET("/item", userHandler.FindAllItem)
		auth.GET("/item", userHandler.FindAllItemByUser)
		auth.DELETE("items/:id", userHandler.RemoveItem)

		auth.GET("/material/:id", userHandler.FindMaterialByID)
		auth.GET("/material", userHandler.FindAllMaterial)

		auth.POST("/order/place", userHandler.PlaceOrder)
		auth.GET("/order/:id", userHandler.FindOrder)
		auth.GET("/order", userHandler.OrderHistory)
		auth.GET("/orders", userHandler.FindOrdersByUser)
		auth.GET("/cuttingresult/:id", userHandler.CuttingResult)

		user.GET("/payment", userHandler.Payment)
		user.POST("/payment/success", userHandler.PaymentSuccess)
		user.GET("/payment-success", userHandler.PaymentSuccessPage)

	}
}
