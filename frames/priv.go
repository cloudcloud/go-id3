package frames

import (
	"bytes"
	"fmt"
)

// PRIV provides a private frame
type PRIV struct {
	Frame

	Owner       string `json:"owner"`
	PrivateData []byte `json:"private_data"`
}

// DisplayContent will comprehensively display known information
func (p *PRIV) DisplayContent() string {
	return fmt.Sprintf("Private owner: %s\nData: %#x\n", p.Owner, p.PrivateData)
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
