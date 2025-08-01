package graph

type GraphExecutionModel struct {
	Policy      string                    `json:"policy"`
	Priority    int                       `json:"priority"`
	Timeout     int                       `json:"timeout"`
	PreProcess  string                    `json:"pre_process"`
	PostProcess string                    `json:"post_process"`
	OnError     string                    `json:"on_error"`
	OnFailure   string                    `json:"on_failure"`
	Retry       GraphExecutionRetryModel  `json:"retry"`
	Access      GraphExecutionAccessModel `json:"access"`
}
