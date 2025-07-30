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

func (mapp Map) Save(path string) error {
	nodePath := path + mapp.Identifier + ".json"
	data, err := json.MarshalIndent(mapp, "", "  ")
	if err != nil {
		log.Printf("[ERROR] Save failed to marshal JSON: %s", err)
		return err
	} else {
		if err := os.WriteFile(nodePath, data, 0644); err != nil {
			log.Printf("[ERROR] Save failed to write file: %s", err)
			return err
		} else {
			log.Printf("[INFO] Map: %s saved successfully at %s", mapp.Identifier, path)
			return nil
		}
	}

}
