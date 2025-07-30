package bootstrap

import (
	config2 "github.com/thisismeamir/kage/internal/bootstrap/config"
	"github.com/thisismeamir/kage/internal/watcher"
)

func InitKage() (string, config2.Config, watcher.Watcher) {
	// Load the configuration
	config := config2.LoadConfiguration(config2.GetConfigPath())
	config2.SetGlobalConfig(config)

	// Load and validate registries
	InitRegistries()
	// Initialize
}
