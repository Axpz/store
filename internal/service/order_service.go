package service

import (
	"fmt"
	"reflect"
	"time"

	"github.com/Axpz/store/internal/storage"
	"github.com/Axpz/store/internal/types"
	"github.com/Axpz/store/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderService struct {
	store           storage.StoreInterface
	paymentProvider PaymentProvider
}

func NewOrderService(store storage.StoreInterface, paymentProvider PaymentProvider) *OrderService {
	return &OrderService{
		store:           store,
		paymentProvider: paymentProvider,
	}
}

func (s *OrderService) CreateOrder(c *gin.Context, order *types.Order) error {
	// Get user_id from context
	userID := utils.GetUserIDFromContext(c)
	if userID == "" {
		return fmt.Errorf("user id is empty")
	}

	// Create payment order
	paymentOrderID, err := s.paymentProvider.CreateOrder(c.Request.Context(), int(order.TotalAmount), order.Currency)
	if err != nil {
		return fmt.Errorf("failed to create payment order: %v", err)
	}
	order.ID = paymentOrderID
	// Set default values
	now := time.Now().Unix()
	order.Created = now
	order.Updated = now
	order.Status = "pending" // Default status
	order.UserID = userID    // Add user_id to order

	return s.store.CreateOrder(storage.Order(*order))
}

func (s *OrderService) CaptureOrder(c *gin.Context, orderID string) error {
	order, err := s.store.GetOrder(orderID)
	if err != nil {
		return err
	}

	if order.UserID != utils.GetUserIDFromContext(c) {
		return fmt.Errorf("order is not owned by current user")
	}

	if order.Status != "pending" {
		return fmt.Errorf("order status is not pending")
	}

	// Capture payment
	err = s.paymentProvider.CaptureOrder(c.Request.Context(), order.ID)
	if err != nil {
		return fmt.Errorf("failed to capture payment: %v", err)
	}

	// Update order status
	order.Status = "completed"
	order.Updated = time.Now().Unix()
	return s.store.UpdateOrder(storage.Order(order))
}

func (s *OrderService) GetOrder(c *gin.Context, id string) (*types.Order, error) {
	order, err := s.store.GetOrder(id)
	if err != nil {
		return nil, err
	}

	if order.UserID != utils.GetUserIDFromContext(c) {
		return nil, fmt.Errorf("order is not owned by current user")
	}

	for i := range order.Products {
		p, err := s.store.GetProduct(order.Products[i].ID)
		if err != nil {
			return nil, err
		}
		order.Products[i].Content = p.Content
		order.Products[i].Name = p.Name
	}

	return &order, nil
}

func (s *OrderService) GetOrders(c *gin.Context, page, pageSize int) ([]types.Order, int, error) {
	userID := utils.GetUserIDFromContext(c)
	logger := utils.LoggerFromContext(c.Request.Context())

	if userID == "" {
		return nil, 0, fmt.Errorf("user id is empty")
	}

	orders, err := s.store.GetOrdersByUserID(userID)

	if err != nil {
		return nil, 0, err
	}

	for i := range orders {
		for j := range orders[i].Products {
			p, err := s.store.GetProduct(orders[i].Products[j].ID)
			if err != nil {
				logger.Error("ignored failed to get product", zap.Error(err))
				continue
			}
			orders[i].Products[j].Name = p.Name
			orders[i].Products[j].Content = p.Content
		}
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
