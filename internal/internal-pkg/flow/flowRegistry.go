package flow

type FlowRegistry struct {
	Flows []FlowRegister `json:"flows"`
}

type FlowRegister struct {
	Identifier string `json:"identifier"`
	Path       string `json:"path"`
}

func (fr *FlowRegistry) AddFlow(f Flow) {
	fr.Flows = append(fr.Flows, FlowRegister{
		Identifier: f.Identifier,
	})

}
