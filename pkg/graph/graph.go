package graph

import (
	"github.com/thisismeamir/kage/pkg/form"
)

type Graph struct {
	form.Form
	Model    GraphModel    `json:"model"`
	Metadata form.Metadata `json:"metadata"`
}
