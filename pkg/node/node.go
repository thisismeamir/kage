package node

import "github.com/thisismeamir/kage/pkg/form"

type Node struct {
	form.Form
	Model    NodeModel     `json:"model"`
	Metadata form.Metadata `json:"metadata"`
}
