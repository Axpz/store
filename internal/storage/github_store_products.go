package storage

import (
	"fmt"
	"sort"
)

// Create 创建新
func (s *GitHubStore) CreateProduct(product Product) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 确保用户表已加载
	if err := s.loadProducts(); err != nil {
		return err
	}

	if _, exists := s.products[product.ID]; exists {
		return fmt.Errorf("商品已存在")
	}

	s.products[product.ID] = product
	return s.saveProducts()
}

// Get 获取商品
func (s *GitHubStore) GetProduct(id string) (Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result Product

	// 确保用户表已加载
	if err := s.loadProducts(); err != nil {
		return result, err
	}

	result, exists := s.products[id]
	if !exists {
		return result, fmt.Errorf("商品不存在")
	}

	return result, nil
}

// GetProducts 获取所有商品
func (s *GitHubStore) GetProducts() ([]Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 确保用户表已加载
	if err := s.loadProducts(); err != nil {
		return nil, err
	}

	var products []Product
	for _, product := range s.products {
		products = append(products, product)
	}

	sort.Slice(products, func(i, j int) bool {
		return products[i].Name < products[j].Name
	})

	return products, nil
}

// Update 更新商品
func (s *GitHubStore) UpdateProduct(product Product) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 确保用户表已加载
	if err := s.loadProducts(); err != nil {
		return err
	}

	if _, exists := s.products[product.ID]; !exists {
		return fmt.Errorf("商品不存在")
	}

	s.products[product.ID] = product
	return s.saveProducts()
}

// Delete 删除商品
func (s *GitHubStore) DeleteProduct(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 确保用户表已加载
	if err := s.loadProducts(); err != nil {
		return err
	}

	if _, exists := s.products[id]; !exists {
		return fmt.Errorf("商品不存在")
	}

	delete(s.products, id)
	return s.saveProducts()
}

func (s *GitHubStore) loadProducts() error {
	if s.loaded["products"] {
		return nil
	}

	if err := s.loadTable("products", &s.products); err != nil {
		return fmt.Errorf("加载商品表失败: %v", err)
	}
	s.loaded["products"] = true
	return nil
}

func (s *GitHubStore) saveProducts() error {
	return s.saveTable("products", s.products)
}
