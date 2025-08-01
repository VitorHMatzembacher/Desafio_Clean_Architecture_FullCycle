package web

import (
	"net/http"

	"github.com/VitorHMatzembacher/Desafio_Clean_Architecture_FullCycle/internal/usecase"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	CreateOrderUseCase *usecase.CreateOrderUseCase
	ListOrdersUseCase  *usecase.ListOrdersUseCase
}

func NewOrderHandler(
	createUc *usecase.CreateOrderUseCase,
	listUc *usecase.ListOrdersUseCase,
) *OrderHandler {
	return &OrderHandler{
		CreateOrderUseCase: createUc,
		ListOrdersUseCase:  listUc,
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req struct {
		Price float64 `json:"price"`
		Tax   float64 `json:"tax"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.CreateOrderUseCase.Execute(c.Request.Context(), req.Price, req.Tax)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) ListOrders(c *gin.Context) {
	orders, err := h.ListOrdersUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}
