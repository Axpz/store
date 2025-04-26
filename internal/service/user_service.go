package service

import (
	"fmt"
	"time"

	"github.com/Axpz/store/internal/storage"
	"github.com/Axpz/store/internal/types"
	"github.com/Axpz/store/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
func (s *UserService) CreateUser(c *gin.Context, user *types.User) error {

	// 创建用户对象
	user.ID = utils.GetUserIDFromEmail(user.Email)
	user.Created = time.Now().Unix()
	user.Updated = time.Now().Unix()

	logger := utils.LoggerFromContext(c.Request.Context())

	// 保存用户
	if err := s.store.Create(storage.User(*user)); err != nil {
		logger.Error("创建用户失败", zap.Error(err))
		return fmt.Errorf("创建用户失败: %v", err)
	}

	return nil
}

// GetUser 获取用户信息
func (s *UserService) GetUser(c *gin.Context, id string) (*types.User, error) {
	logger := utils.LoggerFromContext(c.Request.Context())
	logger.Info("获取用户", zap.String("id", id))
	user, err := s.store.Get(id)
	if err != nil {
		return nil, fmt.Errorf("获取用户失败: %v", err)
	}

	return &user, nil
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(c *gin.Context, id, username, email, plan string) error {
	// 获取现有用户
	existingUser, err := s.store.Get(id)
	if err != nil {
		return fmt.Errorf("获取用户失败: %v", err)
	}

	// 更新用户信息
	existingUser.Username = username
	existingUser.Email = email
	existingUser.Plan = plan
	existingUser.Updated = time.Now().Unix()

	// 保存更新
	if err := s.store.Update(existingUser); err != nil {
		return fmt.Errorf("更新用户失败: %v", err)
	}

	return nil
}

// DeleteUser 删除用户
func (s *UserService) DeleteUser(c *gin.Context, id string) error {
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
