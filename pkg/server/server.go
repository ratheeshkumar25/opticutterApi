package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Server represents the model of the server with a Gin engine.
type Server struct {
	R *gin.Engine
}

// StartServer method starts the server on the specified port.
func (s *Server) StartServer(port string) {
	err := s.R.Run(":" + port)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// NewServer returns a new Server instance with the default Gin engine attached.
func NewServer() *Server {
	engine := gin.Default()

	return &Server{
		R: engine,
	}
}
