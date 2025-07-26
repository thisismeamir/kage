package config

// AddAtomPathResponse is the response structure for checking the existence and validity of an atom path.
type AddAtomPathResponse struct {
	AtomPath string `json:"atomPath"`
	Added    bool   `json:"added"`
	Message  string `json:"message"`
}

// DeleteAtomPathResponse is the response structure for removing an atom path.
type DeleteAtomPathResponse struct {
	AtomPath string `json:"atomPath"`
	Deleted  bool   `json:"removed"`
	Message  string `json:"message,omitempty" jsonschema:"omitempty" jsonschema_extras:"description=Message about the removal status"`
}
