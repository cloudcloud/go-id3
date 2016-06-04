package frames

import (
	"bytes"
	"fmt"
)

// POPM provides a popularity measurement frame for this file individually
type POPM struct {
	Frame

	Email      string `json:"email"`
	Popularity int    `json:"popularity"`
	Counter    int    `json:"counter"`
}

// DisplayContent will comprehensively display known information
func (p *POPM) DisplayContent() string {
	return fmt.Sprintf("Popularimeter\n\tEmail: %s\n\tPopularity: %d\n\tCount: %d\n",
		p.Email,
		p.Popularity,
		p.Counter)
}

// ProcessData will parse bytes for details
func (p *POPM) ProcessData(s int, d []byte) IFrame {
	p.Size = s
	p.Data = d

	idx := bytes.IndexByte(d, '\x00')
	p.Email = GetStr(d[:idx])
	p.Popularity = GetDirectInt(d[idx+1])
	p.Counter = GetSize(d[idx+2:], 8)

	return p
}
