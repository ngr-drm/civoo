package infrastructure

import (
	"github.com/gin-gonic/gin"
	"github/Re44e/civoo/http/routes"
	"log"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)
	log.Print("Server is running...")
	log.Fatal(router.Run(":" + s.port))
}
