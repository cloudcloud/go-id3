package frames

// TEXT houses anything just for a TEXT frame
type TEXT struct {
	Frame
}

// Init will provide the initial values
func (t *TEXT) Init(n, d string, s int) {
	t.Name = n
	t.Description = d
	t.Size = s
}

// DisplayContent will comprehensively display known information
func (t *TEXT) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (t *TEXT) GetExplain() string {
	return t.Description
}

// GetLength will provide the length
func (t *TEXT) GetLength() string {
	return ""
}

// GetName will provide the Name
func (t *TEXT) GetName() string {
	return t.Name
}

// ProcessData will handle the acquisition of all data
func (t *TEXT) ProcessData(s int, d []byte) IFrame {
	t.Size = s
	t.Data = d

	// text encoding is a single byte, 0 for latin, 1 for unicode
	if len(d) > 2 {
		enc := d[0]
		d = d[1:]

		if enc == '\x00' {
			t.Frame.Cleaned = GetStr(d)
		} else if enc == '\x01' {
			t.Utf16 = true

			t.Frame.Cleaned = GetUnicodeStr(d)
		}
	}

	return t
}
