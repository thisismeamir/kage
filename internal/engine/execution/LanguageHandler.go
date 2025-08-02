package execution

import (
	"github.com/thisismeamir/kage/internal/engine/execution/handlers"
	atom "github.com/thisismeamir/kage/pkg/node"
)

var LanguageHandlerMap = map[string]atom.RunHandler{
	"python":     &handlers.PythonHandler{},
	"go":         &handlers.GoHandler{},
	"bash":       &handlers.BashHandler{},
	"javascript": &handlers.JavaScriptHandler{},
}
