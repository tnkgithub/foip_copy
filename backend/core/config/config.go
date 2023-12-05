package config

import (
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"

	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	Server  ServerConfig `yaml:"server"`
	Livekit LivekitConfig
}

type ServerConfig struct {
	Port string                `yaml:"port"`
	CORS middleware.CORSConfig `yaml:"cors"`
}

type LivekitConfig struct {
	Keys map[string]string `yaml:"keys"`

	AccessKey          string
	SecretKey          string
	TokenValidDuration time.Duration
}

var defaultConfig = Config{
	Server: ServerConfig{
		Port: ":8080",
		CORS: middleware.DefaultCORSConfig,
	},
	Livekit: LivekitConfig{
		TokenValidDuration: 30 * 24 * time.Hour,
	},
}

func New(cxt *cli.Context) (*Config, error) {
	var cfg Config = defaultConfig
	if path := cxt.String("config-path"); path != "" {
		b, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		if err := yaml.Unmarshal(b, &cfg.Livekit); err != nil {
			return nil, err
		}
	}

	for key, value := range cfg.Livekit.Keys {
		cfg.Livekit.AccessKey = key
		cfg.Livekit.SecretKey = value
		break
	}

	return &cfg, nil
}
