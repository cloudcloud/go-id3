package frames

import (
	"bytes"
	"fmt"
)

// GRID provides a group registration identifier
type GRID struct {
	Frame

	Owner         string `json:"owner"`
	Symbol        byte   `json:"symbol"`
	DependantData []byte `json:"dependant_data"`
}

// DisplayContent will comprehensively display known information
func (g *GRID) DisplayContent() string {
	return fmt.Sprintf("Group Reg Identifier\n\tOwner: %s\n\tSymbol: %b\n", g.Owner, g.Symbol)
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
