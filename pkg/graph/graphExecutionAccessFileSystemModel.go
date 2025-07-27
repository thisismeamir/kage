package graph

type GraphExecutionAccessFileSystemModel struct {
	Read    bool `json:"read"`
	Write   bool `json:"write"`
	Delete  bool `json:"delete"`
	Execute bool `json:"execute"`
}
