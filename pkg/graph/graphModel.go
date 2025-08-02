package graph

type GraphModel struct {
	Execution    GraphExecutionModel `json:"execution"`
	Dependencies []GraphDependency   `json:"dependencies,omitempty"`
	Structure    []GraphObject       `json:"structure"`
	Attachments  GraphAttachments    `json:"attachments,omitempty"`
}
