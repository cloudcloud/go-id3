package v2

import (
	//"bytes"
	"fmt"
	"os"
)

// FApic provides the structure for specifically processing the
// APIC frame. This frame contains an embedded image.
type FApic struct {
	Frame

	Description string `json:"description"`
	Encoding    int    `json:"encoding"`
	MimeType    string `json:"mime_type"`
	PictureType int    `json:"picture_type"`
}

// NewAPIC will provision a new instance of the FApic
// for processing.
func NewAPIC(n string) *FApic {
	c := new(FApic)

	c.Name = n

	c.TagPreserve = false
	c.FilePreserve = false
	c.ReadOnly = false
	c.Compression = false
	c.Encryption = false
	c.Grouping = false

	return c
}

// Process will complete the full processing of the frame. For
// APIC, this includes extra tricks.
func (t *FApic) Process(b []byte) []byte {
	b = t.processHeader(b)

	t.Encoding = int(b[0])
	t.MimeType, b = getToTerminus(b[1:], false)
	t.PictureType = int(b[0])
	b = b[1:]

	utf := false
	if t.Encoding == 1 {
		utf = true
	}

	t.Description, b = getToTerminus(b, utf)

	outfile, err := os.Create("/tmp/cbm.jpg")
	if err == nil {
		defer outfile.Close()

		outfile.Write(b[:36294])
	}

	pic := b[:36294]
	t.Data = pic
	b = b[len(pic):]

	fmt.Println(b)

	for z := 0; z < len(b); z++ {
		if b[z] < 91 && b[z] > 64 && b[z+1] < 91 && b[z+1] > 64 && b[z+2] < 91 && b[z+2] > 64 && b[z+3] < 91 && b[z+3] > 47 {
			fmt.Println(b[z:z+4], string(b[z:z+4]))
		}
	}

	//fmt.Println(b, len(b), buffer)
	return b
}

// DisplayContent will display a snip, not the picture.
func (t *FApic) DisplayContent() string {
	return "[truncated] [" + t.MimeType + "::" + t.Description + "]"
}

// GetExplain provides the textual description of the Frame.
func (t *FApic) GetExplain() string {
	return "(Attached picture)"
}
