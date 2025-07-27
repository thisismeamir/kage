package graph

type GraphExecutionAccessModel struct {
	FileSystem GraphExecutionAccessFileSystemModel `json:"file_system"`
	Web        GraphExecutionAccessWebModel        `json:"web"`
}
