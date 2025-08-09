package process

import (
	"fmt"
	dependency_handler "github.com/thisismeamir/kage/internal/engine/execution-system/dependency-handler"
	language_handlers "github.com/thisismeamir/kage/internal/engine/execution-system/language-handlers"
	task_manager "github.com/thisismeamir/kage/internal/engine/task-manager"
	"github.com/thisismeamir/kage/internal/internal-pkg/config"
	"github.com/thisismeamir/kage/internal/internal-pkg/registry"
	"github.com/thisismeamir/kage/util"
	"log"
	"os"
)

func ProcessTask(task *task_manager.Task, conf config.Config, reg registry.Registry) int {
	// First off, we make sure that all of the dependencies are satisfied. Otherwise, we skip the task.
	var status int
	if task.Status == 0 {
		if dependency_handler.AllowedByDependencies(*task, conf) {
			status = 1
			deps := dependency_handler.FetchDependencies(*task, conf)
			switch task.ExecutionType {
			case "node":
				log.Printf("[RUNNING] Node Task %s", task.Identifier)
				if task.Level == 0 {
					inputFile := util.LoadJson(task.Input[0])
					status = NodeProcessHandler(*task, inputFile, conf, reg)
				} else {
					status = NodeProcessHandler(*task, deps[0], conf, reg)
				}
			case "map":
				log.Printf("[RUNNING] Map Task %s", task.Identifier)
				status = MapProcessHandler(*task, deps, conf, reg)
			default:
				log.Printf("Undefined task type: %s. Skipping task: %s", task.Type, task.Identifier)
			}
		}
	} else {
		log.Print("Task is Already Done.")
	}
	return status
}

func MapProcessHandler(task task_manager.Task, deps []map[string]interface{}, conf config.Config, reg registry.Registry) int {
	mapp, _ := reg.LoadMap(task.NodeIdentifier)
	output := mapp.MapList(deps)
	util.SaveJson(output, conf.BasePath+"/tmp/"+task.FlowIdentifier+"/"+task.Identifier+".output.json")
	return 2
}

// NodeProcessHandler handles the execution of node-type tasks
func NodeProcessHandler(t task_manager.Task, dep map[string]interface{}, conf config.Config, reg registry.Registry) int {
	// Load the node from registry
	n, err := reg.LoadNode(t.NodeIdentifier)
	if err != nil {
		log.Printf("Failed to load node %s: %v", t.NodeIdentifier, err)
		return -1 // Error status
	}

	// Create language handler registry
	handlerRegistry := language_handlers.NewLanguageHandlerRegistry()

	// Prepare input data from the task's input files and dependencies
	taskInputPath := conf.BasePath + "/tmp/" + t.FlowIdentifier + "/" + t.Identifier + ".input.json"
	util.SaveJson(dep, taskInputPath)

	if err != nil {
		log.Printf("Failed to prepare input for task %s: %v", t.Identifier, err)
		return -1 // Error status
	}

	// Create execution context
	exec := &language_handlers.Execution{
		Node:             *n,
		Task:             t,
		WorkingDirectory: n.Model.WorkingDirectory,
		Input:            taskInputPath,
		Environment:      buildEnvironment(t, conf),
	}

	// Execute the node
	result, err := handlerRegistry.ExecuteTask(exec)
	if err != nil {
		log.Printf("Failed to execute node %s for task %s: %v", t.NodeIdentifier, t.Identifier, err)
		return -1 // Error status
	}
	util.SaveRuntimeDataInCSV(*result, t.NodeIdentifier, conf)
	// Process the execution result
	return processExecutionResult(t, result, conf)
}

// buildEnvironment creates environment variables for the execution
func buildEnvironment(task task_manager.Task, conf config.Config) map[string]string {
	env := make(map[string]string)

	// Add task-specific environment variables
	env["TASK_IDENTIFIER"] = task.Identifier
	env["FLOW_IDENTIFIER"] = task.FlowIdentifier
	env["NODE_IDENTIFIER"] = task.NodeIdentifier
	env["EXECUTION_TYPE"] = task.ExecutionType
	env["TASK_LEVEL"] = fmt.Sprintf("%d", task.Level)
	env["TASK_QUEUE"] = fmt.Sprintf("%d", task.Queue)

	// Add any environment variables from config
	/// TODO: We don't support environment variables right now.
	//if configEnv := nil; configEnv != nil {
	//	for k, v := range configEnv {
	//		env[k] = v
	//	}
	//}

	return env
}

// processExecutionResult processes the execution result and determines the task status
func processExecutionResult(task task_manager.Task, result *language_handlers.ExecutionResult, conf config.Config) int {
	// Log execution details
	log.Printf("Task %s executed in %v", task.Identifier, result.Duration)

	// If there was an execution error, return error status
	if result.Error != "" {
		log.Printf("Task %s execution error: %s", task.Identifier, result.Error)
		if result.Stderr != "" {
			log.Printf("Task %s stderr: %s", task.Identifier, result.Stderr)
		}
		return -1 // Error status
	}

	// If exit code is not 0, consider it a failure
	if result.ExitCode != 0 {
		log.Printf("Task %s exited with code %d", task.Identifier, result.ExitCode)
		if result.Stderr != "" {
			log.Printf("Task %s stderr: %s", task.Identifier, result.Stderr)
		}
		return -1 // Error status
	}

	// Save the output if needed
	if err := saveTaskOutput(task, result, conf); err != nil {
		log.Printf("Failed to save output for task %s: %v", task.Identifier, err)
		// This might not be a fatal error, depending on your requirements
	}

	// Log successful execution
	if result.Stdout != "" {
		log.Printf("Task %s stdout: %s", task.Identifier, result.Stdout)
	}

	return 2 // Success status
}

// saveTaskOutput saves the execution output to a file or storage system
func saveTaskOutput(task task_manager.Task, result *language_handlers.ExecutionResult, conf config.Config) error {
	// Only save if there's actual output
	if result.OutputJsonPath == "" {
		return nil
	}
	outputPath := conf.BasePath + "/tmp/" + task.FlowIdentifier + "/" + task.Identifier + ".output.json"
	os.Rename(result.OutputJsonPath, outputPath)
	return nil
}
