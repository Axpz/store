package storage

import (
	"fmt"

	"github.com/Axpz/store/internal/config"
)

// StorageType 表示存储类型
type StorageType string

const (
	// GitHubStorage 表示 GitHub 存储
	GitHubStorage StorageType = "github"
	// LocalStorage 表示本地存储
	LocalStorage StorageType = "local"
)

// NewStore 创建一个新的存储实例
func New(cfg *config.Config) (StoreInterface, error) {
	switch cfg.Storage.Type {
	case "local":
		return NewLocalStore(cfg)
	case "github":
		return NewGitHubStore(cfg)
	default:
		return nil, fmt.Errorf("不支持的存储类型: %s", cfg.Storage.Type)
	}
}
