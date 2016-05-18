package frames

import "fmt"

// RVAD is the frame for relative volume adjustment
type RVAD struct {
	Frame
}

// Init will provide the initial values
func (r *RVAD) Init(n, d string, v int) {
	r.Name = n
	r.Description = d
	r.Version = v
}

// DisplayContent will comprehensively display known information
func (r *RVAD) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (r *RVAD) GetExplain() string {
	return r.Description
}

// GetLength will provide the length
func (r *RVAD) GetLength() string {
	return ""
}

// GetName will provide the Name
func (r *RVAD) GetName() string {
	return r.Name
}

// ProcessData will handle the acquisition of all data
func (r *RVAD) ProcessData(s int, d []byte) IFrame {
	r.Size = s
	r.Data = d

	// there is quite a structure to this frame
	// all bytes are kept within Data, for processing later
	fmt.Println("Unimplemented: RVAD tag")

	return r
}
