package form

type Form struct {
	Name       string   `json:"name"`
	Identifier string   `json:"identifier"`
	Version    string   `json:"version"`
	Type       string   `json:"type"`
	Path       string   `json:"path"`
	Metadata   Metadata `json:"metadata"`
}
