package frames

import "fmt"

// ETCO provides timing codes for events in the file
type ETCO struct {
	Frame

	Format byte   `json:"encoding"`
	Codes  []byte `json:"codes"`
}

// DisplayContent will comprehensively display known information
func (e *ETCO) DisplayContent() string {
	return fmt.Sprintf("Format: %d\nCode Count: %d", GetDirectInt(e.Format), len(e.Codes)/5)
}

// ProcessData will handle the acquisition of all data
func (e *ETCO) ProcessData(s int, d []byte) IFrame {
	e.Size = s
	e.Data = d

	// there is a format to this, it's just not enhanced here
	// format of 0 for mpeg frames, 1 for milliseconds
	e.Format = d[0]

	// groupings of 5, first byte is the type, 4 bytes for the time stamp
	e.Codes = d[1:]

	return e
}
