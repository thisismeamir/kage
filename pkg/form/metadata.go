package form

type Metadata struct {
	Description string   `json:"description,omitempty"`
	Authors     []Author `json:"authors,omitempty"`
	Manual      string   `json:"manual,omitempty"`
	Repository  string   `json:"repository,omitempty"`
}
