package form

type Form struct {
	Name     string   `json:"name"`
	Version  string   `json:"version"`
	Type     string   `json:"type"`
	Metadata Metadata `json:"metadata"`
}
