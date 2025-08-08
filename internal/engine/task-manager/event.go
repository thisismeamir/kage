package task_manager

type Event struct {
	Identifier      string                 `json:"identifier"`
	GraphIdentifier string                 `json:"graph_identifier"`
	Urgency         int                    `json:"urgency"`
	Input           map[string]interface{} `json:"input"`
}
