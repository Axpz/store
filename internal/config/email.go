package config

import (
	"log"
	"os"
)

// EmailConfig
type EmailConfig struct {
	SMTPServer  string `yaml:"smtp_server"` // e.g. smtp.qq.com
	SMTPPort    int    `yaml:"smtp_port"`   // e.g. 465
	Username    string `yaml:"username"`    // SMTP login username
	Password    string
	PasswordEnv string `yaml:"password_env"` // env var name for auth code
	From        string `yaml:"from"`         // display name and address
}

// validateEmailConfig checks the email sending configuration and environment dependency.
func (c *Config) validateEmailConfig() {
	email := c.Email

	if email.SMTPServer == "" {
		log.Fatal("fatal: email.smtp_server is required")
	}
	if email.SMTPPort <= 0 || email.SMTPPort > 65535 {
		log.Fatalf("fatal: email.smtp_port must be between 1 and 65535, got %d", email.SMTPPort)
	}
	if email.Username == "" {
		log.Fatal("fatal: email.username is required")
	}
	if email.From == "" {
		log.Fatal("fatal: email.from is required")
	}
	if email.PasswordEnv == "" {
		log.Fatal("fatal: email.password_env is required")
	}
	if os.Getenv(email.PasswordEnv) == "" {
		log.Fatalf("fatal: environment variable %s is not set or empty", email.PasswordEnv)
	}
}
