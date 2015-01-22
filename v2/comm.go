package v2

// FComm defines the Frame structure for the COMM tag
// part of Id3 V2.
type FComm struct {
	Frame
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

// GetExplain will return a string representing the type of content.
func (t *FComm) GetExplain() string {
	return "(Comments)"
}
