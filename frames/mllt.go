package frames

// MLLT defines the mpeg location lookup table
type MLLT struct {
	Frame

	FramesBetween       []byte `json:"frames_between"`
	BytesBetween        []byte `json:"bytes_between"`
	MillisecondsBetween []byte `json:"milliseconds_between"`
	BitsForBytes        byte   `json:"bits_for_bytes"`
	BitsForMilliseconds byte   `json:"bits_for_milliseconds"`
	Deviations          []byte `json:"deviations"`
}

// Init will provide the initial values
func (m *MLLT) Init(n, d string, v int) {
	m.Name = n
	m.Description = d
	m.Version = v
}

// DisplayContent will comprehensively display known information
func (m *MLLT) DisplayContent() string {
	return ""
}

// GetExplain will provide output formatting briefly
func (m *MLLT) GetExplain() string {
	return m.Description
}

// GetLength will provide the length
func (m *MLLT) GetLength() string {
	return ""
}

// GetName will provide the Name
func (m *MLLT) GetName() string {
	return m.Name
}

// ProcessData will handle the acquisition of all data
func (m *MLLT) ProcessData(s int, d []byte) IFrame {
	m.Size = s
	m.Data = d

	if m.Size > 10 {
		m.FramesBetween = d[:2]
		m.BytesBetween = d[2:5]
		m.MillisecondsBetween = d[5:8]
		m.BitsForBytes = d[8]
		m.BitsForMilliseconds = d[9]
		m.Deviations = d[10:]
	}

	return m
}
