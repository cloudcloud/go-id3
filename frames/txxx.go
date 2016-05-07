package frames

import "bytes"

// TXXX provides the user string from the file
type TXXX struct {
	Frame

	Type  string `json:"type"`
	Value string `json:"value"`
}

// Init will provide the initial values
func (t *TXXX) Init(n, d string, v int) {
	t.Name = n
	t.Description = d
	t.Version = v
}

// DisplayContent will comprehensively display known information
func (t *TXXX) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (t *TXXX) GetExplain() string {
	return t.Description
}

// GetLength will provide the length
func (t *TXXX) GetLength() string {
	return ""
}

// GetName will provide the Name
func (t *TXXX) GetName() string {
	return t.Name
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
