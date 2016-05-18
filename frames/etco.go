package frames

// ETCO provides timing codes for events in the file
type ETCO struct {
	Frame

	Format byte   `json:"encoding"`
	Codes  []byte `json:"codes"`
}

// Init will provide the initial values
func (e *ETCO) Init(n, d string, v int) {
	e.Name = n
	e.Description = d
	e.Version = v
}

// DisplayContent will comprehensively display known information
func (e *ETCO) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (e *ETCO) GetExplain() string {
	return e.Description
}

// GetLength will provide the length
func (e *ETCO) GetLength() string {
	return ""
}

// GetName will provide the Name
func (e *ETCO) GetName() string {
	return e.Name
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
