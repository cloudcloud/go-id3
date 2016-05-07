package frames

// USER provides terms of use for this file
type USER struct {
	Frame

	Encoding byte   `json:"encoding"`
	Language string `json:"language"`
	Text     string `json:"text"`
}

// Init will provide the initial values
func (u *USER) Init(n, d string, v int) {
	u.Name = n
	u.Description = d
	u.Version = v
}

// DisplayContent will comprehensively display known information
func (u *USER) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (u *USER) GetExplain() string {
	return ""
}

// GetLength will provide the length of frame
func (u *USER) GetLength() string {
	return ""
}

// GetName will provide the Name of EQUA
func (u *USER) GetName() string {
	return u.Name
}

// ProcessData will parse bytes for details
func (u *USER) ProcessData(s int, d []byte) IFrame {
	u.Size = s
	u.Data = d

	u.Encoding = d[0]
	u.Language = GetStr(d[1:4])

	if u.Encoding == '\x00' {
		u.Text = GetStr(d[4:])
	} else if u.Encoding == '\x01' {
		u.Utf16 = true

		u.Text = GetUnicodeStr(d[4:])
	}

	return u
}
