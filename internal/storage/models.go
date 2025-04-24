package storage

import (
	"github.com/Axpz/store/internal/types"
)

// User 表示用户数据模型
type User = types.User

// Comment 表示评论数据模型
type Comment struct {
	ID      string `json:"id"`
	UserID  string `json:"user_id"`
	Content string `json:"content"`
	Created int64  `json:"created"`
	Updated int64  `json:"updated"`
}
