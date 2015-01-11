package v2

import "fmt"

// FText provides the structure for processing the TEXT frame
// type of Id3 V2. This struct is generic for most TEXT style
// frames.
type FText struct {
	*Frame

	Flags        int  `json:"flags"`
	TagPreserve  bool `json:"tag_preserve"`
	FilePreserve bool `json:"file_preserve"`
	ReadOnly     bool `json:"read_only"`
	Compression  bool `json:"compression"`
	Encryption   bool `json:"encryption"`
	Grouping     bool `json:"grouping"`
}

// NewTEXT will provision a new instance of the FText struct
// for processing.
func NewTEXT() *FText {
	c := new(FText)

	c.TagPreserve = false
	c.FilePreserve = false
	c.ReadOnly = false
	c.Compression = false
	c.Encryption = false
	c.Grouping = false

	return c
}

// Process will complete the processing within the provided bytes
// of the full Frame for TEXT.
func (t *FText) Process(b []byte) []byte {
	fmt.Println("Process out")
	return b

	t.Size = int(rune(b[0])<<21 | rune(b[1])<<14 | rune(b[2])<<7 | rune(b[3]))
	t.Flags = int(rune(b[4])<<8 | rune(b[5]))

	if b[4]&128 == 128 {
		t.TagPreserve = true
	}
	if b[4]&64 == 64 {
		t.FilePreserve = true
	}
	if b[4]&32 == 32 {
		t.ReadOnly = true
	}

	if b[5]&128 == 128 {
		t.Compression = true
	}
	if b[5]&64 == 64 {
		t.Encryption = true
	}
	if b[5]&32 == 32 {
		t.Grouping = true
	}

	b = b[6:]
	if b[0] == 0 {
		t.Data = string(b[1:t.Size])
	} else if b[0] == 1 {
		t.Data = GetUtf(b[1:t.Size])
	}

	fmt.Println("Processed")

	b = b[t.Size:]
	return b
}
