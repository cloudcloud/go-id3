package frames

// SYTC defines the synchronised tempo codes
type SYTC struct {
	Frame

	Format    byte   `json:"format"`
	TempoData []byte `json:"tempo_data"`
}

// Init will provide the initial values
func (z *SYTC) Init(n, d string, v int) {
	z.Name = n
	z.Description = d
	z.Version = v
}

// DisplayContent will comprehensively display known information
func (z *SYTC) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (z *SYTC) GetExplain() string {
	return z.Description
}

// GetLength will provide the length
func (z *SYTC) GetLength() string {
	return ""
}

// GetName will provide the Name
func (z *SYTC) GetName() string {
	return z.Name
}

// ProcessData will handle the acquisition of all data
func (z *SYTC) ProcessData(s int, d []byte) IFrame {
	z.Size = s
	z.Data = d

	z.Format = d[0]
	z.TempoData = d[1:]

	return z
}
