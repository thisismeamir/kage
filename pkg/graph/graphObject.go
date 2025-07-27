package graph

type GraphObject struct {
	Id             int           `json:"id"`
	NodeIdentifier string        `json:"node_identifier"`
	Vertices       []GraphVertex `json:"vertices"`
}
