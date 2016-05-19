package frames

import "fmt"

// USER provides terms of use for this file
type USER struct {
	Frame

	Language string `json:"language"`
	Text     string `json:"text"`
}

// DisplayContent will comprehensively display known information
func (u *USER) DisplayContent() string {
	return fmt.Sprintf("Terms of use (%s): %s\n", u.Language, u.Text)
}

// ProcessData will parse bytes for details
func (u *USER) ProcessData(s int, d []byte) IFrame {
	u.Size = s
	u.Data = d

	if d[0] == '\x01' {
		u.Utf16 = true
	}
	d = d[1:]
	u.Language = GetStr(d[:3])

	if !u.Utf16 {
		u.Text = GetStr(d[3:])
	} else {
		u.Text = GetUnicodeStr(d[3:])
	}

	return u
}
