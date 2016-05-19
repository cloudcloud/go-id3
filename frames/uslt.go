package frames

import (
	"bytes"
	"fmt"
)

// USLT defines the Unsynchronised lyrics/text transcription
type USLT struct {
	Frame

	Language   string `json:"language"`
	Descriptor string `json:"descriptor"`
	Lyrics     string `json:"lyrics"`
}

// DisplayContent will comprehensively display known information
func (u *USLT) DisplayContent() string {
	return fmt.Sprintf("Unsynchronised Text (%s)\n"+
		"\t(%s): %s\n",
		u.Language,
		u.Descriptor,
		u.Lyrics)
}

// ProcessData will handle the acquisition of all data
func (u *USLT) ProcessData(s int, d []byte) IFrame {
	u.Size = s
	u.Data = d

	if d[0] == '\x01' {
		u.Utf16 = true
	}
	d = d[1:]

	u.Language = GetStr(d[:3])
	d = d[3:]

	if !u.Utf16 {
		idx := bytes.IndexByte(d, '\x00')
		u.Descriptor = GetStr(d[:idx])
		u.Lyrics = GetStr(d[idx+1:])
	} else {
		idx := bytes.Index(d, []byte{'\x00', '\x00'})
		u.Descriptor = GetUnicodeStr(d[:idx])
		u.Lyrics = GetUnicodeStr(d[idx+2:])
	}

	return u
}
