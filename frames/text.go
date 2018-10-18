package frames

import "fmt"

// TEXT houses anything just for a TEXT frame
type TEXT struct {
	Frame
}

// DisplayContent will comprehensively display known information
func (t *TEXT) DisplayContent() string {
	return fmt.Sprintf("[%s - %d] (%s) %s\n", t.Name, t.Size, t.Description, t.Cleaned)
}

// GetName will add deprecated notes where appropriate based on version
func (t *TEXT) GetName() string {
	out := t.Name
	if t.Version == Version4 {
		deprecated := map[string]bool{
			"TDAT": true,
			"TIME": true,
			"TORY": true,
			"TRDA": true,
			"TSIZ": true,
			"TYER": true,
		}

		if _, found := deprecated[t.Name]; found {
			out += " (deprecated)"
		}
	}

	return out
}

// ProcessData will handle the acquisition of all data
func (t *TEXT) ProcessData(s int, d []byte) IFrame {
	t.Size = s
	t.Data = d

	// text encoding is a single byte, 0 for latin, 1 for unicode
	if len(d) > 1 {
		if d[0] == '\x01' {
			t.Utf16 = true
		}
		d = d[1:]

		if !t.Utf16 {
			t.Frame.Cleaned = GetStr(d)
		} else {
			t.Frame.Cleaned = GetUnicodeStr(d)
		}
	}

	return t
}
