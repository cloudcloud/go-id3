package frames

import "bytes"

// SYLT defines the Synchronised lyrics/text
type SYLT struct {
	Frame

	Language    string   `json:"language"`
	Format      byte     `json:"format"`
	ContentType byte     `json:"content_type"`
	Descriptor  string   `json:"lyrics"`
	Items       [][]byte `json:"items"`
}

// Init will provide the initial values
func (y *SYLT) Init(n, d string, v int) {
	y.Name = n
	y.Description = d
	y.Version = v
}

// DisplayContent will comprehensively display known information
func (y *SYLT) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (y *SYLT) GetExplain() string {
	return y.Description
}

// GetLength will provide the length
func (y *SYLT) GetLength() string {
	return ""
}

// GetName will provide the Name
func (y *SYLT) GetName() string {
	return y.Name
}

// ProcessData will handle the acquisition of all data
func (y *SYLT) ProcessData(s int, d []byte) IFrame {
	y.Size = s
	y.Data = d

	enc := d[0]
	y.Language = GetStr(d[1:4])
	y.Format = d[4]
	y.ContentType = d[5]
	d = d[6:]

	if enc == '\x00' {
		idx := bytes.IndexByte(d, '\x00')
		y.Descriptor = GetStr(d[:idx])
		d = d[idx+1:]
	} else if enc == '\x01' {
		y.Utf16 = true

		idx := bytes.Index(d, []byte{'\x00', '\x00'})
		y.Descriptor = GetUnicodeStr(d[:idx])
		d = d[idx+2:]
	}

	// there's plenty more structure here, and this could probably be done better
	for {
		if len(d) < 1 {
			break
		}

		idx := bytes.IndexByte(d, '\x0A')
		if idx < 1 {
			break
		}

		y.Items = append(y.Items, d[:idx])
		d = d[idx+1:]
	}

	return y
}
