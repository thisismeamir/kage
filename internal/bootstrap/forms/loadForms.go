package forms

import (
	"github.com/thisismeamir/kage/internal/bootstrap"
	"github.com/thisismeamir/kage/internal/runtime/handler/event"
	"github.com/thisismeamir/kage/internal/runtime/handler/file"
	"github.com/thisismeamir/kage/pkg/graph"
	"github.com/thisismeamir/kage/pkg/mapping"
	"github.com/thisismeamir/kage/pkg/node"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func LoadForms() {
	nodePath := bootstrap.GetGlobalConfig().NodePaths
	graphPhath := bootstrap.GetGlobalConfig().GraphPaths
	mapPath := bootstrap.GetGlobalConfig().MapPaths

	paths := append(nodePath, graphPhath...)
	paths = append(paths, mapPath...)
	for _, path := range paths {
		jsonFiles := FindJsonInPath(path.Path)
		for _, jsonFile := range jsonFiles {
			log.Println("Loading form from path:", jsonFile)
			form, err := file.LoadForm(jsonFile)
			if err != nil {
				log.Printf("Error loading form from %s: %v", jsonFile, err)
				continue
			}
			switch v := form.(type) {
			case node.Node:
				event.AddToNodeRegistry(v)
			case mapping.Map:
				event.AddToMapRegister(v)
			case graph.Graph:
				event.AddToGraphRegistry(v)
			default:
				log.Printf("Unknown form type: %T", v)
			}
		}
	}

}

// returns all the json file paths in the given path.
func FindJsonInPath(path string) []string {
	var jsonFiles []string
	files, err := os.ReadDir(path)
	if err != nil {
		return jsonFiles
	}
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".json") {
			jsonFiles = append(jsonFiles, filepath.Join(path, file.Name()))
		}
	}
	log.Printf("Found %d json files", len(jsonFiles))
	return jsonFiles
}
