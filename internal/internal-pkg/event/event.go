package event

type Event struct {
	Identifier   string                 `json:"identifier,omitempty"`
	Date         string                 `json:"date"`
	Graph        string                 `json:"graph"`
	Urgency      int                    `json:"urgency"`
	InitialInput map[string]interface{} `json:"initial_input,omitempty"`
}

func (e Event) GenerateIdentifier() Event {
	e.Identifier = e.Date + ".event"
	return e
}

func (e Event) GetIdentifier() string {
	if e.Identifier == "" {
		e.GenerateIdentifier()
	}
	return e.Identifier
}
