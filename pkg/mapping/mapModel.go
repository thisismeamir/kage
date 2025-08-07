package mapping

type MapModel struct {
	InputSchemas []map[string]interface{} `json:"input_schema"`
	OutputSchema map[string]interface{}   `json:"output_schema"`
}
