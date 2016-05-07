package frames

import "bytes"

// LINK provides linked information for the file
type LINK struct {
	Frame

	Identifier     []byte `json:"identifier"`
	URL            string `json:"url"`
	AdditionalData string `json:"additional_data"`
}

// Init will provide the initial values
func (l *LINK) Init(n, d string, v int) {
	l.Name = n
	l.Description = d
	l.Version = v
}

// DisplayContent will comprehensively display known information
func (l *LINK) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (l *LINK) GetExplain() string {
	return ""
}

// GetLength will provide the length of frame
func (l *LINK) GetLength() string {
	return ""
}

// GetName will provide the Name of EQUA
func (l *LINK) GetName() string {
	return l.Name
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
