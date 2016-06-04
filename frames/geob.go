package frames

import (
	"bytes"
	"fmt"
)

// GEOB is a general encapsulated object frame
type GEOB struct {
	Frame

	MimeType           string `json:"mime_type"`
	ExternalFilename   string `json:"external_filename"`
	ContentDescription string `json:"content_description"`
	Object             []byte `json:"object"`
}

// DisplayContent will comprehensively display known information
func (g *GEOB) DisplayContent() string {
	return fmt.Sprintf(`Mime Type:   %s
Filename:    %s
Description: %s`, g.MimeType, g.ExternalFilename, g.ContentDescription)
}

// ProcessData will parse bytes for details
func (g *GEOB) ProcessData(s int, d []byte) IFrame {
	g.Size = s
	g.Data = d

	if d[0] == '\x01' {
		g.Utf16 = true
	}
	d = d[1:]

	idx := bytes.IndexByte(d, '\x00')
	g.MimeType = GetStr(d[:idx])
	d = d[idx+1:]

	if !g.Utf16 {
		idx = bytes.IndexByte(d, '\x00')
		g.ExternalFilename = GetStr(d[:idx])
		d = d[idx+LengthStandard:]

		idx = bytes.IndexByte(d, '\x00')
		g.ContentDescription = GetStr(d[:idx])
		g.Object = d[idx+LengthStandard:]
	} else {
		idx = bytes.Index(d, []byte{'\x00', '\x00'})
		g.ExternalFilename = GetUnicodeStr(d[:idx])
		d = d[idx+LengthUnicode:]

		idx = bytes.Index(d, []byte{'\x00', '\x00'})
		g.ContentDescription = GetUnicodeStr(d[:idx])
		g.Object = d[idx+LengthUnicode:]
	}

	return g
}
