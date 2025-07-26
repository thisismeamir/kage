package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	"github.com/thisismeamir/kage/internal/server"
	"github.com/thisismeamir/kage/internal/watcher"
	"log"
	"os"
)

func main() {
	config := i.LoadConfiguration(i.GetConfigPath())
	i.SetGlobalConfig(config)
	serverAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	art, _ := os.ReadFile("./util/asci-art")
	fmt.Println(string(art))
	go watcher.WatchAtomPaths(func(event watcher.FileSystemEvent) {
		// Checking if the change is related to .katom extensions:
		if event.Path[len(event.Path)-6:] == ".katom" {
			if event.Event == fsnotify.Create {
				log.Println("New Atom Added")
			} else if event.Event == fsnotify.Write {
				log.Println("Atom Updated")
			} else if event.Event == fsnotify.Remove {
				log.Println("Atom Removed")
			}
		}

	})

	go watcher.WatchModulePaths(func(event watcher.FileSystemEvent) {
		log.Println("module event:", event)
	})
	log.Println("Starting Server in", serverAddr)
	srv := server.New()
	if err := srv.Start(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
