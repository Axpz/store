package storage

import (
	"context"
	"sync"
	"time"

	"github.com/Axpz/store/internal/config"
	"github.com/Axpz/store/internal/pkg/throttle"
	"go.uber.org/zap"
)

// StoreInterface 定义存储接口
type StoreInterface interface {
	// 用户相关操作
	Create(user User) error
	Get(id string) (User, error)
	Update(user User) error
	Delete(id string) error

	// 订单相关操作
	CreateOrder(order Order) error
	GetOrder(id string) (Order, error)
	GetOrdersByUserID(userID string) ([]Order, error)
	UpdateOrder(order Order) error
	DeleteOrder(id string) error

	// 商品相关操作
	CreateProduct(product Product) error
	GetProduct(id string) (Product, error)
	GetProducts() ([]Product, error)
	UpdateProduct(product Product) error
	DeleteProduct(id string) error

	// 评论相关操作
	CreateComment(comment Comment) error
	GetComment(id string) (Comment, error)
	UpdateComment(comment Comment) error
	DeleteComment(id string) error
}

type Store struct {
	mu            sync.RWMutex
	ctx           context.Context
	throttlesaver *throttle.Saver

	config *config.Config
	loaded map[string]bool

	Tables
}

type Tables struct {
	users    map[string]User
	orders   map[string]Order
	products map[string]Product
	comments map[string]Comment
}

func NewStore(cfg *config.Config) Store {
	return Store{
		mu:            sync.RWMutex{},
		ctx:           context.Background(),
		throttlesaver: throttle.NewSaver(10 * time.Second),

		config: cfg,
		loaded: make(map[string]bool),
		Tables: Tables{
			users:    make(map[string]User),
			comments: make(map[string]Comment),
		},
	}
}

func (s *Store) Logger() *zap.Logger {
	return s.config.Logger
}
