package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Axpz/store/internal/config"
	"github.com/Axpz/store/internal/handler"
	"github.com/Axpz/store/internal/middleware"
	"github.com/Axpz/store/internal/service"
	"github.com/Axpz/store/internal/storage"
	"github.com/Axpz/store/internal/utils"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 加载配置
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 创建存储实例
	store, err := storage.New(cfg)
	if err != nil {
		log.Fatalf("创建存储失败: %v", err)
	}

	// 创建服务和处理器
	userService := service.NewUserService(store)
	userHandler := handler.NewUserHandler(userService, cfg.JWT.Secret)

	logger, err := utils.NewLoggerWithoutStacktrace()
	if err != nil {
		panic(err)
	}
	defer logger.Sync() // 确保日志被刷新

	// 设置路由
	r := gin.Default()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
	r.Use(middleware.LoggerMiddleware(logger))
	r.Use(middleware.CORS())
	r.Use(middleware.RateLimit(10, time.Minute))

	// 公开路由组
	public := r.Group("/api")
	{
		// 登录注册等公开接口
		public.POST("/auth/login", userHandler.Login)
		public.POST("/auth/register", userHandler.Register)
	}

	// 需要认证的路由组
	// authorized := r.Group("/api")
	// authorized.Use(middleware.Auth(cfg.JWT.Secret))
	{
		// 用户相关路由
		public.POST("/users", userHandler.CreateUser)
		public.GET("/users/:id", userHandler.GetUser)
		public.PUT("/users/:id", userHandler.UpdateUser)
		public.DELETE("/users/:id", userHandler.DeleteUser)
	}

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("server start at %s", addr)
	if err := r.Run(addr); err != nil {
		logger.Fatal("server start failed", zap.Error(err))
	}
}
