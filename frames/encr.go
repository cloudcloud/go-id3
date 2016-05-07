package frames

import "fmt"

// ENCR contains information for encryption registration
type ENCR struct {
	Frame
}

// Init will provide the initial values
func (e *ENCR) Init(n, d string, v int) {
	e.Name = n
	e.Description = d
	e.Version = v
}

// DisplayContent will comprehensively display known information
func (e *ENCR) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (e *ENCR) GetExplain() string {
	return ""
}

// GetLength will provide the length of frame
func (e *ENCR) GetLength() string {
	return ""
}

// GetName will provide the Name of EQUA
func (e *ENCR) GetName() string {
	return e.Name
}

// ProcessData will parse bytes for details
func (e *ENCR) ProcessData(s int, d []byte) IFrame {
	e.Size = s
	e.Data = d

	fmt.Println("Unimplemented: ENCR frame")

	return e
}
