package task_manager

import (
	"encoding/json"
	"github.com/thisismeamir/kage/internal/engine/context-evaluation/graph-analysis/toposort"
	"github.com/thisismeamir/kage/pkg/graph"
	"os"
)

type Flow struct {
	Identifier      string                    `json:"identifier"`
	Type            string                    `json:"type"`
	GraphIdentifier string                    `json:"graph_identifier"`
	EventIdentifier string                    `json:"event_identifier"`
	Urgency         int                       `json:"urgency"`
	TaskList        map[int][]Task            `json:"task_list"`
	Status          int                       `json:"status"`
	Structure       toposort.TopoSchedule     `json:"structure"`
	Input           map[string]interface{}    `json:"input"`
	ExecutionModel  graph.GraphExecutionModel `json:"execution_policy"`
}

func SaveFlow(fl Flow, path string) {
	data, _ := json.MarshalIndent(fl, "", "  ")
	_ = os.WriteFile(path+fl.Identifier+".json", data, 0644)
}

func LoadFlow(path string) Flow {
	data, _ := os.ReadFile(path)
	var fl Flow
	_ = json.Unmarshal(data, &fl)
	return fl
}

func (fl *Flow) GetTasksLinearized() []Task {
	taskList := make([]Task, 0)
	for _, tasks := range fl.TaskList {
		for _, task := range tasks {
			taskList = append(taskList, task)
		}
	}
	return taskList
}

func (fl *Flow) UpdateTaskStatus(taskIdentifier string, status int) {
	newTaskList := make(map[int][]Task)
	for level, tasks := range fl.TaskList {
		newTasks := make([]Task, len(tasks))
		for i, task := range tasks {
			if task.Identifier == taskIdentifier {
				task.Status = status
			}
			newTasks[i] = task
		}
		newTaskList[level] = newTasks
	}
	fl.TaskList = newTaskList
}
