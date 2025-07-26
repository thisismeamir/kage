package execution

import (
	atom "github.com/thisismeamir/kage/pkg/atom"
	"github.com/thisismeamir/kage/pkg/execution/handlers"
)

var LanguageHandlerMap = map[string]atom.AtomRunHandler{
	"python":     &handlers.PythonHandler{},
	"go":         &handlers.GoHandler{},
	"bash":       &handlers.BashHandler{},
	"javascript": &handlers.JavaScriptHandler{},
}
