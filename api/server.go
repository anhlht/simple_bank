package api

import (
	db "bank_account/db/sqlc"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP request for banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// add routes to router
	router.POST("/accounts", server.createAccount)
	router.GET("accounts/:id", server.getAccount)
	router.GET("accounts", server.listAccount)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
