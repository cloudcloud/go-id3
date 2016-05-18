package frames

import "bytes"

// COMM contains the processing house for Comments
type COMM struct {
	Frame

	Language           string `json:"language"`
	ContentDescription string `json:"content_description"`
	Comment            string `json:"comment"`
}

// Init will provide the initial values
func (c *COMM) Init(n, d string, v int) {
	c.Name = n
	c.Description = d
	c.Version = v
}

// DisplayContent will comprehensively display known information
func (c *COMM) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (c *COMM) GetExplain() string {
	return c.Description
}

// GetLength will provide the length
func (c *COMM) GetLength() string {
	return ""
}

// GetName will provide the Name
func (c *COMM) GetName() string {
	return c.Name
}

// ProcessData will handle the acquisition of all data
func (c *COMM) ProcessData(s int, d []byte) IFrame {
	c.Size = s
	c.Data = d

	// text encoding is a single byte, 0 for latin, 1 for unicode
	if len(d) > 4 {
		enc := d[0]
		c.Language = GetStr(d[1:4])
		d = d[4:]

		if enc == '\x00' {
			idx := bytes.IndexByte(d, '\x00')
			c.ContentDescription = GetStr(d[:idx])
			c.Comment = GetStr(d[idx+1:])
		} else if enc == '\x01' {
			c.Utf16 = true

			idx := bytes.Index(d, []byte{'\x00', '\x00'})
			c.ContentDescription = GetUnicodeStr(d[:idx])
			c.Comment = GetUnicodeStr(d[idx+2:])
		}
	}

	return c
}
