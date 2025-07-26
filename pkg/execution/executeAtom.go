package execution

import (
	"fmt"
	"strings"

	"github.com/thisismeamir/kage/pkg/atom"
)

func ExecuteAtom(atom atom.AtomModel, input map[string]interface{}) (map[string]interface{}, error) {
	handler, ok := LanguageHandlerMap[strings.ToLower(atom.ExecutionModel.Language)]
	if !ok {
		return nil, fmt.Errorf("no handler for language: %s", atom.ExecutionModel.Language)
	}
	return handler.Run(atom.ExecutionModel.Source, input)
}
