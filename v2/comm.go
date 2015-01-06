package v2

import "fmt"

type FComm struct {
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

func NewCOMM(n string) *FComm {
	c := new(FComm)

	c.Name = n

	c.TagPreserve = false
	c.FilePreserve = false
	c.ReadOnly = false
	c.Compression = false
	c.Encryption = false
	c.Grouping = false

	return c
}

func (t *FComm) Process(b []byte) []byte {
	fmt.Println("Have a COMM")

	return b
}
