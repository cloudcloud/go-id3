package frames

import (
	"bytes"
	"fmt"
)

// WXXX provides user defined webpage links
type WXXX struct {
	Frame

	Title string `json:"title"`
	URL   string `json:"url"`
}

// DisplayContent will comprehensively display known information
func (w *WXXX) DisplayContent() string {
	return fmt.Sprintf("User webpage: %s [%s]\n", w.Title, w.URL)
}

// ProcessData will handle the acquisition of all data
func (w *WXXX) ProcessData(s int, d []byte) IFrame {
	w.Size = s
	w.Data = d

	// text encoding is a single byte, 0 for latin, 1 for unicode
	if len(d) > 2 {
		if d[0] == '\x01' {
			w.Utf16 = true
		}
		d = d[1:]

		if !w.Utf16 {
			idx := bytes.IndexByte(d, '\x00')
			w.Title = GetStr(d[:idx])
			w.URL = GetStr(d[idx+LengthStandard:])
		} else {
			idx := bytes.Index(d, []byte{'\x00', '\x00'})
			w.Title = GetUnicodeStr(d[:idx])
			w.URL = GetStr(d[idx+LengthUnicode:])
		}
	}

	return w
}
