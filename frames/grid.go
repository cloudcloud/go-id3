package frames

import "bytes"

// GRID provides a group registration identifier
type GRID struct {
	Frame

	Owner         string `json:"owner"`
	Symbol        byte   `json:"symbol"`
	DependantData []byte `json:"dependant_data"`
}

// Init will provide the initial values
func (g *GRID) Init(n, d string, v int) {
	g.Name = n
	g.Description = d
	g.Version = v
}

// DisplayContent will comprehensively display known information
func (g *GRID) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (g *GRID) GetExplain() string {
	return ""
}

// GetLength will provide the length of frame
func (g *GRID) GetLength() string {
	return ""
}

// GetName will provide the Name of EQUA
func (g *GRID) GetName() string {
	return g.Name
}

// ProcessData will parse bytes for details
func (g *GRID) ProcessData(s int, d []byte) IFrame {
	g.Size = s
	g.Data = d

	idx := bytes.IndexByte(d, '\x00')
	g.Owner = GetStr(d[:idx])
	g.Symbol = d[idx+1]
	g.DependantData = d[idx+2:]

	return g
}
