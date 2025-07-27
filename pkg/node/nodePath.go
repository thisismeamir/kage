package node

// NodePath Data structure for a path of json.
type NodePath struct {
	Path  string `json:"path"`
	Local bool   `json:"local,omitempty"`
}
