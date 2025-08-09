package dependency_handler

import (
	"encoding/json"
	task_manager "github.com/thisismeamir/kage/internal/engine/task-manager"
	"github.com/thisismeamir/kage/internal/internal-pkg/config"
	"os"
)

func CheckTaskDependency(task task_manager.Task, conf config.Config) []map[string]int {
	dependencyEvaluationList := make([]map[string]int, 0)
	// Check if the task has any dependencies
	if len(task.FlowDependency) > 0 {
		for _, dep := range task.FlowDependency {
			_, err := os.Stat(conf.BasePath + "/tmp/" + task.FlowIdentifier + "/" + dep + ".output.json")
			if err != nil {
				if os.IsNotExist(err) {
					// Dependency file does not exist, return an error
					dependencyEvaluationList = append(dependencyEvaluationList, map[string]int{dep: -1}) // Means Dependencies are not satisfied (This is an error because we sorted the tasks, so if
					// we reach here, it means the task cannot run)
				} else {
					// Some other error occurred, log it
					dependencyEvaluationList = append(dependencyEvaluationList, map[string]int{dep: -1})
				}
			} else {
				dependencyEvaluationList = append(dependencyEvaluationList, map[string]int{dep: -1})
			}
		}
	} else {
		dependencyEvaluationList = append(dependencyEvaluationList, map[string]int{"NoDep": 1})
	}
	return dependencyEvaluationList
}

func AllowedByDependencies(task task_manager.Task, conf config.Config) bool {
	depsStatus := CheckTaskDependency(task, conf)
	if len(depsStatus) == 0 || (len(depsStatus) == 1 && depsStatus[0]["NoDep"] == 1) {
		// No dependencies or no dependencies to check, we can run the task.
		return true
	} else if len(depsStatus) > 0 {
		for _, depStatus := range depsStatus {
			for _, status := range depStatus {
				if status == -1 {
					// Dependency is not satisfied, we skip the task.
					// Log the skipped task
					return false // -1 means skipped
				}
			}
		}
	}
	return true
}

func FetchDependencies(task task_manager.Task, conf config.Config) []map[string]interface{} {
	dependencies := make([]map[string]interface{}, 0)
	if len(task.FlowDependency) > 0 {
		for _, dep := range task.FlowDependency {
			dependencyFilePath := conf.BasePath + "/tmp/" + task.FlowIdentifier + "/" + dep + ".output.json"
			if _, err := os.Stat(dependencyFilePath); err == nil {
				// Dependency file exists, we can fetch it
				dependencyData, _ := os.ReadFile(dependencyFilePath)
				var dependency map[string]interface{}
				_ = json.Unmarshal(dependencyData, &dependency)
				dependencies = append(dependencies, dependency)

			}
		}
	}
	return dependencies
}
