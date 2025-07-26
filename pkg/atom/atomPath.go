package atom

// AtomPath Data structure for a path of atoms.
type AtomPath struct {
	Path  string `json:"path"`
	Local bool   `json:"local,omitempty"`
}
