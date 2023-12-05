package config

import (
	"os"
)

var (
	TrustedProxies []string
)

func init() {
	TrustedProxies = nil
	if mode := os.Getenv("GIN_MODE"); mode == "release" {
		if frontAppAddr := os.Getenv("FRONT_APP_ADDR"); frontAppAddr != "" {
			TrustedProxies = []string{frontAppAddr}
		}
	}
}
