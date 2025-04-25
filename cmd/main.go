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
	"github.com/gin-gonic/gin"
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
	userHandler := handler.NewUserHandler(userService)

	// 设置路由
	r := gin.Default()

	// 添加限速中间件
	// 每分钟允许60个请求
	r.Use(middleware.RateLimit(10, time.Minute))

	// 用户相关路由
	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	// 评论相关路由
	r.POST("/comments", userHandler.CreateComment)
	r.GET("/comments/:id", userHandler.GetComment)
	r.PUT("/comments/:id", userHandler.UpdateComment)
	r.DELETE("/comments/:id", userHandler.DeleteComment)

	// 启动服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("服务器启动在 http://localhost%s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
