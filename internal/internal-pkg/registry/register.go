package registry

type NodeRegister struct {
	Identifier   string                 `json:"identifier"`
	Path         string                 `json:"path"`
	InputSchema  map[string]interface{} `json:"input_schema"`
	OutputSchema map[string]interface{} `json:"output_schema"`
}

type MapRegister struct {
	Identifier   string                 `json:"identifier"`
	Path         string                 `json:"path"`
	InputSchema  map[string]interface{} `json:"input_schema"`
	OutputSchema map[string]interface{} `json:"output_schema"`
}

type GraphRegister struct {
	Identifier string `json:"identifier"`
	Path       string `json:"path"`
}
