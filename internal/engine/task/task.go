package task

type Task struct {
	Identifier string `json:"identifier"`   // Unique identifier for the task
	DependsOn  []Task `json:"dependens_on"` // List of tasks that this task depends on
	NeededBy   []Task `json:"needed_by"`    // List of tasks that depend on this task

}
