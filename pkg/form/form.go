package form

type Form struct {
	Name       string   `json:"name"`
	Identifier string   `json:"identifier"`
	Version    string   `json:"version"`
	Type       string   `json:"type"`
	Metadata   Metadata `json:"metadata"`
}
