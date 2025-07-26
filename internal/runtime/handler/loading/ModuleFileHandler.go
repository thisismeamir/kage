package loading

import (
	"encoding/json"
	module "github.com/thisismeamir/kage/pkg/module"
	"io"
	"log"
	"os"
)

func ModuleFileHandler(path string) module.ModuleModel {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file does not exist, return an empty Config
			log.Println("file does not exist:", path)
			return module.ModuleModel{}
		}
		// If there is another error, log it and return an empty Config
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var loadedModule module.ModuleModel
	if err := json.Unmarshal(data, &loadedModule); err != nil {
		return module.ModuleModel{}
	}
	return loadedModule
}
