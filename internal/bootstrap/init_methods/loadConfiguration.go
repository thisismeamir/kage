package init_methods

import (
	"encoding/json"
	"errors"
	"fmt"
	. "github.com/thisismeamir/kage/internal/internal-pkg/config"
	"os"
	"path/filepath"
	"strings"
)

/*
*
Securely loads a configuration file, fills empty information with defaults
*/
func LoadConfiguration(path string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(path)
	if err != nil {
		return cfg, fmt.Errorf("failed to read config file: %w", err)
	}

	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("failed to parse config JSON: %w", err)
	}

	// ---- Set Defaults ----

	if cfg.Name == "" {
		cfg.Name = "Kage"
	}
	if cfg.BasePath == "" {
		return cfg, errors.New("base_path is required")
	}
	if len(cfg.Paths) == 0 {
		cfg.Paths = []string{filepath.Join(cfg.BasePath, "data/sources")}
	}
	if cfg.Version == "" {
		cfg.Version = "1.0.0"
	}

	// Server defaults
	if cfg.Server.Port == 0 {
		cfg.Server.Port = 8080
	}
	if cfg.Server.Host == "" {
		cfg.Server.Host = "localhost"
	}
	if cfg.Server.Api.BaseUrl == "" {
		cfg.Server.Api.BaseUrl = "/api"
	}
	if cfg.Server.Api.Version == "" {
		cfg.Server.Api.Version = "v1"
	}
	if cfg.Server.Logging.Level == "" {
		cfg.Server.Logging.Level = "info"
	}
	// Expand $BasePath
	if cfg.Server.Logging.File == "" {
		cfg.Server.Logging.File = filepath.Join(cfg.BasePath, "logs/kage.log")
	} else {
		cfg.Server.Logging.File = strings.ReplaceAll(
			cfg.Server.Logging.File,
			"$BasePath", cfg.BasePath,
		)
	}

	// Client Web defaults
	web := &cfg.Client.Web
	if web.Port == 0 {
		web.Port = 3000
	}
	if web.Host == "" {
		web.Host = "localhost"
	}
	if web.Path == "" {
		web.Path = "./web"
	}
	if web.Build == "" {
		web.Build = "npm build"
	}
	if web.Run == "" {
		web.Run = "npm run"
	}

	return cfg, nil
}
