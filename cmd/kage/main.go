package main

import (
	"encoding/json"
	"fmt"
	i "github.com/thisismeamir/kage/internal/bootstrap"
	. "github.com/thisismeamir/kage/internal/runtime/handler/file"
	. "github.com/thisismeamir/kage/internal/server"
	. "github.com/thisismeamir/kage/internal/watcher"
	. "github.com/thisismeamir/kage/pkg/graph"
	. "github.com/thisismeamir/kage/pkg/mapping"
	. "github.com/thisismeamir/kage/pkg/node"
	"log"
	"os"
)

func main() {
	config := i.LoadConfiguration(i.GetConfigPath())
	i.SetGlobalConfig(config)
	serverAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	art, _ := os.ReadFile("./util/asci-art")
	fmt.Println(string(art))
	// Initialize and runs the file system watcher to catch any new .katom or .kmodule files
	// This will allow the server to automatically reload the files when they are changed.
	go func() {
		err := WatchPaths(func(event FileSystemEvent) {
			// Check if the event is a json file:
			if event.Path[len(event.Path)-5:] == ".json" {
				log.Println("JSON:", event.Path)
				var hint TypeHint
				data, _ := os.ReadFile(event.Path)
				if err := json.Unmarshal(data, &hint); err != nil {
					log.Printf("JSON has no type. Therefore, skipped. %v", err)
				}

				switch hint.Type {
				case "node":
					var node Node
					if err := json.Unmarshal(data, &node); err != nil {
						log.Fatal("failed to unmarshal node: %w", err)
					} else {
						log.Println("node:", node)
					}

				case "map":
					var m Map
					if err := json.Unmarshal(data, &m); err != nil {
						log.Fatal("failed to unmarshal map: %w", err)
					} else {
						log.Println("map:", m)
					}

				case "graph":
					var g Graph
					if err := json.Unmarshal(data, &g); err != nil {
						log.Fatal("failed to unmarshal graph: %w", err)
					} else {
						log.Println("graph:", g)
					}

				default:
					log.Fatal("unknown type:", hint.Type)
				}
			}
		})
		if err != nil {
		}

	}()

	// Start the server
	log.Println("Starting Server in", serverAddr)
	srv := New()
	if err := srv.Start(serverAddr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
