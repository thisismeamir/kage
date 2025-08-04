package graph

type GraphModel struct {
	Execution GraphExecutionModel `json:"execution,omitempty"`
	Structure []GraphObject       `json:"structure"`
}
