package graph

type GraphExecutionAccessWebModel struct {
	Enabled     bool       `json:"enabled"`
	AllowedUrls []string   `json:"allowed_urls"`
	BlockedUrls []string   `json:"blocked_urls"`
	Proxy       ProxyModel `json:"proxy,omitempty"`
}
