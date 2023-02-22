package api

import (
	db "github.com/NickDiPreta/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for simplebank
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and router
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//routes
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

// Start runs the http server at an address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) map[string]interface{} {
	return gin.H{"error": err.Error()}
}
