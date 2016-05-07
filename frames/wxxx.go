package frames

import "bytes"

// WXXX provides user defined webpage links
type WXXX struct {
	Frame

	Title string `json:"title"`
	URL   string `json:"url"`
}

// Init will provide the initial values
func (w *WXXX) Init(n, d string, v int) {
	w.Name = n
	w.Description = d
	w.Version = v
}

// DisplayContent will comprehensively display known information
func (w *WXXX) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (w *WXXX) GetExplain() string {
	return w.Description
}

// GetLength will provide the length
func (w *WXXX) GetLength() string {
	return ""
}

// GetName will provide the Name
func (w *WXXX) GetName() string {
	return w.Name
}

// ProcessData will handle the acquisition of all data
func (w *WXXX) ProcessData(s int, d []byte) IFrame {
	w.Size = s
	w.Data = d

	// text encoding is a single byte, 0 for latin, 1 for unicode
	if len(d) > 2 {
		enc := d[0]
		d = d[1:]

		if enc == '\x00' {
			idx := bytes.IndexByte(d, '\x00')
			w.Title = GetStr(d[:idx])
			w.URL = GetStr(d[idx+LengthStandard:])
		} else if enc == '\x01' {
			w.Utf16 = true

			idx := bytes.Index(d, []byte{'\x00', '\x00'})
			w.Title = GetUnicodeStr(d[:idx])
			w.URL = GetUnicodeStr(d[idx+LengthUnicode:])
		}
	}

	return w
}
