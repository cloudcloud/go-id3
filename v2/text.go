package v2

import "fmt"

type FText struct {
	*Frame

	Name string `json:"name"`
	Size int    `json:"size"`

	Flags        int  `json:"flags"`
	TagPreserve  bool `json:"tag_preserve"`
	FilePreserve bool `json:"file_preserve"`
	ReadOnly     bool `json:"read_only"`
	Compression  bool `json:"compression"`
	Encryption   bool `json:"encryption"`
	Grouping     bool `json:"grouping"`
}

func NewTEXT(n string) *FText {
	c := new(FText)

	c.Name = n

	c.TagPreserve = false
	c.FilePreserve = false
	c.ReadOnly = false
	c.Compression = false
	c.Encryption = false
	c.Grouping = false

	return c
}

func (t *FText) Process(b []byte) []byte {
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

	fmt.Println("Have a TEXT")
	return b[6:]
}
