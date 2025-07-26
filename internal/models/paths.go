package models

// AtomPath Data structure for a path of atoms.
type AtomPath struct {
	Path  string `json:"path"`
	Local bool   `json:"local,omitempty"`
}

type ModulePath struct {
	Path  string `json:"path"`
	Local bool   `json:"local"`
}
