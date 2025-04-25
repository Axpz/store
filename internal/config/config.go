package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

// Config 表示应用程序配置
type Config struct {
	GitHub  GitHubConfig  `yaml:"github"`
	Server  ServerConfig  `yaml:"server"`
	Storage StorageConfig `yaml:"storage"`
	JWT     JWTConfig     `yaml:"jwt"`
}

// GitHubConfig 表示 GitHub 相关配置
type GitHubConfig struct {
	Repo  RepoConfig `yaml:"repo"`
	Token string
}

// RepoConfig 表示仓库相关配置
type RepoConfig struct {
	Owner  string       `yaml:"owner"`
	Name   string       `yaml:"name"`
	Branch string       `yaml:"branch"`
	Tables TablesConfig `yaml:"tables"`
}

// TablesConfig 表示表配置
type TablesConfig struct {
	Path     string `yaml:"path"`
	Users    string `yaml:"users"`
	Comments string `yaml:"comments"`
}

// ServerConfig 表示服务器配置
type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

// StorageConfig 表示存储配置
type StorageConfig struct {
	Type string `yaml:"type"`
	Path string `yaml:"path"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret string        `yaml:"secret"`
	Expire time.Duration `yaml:"expire"`
}

// Load 从文件加载配置
func Load(configPath string) (*Config, error) {
	// 如果未指定配置文件路径，使用默认路径
	if configPath == "" {
		configPath = "config.yaml"
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析配置
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	// 验证配置
	if err := config.validate(); err != nil {
		return nil, err
	}

	return &config, nil
}

// validate 验证配置是否有效
func (c *Config) validate() error {
	// 检查存储类型
	if c.Storage.Type == "" {
		return fmt.Errorf("存储类型未设置")
	}

	// 如果是 GitHub 存储，检查 GitHub Token
	if c.Storage.Type == "github" {
		token := os.Getenv("GITHUB_API_TOKEN")
		if token == "" {
			return fmt.Errorf("未设置 GITHUB_TOKEN 环境变量")
		}

		c.GitHub.Token = token

		if c.GitHub.Repo.Owner == "" {
			return fmt.Errorf("仓库所有者未设置")
		}

		if c.GitHub.Repo.Name == "" {
			return fmt.Errorf("仓库名称未设置")
		}

		if c.GitHub.Repo.Branch == "" {
			return fmt.Errorf("仓库分支未设置")
		}

		if c.GitHub.Repo.Tables.Path == "" {
			return fmt.Errorf("表目录路径未设置")
		}

		if c.GitHub.Repo.Tables.Users == "" {
			return fmt.Errorf("用户表文件名未设置")
		}

		if c.GitHub.Repo.Tables.Comments == "" {
			return fmt.Errorf("评论表文件名未设置")
		}
	}

	// 如果是本地存储，检查存储路径
	if c.Storage.Type == "local" && c.Storage.Path == "" {
		return fmt.Errorf("本地存储路径未设置")
	}

	return nil
}

// GetTablePath 获取指定表的完整文件路径
func (c *Config) GetTablePath(tableName string) string {
	switch tableName {
	case "users":
		return filepath.Join(c.GitHub.Repo.Tables.Path, c.GitHub.Repo.Tables.Users)
	case "comments":
		return filepath.Join(c.GitHub.Repo.Tables.Path, c.GitHub.Repo.Tables.Comments)
	default:
		return ""
	}
}
