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
