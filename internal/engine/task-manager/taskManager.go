package task_manager

import (
	"encoding/json"
	"github.com/thisismeamir/kage/internal/bootstrap/init_methods"
	"github.com/thisismeamir/kage/internal/internal-pkg/config"
	"github.com/thisismeamir/kage/internal/internal-pkg/flow"
	"log"
	"os"
)

type TaskManager struct {
	Queue Queue               `json:"queue"`
	Flows []flow.FlowRegister `json:"flows"`
}

type Queue struct {
	NodeBasedTasks []flow.TaskRegister `json:"node_based_tasks"`
	MapBasedTasks  []flow.TaskRegister `json:"flow_based_tasks"`
}

func InitializeTaskManager(c config.Config) TaskManager {
	// Check if Queue.json and Flow.registry.json exist, if not create them
	queuePath := c.BasePath + "/data/queue.json"
	flowRegistryPath := c.BasePath + "/data/flows.registry.json"
	tm := LoadTaskManager(queuePath, flowRegistryPath)
	// Find all Flows in c.BasePath /tmp/flows/ *Any json file that has "type": "flow" in it.
	allJsonFiles := init_methods.FindAllJsons([]string{c.BasePath + "/tmp/flows/"})
	for _, file := range allJsonFiles {
		typeOfFile := init_methods.GetTypeOfJson(file)
		if typeOfFile == "flow" {
			// Load the flow from the file
			newFlow := flow.LoadFlow(file)
			// Check if the flow already exists in the registry
			tm.AddFlow(flow.FlowRegister{
				Identifier: newFlow.Identifier,
				Path:       file,
			})

		}
	}
	tm.Clean()
	tm.Save(queuePath, flowRegistryPath)
	return tm
}

func LoadTaskManager(qPath string, frPath string) TaskManager {
	var tm TaskManager
	if _, err := os.Stat(qPath); os.IsNotExist(err) {
		log.Printf("Queue file does not exist at %s, initializing empty queue.", qPath)
		// If the queue file does not exist, initialize an empty queue and make a file
		tm.Queue = Queue{}
		data, _ := json.MarshalIndent(tm.Queue, "", "  ")
		if err := os.WriteFile(qPath, data, 0644); err != nil {
			log.Printf("Error writing empty queue to file: %v", err)
		} else {
			log.Printf("Empty queue initialized and saved to %s", qPath)
		}
	} else {
		data, err := os.ReadFile(qPath)
		if err != nil {
			log.Printf("Error reading queue file: %v", err)
			// If there is an error reading the file, initialize an empty queue and recreate a new queue.json file with empty structure
			tm.Queue = Queue{}
			data, _ := json.MarshalIndent(tm.Queue, "", "  ")
			if err := os.WriteFile(qPath, data, 0644); err != nil {
				log.Printf("Error writing empty queue to file: %v", err)
			} else {
				log.Printf("Empty queue initialized and saved to %s", qPath)
			}
		} else {
			if err := json.Unmarshal(data, &tm.Queue); err != nil {
				log.Printf("Error unmarshalling queue data: %v", err)
				tm.Queue = Queue{}
				data, _ := json.MarshalIndent(tm.Queue, "", "  ")
				if err := os.WriteFile(qPath, data, 0644); err != nil {
					log.Printf("Error writing empty queue to file: %v", err)
				} else {
					log.Printf("Empty queue initialized and saved to %s", qPath)
				}
			} else {
				log.Printf("Queue loaded successfully from %s", qPath)
			}
		}
	}

	if _, err := os.Stat(frPath); os.IsNotExist(err) {
		log.Printf("Flow registry file does not exist at %s, initializing empty flow registry.", frPath)
		tm.Flows = []flow.FlowRegister{}
		data, _ := json.MarshalIndent(tm.Flows, "", "  ")
		if err := os.WriteFile(frPath, data, 0644); err != nil {
			log.Printf("Error writing empty queue to file: %v", err)
		} else {
			log.Printf("Empty queue initialized and saved to %s", frPath)
		}
	} else {
		data, err := os.ReadFile(frPath)
		if err != nil {
			log.Printf("Error reading flow registry file: %v", err)
			tm.Flows = []flow.FlowRegister{}
			data, _ := json.MarshalIndent(tm.Flows, "", "  ")
			if err := os.WriteFile(frPath, data, 0644); err != nil {
				log.Printf("Error writing empty queue to file: %v", err)
			} else {
				log.Printf("Empty queue initialized and saved to %s", frPath)
			}
		} else {
			if err := json.Unmarshal(data, &tm.Flows); err != nil {
				log.Printf("Error unmarshalling flow registry data: %v", err)
				tm.Flows = []flow.FlowRegister{}
				data, _ := json.MarshalIndent(tm.Flows, "", "  ")
				if err := os.WriteFile(frPath, data, 0644); err != nil {
					log.Printf("Error writing empty queue to file: %v", err)
				} else {
					log.Printf("Empty queue initialized and saved to %s", frPath)
				}
			} else {
				log.Printf("Flow registry loaded successfully from %s", frPath)
			}
		}
	}
	return tm
}

