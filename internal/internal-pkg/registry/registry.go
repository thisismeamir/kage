package registry

type Registry struct {
	GraphRegistry GraphRegistry `json:"graph_registry"`
	NodeRegistry  NodeRegistry  `json:"node_registry"`
	MapRegistry   MapRegistry   `json:"map_registry"`
}
