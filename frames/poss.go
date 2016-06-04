package frames

import "fmt"

// POSS provides a frame for position synchronisation
type POSS struct {
	Frame

	Format   string `json:"format"`
	Position []byte `json:"position"`
}

// DisplayContent will comprehensively display known information
func (p *POSS) DisplayContent() string {
	return fmt.Sprintf("Position synchronisation\n\tFormat: %s\n\tPositions: %#x\n",
		p.Format,
		p.Position)
}

// ProcessData will parse bytes for details
func (p *POSS) ProcessData(s int, d []byte) IFrame {
	p.Size = s
	p.Data = d

	format := GetSize([]byte{d[0]}, 1)
	if format == 1 {
		p.Format = "MPEG"
	} else {
		p.Format = "Milliseconds"
	}

	p.Position = d[1:]

	return p
}
