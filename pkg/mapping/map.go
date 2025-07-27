package mapping

import "github.com/thisismeamir/kage/pkg/form"

type Map struct {
	form.Form
	Model    MapModel      `json:"model"`
	Metadata form.Metadata `json:"metadata"`
}
