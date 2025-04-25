package service

import (
	"fmt"
	"time"

	"github.com/Axpz/store/internal/storage"
	"github.com/Axpz/store/internal/types"
)

// UserService 提供用户相关的业务逻辑
type UserService struct {
	store storage.StoreInterface
}

// NewUserService 创建一个新的用户服务
func NewUserService(store storage.StoreInterface) *UserService {
	return &UserService{
		store: store,
	}
}

// CreateUser 创建新用户
func (s *UserService) CreateUser(username, email, plan string) (*types.User, error) {
	// 生成用户ID (在实际应用中，可能需要更复杂的ID生成策略)
	userID := fmt.Sprintf("user_%d", time.Now().UnixNano())

	// 创建用户对象
	user := types.User{
		ID:       userID,
		Username: username,
		Email:    email,
		Plan:     plan,
		Created:  time.Now().Unix(),
		Updated:  time.Now().Unix(),
	}

	// 保存用户
	if err := s.store.Create(user); err != nil {
		return nil, fmt.Errorf("创建用户失败: %v", err)
	}

	return &user, nil
}

// GetUser 获取用户信息
func (s *UserService) GetUser(id string) (*types.User, error) {
	user, err := s.store.Get(id)
	if err != nil {
		return nil, fmt.Errorf("获取用户失败: %v", err)
	}

	return user, nil
}

func (s *UserService) GetUserByEmail(email string) (*types.User, error) {
	user, err := s.store.GetByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("获取用户失败: %v", err)
	}

	return user, nil
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(id, username, email, plan string) (*types.User, error) {
	// 获取现有用户
	existingUser, err := s.store.Get(id)
	if err != nil {
		return nil, fmt.Errorf("获取用户失败: %v", err)
	}

	// 更新用户信息
	existingUser.Username = username
	existingUser.Email = email
	existingUser.Plan = plan
	existingUser.Updated = time.Now().Unix()

	// 保存更新
	if err := s.store.Update(*existingUser); err != nil {
		return nil, fmt.Errorf("更新用户失败: %v", err)
	}

	return existingUser, nil
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(id string) error {
	// 检查用户是否存在
	if _, err := s.store.Get(id); err != nil {
		return fmt.Errorf("用户不存在: %v", err)
	}

	// 删除用户
	if err := s.store.Delete(id); err != nil {
		return fmt.Errorf("删除用户失败: %v", err)
	}

	return nil
}
