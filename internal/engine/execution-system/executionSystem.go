package execution_system

import (
	"github.com/thisismeamir/kage/internal/bootstrap/init_methods"
	system_monitor "github.com/thisismeamir/kage/internal/engine/system-monitor"
	task_manager "github.com/thisismeamir/kage/internal/engine/task-manager"
	"github.com/thisismeamir/kage/internal/internal-pkg/config"
	"sort"
)

type ExecutionSystem struct {
	Flows                   []task_manager.Flow
	CurrentlyAvailableFlows []task_manager.Flow
}

func (ex *ExecutionSystem) FetchFlows(conf config.Config) *ExecutionSystem {
	flowsPath := conf.BasePath + "/cache/flows/"
	jsonFiles := init_methods.FindAllJsons([]string{flowsPath})
	for _, jsonFile := range jsonFiles {
		OfType := init_methods.GetTypeOfJson(jsonFile)
		if OfType == "flow" {
			ex.Flows = append(ex.Flows, task_manager.LoadFlow(jsonFile))
		}
	}
	return ex
}
func (ex *ExecutionSystem) SortByUrgency() *ExecutionSystem {
	flows := ex.Flows
	sort.Slice(flows, func(i, j int) bool {
		return flows[i].Urgency < flows[j].Urgency
	})
	ex.Flows = flows
	return ex
}
func (ex *ExecutionSystem) CreateCurrentlyAvailableFlowsList(sm *system_monitor.SystemMonitor, conf config.Config) *ExecutionSystem {
	ex.CurrentlyAvailableFlows = make([]task_manager.Flow, 0)
	opts := system_monitor.CheckOptions{
		SkipNetwork: true,  // Skip the network check
		SkipCPU:     false, // Don't skip CPU check
		SkipMemory:  false,
		SkipDisk:    false,
	}
	for _, flow := range ex.Flows {
		if system_monitor.IsFlowAbleToRun(flow, sm, conf, opts) && !ex.IsFlowCurrentlyAvailable(flow) {
			ex.CurrentlyAvailableFlows = append(ex.CurrentlyAvailableFlows, flow)
		}
	}
	return ex
}

// For no duplicates
func (ex *ExecutionSystem) IsFlowCurrentlyAvailable(flow task_manager.Flow) bool {
	for _, currentlyAvailableFlow := range ex.CurrentlyAvailableFlows {
		if currentlyAvailableFlow.Identifier == flow.Identifier {
			return true
		}
	}
	return false
}

func (ex *ExecutionSystem) RemoveDuplicateFlows() *ExecutionSystem {
	uniqueFlows := make([]task_manager.Flow, 0)
	flowMap := make(map[string]bool)
	for _, flow := range ex.Flows {
		if _, exists := flowMap[flow.Identifier]; !exists {
			uniqueFlows = append(uniqueFlows, flow)
			flowMap[flow.Identifier] = true
		}
	}
	return ex
}

func (ex *ExecutionSystem) RemoveFlow(identifier string) *ExecutionSystem {
	for i, flow := range ex.CurrentlyAvailableFlows {
		if flow.Identifier == identifier {
			ex.CurrentlyAvailableFlows = append(ex.CurrentlyAvailableFlows[:i], ex.CurrentlyAvailableFlows[i+1:]...)
			break
		}
	}
	return ex
}

func (ex *ExecutionSystem) SetFlowStatus(identifier string, status int) *ExecutionSystem {
	for i, flow := range ex.CurrentlyAvailableFlows {
		if flow.Identifier == identifier {
			ex.CurrentlyAvailableFlows[i].Status = status
			break
		}
	}
	return ex
}

// NewSystemMonitor creates a new SystemMonitor instance
func NewExecutionSystem() *ExecutionSystem {
	return &ExecutionSystem{}
}
