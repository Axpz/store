package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Axpz/store/internal/config"
)

// LocalStore 实现本地文件存储
type LocalStore struct {
	Store
}

// NewLocalStore 创建一个新的本地存储
func NewLocalStore(cfg *config.Config) (StoreInterface, error) {
	store := &LocalStore{
		Store: NewStore(cfg),
	}

	return store, nil
}

// Create 创建新用户
func (s *LocalStore) Create(user User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 确保用户表已加载
	if err := s.loadUsers(); err != nil {
		return err
	}

	if _, exists := s.users[user.ID]; exists {
		return fmt.Errorf("用户已存在")
	}

	s.users[user.ID] = user
	return s.saveUsers()
}

// Get 获取用户
func (s *LocalStore) Get(id string) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result User

	// 确保用户表已加载
	if err := s.loadUsers(); err != nil {
		return result, err
	}

	user, exists := s.users[id]
	if !exists {
		return result, fmt.Errorf("用户不存在")
	}

	result = user
	return result, nil
}

// Update 更新用户
func (s *LocalStore) Update(user User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.loadUsers(); err != nil {
		return err
	}

	if _, exists := s.users[user.ID]; !exists {
		return fmt.Errorf("用户不存在")
	}

	s.users[user.ID] = user
	return s.saveUsers()
}

// Delete 删除用户
func (s *LocalStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.loadUsers(); err != nil {
		return err
	}

	if _, exists := s.users[id]; !exists {
		return fmt.Errorf("用户不存在")
	}

	delete(s.users, id)
	return s.saveUsers()
}

// CreateOrder 创建新订单
func (s *LocalStore) CreateOrder(order Order) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	return nil
}

// GetOrder 获取订单
func (s *LocalStore) GetOrder(id string) (Order, error) {
	return Order{}, nil
}

// GetOrdersByUserID 获取用户订单
func (s *LocalStore) GetOrdersByUserID(userID string) ([]Order, error) {
	return nil, nil
}

// UpdateOrder 更新订单
func (s *LocalStore) UpdateOrder(order Order) error {
	return nil
}

// DeleteOrder 删除订单
func (s *LocalStore) DeleteOrder(id string) error {
	return nil
}

func (s *LocalStore) CreateProduct(product Product) error {
	return nil
}

func (s *LocalStore) GetProduct(id string) (Product, error) {
	return Product{}, nil
}

func (s *LocalStore) GetProducts() ([]Product, error) {
	return nil, nil
}

func (s *LocalStore) UpdateProduct(product Product) error {
	return nil
}

func (s *LocalStore) DeleteProduct(id string) error {
	return nil
}

// CreateComment 创建新评论
func (s *LocalStore) CreateComment(comment Comment) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.loadComments(); err != nil {
		return err
	}

	if _, exists := s.comments[comment.ID]; exists {
		return fmt.Errorf("评论已存在")
	}

	s.comments[comment.ID] = comment
	return s.saveComments()
}

// GetComment 获取评论
func (s *LocalStore) GetComment(id string) (Comment, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result Comment

	if err := s.loadComments(); err != nil {
		return result, err
	}

	result, exists := s.comments[id]
	if !exists {
		return result, fmt.Errorf("评论不存在")
	}

	return result, nil
}

// UpdateComment 更新评论
func (s *LocalStore) UpdateComment(comment Comment) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.loadComments(); err != nil {
		return err
	}

	if _, exists := s.comments[comment.ID]; !exists {
		return fmt.Errorf("评论不存在")
	}

	s.comments[comment.ID] = comment
	return s.saveComments()
}

// DeleteComment 删除评论
func (s *LocalStore) DeleteComment(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.loadComments(); err != nil {
		return err
	}

	if _, exists := s.comments[id]; !exists {
		return fmt.Errorf("评论不存在")
	}

	delete(s.comments, id)
	return s.saveComments()
}

// loadTable 加载指定表的数据
func (s *LocalStore) loadTable(tableName string, data any) error {
	// 构建文件路径
	filePath := filepath.Join(s.config.Storage.Path, tableName)

	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		// 如果文件不存在，创建一个空文件
		emptyData := make(map[string]any)
		if err := s.saveTable(tableName, emptyData); err != nil {
			return fmt.Errorf("创建空文件失败: %v", err)
		}
		return nil
	}

	// 读取文件内容
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("读取文件失败: %v", err)
	}

	// 解析 JSON
	if err := json.Unmarshal(fileData, data); err != nil {
		return fmt.Errorf("解析 JSON 失败: %v", err)
	}

	return nil
}

// saveTable 保存指定表的数据
func (s *LocalStore) saveTable(tableName string, data any) error {
	// 构建文件路径
	filePath := filepath.Join(s.config.Storage.Path, tableName)

	// 序列化为 JSON
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return fmt.Errorf("序列化 JSON 失败: %v", err)
	}

	// 写入文件
	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("写入文件失败: %v", err)
	}

	return nil
}

func (s *LocalStore) loadUsers() error {
	if s.loaded["users"] {
		return nil
	}

	if err := s.loadTable("users", &s.users); err != nil {
		return fmt.Errorf("加载用户表失败: %v", err)
	}

	s.loaded["users"] = true
	return nil
}

func (s *LocalStore) saveUsers() error {
	return s.saveTable("users", s.users)
}

func (s *LocalStore) loadComments() error {
	if s.loaded["comments"] {
		return nil
	}

	if err := s.loadTable("comments", &s.comments); err != nil {
		return fmt.Errorf("加载评论表失败: %v", err)
	}
	s.loaded["comments"] = true
	return nil
}

func (s *LocalStore) saveComments() error {
	return s.saveTable("comments", s.comments)
}
