package system_monitor

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

// SystemMonitor struct for holding system information and resource methods
type SystemMonitor struct{}

// NewSystemMonitor creates a new SystemMonitor instance
func NewSystemMonitor() *SystemMonitor {
	return &SystemMonitor{}
}

// CheckCPU returns the current CPU usage percentage
func (sm *SystemMonitor) CheckCPU() ([]float64, error) {
	cpuUsage, err := cpu.Percent(0, false)
	if err != nil {
		return nil, fmt.Errorf("Error checking CPU usage: %v", err)
	}
	return cpuUsage, nil
}

// CheckMemory returns the current memory usage statistics
func (sm *SystemMonitor) CheckMemory() (*mem.VirtualMemoryStat, error) {
	memUsage, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("Error checking memory usage: %v", err)
	}
	return memUsage, nil
}

// CheckNetwork returns the network I/O stats (bytes received/sent)
func (sm *SystemMonitor) CheckNetwork() ([]net.IOCountersStat, error) {
	netStats, err := net.IOCounters(true)
	if err != nil {
		return nil, fmt.Errorf("Error checking network usage: %v", err)
	}
	return netStats, nil
}

// CheckDisk returns the current disk usage statistics
func (sm *SystemMonitor) CheckDisk() (*disk.UsageStat, error) {
	diskStats, err := disk.Usage("/")
	if err != nil {
		return nil, fmt.Errorf("Error checking disk usage: %v", err)
	}
	return diskStats, nil
}

// CheckAll returns CPU, Memory, Network, and Disk usage stats together
func (sm *SystemMonitor) CheckAll() (map[string]interface{}, error) {
	cpuUsage, err := sm.CheckCPU()
	if err != nil {
		return nil, err
	}

	memUsage, err := sm.CheckMemory()
	if err != nil {
		return nil, err
	}

	netStats, err := sm.CheckNetwork()
	if err != nil {
		return nil, err
	}

	diskStats, err := sm.CheckDisk()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"cpu_usage":     cpuUsage,
		"memory_usage":  memUsage,
		"network_stats": netStats,
		"disk_usage":    diskStats,
	}, nil
}
