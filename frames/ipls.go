package frames

import "bytes"

// IPLS provides the involved people list frame
type IPLS struct {
	Frame

	People map[string]string `json:"people"`
}

// Init will provide the initial values
func (i *IPLS) Init(n, d string, v int) {
	i.Name = n
	i.Description = d
	i.Version = v
}

// DisplayContent will comprehensively display known information
func (i *IPLS) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (i *IPLS) GetExplain() string {
	return i.Description
}

// GetLength will provide the length
func (i *IPLS) GetLength() string {
	return ""
}

// GetName will provide the Name
func (i *IPLS) GetName() string {
	return i.Name
}

// ProcessData will handle the acquisition of all data
func (i *IPLS) ProcessData(s int, d []byte) IFrame {
	i.Size = s
	i.Data = d

	// text encoding is a single byte, 0 for latin, 1 for unicode
	if len(d) > 2 {
		enc := d[0]
		d = d[1:]

		// loop through lines, should be even numbered
		for len(d) > 2 {
			if enc == '\x00' {
				idx := bytes.IndexByte(d, '\x00')
				name := GetStr(d[:idx])
				d = d[:idx+LengthStandard]

				idx = bytes.IndexByte(d, '\x00')
				if idx == -1 {
					break
				}
				i.People[name] = GetStr(d[:idx])
				d = d[idx+LengthStandard:]
			} else if enc == '\x01' {
				i.Utf16 = true

				idx := bytes.Index(d, []byte{'\x00', '\x00'})
				name := GetUnicodeStr(d[:idx])
				d = d[:idx+LengthUnicode]

				idx = bytes.Index(d, []byte{'\x00', '\x00'})
				if idx == -1 {
					break
				}
				i.People[name] = GetUnicodeStr(d[:idx+LengthUnicode])
				d = d[:idx+LengthUnicode]
			}
		}
	}

	return i
}
