package storage

import (
	"sync"

	"github.com/Axpz/store/internal/config"
)

// StoreInterface 定义存储接口
type StoreInterface interface {
	// 用户相关操作
	Create(user User) error
	Get(id string) (User, error)
	Update(user User) error
	Delete(id string) error

	// 评论相关操作
	CreateComment(comment Comment) error
	GetComment(id string) (Comment, error)
	UpdateComment(comment Comment) error
	DeleteComment(id string) error
}

type Store struct {
	mu        sync.RWMutex
	config    *config.Config
	loaded    map[string]bool
	timestamp int64
	waiting   bool // 考虑到github api 限速，我们需要等待一段时间再调用远程存储

	Tables
}

type Tables struct {
	users    map[string]User
	comments map[string]Comment
}

func NewStore(cfg *config.Config) Store {
	return Store{
		config: cfg,
		mu:     sync.RWMutex{},
		loaded: make(map[string]bool),
		Tables: Tables{
			users:    make(map[string]User),
			comments: make(map[string]Comment),
		},
	}
}
