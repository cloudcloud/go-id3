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
	Encryption    []byte `json:"encryption"`
}

// DisplayContent will comprehensively display known information
func (a *AENC) DisplayContent() string {
	r := fmt.Sprintf("Contact: %s\n", a.Contact)
	r += fmt.Sprintf("PreviewStart: %#v\n", a.PreviewStart)
	r += fmt.Sprintf("PreviewLength: %#v\n", a.PreviewLength)
	r += fmt.Sprintf("Encryption: %#v\n", a.Encryption)

	return r
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
	a.PreviewStart = b[term+1 : term+3]
	a.PreviewLength = b[term+3 : term+5]
	a.Encryption = b[term+5:]

	fmt.Println("AENC is untested and possibly wrong")

	return a
}
