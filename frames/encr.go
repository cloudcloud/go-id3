package frames

import (
	"bytes"
	"fmt"
)

// ENCR contains information for encryption registration
type ENCR struct {
	Frame

	Owner          string `json:"owner"`
	Method         byte   `json:"method"`
	EncryptionData []byte `json:"encryption_data"`
}

// DisplayContent will comprehensively display known information
func (e *ENCR) DisplayContent() string {
	return fmt.Sprintf("Owner: %s\nMethod: %v", e.Owner, e.Method)
}

// ProcessData will parse bytes for details
func (e *ENCR) ProcessData(s int, d []byte) IFrame {
	e.Size = s
	e.Data = d

	idx := bytes.IndexByte(d, '\x00')
	e.Owner = GetStr(d[:idx])
	e.Method = d[idx]
	if len(d) > idx {
		e.EncryptionData = d[idx+1:]
	}

	return e
}
