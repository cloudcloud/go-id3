package frames

import "fmt"

// EQUA contains equalisation settings for the file
type EQUA struct {
	Frame

	Adjustment byte `json:"adjustment"`
}

// Init will provide the initial values
func (e *EQUA) Init(n, d string, v int) {
	e.Name = n
	e.Description = d
	e.Version = v
}

// DisplayContent will comprehensively display known information
func (e *EQUA) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (e *EQUA) GetExplain() string {
	return ""
}

// GetLength will provide the length of frame
func (e *EQUA) GetLength() string {
	return ""
}

// GetName will provide the Name of EQUA
func (e *EQUA) GetName() string {
	if e.Version == Version4 && e.Name == "EQUA" {
		return "EQUA (Deprecated)"
	}

	return e.Name
}

// ProcessData will parse bytes for details
func (e *EQUA) ProcessData(s int, d []byte) IFrame {
	e.Size = s

	e.Data = d[1:]
	e.Adjustment = d[0]

	// Data contains all frequency sets in groups 2 or 3 bytes
	fmt.Println("Unimplemented: EQUA frame")

	return e
}
