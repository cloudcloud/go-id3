package frames

import "fmt"

// PCNT provides a frame to store a count of the times this file has been played
type PCNT struct {
	Frame

	Count int `json:"count"`
}

// DisplayContent will comprehensively display known information
func (p *PCNT) DisplayContent() string {
	return fmt.Sprintf("Count: %d\n", p.Count)
}

// ProcessData will parse bytes for details
func (p *PCNT) ProcessData(s int, d []byte) IFrame {
	p.Size = s
	p.Data = d

	p.Count = GetSize(p.Data, 8)

	return p
}
