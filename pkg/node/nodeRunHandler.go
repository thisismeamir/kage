package node

type RunHandler interface {
	Run(source string, input map[string]interface{}) (map[string]interface{}, error)
}
