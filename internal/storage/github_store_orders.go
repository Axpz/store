package storage

import (
	"fmt"

	"go.uber.org/zap"
)

// Create 创建新用户
func (s *GitHubStore) CreateOrder(order Order) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 确保用户表已加载
	if err := s.loadOrders(); err != nil {
		return err
	}

	if _, exists := s.orders[order.ID]; exists {
		return fmt.Errorf("订单已存在")
	}

	s.orders[order.ID] = order
	return s.saveOrders()
}

// Get 获取用户
func (s *GitHubStore) GetOrder(id string) (Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result Order

	// 确保用户表已加载
	if err := s.loadOrders(); err != nil {
		return result, err
	}

	result, exists := s.orders[id]
	if !exists {
		return result, fmt.Errorf("订单不存在")
	}

	return result, nil
}

// Get 获取用户
func (s *GitHubStore) GetOrdersByUserID(userID string) ([]Order, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []Order

	logger := s.Logger()
	logger.Info("GetOrdersByUserID", zap.String("userID", userID))

	// 确保用户表已加载
	if err := s.loadOrders(); err != nil {
		return result, err
	}

	for _, order := range s.orders {
		if order.UserID == userID {
			result = append(result, order)
		}
	}

	return result, nil
}

// Update 更新用户
func (s *GitHubStore) UpdateOrder(order Order) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 确保用户表已加载
	if err := s.loadOrders(); err != nil {
		return err
	}

	if _, exists := s.orders[order.ID]; !exists {
		return fmt.Errorf("订单不存在")
	}

	s.orders[order.ID] = order
	return s.saveOrders()
}

// Delete 删除用户
func (s *GitHubStore) DeleteOrder(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 确保用户表已加载
	if err := s.loadOrders(); err != nil {
		return err
	}

	if _, exists := s.orders[id]; !exists {
		return fmt.Errorf("订单不存在")
	}

	delete(s.orders, id)
	return s.saveOrders()
}

func (s *GitHubStore) loadOrders() error {
	if s.loaded["orders"] {
		return nil
	}

	if err := s.loadTable("orders", &s.orders); err != nil {
		return fmt.Errorf("加载订单表失败: %v", err)
	}
	s.loaded["orders"] = true
	return nil
}

func (s *GitHubStore) saveOrders() error {
	return s.saveTable("orders", s.orders)
}
