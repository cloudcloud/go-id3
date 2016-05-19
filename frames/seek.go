package frames

import (
	"fmt"
	"os"
)

// SEEK provides Seek details for the file
type SEEK struct {
	Frame

	SeekPoint int `json:"seek_point"`
}

// DisplayContent provides a clean display of key information
func (e *SEEK) DisplayContent() string {
	return fmt.Sprintf("Seek point: %d\n", e.SeekPoint)
}

// ProcessData will take bytes and mush into something useful
func (e *SEEK) ProcessData(s int, d []byte) IFrame {
	e.Size = s
	e.Data = d

	if len(e.Data) < 4 {
		fmt.Fprintf(os.Stderr, "SEEK frame invalid length [%d<4]\n", len(e.Data))

		return e
	}

	e.SeekPoint = GetSize(e.Data[:4], 8)

	return e
}
