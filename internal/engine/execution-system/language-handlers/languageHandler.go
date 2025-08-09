package language_handlers

import (
	"bytes"
	"fmt"
	task_manager "github.com/thisismeamir/kage/internal/engine/task-manager"
	"github.com/thisismeamir/kage/pkg/node"
	"log"
	"os"
	execute "os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Execution contains the context for node execution
type Execution struct {
	WorkingDirectory string
	Environment      map[string]string
	Input            string
	Task             task_manager.Task
	Node             node.Node
}

// ExecutionResult contains the result of node execution
type ExecutionResult struct {
	Error          string        `json:"error,omitempty"`
	ExitCode       int           `json:"exit_code"`
	Duration       time.Duration `json:"duration"`
	OutputJsonPath string        `json:"output_json_path"`
	Stdout         string        `json:"stdout,omitempty"`
	Stderr         string        `json:"stderr,omitempty"`
}

// LanguageHandler defines the interface for language-specific execution
type LanguageHandler interface {
	Execute(execCtx *Execution) (*ExecutionResult, error)
	Validate(execCtx *Execution) error
	GetRequiredFiles(node *node.Node) []string
}

// LanguageHandlerRegistry manages different language handlers
type LanguageHandlerRegistry struct {
	handlers map[string]LanguageHandler
}

// NewLanguageHandlerRegistry creates a new registry with default handlers
func NewLanguageHandlerRegistry() *LanguageHandlerRegistry {
	registry := &LanguageHandlerRegistry{
		handlers: make(map[string]LanguageHandler),
	}

	// Register default handlers
	registry.RegisterHandler("python", &PythonHandler{})

	return registry
}

// RegisterHandler registers a new language handler
func (r *LanguageHandlerRegistry) RegisterHandler(language string, handler LanguageHandler) {
	r.handlers[strings.ToLower(language)] = handler
}

// GetHandler returns the handler for a specific language
func (r *LanguageHandlerRegistry) GetHandler(language string) (LanguageHandler, error) {
	handler, exists := r.handlers[strings.ToLower(language)]
	if !exists {
		return nil, fmt.Errorf("unsupported language: %s", language)
	}
	return handler, nil
}

// ExecuteNode executes a node using the appropriate language handler
func (r *LanguageHandlerRegistry) ExecuteTask(exec *Execution) (*ExecutionResult, error) {
	handler, err := r.GetHandler(exec.Node.Model.ExecutionModel.Language.Name)
	log.Printf(exec.Node.Model.ExecutionModel.Language.Name)
	if err != nil {
		return nil, err
	}

	// Validate node before execution
	if err := handler.Validate(exec); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	return handler.Execute(exec)
}

// BaseHandler provides common functionality for all handlers
type BaseHandler struct{}

// prepareWorkingDirectory ensures the working directory exists and is accessible
func (h *BaseHandler) prepareWorkingDirectory(workDir string) error {
	if workDir == "" {
		return nil
	}
	_, err := os.Stat(workDir)
	if err != nil {
		os.MkdirAll(workDir, 0755)
	}
	return nil
}

// executeCommand executes a command with timeout and captures output
func (h *BaseHandler) executeCommand(cmd *execute.Cmd) (*ExecutionResult, error) {
	start := time.Now()
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	duration := time.Since(start)

	result := &ExecutionResult{
		Duration: duration,
		Stdout:   stdout.String(),
		Stderr:   stderr.String(),
	}

	if err != nil {
		if exitErr, ok := err.(*execute.ExitError); ok {
			result.ExitCode = exitErr.ExitCode()
		} else {
			result.ExitCode = -1
		}
		result.Error = err.Error()
	}

	return result, nil
}

// PythonHandler
type PythonHandler struct {
	BaseHandler
}

func (h *PythonHandler) Execute(exec *Execution) (*ExecutionResult, error) {
	workDir := exec.Node.Model.WorkingDirectory
	if exec.WorkingDirectory != "" {
		workDir = exec.WorkingDirectory
	}

	if err := h.prepareWorkingDirectory(workDir); err != nil {
		return nil, err
	}

	// Prepare command
	executable := exec.Node.Model.ExecutionModel.Language.ExecutablePath
	if executable == "" {
		executable = "python3"
	}

	entryFile := filepath.Join(exec.Node.Model.Source, exec.Node.Model.EntryFile)
	outputJsonPath := exec.Node.Model.OutputDirectory + exec.Task.Identifier + ".output.json"

	log.Printf("Output json path: %s", outputJsonPath)
	_ = os.MkdirAll(exec.Node.Model.OutputDirectory, 0755)
	cmd := execute.Command(
		executable,
		entryFile,
		"--input", exec.Input,
		"--output-json", outputJsonPath,
	)
	cmd.Dir = workDir

	// Set environment variables
	cmd.Env = os.Environ()
	for k, v := range exec.Environment {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

	result, _ := h.executeCommand(cmd)

	result.OutputJsonPath = exec.Node.Model.OutputDirectory + exec.Task.Identifier + ".output.json"
	return result, nil
}

func (h *PythonHandler) Validate(exec *Execution) error {
	if exec.Node.Model.EntryFile == "" {
		return fmt.Errorf("entry file is required for Python nodes")
	}
	if !strings.HasSuffix(exec.Node.Model.EntryFile, ".py") {
		return fmt.Errorf("entry file must be a Python file (.py)")
	}
	return nil
}

func (h *PythonHandler) GetRequiredFiles(node *node.Node) []string {
	return []string{node.Model.EntryFile, "requirements.txt"}
}

func (h *PythonHandler) parseOutput(exec *Execution, result *ExecutionResult) error {
	if exec.Node.Model.OutputDirectory != "" {
		outputFile := filepath.Join(exec.Node.Model.OutputDirectory, "output.json")
		result.OutputJsonPath = outputFile
	}

	return nil
}
