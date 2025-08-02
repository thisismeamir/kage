package graph

type GraphObject struct {
	Id                  int    `json:"id"`
	Type                string `json:"type"`
	ExecutionIdentifier string `json:"execution_identifier"`
	Outgoing            []int  `json:"outgoing,omitempty"`
}
