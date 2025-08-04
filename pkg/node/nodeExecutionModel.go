package node

type ExecutionModel struct {
	Language     LanguageModel          `json:"language"`
	InputSchema  map[string]interface{} `json:"input_schema"`
	OutputSchema map[string]interface{} `json:"output_schema"`
}

type LanguageModel struct {
	Name           string `json:"name"`
	ExecutablePath string `json:"executable_path"`
}
