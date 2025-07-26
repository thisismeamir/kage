// pkg/atom/atom.go
package atom

type AtomModel struct {
	Model          AtomMetadata       `json:"model"`
	ExecutionModel AtomExecutionModel `json:"execution_model"`
}

type AtomMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Identifier  string `json:"identifier"`
	Author      string `json:"author"`
	AuthorEmail string `json:"author_email"`
	AuthorURL   string `json:"author_url"`
	ManualLink  string `json:"manual_link"`
}

type AtomExecutionModel struct {
	Language     string                 `json:"language"`
	InputSchema  map[string]interface{} `json:"input_schema"`
	OutputSchema map[string]interface{} `json:"output_schema"`
	Source       string                 `json:"source"`
}

type AtomRunHandler interface {
	Run(source string, input map[string]interface{}) (map[string]interface{}, error)
}
