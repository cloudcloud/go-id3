package frames

import "bytes"

// POPM provides a popularity measurement frame for this file individually
type POPM struct {
	Frame

	Email      string `json:"email"`
	Popularity byte   `json:"popularity"`
	Counter    []byte `json:"counter"`
}

// Init will provide the initial values
func (p *POPM) Init(n, d string, v int) {
	p.Name = n
	p.Description = d
	p.Version = v
}

// DisplayContent will comprehensively display known information
func (p *POPM) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (p *POPM) GetExplain() string {
	return ""
}

// GetLength will provide the length of frame
func (p *POPM) GetLength() string {
	return ""
}

// GetName will provide the Name of EQUA
func (p *POPM) GetName() string {
	return p.Name
}

// ProcessData will parse bytes for details
func (p *POPM) ProcessData(s int, d []byte) IFrame {
	p.Size = s
	p.Data = d

	idx := bytes.IndexByte(d, '\x00')
	p.Email = GetStr(d[:idx])
	p.Popularity = d[idx+1]
	p.Counter = d[idx+2:]

	return p
}
