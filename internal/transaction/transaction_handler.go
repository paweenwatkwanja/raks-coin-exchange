package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/paweenwatkwanja/raks-coin-exchange/config"
	"github.com/paweenwatkwanja/raks-coin-exchange/models"
	broadcast "github.com/paweenwatkwanja/transaction-broadcasting"
	broadcastModels "github.com/paweenwatkwanja/transaction-broadcasting/models"
)

type transactionHandler struct {
	config             *config.Config
	transactionRequest models.Request
}

const (
	BroadcastPath   = "broadcast"
	CheckStatusPath = "check"
)

func NewTransactionHandler(cfg *config.Config, r fiber.Router) {
	handler := &transactionHandler{
		config: cfg,
	}
	r.Post("/", handler.postTransaction)
}

func (h *transactionHandler) postTransaction(c *fiber.Ctx) error {
	h.transactionRequest = models.Request{}
	err := c.BodyParser(&h.transactionRequest)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
	}

	broadcastService := h.initBroadcastService()

	fmt.Println("begin transaction")
	broadcastRequest := h.newBroadcastRequest()
	broadcastURL := h.getBroadcastURL(BroadcastPath)
	txHash, err := broadcastService.BroadcastTransaction(broadcastURL, broadcastRequest)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
	}

	var txStatus string
	monitorURL := h.getCheckURL(CheckStatusPath, txHash)
	if txHash != "" {
		txStatus, err = broadcastService.MonitorTransaction(monitorURL)
		if err != nil {
			fmt.Printf("error: %v\n", err.Error())
		}
	}

	if txStatus != "" {
		txStatus, err = broadcastService.HandleStatus(monitorURL, txStatus)
		if err != nil {
			fmt.Printf("error: %v\n", err.Error())
		}
	}
	fmt.Println("transaction end")

	var response models.Response
	if err != nil {
		response.ErrorMessage = err.Error()
		return c.Status(fiber.StatusInternalServerError).JSON(response)
	}
	response.TransactionStatus = txStatus

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *transactionHandler) initBroadcastService() *broadcast.BroadcastService {
	broadcastService := broadcast.NewBroadcastService()
	broadcastService.WithRetryRequest(&broadcastModels.RetryRequest{
		RetryAttempt:  h.transactionRequest.RetryAttempt,
		RetryDuration: h.transactionRequest.RetryDuration,
	})
	broadcastService.WithCustomHTTPRequest(&broadcastModels.CustomHTTPRequest{
		RetryAttempt:  h.transactionRequest.RetryAttemptHTTP,
		RetryDuration: h.transactionRequest.RetryDurationHTTP,
		Timeout:       h.transactionRequest.Timeout,
	})
	return broadcastService
}

func (h *transactionHandler) newBroadcastRequest() *broadcastModels.BroadcastRequest {
	broadcastRequest := &broadcastModels.BroadcastRequest{
		Symbol:    h.transactionRequest.Symbol,
		Price:     h.transactionRequest.Price,
		Timestamp: h.transactionRequest.Timestamp,
	}
	return broadcastRequest
}

func (h *transactionHandler) getBroadcastURL(path string) string {
	return fmt.Sprintf("%v/%v", h.config.Endpoint, path)
}

func (h *transactionHandler) getCheckURL(path string, param string) string {
	return fmt.Sprintf("%v/%v/%v", h.config.Endpoint, path, param)
}
