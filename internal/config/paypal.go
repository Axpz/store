package config

import (
	"log"
	"os"
)

type PayPalConfig struct {
	ClientIDEnv     string `yaml:"client_id_env"`
	ClientSecretEnv string `yaml:"client_secret_env"`
	WebhookIDEnv    string `yaml:"webhook_id_env"`
	ClientID        string `yaml:"client_id"`
	ClientSecret    string `yaml:"client_secret"`
	WebhookID       string `yaml:"webhook_id"`
	Environment     string `yaml:"environment"`
}

func (c *Config) loadPaypalEnv() {
	pp := &c.PayPal

	if envKey := pp.ClientIDEnv; envKey != "" {
		if val := os.Getenv(envKey); val != "" {
			pp.ClientID = val
		}
	}

	if envKey := pp.ClientSecretEnv; envKey != "" {
		if val := os.Getenv(envKey); val != "" {
			pp.ClientSecret = val
		}
	}

	if envKey := pp.WebhookIDEnv; envKey != "" {
		if val := os.Getenv(envKey); val != "" {
			pp.WebhookID = val
		}
	}
}

func (c *Config) validatePaypalEnv() {
	pp := &c.PayPal

	if pp.ClientID == "" {
		log.Fatalf("missing PayPal ClientID (env: %s)", pp.ClientIDEnv)
	}
	if pp.ClientSecret == "" {
		log.Fatalf("missing PayPal ClientSecret (env: %s)", pp.ClientSecretEnv)
	}
	if pp.WebhookID == "" {
		log.Fatalf("missing PayPal WebhookID (env: %s)", pp.WebhookIDEnv)
	}
	if pp.Environment != "" && pp.Environment != "sandbox" && pp.Environment != "production" {
		log.Fatalf("invalid PayPal Environment: %s (must be 'sandbox' or 'production')", pp.Environment)
	}
}
