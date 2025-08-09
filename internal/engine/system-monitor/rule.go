package system_monitor

import (
	task_manager "github.com/thisismeamir/kage/internal/engine/task-manager"
	"github.com/thisismeamir/kage/internal/internal-pkg/config"
	"log"
)

type CheckOptions struct {
	SkipCPU     bool
	SkipMemory  bool
	SkipDisk    bool
	SkipNetwork bool
}

func IsFlowAbleToRun(fl task_manager.Flow, sm *SystemMonitor, conf config.Config, opts CheckOptions) bool {
	// Default to the execution policy if not available in the flow
	execPolicy := fl.ExecutionModel.Policy
	if execPolicy.CPUUsageThreshold == 0 {
		execPolicy.CPUUsageThreshold = conf.DefaultExecutionPolicy.CPUUsageThreshold
	}
	if execPolicy.MemoryUsageThreshold == 0 {
		execPolicy.MemoryUsageThreshold = conf.DefaultExecutionPolicy.MemoryUsageThreshold
	}
	if execPolicy.DiskUsageThreshold == 0 {
		execPolicy.DiskUsageThreshold = conf.DefaultExecutionPolicy.DiskUsageThreshold
	}
	if execPolicy.NetworkUsageThreshold == 0 {
		execPolicy.NetworkUsageThreshold = conf.DefaultExecutionPolicy.NetworkUsageThreshold
	}

	// Check CPU usage (if not skipped)
	if !opts.SkipCPU {
		cpuUsage, err := sm.CheckCPU()
		if err != nil {
			log.Printf("[ERROR] Failed to check CPU usage: %v", err)
			return false
		}
		if cpuUsage[0] > execPolicy.CPUUsageThreshold {
			log.Printf("[NOT ALLOWED] CPU usage: %f exceeds threshold", cpuUsage[0])
			return false
		}
	}

	// Check memory usage (if not skipped)
	if !opts.SkipMemory {
		memUsage, err := sm.CheckMemory()
		if err != nil {
			log.Printf("[ERROR] Failed to check memory usage: %v", err)
			return false
		}
		if memUsage.UsedPercent > execPolicy.MemoryUsageThreshold {
			log.Printf("[NOT ALLOWED] Memory usage: %f exceeds threshold", memUsage.UsedPercent)
			return false
		}
	}

	// Check disk usage (if not skipped)
	if !opts.SkipDisk {
		diskStats, err := sm.CheckDisk()
		if err != nil {
			log.Printf("[ERROR] Failed to check disk usage: %v", err)
			return false
		}
		if diskStats.UsedPercent > execPolicy.DiskUsageThreshold {
			log.Printf("[NOT ALLOWED] Disk usage: %f exceeds threshold", diskStats.UsedPercent)
			return false
		}
	}

	// Check network usage (if not skipped)
	if !opts.SkipNetwork {
		netStats, err := sm.CheckNetwork()
		if err != nil {
			log.Printf("[ERROR] Failed to check network usage: %v", err)
			return false
		}
		if len(netStats) == 0 || float64(netStats[0].BytesRecv/1024.0) > execPolicy.NetworkUsageThreshold {
			log.Printf("[NOT ALLOWED] Network usage: %f exceeds threshold", netStats[0].BytesRecv)
			return false
		}
	}

	// If all conditions are satisfied, return true
	return true
}
