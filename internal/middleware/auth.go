package middleware

import (
	"net/http"
	"strings"

	"github.com/Axpz/store/internal/pkg/jwt"
	"github.com/gin-gonic/gin"
)

const (
	// TokenHeader 是请求头中存放token的key
	TokenHeader = "Authorization"
	// TokenPrefix 是token的前缀
	TokenPrefix = "Bearer "
	// UserIDKey 是存储用户ID的上下文key
	UserIDKey = "user_id"
	// UsernameKey 是存储用户名的上下文key
	UsernameKey = "username"
)

// Auth 认证中间件
func Auth(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取token
		token := c.GetHeader(TokenHeader)
		if token != "" {
			// 检查token格式
			if !strings.HasPrefix(token, TokenPrefix) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "访问令牌格式错误",
				})
				c.Abort()
				return
			}

			// 提取token
			token = strings.TrimPrefix(token, TokenPrefix)
		} else {
			// 检查cookie
			cookieToken, err := c.Cookie("token")
			if err != nil || cookieToken == "" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "未提供访问令牌",
				})
				c.Abort()
				return
			}
			token = cookieToken
		}

		// 验证token
		claims, err := jwt.ValidateToken(token, jwtSecret)
		if err != nil {
			status := http.StatusUnauthorized
			if err == jwt.ErrExpiredToken {
				status = http.StatusForbidden
			}
			c.JSON(status, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set(UserIDKey, claims.UserID)
		c.Set(UsernameKey, claims.Username)

		c.Next()
	}
}
