package task_manager

type Task struct {
	Identifier     string   `json:"identifier"`
	FlowIdentifier string   `json:"flow_identifier"`
	Type           string   `json:"type"`
	ExecutionType  string   `json:"execution_type"`
	NodeIdentifier string   `json:"execution_identifier"`
	GraphOutgoing  []int    `json:"graph_outgoing"`
	GraphIngoing   []int    `json:"graph_ingoing"`
	FlowDependency []string `json:"flow_dependency"`
	Level          int      `json:"level"`
	Queue          int      `json:"queue"`
	Input          []string `json:"input"`
	Status         int      `json:"status"`
}
