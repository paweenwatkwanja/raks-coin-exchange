package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type transactionHandler struct {
}

func NewTransactionHandler(r fiber.Router) {
	handler := &transactionHandler{}
	r.Post("/", handler.postTransaction)
}

func (h *transactionHandler) postTransaction(c *fiber.Ctx) error {
	fmt.Println("Hello, World!")
	return nil
}
