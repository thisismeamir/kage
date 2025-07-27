package graph

type GraphExecutionRetryModel struct {
	MaxAttempts int `json:"max_attempts"`
	Delay       int `json:"delay"`
}
