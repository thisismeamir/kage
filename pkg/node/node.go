package node

import (
	"encoding/json"
	"github.com/thisismeamir/kage/pkg/form"
	"log"
	"os"
)

type Node struct {
	form.Form
	Model    NodeModel     `json:"model"`
	Metadata form.Metadata `json:"metadata"`
}

func LoadNode(nodePath string) (*Node, error) {
	node := &Node{}
	data, err := os.ReadFile(nodePath)
	if err != nil {
		log.Fatalf("[ERROR] LoadNode failed to read file: %s", err)
		return nil, err
	} else {
		if err := json.Unmarshal(data, node); err != nil {
			log.Fatalf("[ERROR] LoadNode failed to unmarshal JSON: %s", err)
			return nil, err
		}
		return node, nil
	}

}

func (node Node) Save(path string) error {
	nodePath := path + node.Identifier + ".json"
	data, err := json.MarshalIndent(node, "", "  ")
	if err != nil {
		log.Printf("[ERROR] Save failed to marshal JSON: %s", err)
		return err
	} else {
		if err := os.WriteFile(nodePath, data, 0644); err != nil {
			log.Printf("[ERROR] Save failed to write file: %s", err)
			return err
		} else {
			log.Printf("[INFO] Node: %s saved successfully at %s", node.Identifier, path)
			return nil
		}
	}

}
