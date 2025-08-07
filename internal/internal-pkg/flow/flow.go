package flow

import (
	"encoding/json"
	"log"
	"os"
)

type Flow struct {
	Identifier      string                 `json:"identifier,omitempty"`
	Type            string                 `json:"type"`
	RespectiveGraph string                 `json:"respective_graph"`
	RespectiveEvent string                 `json:"respective_event"`
	Tasks           []TaskRegister         `json:"tasks"`
	Urgency         int                    `json:"urgency"`
	Status          int                    `json:"status"`
	InitialInput    map[string]interface{} `json:"initial_input,omitempty"`
}

type TaskRegister struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	GraphId     int    `json:"graph_id"`
	Queue       int    `json:"queue"`
	Level       int    `json:"level"`
	Urgency     int    `json:"urgency"`
	ResourceTag int    `json:"resource_tag,omitempty"`
	Status      int    `json:"status"`
	PathOfTask  string `json:"path_of_task,omitempty"`
}

func (fl Flow) GenerateIdentifier() Flow {
	fl.Identifier = fl.RespectiveEvent[:len(fl.RespectiveEvent)-5] + ".flow"
	return fl
}

func (fl Flow) Save(path string) {
	data, err := json.MarshalIndent(fl, "", "  ")
	if err != nil {
		log.Printf("[FATAL] Could not save flow JSON: %s", err)
	} else if err := os.WriteFile(path+"/"+fl.Identifier+".json", data, 0644); err != nil {
		log.Printf("[FATAL] Could not save flow JSON: %s", err)
	} else {
		log.Printf("Flow saved successfully: %s", fl.Identifier)
	}
}

func LoadFlow(path string) Flow {
	var fl Flow
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("[FATAL] Could not read flow file: %s", err)
		return fl
	} else {
		if err := json.Unmarshal(data, &fl); err != nil {
			log.Printf("[FATAL] Could not unmarshal flow JSON: %s", err)
			return fl
		}
		fl = fl.GenerateIdentifier() // Ensure identifier is generated
		log.Printf("Flow loaded successfully: %s", fl.Identifier)
		return fl
	}
}

func (fl Flow) GetTaskIdentifierByGraphId(graphId int) string {
	for _, task := range fl.Tasks {
		if task.GraphId == graphId {
			return task.Id
		}
	}
	return ""
}
