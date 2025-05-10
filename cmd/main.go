package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Axpz/store/internal/api"
	"github.com/Axpz/store/internal/config"
	"github.com/Axpz/store/internal/middleware"
	"github.com/Axpz/store/internal/service"
	"github.com/Axpz/store/internal/storage"
	"github.com/Axpz/store/internal/utils"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger, err := utils.NewLoggerWithoutStacktrace()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	cfg := config.Load("config.yaml")
	cfg.Logger = logger

	r := gin.Default()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(middleware.LoggerMiddleware(logger))
	r.Use(middleware.CORS())
	r.Use(middleware.RateLimit(60, time.Minute))

	// initialize storage instance
	store, err := storage.New(cfg)
	if err != nil {
		log.Fatalf("get store instace failed: %v", err)
	}

	// 创建支付提供者
	payService, err := service.NewPaymentService(cfg)
	if err != nil {
		logger.Fatal("创建PayPal提供者失败", zap.Error(err))
	}

	// 创建服务和处理器
	userService := service.NewUserService(store)
	emailService := service.NewEmailService(cfg)
	userHandler := api.NewUserHandler(userService, emailService, cfg.JWT.Secret)
	userHandler.RegisterRoutes(r)

	// 订单相关路由
	orderService := service.NewOrderService(store)
	orderHandler := api.NewOrderHandler(orderService, payService, cfg.JWT.Secret)
	orderHandler.RegisterRoutes(r)

	// 商品相关路由
	productService := service.NewProductService(store)
	productHandler := api.NewProductHandler(productService, cfg.JWT.Secret)
	productHandler.RegisterRoutes(r)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("server start at %s", addr)
	if err := r.Run(addr); err != nil {
		logger.Fatal("server start failed", zap.Error(err))
	}
}
