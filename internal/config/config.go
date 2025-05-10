package config

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// Config 表示应用程序配置
type Config struct {
	GitHub  GitHubConfig  `yaml:"github"`
	Server  ServerConfig  `yaml:"server"`
	Storage StorageConfig `yaml:"storage"`
	JWT     JWTConfig     `yaml:"jwt"`
	Email   EmailConfig   `yaml:"email"`
	PayPal  PayPalConfig  `yaml:"paypal"`

	Logger *zap.Logger
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
	Orders   string `yaml:"orders"`
	Products string `yaml:"products"`
}

// ServerConfig
type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

// StorageConfig
type StorageConfig struct {
	Type string `yaml:"type"`
	Path string `yaml:"path"`
}

// JWTConfig
type JWTConfig struct {
	Secret string        `yaml:"secret"`
	Expire time.Duration `yaml:"expire"`
}

// Load reads configuration from the specified file path.
// If the path is empty, it defaults to "config.yaml".
// The function will exit the program immediately if any error occurs.
func Load(configPath string) *Config {
	// Use default config path if none is provided
	if configPath == "" {
		configPath = "config.yaml"
	}

	log.Printf("load config from %s", configPath)

	// Read the configuration file
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("fatal: failed to read config file: %v", err)
	}

	// Parse the YAML data into the Config struct
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		log.Fatalf("fatal: failed to parse config file: %v", err)
	}

	// Load sensitive values from environment variables
	config.Email.Password = os.Getenv(config.Email.PasswordEnv)
	config.GitHub.Token = os.Getenv("GITHUB_API_TOKEN")

	// Validate the full configuration and exit on fatal errors
	config.validate()

	config.loadPaypalEnv()
	config.validatePaypalEnv()

	return &config
}

// validate checks the entire configuration and exits the program if any critical error is found.
func (c *Config) validate() {
	c.validateStorageConfig()
	c.validateGitHubConfig()
	c.validateJWTConfig()
	c.validateEmailConfig()
}

// validateStorageConfig ensures the storage configuration is valid.
func (c *Config) validateStorageConfig() {
	if c.Storage.Type == "" {
		log.Fatal("fatal: storage.type is required")
	}

	switch c.Storage.Type {
	case "github":
		// GitHub-specific validation is handled in validateGitHubConfig
	case "local":
		if c.Storage.Path == "" {
			log.Fatal("fatal: storage.path is required when storage.type is 'local'")
		}
	default:
		log.Fatalf("fatal: invalid storage.type: %s (expected 'github' or 'local')", c.Storage.Type)
	}
}

// validateGitHubConfig checks GitHub storage configuration if storage.type is 'github'.
func (c *Config) validateGitHubConfig() {
	if c.Storage.Type != "github" {
		return
	}

	if c.GitHub.Token == "" {
		log.Fatal("fatal: GITHUB_API_TOKEN is required for GitHub storage")
	}
	if c.GitHub.Repo.Owner == "" {
		log.Fatal("fatal: github.repo.owner is required")
	}
	if c.GitHub.Repo.Name == "" {
		log.Fatal("fatal: github.repo.name is required")
	}
	if c.GitHub.Repo.Branch == "" {
		log.Fatal("fatal: github.repo.branch is required")
	}
	if c.GitHub.Repo.Tables.Path == "" {
		log.Fatal("fatal: github.repo.tables.path is required")
	}
	if c.GitHub.Repo.Tables.Users == "" {
		log.Fatal("fatal: github.repo.tables.users is required")
	}
	if c.GitHub.Repo.Tables.Comments == "" {
		log.Fatal("fatal: github.repo.tables.comments is required")
	}
}

// validateJWTConfig checks the JWT configuration and applies defaults if not set.
func (c *Config) validateJWTConfig() {
	if c.JWT.Secret == "" {
		c.JWT.Secret = c.GitHub.Token
		log.Println("info: jwt.secret not set, using github.token as fallback")
	}
	if c.JWT.Expire == 0 {
		c.JWT.Expire = 7 * 24 * time.Hour
		log.Println("info: jwt.expire not set, using default value of 7 days")
	}
}

// GetTablePath tables path
func (c *Config) GetTablePath(tableName string) string {
	switch tableName {
	case "users":
		return filepath.Join(c.GitHub.Repo.Tables.Path, c.GitHub.Repo.Tables.Users)
	case "comments":
		return filepath.Join(c.GitHub.Repo.Tables.Path, c.GitHub.Repo.Tables.Comments)
	case "orders":
		return filepath.Join(c.GitHub.Repo.Tables.Path, c.GitHub.Repo.Tables.Orders)
	case "products":
		return filepath.Join(c.GitHub.Repo.Tables.Path, c.GitHub.Repo.Tables.Products)
	default:
		return ""
	}
}
