package graph

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/thisismeamir/kage/pkg/form"
	"github.com/thisismeamir/kage/util"
	"log"
	"os"
)

type Graph struct {
	form.Form
	Model    GraphModel    `json:"model"`
	Metadata form.Metadata `json:"metadata"`
}

func LoadGraph(graphPath string) (Graph, error) {
	var graph Graph
	data, err := os.ReadFile(graphPath)
	if err != nil {
		log.Fatalf("[ERROR] LoadGraph failed to read file: %s", err)
		return Graph{}, err
	} else {
		if err := json.Unmarshal(data, &graph); err != nil {
			log.Fatalf("[ERROR] LoadGraph failed to unmarshal JSON: %s", err)
			return Graph{}, err
		}
		return graph, nil
	}
}

func (graph Graph) SaveGraph(path string) error {
	data, err := json.MarshalIndent(graph, "", "  ")
	if err != nil {
		log.Printf("[ERROR] SaveGraph failed to save %s in marshal JSON %s: %s", graph.Name, path, err)
		return err
	} else if err := os.WriteFile(path, data, 0644); err != nil {
		log.Printf("[ERROR] SaveGraph failed to save %s in path %s: %s", graph.Name, path, err)
		return err
	}
	return nil
}

func (graph Graph) GetObject(id int) (*GraphObject, error) {
	for _, obj := range graph.Model.Structure {
		if obj.Id == id {
			log.Printf("[INFO] Found object with ID %d: %+v", id, obj)
			return &obj, nil
		}
	}
	log.Printf("[ERROR] No object found with ID %d", id)
	return nil, errors.New(fmt.Sprintf("No graph with id: %i", id))

}

func (graph Graph) GetLength() int {
	return len(graph.Model.Structure)
}

func (graph Graph) GetDependency(id int) []int {
	var dependencies []int
	for _, m := range graph.Model.Structure {
		if id != m.Id && util.IntInList(id, m.Outgoing) {
			dependencies = append(dependencies, m.Id)
		}
	}
	return dependencies
}
