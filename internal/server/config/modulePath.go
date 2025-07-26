package config

// AddModulePathResponse is the response structure for checking the existence and validity of an atom path.
type AddModulePathResponse struct {
	ModulePath string `json:"module_path"`
	Added      bool   `json:"added"`
	Message    string `json:"message"`
}

// DeleteModulePathResponse is the response structure for removing an atom path.
type DeleteModulePathResponse struct {
	ModulePath string `json:"module_path"`
	Deleted    bool   `json:"removed"`
	Message    string `json:"message,omitempty" jsonschema:"omitempty" jsonschema_extras:"description=Message about the removal status"`
}
