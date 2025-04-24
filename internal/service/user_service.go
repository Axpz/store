package service

import (
	"fmt"
	"time"

	"github.com/Axpz/store/internal/storage"
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
func (s *UserService) CreateUser(username, email, plan string) (*storage.User, error) {
	// 生成用户ID (在实际应用中，可能需要更复杂的ID生成策略)
	userID := fmt.Sprintf("user_%d", time.Now().UnixNano())

	// 创建用户对象
	user := storage.User{
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
func (s *UserService) GetUser(id string) (*storage.User, error) {
	user, err := s.store.Get(id)
	if err != nil {
		return nil, fmt.Errorf("获取用户失败: %v", err)
	}

	return &user, nil
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(id, username, email, plan string) (*storage.User, error) {
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
	if err := s.store.Update(existingUser); err != nil {
		return nil, fmt.Errorf("更新用户失败: %v", err)
	}

	return &existingUser, nil
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

// CreateComment 创建评论
func (s *UserService) CreateComment(userID, content string) (*storage.Comment, error) {
	// 检查用户是否存在
	if _, err := s.store.Get(userID); err != nil {
		return nil, fmt.Errorf("用户不存在: %v", err)
	}

	// 生成评论ID
	commentID := fmt.Sprintf("comment_%d", time.Now().UnixNano())

	// 创建评论对象
	comment := storage.Comment{
		ID:      commentID,
		UserID:  userID,
		Content: content,
		Created: time.Now().Unix(),
		Updated: time.Now().Unix(),
	}

	// 保存评论
	if err := s.store.CreateComment(comment); err != nil {
		return nil, fmt.Errorf("创建评论失败: %v", err)
	}

	return &comment, nil
}

// GetComment 获取评论
func (s *UserService) GetComment(id string) (*storage.Comment, error) {
	comment, err := s.store.GetComment(id)
	if err != nil {
		return nil, fmt.Errorf("获取评论失败: %v", err)
	}

	return &comment, nil
}

// UpdateComment 更新评论
func (s *UserService) UpdateComment(id, content string) (*storage.Comment, error) {
	// 获取现有评论
	existingComment, err := s.store.GetComment(id)
	if err != nil {
		return nil, fmt.Errorf("获取评论失败: %v", err)
	}

	// 更新评论内容
	existingComment.Content = content
	existingComment.Updated = time.Now().Unix()

	// 保存更新
	if err := s.store.UpdateComment(existingComment); err != nil {
		return nil, fmt.Errorf("更新评论失败: %v", err)
	}

	return &existingComment, nil
}

// DeleteComment 删除评论
func (s *UserService) DeleteComment(id string) error {
	// 检查评论是否存在
	if _, err := s.store.GetComment(id); err != nil {
		return fmt.Errorf("评论不存在: %v", err)
	}

	// 删除评论
	if err := s.store.DeleteComment(id); err != nil {
		return fmt.Errorf("删除评论失败: %v", err)
	}

	return nil
}
