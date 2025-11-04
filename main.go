package main

import (
	"github.com/paweenwatkwanja/raks-coin-exchange/config"
	"github.com/paweenwatkwanja/raks-coin-exchange/server"
)

func main() {
	cfg := config.LoadConfig()

	s := server.NewServer(cfg)
	s.Start()
}
