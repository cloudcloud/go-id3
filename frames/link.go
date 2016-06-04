package frames

import (
	"bytes"
	"fmt"
)

// LINK provides linked information for the file
type LINK struct {
	Frame

	Identifier     []byte `json:"identifier"`
	URL            string `json:"url"`
	AdditionalData string `json:"additional_data"`
}

// DisplayContent will comprehensively display known information
func (l *LINK) DisplayContent() string {
	return fmt.Sprintf("Linked information\n\tIdentifier: %s\n\tURL: %s\n", l.Identifier, l.URL)
}

// ProcessData will parse bytes for details
func (l *LINK) ProcessData(s int, d []byte) IFrame {
	l.Size = s
	l.Data = d

	l.Identifier = d[:3]
	d = d[3:]

	idx := bytes.IndexByte(d, '\x00')
	l.URL = GetStr(d[:idx])
	l.AdditionalData = GetStr(d[idx+1:])

	return l
}
