package execution_system

import (
	"github.com/thisismeamir/kage/internal/engine/execution-system/process"
	task_manager "github.com/thisismeamir/kage/internal/engine/task-manager"
	"github.com/thisismeamir/kage/internal/internal-pkg/config"
	"github.com/thisismeamir/kage/internal/internal-pkg/registry"
)

func RunFlow(identifier string, conf config.Config, reg registry.Registry) {
	flowPath := conf.BasePath + "/cache/flows/"
	flow := task_manager.LoadFlow(flowPath + identifier + ".json")
	flow.Status = 1
	for _, tasks := range flow.TaskList {
		for _, task := range tasks {
			status := process.ProcessTask(&task, conf, reg)
			flow.UpdateTaskStatus(task.Identifier, status)
			task_manager.SaveFlow(flow, flowPath)
		}
	}
	var flowDone = true
	for _, tts := range flow.TaskList {
		for _, t := range tts {
			// Checking if all ts are
			if t.Status != 2 {
				if t.Status >= 0 {
					flowDone = false
				}
			}
		}
	}
	if flowDone {
		flow.Status = 2
		task_manager.SaveFlow(flow, flowPath)
	}
}
