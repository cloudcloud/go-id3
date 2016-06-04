package frames

import (
	"bytes"
	"fmt"
)

// TXXX provides the user string from the file
type TXXX struct {
	Frame

	Type  string `json:"type"`
	Value string `json:"value"`
}

// DisplayContent will comprehensively display known information
func (t *TXXX) DisplayContent() string {
	return fmt.Sprintf("User text (%s):(%s)\n", t.Type, t.Value)
}

// ProcessData will handle the acquisition of all data
func (t *TXXX) ProcessData(s int, d []byte) IFrame {
	t.Size = s
	t.Data = d

	// text encoding is a single byte, 0 for latin, 1 for unicode
	if len(d) > 2 {
		enc := d[0]
		d = d[1:]

		if enc == '\x00' {
			idx := bytes.IndexByte(d, '\x00')
			t.Type = GetStr(d[:idx])
			t.Value = GetStr(d[idx+LengthStandard:])
		} else if enc == '\x01' {
			t.Utf16 = true

			idx := bytes.Index(d, []byte{'\x00', '\x00'})
			t.Type = GetUnicodeStr(d[:idx])
			t.Value = GetUnicodeStr(d[idx+LengthUnicode:])
		}
	}

	return t
}
