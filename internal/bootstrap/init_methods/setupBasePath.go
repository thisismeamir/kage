package init_methods

import (
	"fmt"
	. "github.com/thisismeamir/kage/internal/internal-pkg/config"
	"os"
	"path/filepath"
)

func SetupBasePath(cfg Config) error {
	base := cfg.BasePath
	if base == "" {
		return fmt.Errorf("base path is empty")
	}

	// Directories to create under base path
	subdirs := []string{
		"data",
		"data/sources",
		"data/sources/nodes",
		"data/sources/maps",
		"data/sources/graphs",
		"logs",
		"cache/flows",
		"tmp",
	}

	for _, sub := range subdirs {
		fullPath := filepath.Join(base, sub)
		if err := os.MkdirAll(fullPath, 0755); err != nil {
			return fmt.Errorf("failed to create subdir '%s': %w", sub, err)
		}
	}

	// Ensure registry files exist (even if empty)
	registryFiles := []string{
		"data/registry.json",
	}

	for _, file := range registryFiles {
		fullPath := filepath.Join(base, file)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			if err := os.WriteFile(fullPath, []byte(`{"items": []}`), 0644); err != nil {
				return fmt.Errorf("failed to create registry file '%s': %w", file, err)
			}
		}
	}

	return nil
}
