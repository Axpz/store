package service

import (
	"time"

	"github.com/Axpz/store/internal/storage"
	"github.com/Axpz/store/internal/types"
	"github.com/Axpz/store/internal/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// UserService
type UserService struct {
	store storage.StoreInterface
}

// NewUserService
func NewUserService(store storage.StoreInterface) *UserService {
	return &UserService{
		store: store,
	}
}

// CreateUser
func (s *UserService) CreateUser(c *gin.Context, user *types.User) error {

	// create user object
	user.ID = utils.GetUserIDFromEmail(user.Email)
	user.Created = time.Now().Unix()
	user.Updated = time.Now().Unix()
	user.LastLogin = time.Now().Unix()
	user.Verified = &[]bool{false}[0]

	logger := utils.LoggerFromContext(c.Request.Context())

	// save user
	if err := s.store.Create(storage.User(*user)); err != nil {
		logger.Error("create user failed", zap.Error(err))
		return err
	}

	return nil
}

// GetUser
func (s *UserService) GetUser(c *gin.Context, id string) (*types.User, error) {
	logger := utils.LoggerFromContext(c.Request.Context())
	user, err := s.store.Get(id)
	if err != nil {
		logger.Error("get user failed", zap.Error(err))
		return nil, err
	}

	return &user, nil
}

// UpdateUser
func (s *UserService) UpdateUser(c *gin.Context, user *types.User) error {
	logger := utils.LoggerFromContext(c.Request.Context())

	existingUser, err := s.store.Get(user.ID)
	if err != nil {
		logger.Error("get user failed", zap.Error(err))
		return err
	}

	if user.Username != "" {
		existingUser.Username = user.Username
	}
	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.Password != "" {
		existingUser.Password = user.Password
	}
	if user.Verified != nil && *user.Verified {
		existingUser.Verified = user.Verified
	}

	existingUser.Plan = user.Plan
	existingUser.Updated = time.Now().Unix()

	// save update
	if err := s.store.Update(existingUser); err != nil {
		logger.Error("update user failed", zap.Error(err))
		return err
	}

	return nil
}

// UpdateUserLastLogin
func (s *UserService) UpdateUserLastLogin(c *gin.Context, id string) error {
	logger := utils.LoggerFromContext(c.Request.Context())

	existingUser, err := s.store.Get(id)
	if err != nil {
		logger.Error("get user failed", zap.Error(err))
		return err
	}

	existingUser.LastLogin = time.Now().Unix()

	// save update
	if err := s.store.Update(existingUser); err != nil {
		logger.Error("update user failed", zap.Error(err))
		return err
	}

	return nil
}

// DeleteUser
func (s *UserService) DeleteUser(c *gin.Context, id string) error {
	logger := utils.LoggerFromContext(c.Request.Context())

	if _, err := s.store.Get(id); err != nil {
		logger.Error("user not found", zap.Error(err))
		return err
	}

	// delete user
	if err := s.store.Delete(id); err != nil {
		logger.Error("delete user failed", zap.Error(err))
		return err
	}

	return nil
}
