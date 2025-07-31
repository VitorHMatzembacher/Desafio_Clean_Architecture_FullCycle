package web

import (
	"net/http"

	"ordersystem/internal/usecase"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	ListOrdersUseCase *usecase.ListOrdersUseCase
}

func NewOrderHandler(listOrders *usecase.ListOrdersUseCase) *OrderHandler {
	return &OrderHandler{
		ListOrdersUseCase: listOrders,
	}
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	orders, err := h.ListOrdersUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}
