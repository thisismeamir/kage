package node

type ExecutionModel struct {
	Language     string                 `json:"language"`
	InputSchema  map[string]interface{} `json:"input_schema"`
	OutputSchema map[string]interface{} `json:"output_schema"`
	Artifacts    []string               `json:"artifacts"`
}
