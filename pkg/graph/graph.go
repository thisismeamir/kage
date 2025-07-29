package graph

import (
	"encoding/json"
	"github.com/thisismeamir/kage/pkg/form"
	"log"
	"os"
)

type Graph struct {
	form.Form
	Model    GraphModel    `json:"model"`
	Metadata form.Metadata `json:"metadata"`
}

func LoadGraph(graphPath string) (*Graph, error) {
	graph := &Graph{}
	data, err := os.ReadFile(graphPath)
	if err != nil {
		log.Fatalf("[ERROR] LoadGraph failed to read file: %s", err)
		return nil, err
	} else {
		if err := json.Unmarshal(data, graph); err != nil {
			log.Fatalf("[ERROR] LoadGraph failed to unmarshal JSON: %s", err)
			return nil, err
		}
		return graph, nil
	}
}