func (tm *TaskManager) ContainsFlow(identifier string) bool {
	for _, flow := range tm.Flows {
		if flow.Identifier == identifier {
			return true
		}
	}
	return false
}

func (tm *TaskManager) AddFlow(flow flow.FlowRegister) TaskManager {
	if !tm.ContainsFlow(flow.Identifier) {
		tm.Flows = append(tm.Flows, flow)
		log.Printf("Flow %s added to task manager.", flow.Identifier)
	} else {
		log.Printf("Flow %s already exists in task manager.", flow.Identifier)
	}
	return *tm
}
func (tm *TaskManager) Save(queuePath string, flowRegistryPath string) {
	data, err := json.MarshalIndent(tm.Queue, "", "  ")
	if err != nil {
		log.Printf("Error serializing queue: %v", err)
	} else {
		if err := os.WriteFile(queuePath, data, 0644); err != nil {
			log.Printf("Error saving queue to file: %v", err)
		} else {
			log.Printf("Queue saved successfully to %s", queuePath)
		}
	}

	flowData, err := json.MarshalIndent(tm.Flows, "", "  ")
	if err != nil {
		log.Printf("Error serializing flow registry: %v", err)
	} else {
		if err := os.WriteFile(flowRegistryPath, flowData, 0644); err != nil {
			log.Printf("Error saving flow registry to file: %v", err)
		} else {
			log.Printf("Flow registry saved successfully to %s", flowRegistryPath)
		}
	}
}

func (tm *TaskManager) Clean() {
	tm.CleanNonExistingFlows()
	tm.CleanDuplicatedFlows()
}

func (tm *TaskManager) CleanDuplicatedFlows() {
	flowMap := make(map[string]flow.FlowRegister)
	for _, flow := range tm.Flows {
		if _, exists := flowMap[flow.Identifier]; !exists {
			flowMap[flow.Identifier] = flow
		} else {
			log.Printf("Duplicate flow found: %s, removing it from task manager.", flow.Identifier)
		}
	}
	tm.Flows = []flow.FlowRegister{}
	for _, uniqueFlow := range flowMap {
		tm.Flows = append(tm.Flows, uniqueFlow)
	}
	log.Printf("Cleaned duplicated flows, remaining flows: %d", len(tm.Flows))
}

func (tm *TaskManager) CleanNonExistingFlows() {
	for _, flow := range tm.Flows {
		if _, err := os.Stat(flow.Path); os.IsNotExist(err) {
			log.Printf("Flow %s does not exist at path %s, removing from task manager.", flow.Identifier, flow.Path)
			tm.RemoveFlow(flow.Identifier)
		}
	}
}

func (tm *TaskManager) RemoveFlow(identifier string) {
	for i, flow := range tm.Flows {
		if flow.Identifier == identifier {
			tm.Flows = append(tm.Flows[:i], tm.Flows[i+1:]...)
			log.Printf("Flow %s removed from task manager.", identifier)
			return
		}
	}
	log.Printf("Flow %s not found in task manager.", identifier)
}

func (tm *TaskManager) CleanFinishedOrCanceledTasks() {
	for _, task := range tm.Queue.NodeBasedTasks {
		if task.Status == 3 || task.Status == 4 { // Assuming 2 is finished and 3 is canceled
			log.Printf("Removing finished or canceled task %s from queue.", task.Id)
			tm.RemoveTaskFromQueue(task.Id)
		}
	}
}

func (tm *TaskManager) RemoveTaskFromQueue(taskId string) {
	for i, task := range tm.Queue.NodeBasedTasks {
		if task.Id == taskId {
			tm.Queue.NodeBasedTasks = append(tm.Queue.NodeBasedTasks[:i], tm.Queue.NodeBasedTasks[i+1:]...)
			log.Printf("Task %s removed from queue.", taskId)
			return
		}
	}
	log.Printf("Task %s not found in queue.", taskId)
}
