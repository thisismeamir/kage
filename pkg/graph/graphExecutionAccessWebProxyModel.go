package graph

type ProxyModel struct {
	Enabled  bool   `json:"enabled"`
	Url      string `json:"url"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
