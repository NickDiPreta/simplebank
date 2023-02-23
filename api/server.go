package api

import (
	db "github.com/NickDiPreta/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	//routes
	router.POST("/accounts", server.createAccount)
	router.POST("/transfers", server.createTransfer)
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
