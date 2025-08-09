package engine

import (
	"context"
	execution_system "github.com/thisismeamir/kage/internal/engine/execution-system"
	system_monitor "github.com/thisismeamir/kage/internal/engine/system-monitor"
	"github.com/thisismeamir/kage/internal/internal-pkg/config"
	"github.com/thisismeamir/kage/internal/internal-pkg/registry"
	"log"
	"sync"
	"time"
)

// Engine ties together SystemMonitor, ExecutionSystem, Registry, and Config
type Engine struct {
	SystemMonitor   system_monitor.SystemMonitor     // The system monitor keeps track of the current state and available flows
	ExecutionSystem execution_system.ExecutionSystem // Handles the execution of the flows
	Registry        registry.Registry                // Stores all available flows and their configurations
	Configuration   config.Config                    // The Configuration for the execution system
}

// NewEngine initializes the Engine with the provided components
func NewEngine(conf config.Config, reg registry.Registry, exe execution_system.ExecutionSystem, sysm system_monitor.SystemMonitor) *Engine {
	return &Engine{
		SystemMonitor:   sysm,
		ExecutionSystem: exe,
		Registry:        reg,
		Configuration:   conf,
	}
}

// Routine orchestrates the flow execution cycle in a background process
func (e *Engine) Routine(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure the wait group counter is decremented when the goroutine completes

	ticker := time.NewTicker(time.Second * 5) // Periodically check for available flows every 5 seconds
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done(): // If the context is canceled (e.g., system shutdown)
			log.Println("Routine stopping due to context cancellation.")
			return
		case <-ticker.C: // Triggered by the ticker for periodic execution
			e.executeFlows()
		}
	}
}

// executeFlows fetches and executes the flows based on available system state
func (e *Engine) executeFlows() {
	// Step 1: Fetch new flows from the system monitor (e.g., check for available tasks)
	e.ExecutionSystem.FetchFlows(e.Configuration)

	// Step 2: Identify flows that can be executed from the available flows list
	// You'll need to write logic to filter out flows that are not ready for execution
	e.ExecutionSystem.CreateCurrentlyAvailableFlowsList(&e.SystemMonitor, e.Configuration)

	// Step 3: For each executable flow, execute it using the ExecutionSystem and Registry
	for _, flow := range e.ExecutionSystem.CurrentlyAvailableFlows {
		// Run the flow using the ExecutionSystem
		execution_system.RunFlow(flow.Identifier, e.Configuration, e.Registry)
	}
}
