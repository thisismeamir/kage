package mapping

import (
	"encoding/json"
	"github.com/thisismeamir/kage/pkg/form"
	"log"
	"os"
)

type Map struct {
	form.Form
	Model    MapModel      `json:"model"`
	Metadata form.Metadata `json:"metadata"`
}

func LoadMap(mapPath string) (*Map, error) {
	mapping := &Map{}
	data, err := os.ReadFile(mapPath)
	if err != nil {
		log.Fatalf("[ERROR] LoadMap failed to read file: %s", err)
		return nil, err
	} else {
		if err := json.Unmarshal(data, mapping); err != nil {
			log.Fatalf("[ERROR] LoadMap failed to unmarshal: %s", err)
			return nil, err
		}
		return mapping, nil
	}

}
