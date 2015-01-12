package v2

import "fmt"

// FComm defines the Frame structure for the COMM tag
// part of Id3 V2.
type FComm struct {
	Frame

	Flags        int  `json:"flags"`
	TagPreserve  bool `json:"tag_preserve"`
	FilePreserve bool `json:"file_preserve"`
	ReadOnly     bool `json:"read_only"`
	Compression  bool `json:"compression"`
	Encryption   bool `json:"encryption"`
	Grouping     bool `json:"grouping"`
}

// NewCOMM will provision a new instance of the FComm struct
// and make it available for processing.
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

// Process completes the processing of the proceeding bytes
// for the COMM frame type.
func (t *FComm) Process(b []byte) []byte {
	fmt.Println("Have a COMM")

	return b
}
