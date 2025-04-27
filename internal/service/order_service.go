package service

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"time"

	"github.com/Axpz/store/internal/storage"
	"github.com/Axpz/store/internal/types"
	"github.com/Axpz/store/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrderService struct {
	store storage.StoreInterface
}

func NewOrderService(store storage.StoreInterface) *OrderService {
	return &OrderService{
		store: store,
	}
}

func (s *OrderService) CreateOrder(c *gin.Context, order *types.Order) error {
	// Get user_id from context
	userID := utils.GetUserIDFromContext(c)
	if userID == "" {
		return fmt.Errorf("user id is empty")
	}

	// Set default values
	now := time.Now().Unix()
	order.Created = now
	order.Updated = now
	order.Status = "pending" // Default status
	u := uuid.New()
	order.ID = base64.RawURLEncoding.EncodeToString(u[:])
	order.UserID = userID // Add user_id to order

	return s.store.CreateOrder(storage.Order(*order))
}

func (s *OrderService) GetOrder(c *gin.Context, id string) (*types.Order, error) {
	order, err := s.store.GetOrder(id)
	if err != nil {
		return nil, err
	}

	if order.UserID != utils.GetUserIDFromContext(c) {
		return nil, fmt.Errorf("order is not owned by current user")
	}

	return &order, nil
}

func (s *OrderService) GetOrders(c *gin.Context, page, pageSize int) ([]types.Order, int, error) {
	userID := utils.GetUserIDFromContext(c)
	if userID == "" {
		return nil, 0, fmt.Errorf("user id is empty")
	}

	orders, err := s.store.GetOrdersByUserID(userID)
	if err != nil {
		return nil, 0, err
	}

	return orders, len(orders), nil
}

func (s *OrderService) UpdateOrder(c *gin.Context, order *types.Order) error {
	// Update timestamp
	dbOrder, err := s.store.GetOrder(order.ID)
	if err != nil {
		return err
	}

	if reflect.DeepEqual(dbOrder, *order) {
		return fmt.Errorf("no changes")
	}

	if order.UserID != dbOrder.UserID {
		return fmt.Errorf("order is not owned by current user")
	}

	order.Updated = time.Now().Unix()

	return s.store.UpdateOrder(storage.Order(*order))
}

func (s *OrderService) DeleteOrder(c *gin.Context, id string) error {
	return s.store.DeleteOrder(id)
}
