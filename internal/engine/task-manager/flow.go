package task_manager

import (
	"encoding/json"
	"github.com/thisismeamir/kage/internal/engine/context-evaluation/graph-analysis/toposort"
	"os"
)

type Flow struct {
	Identifier      string                 `json:"identifier"`
	Type            string                 `json:"type"`
	GraphIdentifier string                 `json:"graph_identifier"`
	EventIdentifier string                 `json:"event_identifier"`
	Urgency         int                    `json:"urgency"`
	TaskList        map[int][]Task         `json:"task_list"`
	Status          int                    `json:"status"`
	Structure       toposort.TopoSchedule  `json:"structure"`
	Input           map[string]interface{} `json:"input"`
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
