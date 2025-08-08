package task_manager

type TaskManager struct {
	Queue []FlowRegister `json:"flows"`
}

type FlowRegister struct {
	Identifier string `json:"identifier"`
	Status     string `json:"status"`
	Path       string `json:"path"`
}
