package graph

import (
	"github.com/thisismeamir/kage/pkg/mapping"
	"github.com/thisismeamir/kage/pkg/node"
)

type GraphAttachments struct {
	GraphNodes []node.Node   `json:"graph_nodes"`
	GraphMaps  []mapping.Map `json:"graph_maps"`
}
