package system_resource_management

type SystemResource struct {
	CpuUsage      float64 `json:"cpu_usage"` // CPU usage percentage
	RamUsage      float64 `json:"ram_usage"`
	GraphicsUsage float64 `json:"graphics_usage"`
	DiskUsage     float64 `json:"disk_usage"`
}
