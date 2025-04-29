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
	defer logger.Sync() // 确保日志被刷新

	// 加载配置
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}
	cfg.Logger = logger

	// 设置路由
	r := gin.Default()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(middleware.LoggerMiddleware(logger))
	r.Use(middleware.CORS())
	r.Use(middleware.RateLimit(60, time.Minute))

	// 创建存储实例
	store, err := storage.New(cfg)
	if err != nil {
		log.Fatalf("创建存储失败: %v", err)
	}

	// 创建服务和处理器
	userService := service.NewUserService(store)
	userHandler := api.NewUserHandler(userService, cfg.JWT.Secret)
	userHandler.RegisterRoutes(r)

	// 订单相关路由
	orderService := service.NewOrderService(store)
	orderHandler := api.NewOrderHandler(orderService, cfg.JWT.Secret)
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
