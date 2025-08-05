package scheduler

import (
	"github.com/thisismeamir/kage/internal/engine/execution/context-evaluation/graph-analysis/toposort"
	"github.com/thisismeamir/kage/internal/internal-pkg/config"
	"github.com/thisismeamir/kage/internal/internal-pkg/event"
	"github.com/thisismeamir/kage/internal/internal-pkg/flow"
	"github.com/thisismeamir/kage/internal/internal-pkg/registry"
	"github.com/thisismeamir/kage/internal/internal-pkg/task"
	"log"
	"os"
	"time"
)

func Scheduler(e event.Event, r registry.Registry, c config.Config) flow.Flow {
	var f = flow.Flow{
		Type:            "flow",
		RespectiveGraph: e.Graph,
		RespectiveEvent: e.GetIdentifier(),
		Urgency:         e.Urgency,
		Status:          0,
		InitialInput:    e.InitialInput,
	}
	f = f.GenerateIdentifier()
	flowPath := c.BasePath + "/tmp/" + "flows/" + e.Identifier
	_ = os.MkdirAll(flowPath, os.ModePerm)
	_ = os.Mkdir(flowPath+"/tasks", os.ModePerm)
	graph, err := r.LoadGraph(e.Graph)
	if err != nil {
		log.Printf("graph %v does not exist", e.Graph)
	} else {
		sortedGraphStructure, err2 := toposort.TopoSort(graph.Model.Structure)
		if err2 != nil {
			log.Printf("graph structure %v cannot be sorted (possibly a cycle exists)", e.Graph)
		} else {
			count := 0
			for _, level := range sortedGraphStructure.Levels {
				for _, n := range level.Nodes {
					respectiveNode, _ := graph.GetObject(n)
					depsGraphId := graph.GetDependency(n)
					depsTaskId := make([]string, 0)
					for _, id := range depsGraphId {
						depsTaskId = append(depsTaskId, f.GetTaskIdentifierByGraphId(id))
					}
					t := task.Task{
						RespectiveNode:      respectiveNode.ExecutionIdentifier,
						IdInGraph:           n,
						Type:                "task",
						DependenciesInGraph: depsGraphId,
						DependenciesInFlow:  depsTaskId,
						ExecutionType:       respectiveNode.Type,
						RespectiveFlow:      f.Identifier,
						Priority:            level.Level,
						Urgency:             f.Urgency,
						Status:              0,
						CreatedAt:           time.Now().Format("2006.01.02.15.04.05"),
						FlowQueue:           count,
						ResourceTag:         0,
					}
					t = t.GenerateIdentifier()
					t.SaveTask(flowPath + "/tasks/")
					f.Tasks = append(f.Tasks, flow.TaskRegister{
						Id:          t.GetIdentifier(),
						Queue:       t.FlowQueue,
						GraphId:     t.IdInGraph,
						Type:        respectiveNode.Type,
						Level:       t.Priority,
						Urgency:     t.Urgency,
						ResourceTag: t.ResourceTag,
						Status:      t.Status,
					})
					count++
					log.Printf("[INFO] Task %s for node %s added to flow %s", t.Identifier, respectiveNode.ExecutionIdentifier, f.Identifier)
				}
			}
		}
	}
	f.Save(flowPath)
	return f
}
