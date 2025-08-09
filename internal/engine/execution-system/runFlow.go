package execution_system

import (
	"github.com/thisismeamir/kage/internal/engine/execution-system/process"
	task_manager "github.com/thisismeamir/kage/internal/engine/task-manager"
	"github.com/thisismeamir/kage/internal/internal-pkg/config"
	"github.com/thisismeamir/kage/internal/internal-pkg/registry"
)

func RunFlow(identifier string, conf config.Config, reg registry.Registry) {
	flow := task_manager.LoadFlow(conf.BasePath + "/cache/flows/" + identifier + ".json")
	tasks := flow.GetTasksLinearized()
	for _, task := range tasks {
		process.ProcessTask(&task, conf, reg)
	}

}
