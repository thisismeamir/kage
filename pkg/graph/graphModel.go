package graph

type GraphModel struct {
	Execution      GraphExecutionModel `json:"execution"`
	Dependencies   []GraphDependency   `json:"dependencies"`
	GraphStructure []GraphObject       `json:"graph_structure"`
	Attachments    GraphAttachments    `json:"attachments,omitempty"`
}
