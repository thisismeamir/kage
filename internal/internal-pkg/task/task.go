package task

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Task struct {
	FlowQueue           int             `json:"flow_queue"`
	Identifier          string          `json:"identifier"`
	RespectiveNode      string          `json:"respective_node"`
	IdInGraph           int             `json:"id_in_graph"`
	DependenciesInGraph []int           `json:"dependencies_in_graph,omitempty"`
	DependenciesInFlow  []string        `json:"dependencies_in_flow,omitempty"`
	ExecutionType       string          `json:"execution_type"`
	Type                string          `json:"type"`
	RespectiveFlow      string          `json:"respective_flow"`
	Input               string          `json:"input,omitempty"`
	Output              string          `json:"output,omitempty"`
	Priority            int             `json:"priority"`
	Urgency             int             `json:"urgency"`
	Status              int             `json:"status"`
	CreatedAt           string          `json:"created_at"`
	ResolvedAt          string          `json:"resolved_at"`
	RuntimeData         TaskRuntimeData `json:"runtime_data,omitempty"`
	ResourceTag         int             `json:"resource_tag,omitempty"`
}

type TaskRuntimeData struct {
	ExecutionStartTime string `json:"execution_start_time,omitempty"`
	ExecutionEndTime   string `json:"execution_end_time,omitempty"`
	ExecutionDuration  string `json:"execution_duration,omitempty"`
	ExecutionError     string `json:"execution_error,omitempty"`
}

func (t Task) GetIdentifier() string {
	if t.Identifier == "" {
		t = t.GenerateIdentifier()
	}
	return t.Identifier
}

func (t Task) GenerateIdentifier() Task {
	t.Identifier = fmt.Sprintf("%d.%s.task", t.FlowQueue, t.RespectiveFlow[:len(t.RespectiveFlow)-5])
	return t
}

func (t Task) SaveTask(path string) {
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		log.Printf("Error serializing task: %v", err)
	} else {
		if err := os.WriteFile(path+t.GetIdentifier()+".json", data, 0644); err != nil {
			log.Printf("Error saving task to file: %v", err)
		} else {
			log.Printf("Task saved successfully: %s", t.GetIdentifier())
		}
	}

}
