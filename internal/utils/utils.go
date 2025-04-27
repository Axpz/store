package utils

import (
	"crypto/md5"
	"encoding/base64"

	"github.com/gin-gonic/gin"
)

func GetUserIDFromEmail(email string) string {
	hash := md5.Sum([]byte(email))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

func GetUserIDFromContext(c *gin.Context) string {
	userID, exists := c.Get("user_id")
	if !exists {
		return ""
	}
	return userID.(string)
}
