package flow

import (
	"encoding/json"
	"log"
	"os"
)

type Flow struct {
	Identifier      string                 `json:identifier,omitempty"`
	RespectiveGraph string                 `json:"respective_graph"`
	RespectiveEvent string                 `json:"respective_event"`
	Tasks           []TaskRegister         `json:"tasks"`
	Urgency         int                    `json:"urgency"`
	Status          int                    `json:"status"`
	InitialInput    map[string]interface{} `json:"initial_input,omitempty"`
}

type TaskRegister struct {
	Id          string `json:"id"`
	Queue       int    `json:"queue"`
	Level       int    `json:"level"`
	ResourceTag int    `json:"resource_tag,omitempty"`
	Status      int    `json:"status"`
}

func (fl Flow) GenerateIdentifier() Flow {
	fl.Identifier = fl.RespectiveEvent + ".flow"
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
