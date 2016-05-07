package frames

// POSS provides a frame for position synchronisation
type POSS struct {
	Frame

	Format   byte   `json:"format"`
	Position []byte `json:"position"`
}

// Init will provide the initial values
func (p *POSS) Init(n, d string, v int) {
	p.Name = n
	p.Description = d
	p.Version = v
}

// DisplayContent will comprehensively display known information
func (p *POSS) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (p *POSS) GetExplain() string {
	return ""
}

// GetLength will provide the length of frame
func (p *POSS) GetLength() string {
	return ""
}

// GetName will provide the Name of EQUA
func (p *POSS) GetName() string {
	return p.Name
}

// ProcessData will parse bytes for details
func (p *POSS) ProcessData(s int, d []byte) IFrame {
	p.Size = s
	p.Data = d

	p.Format = d[0]
	p.Position = d[1:]

	return p
}
