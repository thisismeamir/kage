package node

type NodeModel struct {
	ExecutionModel ExecutionModel `json:"execution_model"`
	Source         string         `json:"source"`
}
