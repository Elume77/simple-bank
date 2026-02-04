package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	db "tutorial.sqlc.dev/app/db/sqlc"
	"tutorial.sqlc.dev/app/token"
	"tutorial.sqlc.dev/app/utils"
)

type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

//type transferRequest struct {
//	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
//	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
//	Amount        int64  `json:"amount" binding:"required,gt=0"`
//	Curency       string `json:"currency" binding:"required,currency"` // Corrected typo
//}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	// Note: Check your config field name (SymmetricKey vs TokenSymmetricKeysymmetricKey)
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.POST("/users", server.createUser)

	router.POST("/users/login", server.loginUser)

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)

	server.router = router
}

//                          //                   //

//                          //

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
