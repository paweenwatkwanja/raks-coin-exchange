package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/paweenwatkwanja/raks-coin-exchange/config"
	handler "github.com/paweenwatkwanja/raks-coin-exchange/internal/transaction"
)

type Server struct {
	App *fiber.App
	Cfg *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{App: fiber.New(), Cfg: cfg}
}

func (s *Server) Start() {
	router := s.App.Group("/v1")
	group := router.Group("/transactions")
	handler.NewTransactionHandler(group)

	appConnString := fmt.Sprintf("%s:%s", s.Cfg.AppHost, s.Cfg.AppPort)

	fmt.Println("Server is running on " + appConnString)
	if err := s.App.Listen(appConnString); err != nil {
		panic(err)
	}
}
