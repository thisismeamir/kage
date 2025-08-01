package node

import (
	"encoding/json"
	"errors"
	"github.com/thisismeamir/kage/pkg/form"
	"io/ioutil"
	"os"
)

type Node struct {
	form.Form
	Model    NodeModel     `json:"model"`
	Metadata form.Metadata `json:"metadata"`
}

// LoadNode loads a Node from a JSON file and validates it.
func LoadNode(path string) (*Node, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if err := ValidateNodeJSON(data); err != nil {
		return nil, err
	}
	var n Node
	if err := json.Unmarshal(data, &n); err != nil {
		return nil, err
	}
	return &n, nil
}

// SaveNode validates the Node and saves it as JSON to a file.
func SaveNode(n *Node, path string) error {
	data, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

// ValidateNodeJSON checks if the JSON is a valid node type.
func ValidateNodeJSON(data []byte) error {
	var temp struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	if temp.Type != "node" {
		return errors.New("invalid type: not a node")
	}
	return nil
}
