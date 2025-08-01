package form

// FormPath Data structure for a path of json.
type FormPath struct {
	Path  string `json:"path"`
	Local bool   `json:"local,omitempty"`
}
