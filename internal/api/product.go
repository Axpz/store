package api

import (
	"net/http"
	"strconv"

	"github.com/Axpz/store/internal/middleware"
	"github.com/Axpz/store/internal/service"
	"github.com/Axpz/store/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductHandler struct {
	productService *service.ProductService
	jwtSecret      string
}

func NewProductHandler(productService *service.ProductService, jwtSecret string) *ProductHandler {
	return &ProductHandler{
		productService: productService,
		jwtSecret:      jwtSecret,
	}
}

func (h *ProductHandler) RegisterRoutes(router *gin.Engine) {
	products := router.Group("/api/products")
	products.Use(middleware.Auth(h.jwtSecret))
	{
		// products.POST("", h.CreateProduct)
		products.GET("", h.GetProducts)
		products.GET("/:id", h.GetProduct)
		// products.PUT("/:id", h.UpdateProduct)
		// products.DELETE("/:id", h.DeleteProduct)
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	return
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := h.productService.GetProduct(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	logger := utils.LoggerFromContext(c.Request.Context())

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "100"))

	logger.Info("GetProducts", zap.Int("page", page), zap.Int("pageSize", pageSize))

	products, err := h.productService.GetProducts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  products,
		"total": len(products),
		"page":  page,
		"size":  pageSize,
	})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	return
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	return
}
