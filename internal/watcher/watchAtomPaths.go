package watcher

import (
	i "github.com/thisismeamir/kage/internal/bootstrap"
	"github.com/thisismeamir/kage/pkg/atom"
	"log"
)

func WatchAtomPaths(callback func(FileSystemEvent)) error {
	var localPaths []atom.AtomPath

	for _, path := range i.GetGlobalConfig().AtomPaths {
		if path.Local {
			localPaths = append(localPaths, path)
		}
	}

	// Map localPaths to their Path property
	var paths []string
	for _, item := range localPaths {
		paths = append(paths, item.Path)
	}
	log.Printf("Local Paths: %v", paths)
	watcher, err := NewWatcher(paths, callback)
	if err != nil {
		return err
	}

	// Start watching (if needed)
	return watcher.Start()
}
