package frames

import "bytes"

// PRIV provides a private frame
type PRIV struct {
	Frame

	Owner       string `json:"owner"`
	PrivateData []byte `json:"private_data"`
}

// Init will provide the initial values
func (p *PRIV) Init(n, d string, v int) {
	p.Name = n
	p.Description = d
	p.Version = v
}

// DisplayContent will comprehensively display known information
func (p *PRIV) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (p *PRIV) GetExplain() string {
	return ""
}

// GetLength will provide the length of frame
func (p *PRIV) GetLength() string {
	return ""
}

// GetName will provide the Name of EQUA
func (p *PRIV) GetName() string {
	return p.Name
}

// ProcessData will parse bytes for details
func (p *PRIV) ProcessData(s int, d []byte) IFrame {
	p.Size = s
	p.Data = d

	idx := bytes.IndexByte(d, '\x00')
	p.Owner = GetStr(d[:idx])
	p.PrivateData = d[idx+1:]

	return p
}
