package watcher

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

type FileSystemEvent struct {
	Path  string
	Event fsnotify.Op
}

type Watcher struct {
	paths    []string
	callback func(FileSystemEvent)
	watcher  *fsnotify.Watcher
}

func NewWatcher(paths []string, callback func(FileSystemEvent)) (*Watcher, error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return &Watcher{
		paths:    paths,
		callback: callback,
		watcher:  w,
	}, nil
}

func (w *Watcher) Start() error {
	for _, path := range w.paths {
		if err := w.watcher.Add(path); err != nil {
			return err
		}
	}
	go func() {
		for {
			select {
			case event, ok := <-w.watcher.Events:
				if !ok {
					return
				}
				if w.callback != nil {
					w.callback(FileSystemEvent{Path: event.Name, Event: event.Op})
				}
			case err, ok := <-w.watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	return nil
}

func (w *Watcher) Close() error {
	return w.watcher.Close()
}
