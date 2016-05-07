package frames

// PCNT provides a frame to store a count of the times this file has been played
type PCNT struct {
	Frame

	Count []byte `json:"count"`
}

// Init will provide the initial values
func (p *PCNT) Init(n, d string, v int) {
	p.Name = n
	p.Description = d
	p.Version = v
}

// DisplayContent will comprehensively display known information
func (p *PCNT) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (p *PCNT) GetExplain() string {
	return ""
}

// GetLength will provide the length of frame
func (p *PCNT) GetLength() string {
	return ""
}

// GetName will provide the Name of EQUA
func (p *PCNT) GetName() string {
	return p.Name
}

// ProcessData will parse bytes for details
func (p *PCNT) ProcessData(s int, d []byte) IFrame {
	p.Size = s
	p.Data = d
	p.Count = p.Data

	return p
}
