package form

type Metadata struct {
	Description string   `json:"description"`
	Authors     []Author `json:"authors"`
	Manual      string   `json:"manual"`
	Repository  string   `json:"repository"`
}
