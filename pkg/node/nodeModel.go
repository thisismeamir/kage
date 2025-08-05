package node

type NodeModel struct {
	ExecutionModel   ExecutionModel `json:"execution_model"`
	Source           string         `json:"source"`
	WorkingDirectory string         `json:"working_directory"`
	EntryFile        string         `json:"entry_file"`
	OutputDirectory  string         `json:"output_directory"`
}
