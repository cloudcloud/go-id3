package frames

import (
	"bytes"
	"fmt"
)

// AENC contains the AENC frame for Audio encryption
type AENC struct {
	Frame

	Contact       string `json:"contact"`
	PreviewStart  []byte `json:"preview_start"`
	PreviewLength []byte `json:"preview_length"`
}

// Init will provide the initial values
func (a *AENC) Init(n, d string, v int) {
	a.Name = n
	a.Description = d
	a.Version = v
}

// DisplayContent will comprehensively display known information
func (a *AENC) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (a *AENC) GetExplain() string {
	return ""
}

// GetLength will provide the length of frame
func (a *AENC) GetLength() string {
	return ""
}

// GetName will provide the Name of AENC
func (a *AENC) GetName() string {
	return a.Name
}

// ProcessData will accept the incoming chunk and process it for the frame specifically
func (a *AENC) ProcessData(l int, b []byte) IFrame {
	a.Size = l
	a.Data = b

	// <text string> \x00 \xXX \xXX \xXX \xXX <binary data>
	term := bytes.IndexByte(b, '\x00')
	if term == 0 {
		fmt.Println("This file is encrypted and has no decryption detail")
		return a
	}

	a.Contact = GetStr(b[:term])
	a.Frame.Cleaned = GetStr(b[term:])

	fmt.Println("AENC is untested and possibly wrong")

	return a
}
