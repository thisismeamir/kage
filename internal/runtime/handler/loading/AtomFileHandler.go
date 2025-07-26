package loading

import (
	"encoding/json"
	"github.com/thisismeamir/kage/pkg/atom"
	"io"
	"log"
	"os"
)

func AtomFileHandler(path string) atom.AtomModel {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			// If the file does not exist, return an empty Config
			log.Println("file does not exist:", path)
			return atom.AtomModel{}
		}
		// If there is another error, log it and return an empty Config
		log.Fatal(err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var loadedAtom atom.AtomModel
	if err := json.Unmarshal(data, &loadedAtom); err != nil {
		return atom.AtomModel{}
	}
	return loadedAtom
}
