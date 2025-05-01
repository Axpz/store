package api

import (
	"net/http"
	"strconv"

	"github.com/Axpz/store/internal/middleware"
	"github.com/Axpz/store/internal/service"
	"github.com/Axpz/store/internal/types"
	"github.com/Axpz/store/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderHandler struct {
	orderService *service.OrderService
	jwtSecret    string
}

func NewOrderHandler(orderService *service.OrderService, jwtSecret string) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
		jwtSecret:    jwtSecret,
	}
}

func (h *OrderHandler) RegisterRoutes(router *gin.Engine) {
	orders := router.Group("/api/orders")
	orders.Use(middleware.Auth(h.jwtSecret))
	{
		orders.POST("", h.CreateOrder)
		orders.POST("/:id/capture", h.CaptureOrder)
		orders.GET("", h.GetOrders)
		orders.GET("/:id", h.GetOrder)
		orders.PUT("/:id", h.UpdateOrder)
		orders.DELETE("/:id", h.DeleteOrder)
	}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req types.CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logger := utils.LoggerFromContext(c.Request.Context())
	logger.Info("CreateOrder", zap.Any("req", req))

	order := types.Order{
		Currency:    req.Currency,
		Products:    req.Products,
		TotalAmount: req.TotalAmount,
		Description: req.Description,
	}

	err := h.orderService.CreateOrder(c, &order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dbOrder, err := h.orderService.GetOrder(c, order.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dbOrder)
}

func (h *OrderHandler) CaptureOrder(c *gin.Context) {
	orderID := c.Param("id")
	if orderID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order id is required"})
		return
	}

	if err := h.orderService.CaptureOrder(c, orderID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := h.orderService.GetOrder(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if order == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) GetOrders(c *gin.Context) {
	logger := utils.LoggerFromContext(c.Request.Context())

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "100"))

	logger.Info("GetOrders", zap.Int("page", page), zap.Int("pageSize", pageSize))

	orders, total, err := h.orderService.GetOrders(c, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  orders,
		"total": total,
		"page":  page,
		"size":  pageSize,
	})
}

func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var req types.UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbOrder, err := h.orderService.GetOrder(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updateOrder := dbOrder
	updateOrder.Currency = req.Currency
	updateOrder.Products = req.Products
	updateOrder.TotalAmount = req.TotalAmount
	updateOrder.Description = req.Description

	err = h.orderService.UpdateOrder(c, updateOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateOrder)
}

func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	userID := utils.GetUserIDFromContext(c)
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	order, err := h.orderService.GetOrder(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if order.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "forbidden: you are not allowed to delete this order"})
		return
	}

	err = h.orderService.DeleteOrder(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order deleted successfully"})
}
