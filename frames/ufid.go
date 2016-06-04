package frames

import (
	"bytes"
	"fmt"
)

// UFID provides descriptor of uniqueness for the file
type UFID struct {
	Frame

	Owner      string `json:"owner"`
	Identifier []byte `json:"identifier"`
}

// DisplayContent will comprehensively display known information
func (u *UFID) DisplayContent() string {
	return fmt.Sprintf("Owner: (%s) Identifier: (%x)\n", u.Owner, u.Identifier)
}

// ProcessData will handle the acquisition of all data
func (u *UFID) ProcessData(s int, d []byte) IFrame {
	u.Size = s
	u.Data = d

	if len(d) > 2 {
		idx := bytes.IndexByte(d, '\x00')
		u.Owner = GetStr(d[:idx])
		u.Identifier = d[idx+1:]
	}

	return u
}
