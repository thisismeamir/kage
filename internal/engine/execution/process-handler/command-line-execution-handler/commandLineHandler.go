package command_line_execution_handler

import (
	"os/exec"
	"time"
)

// CommandResult holds the result of a command execution
type CommandResult struct {
	Stdout    string
	Stderr    string
	StartTime time.Time
	EndTime   time.Time
	Duration  time.Duration
	Error     error
}

// RunCommand executes a terminal command and returns detailed timing information
func RunCommand(command string, args ...string) *CommandResult {
	result := &CommandResult{
		StartTime: time.Now(),
	}

	// Create the command
	cmd := exec.Command(command, args...)

	// Capture stdout and stderr
	stdout, err := cmd.Output()
	if err != nil {
		// If there's an error, try to get stderr as well
		if exitError, ok := err.(*exec.ExitError); ok {
			result.Stderr = string(exitError.Stderr)
		}
		result.Error = err
	}

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	result.Stdout = string(stdout)

	return result
}

// RunCommandWithCombinedOutput executes a command and captures both stdout and stderr together
func RunCommandWithCombinedOutput(command string, args ...string) *CommandResult {
	result := &CommandResult{
		StartTime: time.Now(),
	}

	// Create the command
	cmd := exec.Command(command, args...)

	// Capture combined output
	output, err := cmd.CombinedOutput()

	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	result.Stdout = string(output)
	result.Error = err

	return result
}
