package frames

import "fmt"

// SIGN provides signature details for the file
type SIGN struct {
	Frame

	Symbol    byte   `json:"symbol"`
	Signature []byte `json:"signature"`
}

// DisplayContent provides a clean display of key information
func (i *SIGN) DisplayContent() string {
	return fmt.Sprintf("Signature (%x): %x\n", i.Symbol, i.Signature)
}

// ProcessData will take bytes and mush into something useful
func (i *SIGN) ProcessData(s int, d []byte) IFrame {
	i.Size = s
	i.Data = d

	i.Symbol = d[0]
	i.Signature = d[1:]

	return i
}
