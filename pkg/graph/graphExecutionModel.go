package graph

type GraphExecutionModel struct {
	Policy      GraphExecutionPolicy      `json:"policy,omitempty"`
	Priority    int                       `json:"priority"`
	Timeout     int                       `json:"timeout"`
	PreProcess  string                    `json:"pre_process"`
	PostProcess string                    `json:"post_process"`
	OnError     string                    `json:"on_error"`
	OnFailure   string                    `json:"on_failure"`
	Retry       GraphExecutionRetryModel  `json:"retry"`
	Access      GraphExecutionAccessModel `json:"access"`
}

// GraphExecutionPolicy defines the conditions under which a flow can run
type GraphExecutionPolicy struct {
	CPUUsageThreshold     float64 `json:"cpu_usage_threshold"`     // Maximum CPU usage to allow execution
	MemoryUsageThreshold  float64 `json:"memory_usage_threshold"`  // Maximum memory usage to allow execution
	DiskUsageThreshold    float64 `json:"disk_usage_threshold"`    // Maximum disk usage to allow execution
	NetworkUsageThreshold float64 `json:"network_usage_threshold"` // Maximum network usage to allow execution
}
