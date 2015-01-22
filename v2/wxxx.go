package v2

// FWxxx provides the structure for processing the W-based frames of
// the Id3 V2 spec. These frames are all containing of basic URLs.
type FWxxx struct {
	Frame
}

// NewWXXX will provision a new instance of the FWxxx struct
// for processing.
func NewWXXX(n string) *FWxxx {
	c := new(FWxxx)

	c.Name = n

	c.TagPreserve = false
	c.FilePreserve = false
	c.ReadOnly = false
	c.Compression = false
	c.Encryption = false
	c.Grouping = false

	return c
}

// GetExplain provides a description of the specific W-based frame.
func (t *FWxxx) GetExplain() string {
	a := "("

	switch t.Name {
	case "WXXX":
		a += "User defined URL link"
	}

	return a + ")"
}
