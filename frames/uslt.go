package frames

import "bytes"

// USLT defines the Unsynchronised lyrics/text transcription
type USLT struct {
	Frame

	Language   string `json:"language"`
	Descriptor string `json:"descriptor"`
	Lyrics     string `json:"lyrics"`
}

// Init will provide the initial values
func (u *USLT) Init(n, d string, v int) {
	u.Name = n
	u.Description = d
	u.Version = v
}

// DisplayContent will comprehensively display known information
func (u *USLT) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (u *USLT) GetExplain() string {
	return u.Description
}

// GetLength will provide the length
func (u *USLT) GetLength() string {
	return ""
}

// GetName will provide the Name
func (u *USLT) GetName() string {
	return u.Name
}

// ProcessData will handle the acquisition of all data
func (u *USLT) ProcessData(s int, d []byte) IFrame {
	u.Size = s
	u.Data = d

	enc := d[0]
	u.Language = GetStr(d[1:4])
	d = d[4:]

	term := []byte{'\x00'}
	if enc == '\x00' {
		idx := bytes.Index(d, term)
		u.Descriptor = GetStr(d[:idx])
		u.Lyrics = GetStr(d[idx+1:])
	} else if enc == '\x01' {
		u.Utf16 = true

		term = append(term, '\x00')
		idx := bytes.Index(d, term)
		u.Descriptor = GetUnicodeStr(d[:idx])
		u.Lyrics = GetUnicodeStr(d[idx+2:])
	}

	return u
}
