package watcher

import (
	i "github.com/thisismeamir/kage/internal/bootstrap"
	module "github.com/thisismeamir/kage/pkg/module"
	"log"
)

func WatchModulePaths(callback func(FileSystemEvent)) error {
	var localPaths []module.ModulePath

	for _, path := range i.GetGlobalConfig().ModulePaths {
		if path.Local {
			localPaths = append(localPaths, path)
		}
	}

	// Map localPaths to their Path property
	var paths []string
	for _, item := range localPaths {
		paths = append(paths, item.Path)
	}
	log.Printf("Local Module Paths: %v", paths)
	watcher, err := NewWatcher(paths, callback)
	if err != nil {
		return err
	}

	// Start watching (if needed)
	return watcher.Start()
}
