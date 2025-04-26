package utils

import (
	"crypto/md5"
	"encoding/base64"
)

func GetUserIDFromEmail(email string) string {
	hash := md5.Sum([]byte(email))
	return base64.RawURLEncoding.EncodeToString(hash[:])
}
