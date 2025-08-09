package task_manager

import (
	"github.com/thisismeamir/kage/internal/engine/context-evaluation/graph-analysis/toposort"
	"github.com/thisismeamir/kage/internal/internal-pkg/config"
	"github.com/thisismeamir/kage/internal/internal-pkg/registry"
	"github.com/thisismeamir/kage/util"
	"os"
)

func (e *Event) ScheduleFlow(conf config.Config, reg registry.Registry) Flow {
	// Creating a new flow based on the event information.
	fl := Flow{
		Identifier:      IdentifierGeneration("flow"),
		Type:            "flow",
		GraphIdentifier: e.GraphIdentifier,
		EventIdentifier: e.Identifier,
		Status:          0,
		Urgency:         e.Urgency,
		TaskList:        make(map[int][]Task),
		Input:           e.Input,
	}
	_ = os.MkdirAll(conf.BasePath+"/tmp/"+fl.Identifier, os.ModePerm)
	// Loading Graph
	gr, _ := reg.LoadGraph(fl.GraphIdentifier)
	// Sorting the graph with Topological sort to know which nodes (tasks) comes first
	sortedStructure, _ := toposort.TopoSort(gr.Model.Structure)
	fl.Structure = sortedStructure
	fl.ExecutionModel = gr.Model.Execution
	// Counting from zero so that each task has a unique number for queue
	count := 0
	// Iterating through the sorted structure to create tasks
	for level, nodeIds := range sortedStructure.Levels {
		for _, id := range nodeIds.Nodes {
			taskId := IdentifierGeneration("task")
			obj, _ := gr.GetObject(id)
			var deps []int
			flowDeps := make([]string, 0)
			input := make([]string, 0)
			if level == 0 {
				// Having a node id -1 in task deps means that it doesn't have any deps (Should be injected with the event input instead).
				deps = append(deps, -1)
				initialInputPath := conf.BasePath + "/tmp/" + fl.Identifier + "/" + taskId + ".input.json"
				util.SaveJson(e.Input, initialInputPath)
				input = append(input, initialInputPath)
			} else {
				deps = gr.GetDependency(obj.Id)
				for i := level - 1; i >= 0; i-- {
					for _, task := range fl.TaskList[i] {
						if util.IntInList(obj.Id, task.GraphOutgoing) {
							flowDeps = append(flowDeps, task.Identifier)
						}
					}
				}
			}
			newTask := Task{
				Identifier:     taskId,
				Type:           "task",
				ExecutionType:  obj.Type,
				NodeIdentifier: obj.ExecutionIdentifier,
				GraphOutgoing:  obj.Outgoing,
				GraphIngoing:   deps,
				FlowDependency: flowDeps,
				Level:          level,
				Queue:          count,
				Input:          input,
				FlowIdentifier: fl.Identifier,
				Status:         0,
			}
			fl.TaskList[level] = append(fl.TaskList[level], newTask)

			count++
		}
	}

	SaveFlow(fl, conf.BasePath+"/cache"+"/flows/")
	return fl
}
