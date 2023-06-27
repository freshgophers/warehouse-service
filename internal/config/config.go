package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	defaultHTTPPort               = "80"
	defaultHTTPHost               = "localhost"
	defaultHTTPSchema             = "http"
	defaultHTTPReadTimeout        = 15 * time.Second
	defaultHTTPWriteTimeout       = 15 * time.Second
	defaultHTTPIdleTimeout        = 60 * time.Second
	defaultHTTPMaxHeaderMegabytes = 1
)

type (
	Configs struct {
		HTTP     HTTPConfig
		POSTGRES DatabaseConfig
	}

	HTTPConfig struct {
		Port               string
		Host               string
		Schema             string
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
		IdleTimeout        time.Duration
		MaxHeaderMegabytes int
	}

	ClientConfig struct {
		Endpoint string
		Username string
		Password string
	}

	DatabaseConfig struct {
		DSN string
	}
)

// New populates Configs struct with values from config file
// located at filepath and environment variables.
func New() (cfg Configs, err error) {
	root, err := os.Getwd()
	if err != nil {
		return
	}
	godotenv.Load(filepath.Join(root, ".env"))

	cfg.HTTP = HTTPConfig{
		Port:               defaultHTTPPort,
		Host:               defaultHTTPHost,
		Schema:             defaultHTTPSchema,
		ReadTimeout:        defaultHTTPReadTimeout,
		WriteTimeout:       defaultHTTPWriteTimeout,
		IdleTimeout:        defaultHTTPIdleTimeout,
		MaxHeaderMegabytes: defaultHTTPMaxHeaderMegabytes,
	}

	err = envconfig.Process("HTTP", &cfg.HTTP)
	if err != nil {
		return
	}

	err = envconfig.Process("POSTGRES", &cfg.POSTGRES)
	if err != nil {
		return
	}

	return
}
