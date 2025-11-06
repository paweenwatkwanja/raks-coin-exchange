package handler

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/paweenwatkwanja/raks-coin-exchange/config"
	broadcast "github.com/paweenwatkwanja/transaction-broadcasting"
	"github.com/paweenwatkwanja/transaction-broadcasting/models"
)

type transactionHandler struct {
	config config.Config
}

func NewTransactionHandler(r fiber.Router) {
	handler := &transactionHandler{}
	r.Post("/", handler.postTransaction)
}

func (h *transactionHandler) postTransaction(c *fiber.Ctx) error {
	request := models.BroadcastRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
	}

	broadcastService := broadcast.NewBroadcastService()

	broadcastService.WithRetryAttempt(5)
	broadcastService.WithRetryDuration(time.Duration(10))

	endpoint := h.config.Endpoint
	broadcastURL := fmt.Sprintf("%v/%v", endpoint, "broadcast")

	txHash, err := broadcastService.BroadcastTransaction(broadcastURL, request)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
	}
	fmt.Printf("txHash: %v\n", txHash)

	monitorURL := fmt.Sprintf("%v/%v/%v", endpoint, "check", txHash)
	txStatus, err := broadcastService.MonitorTransaction(monitorURL)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
	}
	fmt.Printf("txStatus: %v\n", txStatus)

	err = broadcastService.HandleStatus(monitorURL, txStatus)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
	}
	fmt.Println("done")
	return nil
}
