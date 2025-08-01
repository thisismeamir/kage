package init_methods

import (
	"encoding/json"
	"fmt"
	"github.com/thisismeamir/kage/internal/internal-pkg/registry"
	"github.com/thisismeamir/kage/pkg/graph"
	"github.com/thisismeamir/kage/pkg/mapping"
	"github.com/thisismeamir/kage/pkg/node"
	"io/fs"
	"os"
	"strings"

	//"github.com/thisismeamir/kage/internal/internal-pkg/registry"
	"log"
	"path/filepath"
)

func InitializeRegistries(paths []string, registryPath string) error {
	reg, err := registry.LoadRegistry(registryPath)
	if err != nil {
		reg = &registry.Registry{
			GraphRegistry: make([]registry.GraphRegister, 0),
			NodeRegistry:  make([]registry.NodeRegister, 0),
			MapRegistry:   make([]registry.MapRegister, 0),
		}
	}
	// Zeroth step: before we start adding, let's remove registers that are not available anymore:
	reg.CleanMissingFiles()
	// First we find all the json files in the paths that's been set in config file:
	files := FindAllJsons(paths)
	// For each json file we do the following:
	for _, file := range files {
		// setting a pointer that would hold the value of the json "type" key if in existence.
		var jsonType *string
		jsonType = GetTypeOfJson(file)
		// Going case by case with jsonType value:
		switch *jsonType {
		// If it is a node:
		case "node":
			// First we load the node to make sure it's loadable and fine.
			newNode, err := node.LoadNode(file)
			// if an error occurred:
			if err != nil {
				log.Println("Error loading node: ", err)
			} else {
				// otherwise, we generate an identifier for it which is a simple naming convention
				newNodeIdentifier := GenerateNodeIdentifier(*newNode)
				// Checking to see if the node already is there in the registry
				if reg.Contains(newNodeIdentifier, file) {
					log.Printf("Node %s, already exists in registry with path: %s", newNodeIdentifier, file)
				} else {
					// If the node is new we add this to our registry
					log.Printf("Node %s does not exist in registry with path: %s", newNode.Name, file)
					log.Printf("Adding Node with identifier %s to registry.", newNodeIdentifier)
					reg.AddNode(*newNode, newNodeIdentifier, file)
				}
			}
		// If it is a map:
		case "map":
			// It's the same as the nodes
			newMap, err := mapping.LoadMap(file)
			if err != nil {
				log.Println("Error loading map: ", err)
			} else {
				newMapIdentifier := GenerateMapIdentifier(*newMap)
				if reg.Contains(newMapIdentifier, file) {
					log.Printf("Map %s already exists in registry with path: %s", newMapIdentifier, file)
				} else {
					log.Printf("Map %s does not exist in registry with path: %s", newMap.Name, file)
					log.Printf("Adding Map with identifier %s to registry.", newMapIdentifier)
					reg.AddMap(*newMap, newMapIdentifier, file)
				}
			}
		case "graph":
			newGraph, err := graph.LoadGraph(file)
			if err != nil {
				log.Println("Error loading graph: ", err)
			} else {
				newGraphIdentifier := GenerateGraphIdentifier(*newGraph)
				if reg.Contains(newGraphIdentifier, file) {
					log.Printf("Graph %s already exists in registry with path: %s", newGraphIdentifier, file)
				} else {
					reg.AddGraph(newGraphIdentifier, file)
				}
			}
		default:
			log.Println("Unsupported json type: ", *jsonType)
		}
	}
	fmt.Println("Registry initialization complete ")
	fmt.Printf("Registry nodes: %s ", reg)
	reg.Save(registryPath)
	return nil
}

// FindAllJsons : This function would go inside all the paths that are set in config file and finds files that are json.
func FindAllJsons(paths []string) []string {
	files := make([]string, 0)
	for _, path := range paths {
		log.Printf("Checking path: %s", path)
		err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			} else {
				if !d.IsDir() && filepath.Ext(path) == ".json" {
					files = append(files, path)
				}
			}
			return nil
		})
		if err != nil {
			log.Printf("Error walking path %s: %v", path, err)
		}
	}
	log.Printf("Found %d files", len(files))
	log.Printf("files: %s", files)
	return files
}

func GetTypeOfJson(path string) *string {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error reading %s: %v", path, err)
		return nil
	}

	var jsonObject map[string]interface{}
	err = json.Unmarshal(data, &jsonObject)
	if err != nil {
		log.Printf("Error parsing %s: %v", path, err)
		return nil
	}

	if val, exists := jsonObject["type"]; exists {
		strVal := fmt.Sprintf("%v", val)
		return &strVal
	}

	return nil
}

func GenerateNodeIdentifier(n node.Node) string {
	return strings.ReplaceAll(n.Name, " ", "-") + "." + n.Version + ".node"
}

func GenerateMapIdentifier(m mapping.Map) string {
	return strings.ReplaceAll(m.Name, " ", "-") + "." + m.Version + ".map"
}

func GenerateGraphIdentifier(g graph.Graph) string {
	return strings.ReplaceAll(g.Name, " ", "-") + "." + g.Version + ".graph"
}
