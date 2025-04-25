package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	ginmiddleware "github.com/ulule/limiter/v3/drivers/middleware/gin"
	memorystore "github.com/ulule/limiter/v3/drivers/store/memory"
)

// RateLimit 返回一个基于 IP 的限速中间件
func RateLimit(limit int, duration time.Duration) gin.HandlerFunc {
	rate := limiter.Rate{
		Period: duration,
		Limit:  int64(limit),
	}

	store := memorystore.NewStore()
	instance := limiter.New(store, rate)

	return ginmiddleware.NewMiddleware(instance)
}
