package event

import (
	"github.com/fsnotify/fsnotify"
	. "github.com/thisismeamir/kage/internal/runtime/handler/file"
	. "github.com/thisismeamir/kage/internal/watcher"
	"log"
)

func FileAddedEvent(event FileSystemEvent, formType string) {
	switch formType {
	case "node":
		switch event.Event {
		case fsnotify.Create:
			log.Println("Node Created/Updated:", event.Path)
			node, err := LoadForm(event.Path)
			if err != nil {
				log.Fatal("Error loading node:", err)
			} else {
				log.Println("Node:", node)
			}
		case fsnotify.Remove:
			log.Println("Node Removed:", event.Path)
		default:
			log.Println("Unknown Event for Node:", event.Event)
		}
	case "map":
		switch event.Event {
		case fsnotify.Create:
			log.Println("Map Created/Updated:", event.Path)
		case fsnotify.Remove:
			log.Println("Map Removed:", event.Path)
		default:
			log.Println("Unknown Event for Map:", event.Event)
		}
	case "graph":
		switch event.Event {
		case fsnotify.Create:
			log.Println("Graph Created/Updated:", event.Path)
		case fsnotify.Remove:
			log.Println("Graph Removed:", event.Path)
		default:
			log.Println("Unknown Event for Graph:", event.Event)
		}
	default:
		log.Println("Unknown Form Type:", formType)

	}
}
