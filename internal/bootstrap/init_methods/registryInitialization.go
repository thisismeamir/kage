package init_methods

import (
	//"github.com/thisismeamir/kage/internal/internal-pkg/registry"
	"log"
	"os"
	"path/filepath"
)

func InitializeRegistries() {
	// TODO
}

// GetPathsObjects : This function would go inside all the paths that are set in config file and finds files that are:
// 1. <some-name>.node.json
// 2. <some-name>.map.json
// 3. <some-name>.graph.json
func GetPathsObjects(paths []string) {
	for _, path := range paths {
		log.Printf("Checking path: %s", path)
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			// Checking the suffix of each file and adding them in case of match:
			if info.Name()[len(info.Name())-4:] == ".json" {
				print(info.Name())
			}
			return err
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}
