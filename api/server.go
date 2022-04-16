package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/renatoviolin/simplebank/db/sqlc"
	"github.com/renatoviolin/simplebank/token"
	"github.com/renatoviolin/simplebank/util"
)

type Server struct {
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
	config     util.Config
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	// custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	setupRouter(server)
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	fmt.Println(err.Error())
	return gin.H{"error": err.Error()}
}

func setupRouter(server *Server) {
	router := gin.Default()

	// Middleware Authentication
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// add users to router
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	authRoutes.GET("/users/:username", server.getUser)

	// add routes to router
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts/", server.listAccount)
	authRoutes.PUT("/accounts/", server.updateAccount)
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)

	// add transfer to router
	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}
