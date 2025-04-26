// middleware/logger.go
package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Axpz/store/internal/utils"
	"github.com/Axpz/store/internal/utils/base62"
)

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now().Unix()
		requestID := base62.Encode(now)

		// 创建子 logger
		childLogger := logger.With(
			zap.String("request_id", requestID),
		)

		// 注入 context
		ctx := utils.WithLogger(c.Request.Context(), childLogger)
		c.Request = c.Request.WithContext(ctx)

		// 添加响应头
		c.Header("X-Request-ID", requestID)

		c.Next()
	}
}
