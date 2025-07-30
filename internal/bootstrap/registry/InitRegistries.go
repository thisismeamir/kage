package registry

import (
	"log"
)

func InitRegistries() error {
	// Initialize the graph registry
	InitGraphRegistry()

	// Initialize the node registry
	InitNodeRegistry()

	// Initialize the map registry
	InitMapRegistry()

}
