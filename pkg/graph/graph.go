package graph

import (
	"encoding/json"
	"github.com/thisismeamir/kage/pkg/form"
	"github.com/thisismeamir/kage/pkg/mapping"
	"github.com/thisismeamir/kage/pkg/node"
	"log"
	"os"
)

type Graph struct {
	form.Form
	Model    GraphModel    `json:"model"`
	Metadata form.Metadata `json:"metadata"`
}

func LoadGraph(graphPath string) (*Graph, error) {
	var graph *Graph
	data, err := os.ReadFile(graphPath)
	if err != nil {
		log.Fatalf("[ERROR] LoadGraph failed to read file: %s", err)
		return nil, err
	} else {
		if err := json.Unmarshal(data, &graph); err != nil {
			log.Fatalf("[ERROR] LoadGraph failed to unmarshal JSON: %s", err)
			return nil, err
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

func (graph Graph) AddNodeToGraph(node node.Node) Graph {
	graph.Model.Attachments.GraphNodes = append(graph.Model.Attachments.GraphNodes, node)
	return graph
}

func (graph Graph) AddMapToGraph(mapp mapping.Map) Graph {
	graph.Model.Attachments.GraphMaps = append(graph.Model.Attachments.GraphMaps, mapp)
	return graph
}

func (graph Graph) AddGraphObject(object GraphObject) Graph {
	graph.Model.GraphStructure = append(graph.Model.GraphStructure, object)
	return graph
}

func (graph Graph) AddDependency(dep GraphDependency) Graph {
	graph.Model.Dependencies = append(graph.Model.Dependencies, dep)
	return graph
}

func (graph Graph) RemoveFromGraph(identifier string) Graph {
	newGraphDependencies := make([]GraphDependency, 0)
	newGraphAttachments := GraphAttachments{
		make([]node.Node, 0),
		make([]mapping.Map, 0),
	}
	for _, existingNode := range graph.Model.Attachments.GraphNodes {
		if existingNode.Name != identifier {
			newGraphAttachments.GraphNodes = append(newGraphAttachments.GraphNodes, existingNode)
		}
	}
	for _, existingMap := range graph.Model.Attachments.GraphMaps {
		if existingMap.Name != identifier {
			newGraphAttachments.GraphMaps = append(newGraphAttachments.GraphMaps, existingMap)
		}
	}
	for _, dependency := range graph.Model.Dependencies {
		if dependency.Identifier != identifier {
			newGraphDependencies = append(newGraphDependencies, dependency)
		}
	}
	newModel := GraphModel{
		graph.Model.Execution,
		newGraphDependencies,
		graph.Model.GraphStructure,
		newGraphAttachments,
	}
	newGraph := Graph{
		graph.Form,
		newModel,
		graph.Metadata,
	}
	return newGraph
}
