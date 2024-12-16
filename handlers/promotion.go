package handlers

import (
	"gotest/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PromotionHandler interface {
	CalculateDiscount(c *fiber.Ctx) error
}

type promotionHandler struct {
	promotionService services.PromotionService
}

func NewPromotionHandler(promotionService services.PromotionService) PromotionHandler {
	return promotionHandler{promotionService}
}

func (h promotionHandler) CalculateDiscount(c *fiber.Ctx) error {
	// http://localhost:8000/calculate?amount=100

	amountStr := c.Query("amount")
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	discount, err := h.promotionService.CalculateDiscount(amount)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendString(strconv.Itoa(discount))
}
