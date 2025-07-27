package graph

import "github.com/thisismeamir/kage/pkg/mapping"

type GraphVertex struct {
	ToId int         `json:"to_id"`
	Map  mapping.Map `json:"map"`
}
