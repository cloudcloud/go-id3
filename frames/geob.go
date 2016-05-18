package frames

import "bytes"

// GEOB is a general encapsulated object frame
type GEOB struct {
	Frame

	Encoding           byte   `json:"encoding"`
	MimeType           string `json:"mime_type"`
	ExternalFilename   string `json:"external_filename"`
	ContentDescription string `json:"content_description"`
	Object             []byte `json:"object"`
}

// Init will provide the initial values
func (g *GEOB) Init(n, d string, v int) {
	g.Name = n
	g.Description = d
	g.Version = v
}

// DisplayContent will comprehensively display known information
func (g *GEOB) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (g *GEOB) GetExplain() string {
	return ""
}

// GetLength will provide the length of frame
func (g *GEOB) GetLength() string {
	return ""
}

// GetName will provide the Name of EQUA
func (g *GEOB) GetName() string {
	return g.Name
}

// ProcessData will parse bytes for details
func (g *GEOB) ProcessData(s int, d []byte) IFrame {
	g.Size = s
	g.Data = d

	g.Encoding = d[0]
	idx := bytes.IndexByte(d, '\x00')
	g.MimeType = GetStr(d[1:idx])
	d = d[idx+1:]

	if g.Encoding == '\x00' {
		idx = bytes.IndexByte(d, '\x00')
		g.ExternalFilename = GetStr(d[:idx])
		d = d[idx+1:]

		idx = bytes.IndexByte(d, '\x00')
		g.ContentDescription = GetStr(d[:idx])
		g.Object = d[idx+1:]
	} else if g.Encoding == '\x01' {
		g.Utf16 = true

		idx = bytes.Index(d, []byte{'\x00', '\x00'})
		g.ExternalFilename = GetUnicodeStr(d[:idx])
		d = d[idx+2:]

		idx = bytes.Index(d, []byte{'\x00', '\x00'})
		g.ContentDescription = GetUnicodeStr(d[:idx])
		g.Object = d[idx+2:]
	}

	return g
}
